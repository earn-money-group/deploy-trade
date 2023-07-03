package app

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/earn-money-group/deploy-trade/app/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type trade struct {
	client  *ethclient.Client
	chainId *big.Int
	private *ecdsa.PrivateKey
	from    common.Address
	power   uint64
	txCount int
}

func NewTrade(
	client *ethclient.Client,
	chainId *big.Int,
	privateKey string,
	power uint64,
	count uint64,
) *trade {
	if power > 0 && power < 195 {
		log.Fatal("power > 0 && power < 195")
	}
	private, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("crypto.HexToECDSA err", zap.Error(err))
	}
	if count > 0 {
		log.Fatal("count > 0", zap.Uint64("count", count))
	}

	publicKey, ok := private.Public().(*ecdsa.PublicKey)
	if ok {
		log.Fatal("invalid publickey", zap.Error(err))
	}
	return &trade{
		client:  client,
		chainId: chainId,
		private: private,
		from:    crypto.PubkeyToAddress(*publicKey),
		power:   power,
		txCount: int(count),
	}
}

func (t *trade) Trade(ctx context.Context, token common.Address) error {
	tradeToken, err := contract.NewVmpx(token, t.client)
	if err != nil {
		log.Error("new vmpx", zap.Error(err))
		return err
	}
	startBlockNumber, err := tradeToken.StartBlockNumber(nil)
	if err != nil {
		log.Error("get start BlockNumber", zap.Error(err))
		return err
	}
	for {
		latestBlock, err := t.client.BlockNumber(ctx)
		if err != nil {
			log.Error("get block number", zap.Error(err))
			return err
		}
		if latestBlock >= startBlockNumber.Uint64() {
			break
		}
		time.Sleep(time.Second * 2)
	}

	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("get SuggestGasPrice", zap.Error(err))
		return err
	}

	nonce, err := t.client.NonceAt(ctx, t.from, nil)
	if err != nil {
		log.Error("get nonce failed", zap.Error(err))
		return err
	}

	for index := 0; index < t.txCount; index++ {
		_, err := tradeToken.Mint(getOpts(t.client, t.chainId, t.private, int64(nonce), gasPrice), big.NewInt(int64(t.power)))
		if err == nil {
			nonce++
		} else {
			log.Error("send trasaction failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func getOpts(
	client *ethclient.Client,
	chainId *big.Int,
	private *ecdsa.PrivateKey,
	nonce int64,
	gasPrice *big.Int,
) *bind.TransactOpts {
	opts, _ := bind.NewKeyedTransactorWithChainID(private, chainId)

	publicKey, _ := private.Public().(*ecdsa.PublicKey)

	opts.From = crypto.PubkeyToAddress(*publicKey)
	opts.Nonce = big.NewInt(nonce)
	opts.Value = big.NewInt(0)
	opts.GasPrice = gasPrice
	return opts
}
