package cmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ArtosSystems/plasma/pkg/chain"
	"github.com/ArtosSystems/plasma/pkg/eth"
	"github.com/ArtosSystems/plasma/pkg/log"
	"github.com/ArtosSystems/plasma/pkg/rpc/pb"
	"github.com/ArtosSystems/plasma/util"
	"github.com/spf13/cobra"
	"math/big"
	"sort"
	"time"
)

type sendCmdOutput struct {
	Value            string   `json:"value"`
	To               string   `json:"to"`
	BlockNumber      uint64   `json:"blockNumber"`
	TransactionIndex uint32   `json:"transactionIndex"`
	DepositNonce     string   `json:"depositNonce"`
	MerkleRoot       string   `json:"merkleRoot"`
	ConfirmSigs      []string `json:"confirmSigs"`
}

var sendCmdLog = log.ForSubsystem("SendCmd")

type outpoint struct {
	tx  chain.ConfirmedTransaction
	idx uint8
}

var sendCmd = &cobra.Command{
	Use:   "send to value [depositNonce] [contractAddr]",
	Short: "Sends funds",
	RunE: func(cmd *cobra.Command, args []string) error {
		privKey, err := ParsePrivateKey(cmd)
		if err != nil {
			return err
		}
		from := crypto.PubkeyToAddress(privKey.PublicKey)
		to := common.HexToAddress(args[0])
		value, ok := new(big.Int).SetString(args[1], 10)
		if !ok {
			return errors.New("invalid send value")
		}

		url := cmd.Flag(FlagNodeURL).Value.String()
		if url == "" {
			return errors.New("no node url set")
		}
		client, conn, err := CreateRootClient(url)
		if err != nil {
			return err
		}
		defer conn.Close()

		if len(args) == 4 {
			depositNonce, ok := new(big.Int).SetString(args[2], 10)
			if !ok {
				return errors.New("invalid deposit nonce")
			}
			contractAddr := common.HexToAddress(args[3])
			contract, err := eth.NewClient(cmd.Flag(FlagEthereumNodeUrl).Value.String(), contractAddr.Hex(), privKey)
			if err != nil {
				return err
			}
			return SpendDeposit(client, contract, privKey, from, to, value, depositNonce)
		}

		return SpendTx(client, privKey, from, to, value)
	},
}

func SpendDeposit(client pb.RootClient, contract eth.Client, privKey *ecdsa.PrivateKey, from common.Address, to common.Address, value *big.Int, depositNonce *big.Int) error {
	sendCmdLog.Info("spending deposit")
	total, owner, err := contract.LookupDeposit(depositNonce)
	if err != nil {
		return err
	}
	if owner != from {
		return errors.New("you don't own this deposit")
	}
	if total.Cmp(value) < 0 {
		return errors.New("cannot send more than the deposit amount")
	}

	body := chain.ZeroBody()
	body.Input0.DepositNonce = depositNonce
	body.Output0.Amount = value
	body.Output0.Owner = to
	if total.Cmp(body.Output0.Amount) > 0 {
		totalClone := new(big.Int).Set(total)
		body.Output1.Amount = totalClone.Sub(totalClone, value)
		body.Output1.Owner = from
	}

	sig, err := eth.Sign(privKey, body.SignatureHash())
	if err != nil {
		return err
	}

	// no confirm sigs on deposits
	tx := &chain.Transaction{
		Body: body,
		Sigs: [2]chain.Signature{
			sig,
			sig,
		},
	}

	sendCmdLog.Info("sending spend message")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	sendRes, err := client.Send(ctx, &pb.SendRequest{
		Transaction: tx.Proto(),
	})
	if err != nil {
		return err
	}

	sendCmdLog.Info("confirming spend")

	tx.Body.BlockNumber = sendRes.Inclusion.BlockNumber
	tx.Body.TransactionIndex = sendRes.Inclusion.TransactionIndex
	var buf bytes.Buffer
	buf.Write(tx.RLPHash(util.Sha256))
	buf.Write(sendRes.Inclusion.MerkleRoot)
	sigHash := util.Sha256(buf.Bytes())
	confirmSig, err := eth.Sign(privKey, sigHash)
	if err != nil {
		return err
	}

	var emptySig chain.Signature
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	_, err = client.Confirm(ctx, &pb.ConfirmRequest{
		BlockNumber:      sendRes.Inclusion.BlockNumber,
		TransactionIndex: sendRes.Inclusion.TransactionIndex,
		ConfirmSig0:      confirmSig[:],
		ConfirmSig1:      emptySig[:],
	})
	if err != nil {
		return err
	}

	out := &sendCmdOutput{
		Value:            value.Text(10),
		To:               to.Hex(),
		BlockNumber:      sendRes.Inclusion.BlockNumber,
		TransactionIndex: sendRes.Inclusion.TransactionIndex,
		MerkleRoot:       hexutil.Encode(sendRes.Inclusion.MerkleRoot),
		ConfirmSigs: []string{
			hexutil.Encode(confirmSig[:]),
			hexutil.Encode(emptySig[:]),
		},
	}

	return PrintJSON(out)
}

func SpendTx(client pb.RootClient, privKey *ecdsa.PrivateKey, from common.Address, to common.Address, value *big.Int) error {
	sendCmdLog.Info("selecting outputs")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	res, err := client.GetOutputs(ctx, &pb.GetOutputsRequest{
		Address:   from.Bytes(),
		Spendable: true,
	})
	if err != nil {
		return err
	}

	var utxos []chain.ConfirmedTransaction
	for _, utxoProto := range res.ConfirmedTransactions {
		confirmedTx, err := chain.ConfirmedTransactionFromProto(utxoProto)
		if err != nil {
			return err
		}
		utxos = append(utxos, *confirmedTx)
	}
	if len(utxos) == 0 {
		return errors.New("no spendable outputs")
	}
	selectedTransactions, outputIndices, err := selectUTXOs(utxos, from, value)
	if err != nil {
		return err
	}

	total := big.NewInt(0)
	tx := &chain.Transaction{
		Body: chain.ZeroBody(),
	}
	for i, utxo := range selectedTransactions {
		txBody := utxo.Transaction.Body
		var input *chain.Input
		if i == 0 {
			input = tx.Body.Input0
		} else if i == 1 {
			input = tx.Body.Input1
		} else {
			panic("too many inputs!")
		}

		input.BlockNumber = txBody.BlockNumber
		input.TransactionIndex = txBody.TransactionIndex
		input.OutputIndex = outputIndices[i]

		if i == 0 {
			tx.Body.Input0ConfirmSigs = [2]chain.Signature{
				utxo.ConfirmSigs[0],
				utxo.ConfirmSigs[1],
			}
		} else {
			tx.Body.Input1ConfirmSigs = [2]chain.Signature{
				utxo.ConfirmSigs[0],
				utxo.ConfirmSigs[1],
			}
		}

		total = total.Add(total, txBody.OutputAt(outputIndices[i]).Amount)
	}

	tx.Body.Output0.Amount = value
	tx.Body.Output0.Owner = to

	if total.Cmp(value) > 0 {
		totalClone := new(big.Int).Set(total)
		tx.Body.Output1.Amount = totalClone.Sub(totalClone, value)
		tx.Body.Output1.Owner = from
	}

	sig, err := eth.Sign(privKey, tx.Body.SignatureHash())
	if err != nil {
		return err
	}
	tx.Sigs[0] = sig
	tx.Sigs[1] = sig

	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	sendRes, err := client.Send(ctx, &pb.SendRequest{
		Transaction: tx.Proto(),
	})
	if err != nil {
		return err
	}

	tx.Body.BlockNumber = sendRes.Inclusion.BlockNumber
	tx.Body.TransactionIndex = sendRes.Inclusion.TransactionIndex
	var buf bytes.Buffer
	buf.Write(tx.RLPHash(util.Sha256))
	buf.Write(sendRes.Inclusion.MerkleRoot)
	sigHash := util.Sha256(buf.Bytes())
	confirmSig0, err := eth.Sign(privKey, sigHash)
	if err != nil {
		return err
	}

	var confirmSig1 chain.Signature
	if !tx.Body.Input1.IsZero() {
		confirmSig1 = confirmSig0
	}

	sendCmdLog.Info("confirming transaction")

	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	_, err = client.Confirm(ctx, &pb.ConfirmRequest{
		BlockNumber:      sendRes.Inclusion.BlockNumber,
		TransactionIndex: sendRes.Inclusion.TransactionIndex,
		ConfirmSig0:      confirmSig0[:],
		ConfirmSig1:      confirmSig1[:],
	})
	if err != nil {
		return err
	}

	out := &sendCmdOutput{
		Value:            value.Text(10),
		To:               to.Hex(),
		BlockNumber:      sendRes.Inclusion.BlockNumber,
		TransactionIndex: sendRes.Inclusion.TransactionIndex,
		MerkleRoot:       hexutil.Encode(sendRes.Inclusion.MerkleRoot),
		ConfirmSigs: []string{
			hexutil.Encode(confirmSig0[:]),
			hexutil.Encode(confirmSig1[:]),
		},
	}

	return PrintJSON(out)
}

func selectUTXOs(confirmedTxs []chain.ConfirmedTransaction, addr common.Address, total *big.Int) ([]chain.ConfirmedTransaction, []uint8, error) {
	var outpoints []outpoint
	for _, tx := range confirmedTxs {
		indices := tx.Transaction.Body.OutputIndicesFor(&addr)
		for _, idx := range indices {
			outpoints = append(outpoints, outpoint{
				tx:  tx,
				idx: idx,
			})
		}
	}

	sort.Slice(outpoints, func(i, j int) bool {
		a := outpoints[i].tx.Transaction.Body.OutputAt(outpoints[i].idx).Amount
		b := outpoints[j].tx.Transaction.Body.OutputAt(outpoints[j].idx).Amount
		return a.Cmp(b) > 0
	})

	first := outpoints[0]
	firstBody := first.tx.Transaction.Body

	if firstBody.OutputAt(first.idx).Amount.Cmp(total) >= 0 {
		return []chain.ConfirmedTransaction{first.tx}, []uint8{first.idx}, nil
	}

	for i := len(confirmedTxs) - 1; i >= 0; i-- {
		second := outpoints[i]
		secondBody := second.tx.Transaction.Body
		sum := big.NewInt(0)
		sum = sum.Add(sum, firstBody.OutputAt(first.idx).Amount)
		sum = sum.Add(sum, secondBody.OutputAt(second.idx).Amount)
		if sum.Cmp(total) >= 0 {
			return []chain.ConfirmedTransaction{
				first.tx,
				second.tx,
			}, []uint8 { first.idx, second.idx }, nil
		}
	}

	return nil, nil, errors.New("no suitable UTXOs found")
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringP(FlagEthereumNodeUrl, "e", "http://34.241.211.99:8545", "URL to a running Ethereum node.")
}
