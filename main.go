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

// https://goerli.etherscan.io/block/5410648
const blockNum = 5410648

var blockHash = common.HexToHash("0xb85f4e8338828821a821217fbc8c501b21b8752d71b77fe0dae35783d74423d3")

var nodeUrl = flag.String("n", "", "Ethereum node URL")

func main() {
	flag.Parse()
	if nodeUrl == nil || *nodeUrl == "" {
		log.Fatal("Node URL must be provided.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := ethclient.DialContext(ctx, *nodeUrl)
	failOnError("Dialing node", err)

	byNum, err := cl.HeaderByNumber(ctx, big.NewInt(blockNum))
	failOnError("HeaderByNumber", err)
	checkHeader("HeaderByNumber", byNum)
	printHeader(byNum)

	byHash, err := cl.HeaderByHash(ctx, blockHash)
	failOnError("HeaderByHash", err)
	checkHeader("HeaderByHash", byHash)
	printHeader(byHash)

	if byNum.Hash() == byHash.Hash() {
		log.Print("✅ HeaderbyNumber/Hash return headers with same hash value.")
	}

	_, err = cl.HeaderByHash(ctx, byHash.Hash())
	if err != nil {
		log.Printf("(Expected) error getting header with wrong hash: %v", err)
	}
}

func failOnError(desc string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", desc, err)
	}
}

func checkHeader(desc string, h *types.Header) {
	if num := h.Number.Int64(); num != blockNum {
		log.Printf("🐛 %s Number mismatch, expected %d, got %d",
			desc, blockNum, num)
	}
	if hash := h.Hash(); hash != blockHash {
		log.Printf("🐛 %s Hash mismatch, expected %s, got %s",
			desc, blockHash.String(), hash.String())
	}
}

func printHeader(h *types.Header) {
	log.Printf("Num: %v, Hash: %s, Header:\n%+v\n", h.Number, h.Hash().String(), h)
}
