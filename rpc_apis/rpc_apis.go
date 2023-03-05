package rpcapis

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// *** see GETH's internal/ethapi, for examples of how to register APIs
// services are actual implementations, registered with a namespace
func GetAPIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "eth",
			Service:   NewMockEthApis(),
		},
	}
}

type MockEthApis struct {

}

func NewMockEthApis() *MockEthApis {
	return &MockEthApis{}
}

func (e *MockEthApis) ChainId() *hexutil.Big {
	log.Println(`
	===================================================================

	YAYYYYYY MOTHER FUCKER!!!!!!!!!!!!!!!!

	===================================================================
	`)
	return (*hexutil.Big)(big.NewInt(1))
}