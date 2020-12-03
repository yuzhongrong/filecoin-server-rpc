package test

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/client"
	"github.com/myxtype/filecoin-client/crypto"
	"log"
	"testing"
)

// 查询节点上账号余额
func TestClient_WalletBalance(t *testing.T) {
	c := testClient()
	list,_:=c.WalletList(context.Background())
	for _,item:=range list{
		addr, _ := address.NewFromString(item.String())
		b, err := c.WalletBalance(context.Background(), addr)
		if err != nil {
			t.Error(err)
		}
		t.Log(item.String(),b.String())
		t.Log(item.String(),client.ToFil(b).String())
		t.Log(item.String(),client.FromFil(client.ToFil(b)).String())
	}

}

// 查询节点上账号余额
func TestClient_WalletBalance1(t *testing.T) {
	c := testClient()
	addr, _ := address.NewFromString("t3s6aokaqafigv44ai4ojforfsvpnhhzcr6ct5cwfjgpd45x3ti77sntvaa42pogm7jmwrgj3vevseoog3ny5q")
	b, err := c.WalletBalance(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}
	t.Log(addr.String(),b.String())
	t.Log(addr.String(),client.ToFil(b).String())
	t.Log(addr.String(),client.FromFil(client.ToFil(b)).String())
}


//节点上列出地址列表
func TestClient_WalletAddressList(t *testing.T){
	c := testClient()
	list,_:=c.WalletList(context.Background())
	for _,item:=range list{
		t.Log(item.String())
	}

}

//节点上判断钱包地址合法性
func TestClient_WalletAddressVerify(t *testing.T){
	c := testClient()
	validate,err:=c.WalletValidateAddress(context.Background(),"f1j2r33ztpldiugtytswhwcooxjjpraond3fmnooy")
	if err != nil {
		t.Error(err)
	}
	t.Log(validate)
}


//节点上生成一个账号
func TestClient_WalletNew(t *testing.T) {
	c := testClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), crypto.SigTypeSecp256k1)
	if err != nil {
		t.Error(err)
	}
	t.Log(addr)

	ki, err := c.WalletExport(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}
	// secp256k1 fd1d429f2e0744f5dbcc361796e1a6f5cf4b59ecca92c15c27f837401c12a3da
	t.Log(ki.Type, hex.EncodeToString(ki.PrivateKey))
}



//判断是不是这个节点上的钱包的地址
func TestClient_HasAddress(t *testing.T){
	c := testClient()
	account:="t122bhm536gh25zycdqshl43dh7t2lkdtwprd2fmi"
	account_addr, _ := address.NewFromString(account)
	if b, err :=c.WalletHas(context.Background(),account_addr);err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(b)
	}
}

//测试创建新账号
func TestClient_NewAccount(t *testing.T){
	c := testClient()
	addr, err := c.WalletNew(context.Background(), crypto.SigTypeSecp256k1)
	if err != nil {
		log.Fatal(err)
	}
	println(addr.String())
}
