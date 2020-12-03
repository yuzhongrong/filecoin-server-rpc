package test

import (
	"github.com/myxtype/filecoin-client/client"
	"github.com/shopspring/decimal"
	"testing"
)

func TestToFil(t *testing.T) {
	v := decimal.RequireFromString("79999999877989641883")
	t.Log(client.ToFil(v).String())
}

func TestFromFil(t *testing.T) {
	v := decimal.RequireFromString("79.9999998779896419")
	t.Log(client.FromFil(v))
}
