package chain

import (
	"math/big"
	"github.com/ArtosSystems/plasma/util"
)

type BlockResult struct {
	MerkleRoot         util.Hash
	NumberTransactions uint32
	BlockFees          *big.Int
	BlockNumber        *big.Int
}
