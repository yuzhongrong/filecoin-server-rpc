package controller

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/myxtype/filecoin-client/api"
	"github.com/myxtype/filecoin-client/client"
	"github.com/myxtype/filecoin-client/manager"
	"github.com/myxtype/filecoin-client/models"
	"github.com/myxtype/filecoin-client/pkg/setting"
	"github.com/myxtype/filecoin-client/types"
	"github.com/robfig/cron"
	"log"
	"math/big"
	"strconv"
	"time"
)

func Fil_InitCron( db *gorm.DB)  error{

	//启动定时任务定时更新区块
	 go InitCron(db)
	 return nil
}

func Fil_InitOrderCron(db *gorm.DB) error {

	go InitTran2Cron(db)
	go QueryConfirms(db)
	return nil
}


func InitWallet(db *gorm.DB) string {
	//如果没有账号创建一个账号
	var wallet models.Wallet
	db.Last(&wallet)
	if wallet.Address == "" {
		if inaccount,outaccount, err := api.NewAccount1(); err == nil {
			db.Save(&models.Wallet{Address: inaccount,Outaddress:outaccount})
			log.Println("新创建入金账号:" + inaccount)
			log.Println("新创建出金账号:" + outaccount)
			return inaccount
		}
	} else {

		log.Println("钱包入金账号:" + wallet.Address)
		return wallet.Address

	}

	return ""
}





func fil_insertDb1(transactions *types.BlockMessages1) error {
	//查询这笔交易的确认数
	currentHeight, err := api.GetLastHeight()
	db := manager.GetDbInstance()
	if err != nil {
		fmt.Printf("查询最新高度出错")
		return err
	}

	//构造cmctransation对象并且加入到集合 -------------------------------------------------
	var transDatas []*models.CmcTransation
	for index, item := range transactions.BlsMessages {
		//if item.Generated && item.Category == "generate" { //非coinbase 不处理

		//构造cmc交易
		amount := client.ToFil(item.Value).String()
		amount64, _ := strconv.ParseFloat(amount, 64)
		t := &models.CmcTransation{
			Coinbase:      false, //在filecoin中无用
			BlockHight:    uint64(transactions.Height),
			Confirmations: currentHeight - transactions.Height,
			Amount:        amount64,
			Txid:          transactions.Cids[index].String(),
			Address:       item.From.String(),
			To:            item.To.String(),
		}

		transDatas = append(transDatas, t)
		//按照区块高度排序
		//	sort.Slice(transDatas, func(i, j int) bool { return transDatas[i].BlockHight > transDatas[j].BlockHight })

	}
	//构造cmctransation对象并且加入到集合 ------------------------end-------------------------

	//遍历集合插入到数据库------------------------------------------------
	for _, sortitem := range transDatas {
		sortitem.AddTransation(db)
		fil_relateCmcRecord(db, sortitem)
		time.Sleep(100 * time.Millisecond)
	}

	//获取数据库中最新块的高度记录下来

	return nil

}



func fil_insertDb2(transactions *types.BlockMessages1) error {
	//查询这笔交易的确认数
	currentHeight, err := api.GetLastHeight()
	db := manager.GetDbInstance()
	if err != nil {
		fmt.Printf("查询最新高度出错")
		return err
	}

	//构造cmctransation对象并且加入到集合 -------------------------------------------------
	var transDatas []*models.FilTransation
	for index, item := range transactions.BlsMessages {
		//if item.Generated && item.Category == "generate" { //非coinbase 不处理

		//构造cmc交易
		amount := client.ToFil(item.Value).String()
		amount64, _ := strconv.ParseFloat(amount, 64)
		t := &models.FilTransation{
			Coinbase:      false, //在filecoin中无用
			BlockHight:    uint64(transactions.Height),
			Confirmations: currentHeight - transactions.Height,
			Amount:        amount64,
			Txid:          transactions.Cids[index].String(),
			Address:       item.From.String(),
			To:            item.To.String(),
			Direct:        item.Direct,
		}

		transDatas = append(transDatas, t)
		//按照区块高度排序
		//	sort.Slice(transDatas, func(i, j int) bool { return transDatas[i].BlockHight > transDatas[j].BlockHight })

	}
	//构造cmctransation对象并且加入到集合 ------------------------end-------------------------

	//遍历集合插入到数据库------------------------------------------------
	for _, sortitem := range transDatas {
		sortitem.AddTransation(db)
		sortitem.RelatedFilIntegral(db)//关联积分表-提供让其他人操纵这个表 其他人不能直接操作sys_filtransation表
		time.Sleep(100 * time.Millisecond)
	}

	//获取数据库中最新块的高度记录下来

	return nil

}

func fil_relateCmcRecord(db *gorm.DB, trans *models.CmcTransation) {
	//查询上一个recod 记录
	var temp models.CmcRecord
	db.Last(&temp)
	result := new(big.Float)
	amount := big.NewFloat(trans.Amount)
	nowallamount := big.NewFloat(temp.NowAllamount)
	a, _ := result.Add(amount, nowallamount).Float64()

	db.Where("txid = ?", trans.Txid).FirstOrCreate(&models.CmcRecord{
		Amount:        trans.Amount,
		NowAllamount:  a,
		LastAllamount: temp.NowAllamount,
		From:          trans.Address,
		To:            trans.To,
		Txid:          trans.Txid,
		Timestamp:     time.Now().Unix(),
	})

}



func InitCron(db *gorm.DB) {
	//初始化直接同步最新块
	if err := AsynChain(db); err != nil {
		log.Println(err)
	}else{//同步完成再启动定时任务
		c := cron.New()
		c.AddFunc(setting.CRONRULE, func() {
			if err := AsynChain(db); err != nil {
				log.Println(err)
			}
		})
		c.Run()
	}




}







//初始化转账任务+查询确认数任务
func InitTran2Cron(db *gorm.DB)  {
	c := cron.New()
	c.AddFunc(setting.CRONRULE_TRANS, func() {
		if err := QueryOneTransRecored(db); err != nil {
			log.Println(err)
		}
	})
	c.Run()
}


func QueryConfirms(db *gorm.DB){

	c := cron.New()
	c.AddFunc(setting.CRONRULE_CONFIRM, func() {
		if err := QueryOrderConfirmsFromChain(db); err != nil {
			log.Println(err)
		}
	})
	c.Run()

}

func QueryOneTransRecored(db *gorm.DB) error {
	//过去最新一条符合条件的订单
	var order models.Order
	if action:=db.Last(&order,"state = ? AND txid is null",1);action.RowsAffected==1{//找到一条待转账记录
	   fmt.Printf("查询出一条待转账订单: %s %s\n",order.Orderid,order.Bindaddress)
	   //开始转账
	   cid,err:=api.SendFilCoin(order.Bindaddress,order.Amount)
	   if err!=nil{//转账异常
	    return err
	   }

	   if cid!=""{
		   order.Txid=cid
		   order.State=2
		   db.Save(&order)

	   }

	}

	return nil

}

//查询一个订单的确认数
func QueryOrderConfirmsFromChain(db *gorm.DB)  error{
	var order models.Order
	if action:=db.Last(&order,"state=?",2);action.RowsAffected==1{
		code,err:=api.GetCidState(order.Txid)
		if err!=nil{ //查询出现异常
			return err
		}
		if code==0{//ok
			order.State=3
			db.Save(&order)
		}
	}
	return nil
}


func AsynChain(db *gorm.DB) error {
	//判断current高度
	var wallet models.Wallet
	db.Last(&wallet)
	if wallet.Current == 0 { //初始化第一次
		//获取当前最新的块高度
		if height, err := api.GetLastHeight(); err != nil {
			return err
		} else {
			wallet.Starthight = height
			wallet.Current = height
			db.Save(&wallet)
			//同步数据
			fmt.Printf("初始化开始同步第%s个块数据", height)
			err := InitAsyn(db, wallet.Address, wallet.Starthight)
			if err != nil {
				fmt.Println("初始化同步区块异常", err.Error())
			}
		}

	} else { //第二次
		fmt.Printf("开始同步第%s个块数据\n", wallet.Current)
		//同步数据
		err := UpdateAsynLastChain(db, &wallet)
		if err != nil { //同步过程中不出错 才去保存
			return err
		}
	}
	return nil
}

func InitAsyn(db *gorm.DB, account string, height int64) error {
	//trans, _err := api.UpgradeTransations(account, height)
	trans, _err := api.UpgradeTransations1( height)

	if _err != nil {
		log.Println("------InitAsyn error------->", _err.Error())
		return _err
	}

	if len(trans.BlsMessages) > 0 {
		if err := fil_insertDb2(trans); err != nil {
			return err
		}
	}
	return nil
}

//同步filecoin 一个块区块
//如果报错返回报错的块的高度
func UpdateAsynOneChain(account string, height int64) (int64, error) {
	//获取这个高度上饿的所有cid消息
	//trans, _err := api.UpgradeTransations(account, height)
	trans, _err := api.UpgradeTransations1( height)
	if _err != nil { //同步这个块报错
		return height, _err //直接返回和这个高度和报错信息
	} else { //同步该块正常
		if len(trans.BlsMessages) > 0 {
			if err := fil_insertDb2(trans); err != nil {
				return height, err
			}
		}
	}
	return height, nil
}

//从当前块同步到最新块
//需要返回中断时候的那个块的高度
func UpdateAsynLastChain(db *gorm.DB, wallet *models.Wallet) error {

	//数据库获取当前
	currentAsynedHeight := wallet.Current
	tempheight:=wallet.Current
	lastChainHeight, err := api.GetLastHeight()
	if err != nil {
		return err
	}


	fmt.Println("********当前高度-链上最新高度********",currentAsynedHeight,lastChainHeight)

	for i := currentAsynedHeight; i <= lastChainHeight; i++ {
		fmt.Println("********开始索引块高度********",i)
		chainHeight, err := UpdateAsynOneChain(wallet.Address, i)
		tempheight= chainHeight
		if err != nil { //报错后返回报错的块的高度
			break
		}
	}

	//保存这次任务执行后最终同步到第几个块或者说在第几个块同步报错
	wallet.Current = tempheight
	fmt.Println("********同步完成保存当前高度********",currentAsynedHeight,tempheight)
	db.Save(&wallet)
	return nil

}
