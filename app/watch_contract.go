package app

import (
	"bytes"
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type watchDeployedContract struct {
	client  *ethclient.Client
	chainId *big.Int
}

func NewWatch(
	client *ethclient.Client,
) *watchDeployedContract {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("get chainId", zap.Error(err))
	}

	return &watchDeployedContract{
		client:  client,
		chainId: chainId,
	}
}

func (w *watchDeployedContract) Watch(ctx context.Context) {
	log.Info("watch deployed contract...")
	var lastBlock uint64
	for {
		blockNumber, err := w.client.BlockNumber(ctx)
		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		}
		if blockNumber <= lastBlock {
			time.Sleep(3 * time.Second)
			continue
		}

		lastBlock = blockNumber

		block, err := w.client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
		if err != nil {
			log.Error("get block", zap.Uint64("block", blockNumber), zap.Error(err))
			time.Sleep(2 * time.Second)
			continue
		}
		index := 0
		txs := block.Transactions()
		len := txs.Len()
		for index < len {
			if txs[index].To() == nil {
				contractAddr := crypto.CreateAddress(getFrom(txs[index], w.chainId), txs[index].Nonce())
				log.Debug("new contract", zap.String("contract", contractAddr.Hex()))

				code, err := w.client.CodeAt(ctx, contractAddr, nil)
				if err != nil {
					continue
				}
				if bytes.Contains(code, common.HexToHash("0x404d724a61636b4c6576696e204061636b65626f6d20406c62656c7961657620404a616d6d614265616e73206661697263727970746f2e6f7267").Bytes()) {
					log.Info(
						"watch vmpx contract",
						zap.String("contract", contractAddr.Hex()),
					)
				}
			}
			index++
		}
	}
}

func getFrom(tx *types.Transaction, chainId *big.Int) common.Address {
	signer := types.LatestSignerForChainID(chainId)
	from, _ := types.Sender(signer, tx)
	return from
}
