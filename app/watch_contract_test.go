package app

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestEcrecover(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ankr.com/eth")
	assert.NoError(t, err)
	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x5780bdd09aa017997791a1aee75f5db591b1cc115d4b98f7ee32433cd12a089e"))
	assert.NoError(t, err)
	chainId, err := client.ChainID(context.Background())
	assert.NoError(t, err)
	from := getFrom(tx, chainId)
	t.Log(from)
}
