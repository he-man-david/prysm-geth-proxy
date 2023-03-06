package rpcapis

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// *** see GETH's internal/ethapi, for examples of how to register APIs
// services are actual implementations, registered with a namespace
func GetAPIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "eth",
			Service:   &EthInterceptors{},
		},
		{
			Namespace: "engine",
			Service:   &EngineInterceptors{},
		},
	}
}

type EthInterceptors struct {

}

func (r *EthInterceptors) ChainId() *hexutil.Big {
	log.Println(`[EthInterceptors] intercepted --------------- ChainId`)
	return (*hexutil.Big)(big.NewInt(1))
}

func (r * EthInterceptors) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	log.Println(`[EthInterceptors] intercepted --------------- GetBlockByNumber`)
	return nil, nil
}


type EngineInterceptors struct {

}

func (r * EngineInterceptors) ExchangeTransitionConfigurationV1(config engine.TransitionConfigurationV1) (*engine.TransitionConfigurationV1, error) {
	log.Println(`[EngineInterceptors] intercepted --------------- ExchangeTransitionConfigurationV1`)
	return nil, nil
}