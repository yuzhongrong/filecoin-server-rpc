package main

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/controller"
	"github.com/myxtype/filecoin-client/manager"
	"github.com/myxtype/filecoin-client/pkg/setting"
	routers "github.com/myxtype/filecoin-client/routers/api"
	"log"
	"net/http"
)

func main() {

	//使用主网参数去解析链上交易
	address.CurrentNetwork = address.Mainnet
	//初始化钱包账号
	account:=controller.InitWallet(manager.GetDbInstance())
	log.Println("初始化钱包成功",account)
	//主矿初始化同步矿厂定时任务
	if setting.ISMINER{
		controller.Fil_InitCron(manager.GetDbInstance())


	}
	//初始化转账和查询订单确认数任务
	//controller.Fil_InitOrderCron(manager.GetDbInstance())
	//log.Println("初始化转账和查询订单确认数任务成功\n")

	//初始化api服务
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	 _err := s.ListenAndServe()
	if _err != nil {
		fmt.Printf("监听报错 %s", _err.Error())
	}else{
		fmt.Println("初始化api服务成功")
	}



}
