package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// Default Metamask Infura url
	nodeUrl = flag.String("u", "wss://goerli.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161", "Ethereum node URL")
	// Default Goerli block to check https://goerli.etherscan.io/block/5410648
	blockNum  = flag.Int64("n", 5410648, "block number to check")
	blockHash = HashValue{
		Hash: common.HexToHash("0xb85f4e8338828821a821217fbc8c501b21b8752d71b77fe0dae35783d74423d3"),
	}
)

type HashValue struct {
	common.Hash
}

func (v *HashValue) Set(s string) error {
	v.Hash = common.HexToHash(s)
	return nil
}

func main() {
	flag.Var(&blockHash, "h", "block hash to check")
	flag.Parse()
	if nodeUrl == nil || *nodeUrl == "" {
		log.Fatal("Node URL must be provided.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := ethclient.DialContext(ctx, *nodeUrl)
	failOnError("Dialing node", err)

	byNum, err := cl.HeaderByNumber(ctx, big.NewInt(*blockNum))
	failOnError("HeaderByNumber", err)
	checkHeader("HeaderByNumber", byNum)
	printHeader(byNum)

	byHash, err := cl.HeaderByHash(ctx, blockHash.Hash)
	failOnError("HeaderByHash", err)
	checkHeader("HeaderByHash", byHash)
	printHeader(byHash)

	if byNum.Hash() == byHash.Hash() {
		log.Print("✅ HeaderbyNumber/Hash return headers with same hash value.")
	}

	_, err = cl.HeaderByHash(ctx, byHash.Hash())
	if err != nil {
		log.Printf("🐛 (Expected) error getting header again with own hash: %v", err)
	}

	nextHeader, err := cl.HeaderByNumber(ctx, big.NewInt(*blockNum+1))
	failOnError("HeaderByNumber+1", err)
	if nextHeader.ParentHash != byNum.Hash() {
		log.Printf("🐛 ParentHash mismatch, expected %s, got %s",
			nextHeader.ParentHash.String(), byNum.Hash())
	}
}

func failOnError(desc string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", desc, err)
	}
}

func checkHeader(desc string, h *types.Header) {
	if num := h.Number.Int64(); num != *blockNum {
		log.Printf("🐛 %s Number mismatch, expected %d, got %d",
			desc, blockNum, num)
	}
	if hash := h.Hash(); hash != blockHash.Hash {
		log.Printf("🐛 %s Hash mismatch, expected %s, got %s",
			desc, blockHash.String(), hash.String())
	}
}

func printHeader(h *types.Header) {
	log.Printf("Num: %v, Hash: %s, Header:\n%+v\n", h.Number, h.Hash().String(), h)
}
