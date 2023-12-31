package cmd

import (
	"context"
	"errors"
	"github.com/ArtosSystems/plasma/pkg/chain"
	"github.com/ArtosSystems/plasma/pkg/rpc/pb"
	"github.com/spf13/cobra"
	"time"
)

type utxoCmdOutput struct {
	BlockNumber      uint64 `json:"blockNumber"`
	TransactionIndex uint32 `json:"transactionIndex"`
	OutputIndex      uint8  `json:"outputIndex"`
	Amount           string `json:"amount"`
}

var utxosCmd = &cobra.Command{
	Use:   "utxos [addr]",
	Short: "Returns the UTXOs for a given address",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, err := AddrOrPrivateKeyAddr(cmd, args, 0)
		if err != nil {
			return err
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

		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		res, err := client.GetOutputs(ctx, &pb.GetOutputsRequest{
			Address:   addr.Bytes(),
			Spendable: true,
		})
		if err != nil {
			return err
		}

		out := make([]utxoCmdOutput, 0)

		for _, conf := range res.ConfirmedTransactions {
			deser, err := chain.ConfirmedTransactionFromProto(conf)
			if err != nil {
				return err
			}
			tx := deser.Transaction.Body

			indices := tx.OutputIndicesFor(&addr)
			for _, idx := range indices {
				out = append(out, utxoCmdOutput{
					BlockNumber:      tx.BlockNumber,
					TransactionIndex: tx.TransactionIndex,
					OutputIndex:      idx,
					Amount:           tx.OutputAt(idx).Amount.Text(10),
				})
			}
		}

		return PrintJSON(out)
	},
}

func init() {
	rootCmd.AddCommand(utxosCmd)
}
