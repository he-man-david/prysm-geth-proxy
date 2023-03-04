package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	ctx = context.Background()
	testKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddr    = crypto.PubkeyToAddress(testKey.PublicKey)
)

func main() {
	// TEST again when light clients are available. Don't have DISK space to run Full Node myself now.
	ethcl, err := ethclient.Dial("https://weathered-broken-pond.discover.quiknode.pro/92d7a1da0e2c55ffa42c3e677973b841c9d7f1b4")
	if err != nil {
		log.Fatalf("[testclient] failed to dial via ethclient :: 8545, ERR: %v", err)
	}

	// Retrieve the pending nonce for an account
	chainID, err := ethcl.ChainID(ctx)
	nonce, err := ethcl.NonceAt(ctx, testAddr, nil)
	to := common.HexToAddress("0xABCD")
	amount := big.NewInt(10 * params.GWei)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(10 * params.GWei)
	data := []byte{}
	// Create a raw unsigned transaction
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), testKey)
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		log.Fatal(err)
	}
	// Send transaction
	err = ethcl.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
}
