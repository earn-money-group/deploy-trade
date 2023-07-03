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
	trade   *trade
}

func NewWatch(
	client *ethclient.Client,
	privateKey string,
	power uint64,
	count uint64,
) *watchDeployedContract {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("get chainId", zap.Error(err))
	}
	trade := NewTrade(client, chainId, privateKey, power, count)

	return &watchDeployedContract{
		client:  client,
		chainId: chainId,
		trade:   trade,
	}
}

func (w *watchDeployedContract) Watch(ctx context.Context) {
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
		for index < txs.Len() {
			if txs[index].To() == nil {

				contractAddr := crypto.CreateAddress(getFrom(txs[index], w.chainId), txs[index].Nonce())
				code, err := w.client.CodeAt(ctx, contractAddr, nil)
				if err != nil {
					continue
				}
				if bytes.Contains(code, common.HexToHash("0x9af3d88fa3cc7ccaebec09593bdc0498b09e10865caff051172d4c36489d913c").Bytes()) {
					log.Info(
						"watch vmpx contract",
						zap.String("contract", contractAddr.Hex()),
					)
					err = w.trade.Trade(ctx, contractAddr)
					if err == nil {
						log.Info("mint completed")
						return
					} else {
						log.Error("trade failed", zap.Error(err))
					}
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
