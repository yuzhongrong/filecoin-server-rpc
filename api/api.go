package api

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/client"
	"github.com/myxtype/filecoin-client/crypto"
	"github.com/myxtype/filecoin-client/manager"
	"github.com/myxtype/filecoin-client/models"
	"github.com/myxtype/filecoin-client/pkg/setting"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
	//"time"

	"log"
)




func ConnectClient() *client.Client {

	return client.NewClient(setting.SERVER_HOST, setting.TOKEN)
}

//upgrade transations from chain
func UpgradeTransations(account string, height int64) (*types.BlockMessages1, error) {
	fmt.Sprint(height)
	var Messages []*types.Message
	var Cids []cid.Cid
	c := ConnectClient()
	ts, err := c.ChainGetTipSetByHeight(context.Background(), height, nil)

	if err != nil {
		return nil, err
	}
	for _, n := range ts.Cids {
		log.Println("开始遍历子块", height, n.String())
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			return nil, err
		}
		//fmt.Println("子块SecpkMessages长度", len(bm.SecpkMessages))
		for subindex, item := range bm.SecpkMessages {
			fmt.Println(item.Message)
			if item.Message.To.String() == account { //找和自己有关的交易
				fmt.Println("找到一笔自己的交易")
				subcid := len(bm.BlsMessages) + subindex
				cid := bm.Cids[subcid]
				fmt.Println(subcid, cid, item.Message)
				Messages = append(Messages, item.Message)
				Cids = append(Cids, cid)
			}

		}
	}
	return &types.BlockMessages1{BlsMessages: Messages, Cids: Cids, Height: ts.Height}, nil

}

func UpgradeTransations1(height int64) (*types.BlockMessages1, error) {

	var Messages []*types.Message
	var Cids []cid.Cid
	c := ConnectClient()
	ts, err := c.ChainGetTipSetByHeight(context.Background(), height, nil)

	if err != nil {
		return nil, err
	}
	for _, n := range ts.Cids {
		log.Println("开始遍历子块", height, n.String())
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			return nil, err
		}

		//fmt.Println("********开始查找 BlsMessages********",len(bm.BlsMessages))
		for index, item := range bm.BlsMessages{

			//fmt.Println("********输出所有的bm.BlsMessages-cid********",len(bm.BlsMessages),index,bm.Cids[index])
			//log.Println("********输出所有的bm.BlsMessages-cid********",len(bm.BlsMessages),index,bm.Cids[index])
			if b, direct, err := hasAddress(c, item.From.String(), item.To.String()); err != nil { //找属于本钱包节点的冲币地址
				log.Println("*****解析bm.BlsMessages出错*****" + err.Error())
			} else {
				if b {
					log.Println("找到一笔bm.BlsMessages出错交易",bm.Cids[index],item.From.String(),item.To.String())
					item.Direct = direct //设置方向
					Messages = append(Messages, item)
					Cids = append(Cids, bm.Cids[index])
				}
			}

		}

		//fmt.Println("子块SecpkMessages长度", len(bm.SecpkMessages))
		//log.Println("********开始查找 SecpkMessages********",len(bm.SecpkMessages))
		for subindex, item := range bm.SecpkMessages {

			if b, direct, err := hasAddress(c, item.Message.From.String(), item.Message.To.String()); err != nil { //找属于本钱包节点的冲币地址
				log.Println("*****解析bm.SecpkMessages出错*****" + err.Error())
			} else {
				if b {
					log.Println("找到一笔bm.SecpkMessages交易")
					subcid := len(bm.BlsMessages) + subindex
					cid := bm.Cids[subcid]
					log.Println(subcid, cid, item.Message)
					item.Message.Direct = direct //设置方向
					Messages = append(Messages, item.Message)
					Cids = append(Cids, cid)
				}

			}


		}




	}
	return &types.BlockMessages1{BlsMessages: Messages, Cids: Cids, Height: ts.Height}, nil

}

func NewAccount() (string, error) {
	c := ConnectClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), crypto.SigTypeSecp256k1)
	if err != nil {
		return "", nil
	}
	println(addr.String())
	ki, err := c.WalletExport(context.Background(), addr)
	if err != nil {
		return "", nil
	}
	// secp256k1 fd1d429f2e0744f5dbcc361796e1a6f5cf4b59ecca92c15c27f837401c12a3da
	//t.Log(ki.Type, hex.EncodeToString(ki.PrivateKey))
	println(ki.Type, hex.EncodeToString(ki.PrivateKey))
	return addr.String(), nil
}



func NewAccount1() (string,string, error) {
	c := ConnectClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), crypto.SigTypeSecp256k1)

	addr1, err1 := c.WalletNew(context.Background(), crypto.SigTypeSecp256k1)
	if err != nil {
		return "","", nil
	}

	if err1 != nil {
		return "","", nil
	}
	println("入金账号"+addr.String())
	println("出金账号"+addr1.String())
	ki, err := c.WalletExport(context.Background(), addr)
	if err != nil {
		return "","", nil
	}

	ki1, err1 := c.WalletExport(context.Background(), addr)
	if err1 != nil {
		return "","", nil
	}

	// secp256k1 fd1d429f2e0744f5dbcc361796e1a6f5cf4b59ecca92c15c27f837401c12a3da
	//t.Log(ki.Type, hex.EncodeToString(ki.PrivateKey))
	println(ki.Type, hex.EncodeToString(ki.PrivateKey))
	println(ki1.Type, hex.EncodeToString(ki1.PrivateKey))
	return addr.String(),addr1.String(), nil
}



//获取当前块高度
func GetLastHeight() (int64, error) {
	c := ConnectClient()
	ts, err := c.ChainHead(context.Background())
	if err != nil {
		return 0, err
	}
	return ts.Height, nil
}

func GetWalletList() []address.Address {
	c := ConnectClient()
	list, _ := c.WalletList(context.Background())
	for _, item := range list {
		log.Println(item.String())
	}
	return list
}

//根据cid获取msg
func GetoneTransation(id string) (*types.Message, error) {
	c := ConnectClient()

	cid, err := cid.Parse(id)
	if err != nil {
		fmt.Println(err)
	}

	msg, err := c.ChainGetMessage(context.Background(), cid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return msg, nil

}

func GetHeightByCid(id string) {
	c := ConnectClient()
	mycid, _ := cid.Parse(id)
	t, _ := c.ChainGetBlock(context.Background(), mycid)
	log.Println(t.Height)

}

func GetFileWalletBalance(account string) (string, error) {
	c := ConnectClient()
	addr, _ := address.NewFromString(account)
	b, err := c.WalletBalance(context.Background(), addr)
	if err != nil {
		return "", err
	}
	return client.ToFil(b).String(), nil
}

func SendFilCoin(toAddress string, amount float64) (string, error) {

	c := ConnectClient()
	db := manager.GetDbInstance()
	var wallet models.Wallet
	db.Last(&wallet)
	from, _ := address.NewFromString(wallet.Address)
	to, _ := address.NewFromString(toAddress)

	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      client.FromFil(decimal.NewFromFloat(amount)),
		GasLimit:   0,
		GasFeeCap:  decimal.Zero,
		GasPremium: decimal.Zero,
		Method:     0,
		Params:     nil,
	}


	maxFee := client.FromFil(decimal.NewFromFloat(0.01))
	msg, err := c.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		return "", err
	}



	actor, err := c.StateGetActor(context.Background(), msg.From,  nil)
	if err != nil {
		return "", err
	}

	msg.Nonce = actor.Nonce
	fmt.Println(msg)

	sm, err := c.WalletSignMessage(context.Background(), msg.From, msg)
	if err != nil {
		return "", err
	}

	id, err := c.MpoolPush(context.Background(), sm)
	if err != nil {
		return "", err
	}
	fmt.Println(id.Version())
	return id.String(), nil

}


func SendFromOutAccount(toAddress string, amount float64) (string, error) {

	c := ConnectClient()
	db := manager.GetDbInstance()
	var wallet models.Wallet
	db.Last(&wallet)
	from, _ := address.NewFromString(wallet.Outaddress)
	to, _ := address.NewFromString(toAddress)

	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      client.FromFil(decimal.NewFromFloat(amount)),
		GasLimit:   0,
		GasFeeCap:  decimal.Zero,
		GasPremium: decimal.Zero,
		Method:     0,
		Params:     nil,
	}


	maxFee := client.FromFil(decimal.NewFromFloat(0.01))
	msg, err := c.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		return "", err
	}


	actor, err := c.StateGetActor(context.Background(), msg.From,  nil)
	if err != nil {
		return "", err
	}

	msg.Nonce = actor.Nonce
	fmt.Println(msg)

	sm, err := c.WalletSignMessage(context.Background(), msg.From, msg)
	if err != nil {
		return "", err
	}

	id, err := c.MpoolPush(context.Background(), sm)
	if err != nil {
		return "", err
	}
	fmt.Println(id.Version())
	return id.String(), nil

}

func SendFilCoin1(fromAddress string, toAddress string, amount float64) (string, error) {

	c := ConnectClient()
	from, _ := address.NewFromString(fromAddress)
	to, _ := address.NewFromString(toAddress)

	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      client.FromFil(decimal.NewFromFloat(amount)),
		GasLimit:   0,
		GasFeeCap:  decimal.Zero,
		GasPremium: decimal.Zero,
		Method:     0,
		Params:     nil,
	}



	maxFee := client.FromFil(decimal.NewFromFloat(0.01))
	msg, err := c.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		return "", err
	}



	actor, err := c.StateGetActor(context.Background(),msg.From, nil)
	if err != nil {
		return "", err
	}

	msg.Nonce = actor.Nonce
	fmt.Println(msg)

	sm, err := c.WalletSignMessage(context.Background(),msg.From, msg)
	if err != nil {
		return "", err
	}

	id, err := c.MpoolPush(context.Background(), sm)
	if err != nil {
		return "", err
	}
	fmt.Println(id.Version())
	return id.String(), nil

}

func GetCidState(mycid string) (int64, error) {
	//nil 标示正在打包
	c := ConnectClient()
	id, err := cid.Parse(mycid)
	if err != nil {
		return -1, err
	}
	mr, err := c.StateGetReceipt(context.Background(), id, nil)
	if err != nil {
		return -1, err
	}

	fmt.Println(mr)
	if mr == nil {
		return -1, errors.New("待打包")
	}
	return mr.ExitCode, nil
}

//检测FIL 地址格式有效性
func VerifyAccount(account string) (string, bool, error) {
	c := ConnectClient()
	addres, err := c.WalletValidateAddress(context.Background(), account)
	if err != nil {
		return "", false, err
	} else {
		return addres.String(), true, nil
	}

	return addres.String(), false, nil
}

func HasAddress(account string) (bool, error) {
	c := ConnectClient()
	account_addr, _ := address.NewFromString(account)
	if b, err := c.WalletHas(context.Background(), account_addr); err != nil {
		return false, err
	} else {
		return b, nil
	}
}

func hasAddress(c *client.Client, from string, to string) (bool, string, error) {
	from_str, _ := address.NewFromString(from)
	to_str, _ := address.NewFromString(to)
	from_b, err := c.WalletHas(context.Background(), from_str)
	to_b, err1 := c.WalletHas(context.Background(), to_str)
	//异常情况
	if err != nil {
		return false, "", err
	} else if err1 != nil {
		return false, "", err1
	}

	//首先判断和自己有没有关系 后面再写
	if from_b == false && to_b == false { //和钱包没有关系

		return false, "", nil
	} else {
		fmt.Println(">>>>>>>>>>>>本钱包交易<<<<<<<<<<<<<",from_str,to_str)
		//属于钱包 的地址 分3种情况
		if from_b && to_b == false { //out 转出
			return true, "out", nil
		} else if from_b == false && to_b { //in  转入
			return true, "in", nil
		} else if from_b && to_b { //inner 内部转账
			return true, "inner", nil
		}

	}

	return false, "", nil
}

func GetTransConfirms(txid string) (int64, error) {
	db := manager.GetDbInstance()
	var filModel models.FilTransation
	if db := db.Where("txid = ?", txid).Find(&filModel); db.RowsAffected == 1 {
		if last, err := GetLastHeight(); err != nil {
			return -1, err
		} else {
			result := last - int64(filModel.BlockHight)
			return result, nil
		}
	}
	return -1, errors.New("等待打包（手续费比较低）：" + txid)

}

/**获取一笔状态交易
  state:0 表示新交易
  state:1 表示加积分成功-后台使用
 */
func GetOneTransation(state string) (*models.FilIntegral, error) {
	db := manager.GetDbInstance()
	var order models.FilIntegral
	if db := db.Where("state = ?", state).First(&order); db.RowsAffected == 1 {
		return &order,nil
	}
	return nil, nil
}
/**
  加完积分后调用资金归集到总账号
 */

func ChangeIntegralState(txid string,state int64)  error{
	db := manager.GetDbInstance()
	if db:=db.Model(&models.FilIntegral{}).Where("txid = ?",txid ).Update("state", state);db.RowsAffected==1{
		return nil
	}
	return errors.New("更新状态失败")
	
}