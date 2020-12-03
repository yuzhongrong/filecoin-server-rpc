package test

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/client"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
	"testing"
)

// 发送FileCoin
func TestClient_MpoolPush(t *testing.T) {
	c := testClient()

	from, _ := address.NewFromString("t1blnsfsy2v4x6bi2bkr7aakkxmmel3igfpuxq36i")
	to, _ := address.NewFromString("t1au7uretcjjtjbu42p7ivtn5ebpn7kazl7pg4poa")

	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      client.FromFil(decimal.NewFromFloat(0.01)),
		GasLimit:   0,
		GasFeeCap:  decimal.Zero,
		GasPremium: decimal.Zero,
		Method:     0,
		Params:     nil,
	}

	maxFee := client.FromFil(decimal.NewFromFloat(0.01))
	msg, err := c.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		t.Error(err)
	}

	actor, err := c.StateGetActor(context.Background(), msg.From, nil)
	if err != nil {
		t.Error(err)
	}

	msg.Nonce = actor.Nonce
	t.Log(msg)

	sm, err := c.WalletSignMessage(context.Background(), msg.From, msg)
	if err != nil {
		t.Error(err)
	}

	id, err := c.MpoolPush(context.Background(), sm)
	if err != nil {
		t.Error(err)
	}

	t.Log(id.Version())
	t.Log(id)
}
