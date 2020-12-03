package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/myxtype/filecoin-client/middleware/jwt"
	"github.com/myxtype/filecoin-client/pkg/setting"
	api "github.com/myxtype/filecoin-client/routers"
	v1 "github.com/myxtype/filecoin-client/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/getRawTransaction/:cid", v1.GetFilTransaction)//rpc

		apiv1.GET("/pay/sendtoaddress/:from/:to/:amount", v1.SendFilCoin1)//rpc
		//apiv1.GET("/GetNewAddress", v1.GetNewAddress)//rpc
		apiv1.GET("/getwalletinfo/:account", v1.GetFilWalletBalance)//rpc
		apiv1.GET("/validateAddress/:address", v1.Fil_ValidateAddress)//rpc
		apiv1.GET("/newaccount/", v1.Fil_NewAccount)//rpc

		apiv1.GET("/isnodeaddres/:address", v1.Fil_wallethas)//rpc 判断是不是钱包节点地址

		apiv1.GET("/getlastHeight", v1.Fil_GetlastHeight)//rpc 获取当前最新块高度

		apiv1.GET("/getComfirm/:txid", v1.Fil_TransConfirms)//rpc 获取一笔交易的确认数

		apiv1.GET("/pay/sendtowallet/:address/:amount", v1.SendFilCoin)//rpc 统一转账到入金账号地址

		apiv1.GET("/pay/sendfromoutaccount/:to/:amount", v1.SendFromOutAccount)//rpc 统一转从出金账号提币

		apiv1.GET("/getOneTranstion/:state", v1.Fil_GetOneTransation)//rpc 根据状态检索一笔交易

		apiv1.GET("/changeIntegralState/:txid/:state", v1.Fil_ChangeIntegralState)//rpc 修改一笔交易状态

	}
	return r

}
