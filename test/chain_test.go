package test

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

// 根据消息Cid获取消息
func TestClient_ChainGetMessage(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacebv7h5xsceyz6aqa4robpcjitdmldzkwydklmqzdvhw76d5zikono")
	if err != nil {
		t.Error(err)
	}

	msg, err := c.ChainGetMessage(context.Background(), id)
	if err != nil {
		t.Error(err)
	}

	t.Log(msg)
}


// 获取当前头部高度
func TestClient_ChainHead(t *testing.T) {
	c := testClient()
	ts, err := c.ChainHead(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Log(ts.Height)

	for _, n := range ts.Cids {
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			t.Error(err)
		}
		for index, msg := range bm.BlsMessages {
			t.Log(ts.Height,bm.Cids[index], msg)
		}
	}
}

// 根据高度遍历区块所有交易
func TestClient_ChainGetTipSetByHeight(t *testing.T) {
	c := testClient()
	ts, err := c.ChainGetTipSetByHeight(context.Background(), 143366, nil)
	if err != nil {
		t.Error(err)
	}
	for _, n := range ts.Cids {//遍历矿工cid集合
		t.Log("遍历子块",n)
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			t.Error(err)
		}

		t.Log("子块SecpkMessages长度",len(bm.SecpkMessages))
		//
		//for cidindex,cid:=range bm.Cids{
		//	t.Log(cidindex,cid)
		//}
		for subindex, msg := range bm.SecpkMessages {
			subcid:= len(bm.BlsMessages)+subindex
			cid:=bm.Cids[subcid]
			t.Log(subcid,cid,msg.Message)
		}
	}
}


func TestClient_ChainGetLastHeight(t *testing.T){

	c := testClient()
	ts, err := c.ChainHead(context.Background())
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(ts.Height)

}

func TestClient_ChainGetTipSetByHeight1(t *testing.T){
	url := "http://202.122.108.16:1234/rpc/v0"
	method := "POST"

	payload := strings.NewReader("{\n\"jsonrpc\": \"2.0\",\n\"method\": \"Filecoin.ChainGetBlockMessages\",\n\"id\": 1,\n\"params\": [{\"/\": \"bafy2bzacecjsjcw5qkptuy4ams62zefuw4cgkbsahnpuqofwwop7cw7e6zwwu\" }]\n}\n")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.kKPBuX4GW81hMze3OksFMpPPWI5l8r_h2vDr0CshDs0")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}




