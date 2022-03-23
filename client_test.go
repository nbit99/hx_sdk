package hx_sdk

import (
	"github.com/nbit99/hx_sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	testNet = "wss://node.testnet.bitshares.eu"
	mainNet = "ws://47.57.233.50:19099"
)

func TestClient(t *testing.T) {
	t.Run("valid ws url", func(t *testing.T) {
		_, err := NewClient(mainNet)
		require.NoError(t, err)
	})

	t.Run("invalid ws url", func(t *testing.T) {
		_, err := NewClient("wss://invalid")
		require.Error(t, err)
	})
}

func Test_getblock(t *testing.T) {
	cli, err := NewClient(mainNet)
	t.Log(err)

	blk, err := cli.Database.GetBlock(19969663)

	op := blk.Transactions[0].Operations[0]
	tx := op.(*types.TransferOperation)
	t.Log(tx.Amount.Amount)
	t.Log(err)
}

func Test_getblockerheader(t *testing.T) {
	cli, err := NewClient(mainNet)
	t.Log(err)

	blk, err := cli.Database.GetBlock(19969663)
	op := blk.Transactions[0].Operations[0]
	tx := op.(*types.TransferOperation)
	t.Log(tx.Amount.Amount)
	t.Log(err)


	blkh, err := cli.Database.GetBlockHeader(19969663)
	t.Log(blkh.Previous)
}

