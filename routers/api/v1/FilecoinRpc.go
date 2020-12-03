package v1

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	_ "github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/api"
	"github.com/myxtype/filecoin-client/pkg/e"
	"github.com/myxtype/filecoin-client/util"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"strconv"
)

func GetFilTransaction(c *gin.Context) {
	//log.Panicln("------GetRawTransaction---->")
	cid:= util.Decrypt(c.Param("cid"))
	valid := validation.Validation{}

	valid.Length(cid, 62, "cid").Message("cid必须是62位")
	data := make(map[string]interface{})
	code := e.ERROR
	if !valid.HasErrors() { //参数校验ok
		if msg, err := api.GetoneTransation(cid); err != nil {
			code = e.ERROR
		} else {
			code = e.SUCCESS
			result,_:=json.Marshal(msg)
			data["result"]=string(result)
		}

	} else { //参数校验fail
		for _, err := range valid.Errors {
			log.Printf("err.key: %s,err.message: %s", err.Error(), err.Message)

		}

	}
	fmt.Printf("%+v\n", data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetFilWalletBalance(c *gin.Context) {
	account := util.Decrypt(c.Param("account"))
	//valid := validation.Validation{}
	//valid.Length(account, 41, "account").Message("account必须是41位")
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	//if !valid.HasErrors() { //参数校验ok
		if balance, err := api.GetFileWalletBalance(account); err != nil {
			code = e.ERROR
			data["result"] = nil
			fmt.Println(err.Error())
		} else {
			code = e.SUCCESS
			data["result"] = balance
		}

	//}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func SendFilCoin(c *gin.Context) {
	address:= util.Decrypt(c.Param("address"))
	amountStr :=util.Decrypt(c.Param("amount"))
	amount := com.StrTo(amountStr).MustFloat64()
	code := e.ERROR
	data := make(map[string]interface{})
	cid, err := api.SendFilCoin(address, amount)
	if err != nil {
		code = e.ERROR
		log.Println("------SendToAddress err------>", err.Error())
		data["result"] = nil
	} else {
		log.Printf("cid result: %s", cid)
		code = e.SUCCESS
		data["result"] = cid
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}





func SendFromOutAccount(c *gin.Context) {
	to:= util.Decrypt(c.Param("to"))
	amountStr :=util.Decrypt(c.Param("amount"))
	amount := com.StrTo(amountStr).MustFloat64()
	code := e.ERROR
	data := make(map[string]interface{})
	cid, err := api.SendFromOutAccount(to, amount)
	if err != nil {
		code = e.ERROR
		log.Println("------SendFromOutAccount err------>", err.Error())
		data["result"] = nil
		data["msg"] = err.Error()
	} else {
		log.Printf("cid result: %s", cid)
		code = e.SUCCESS
		data["result"] = cid
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}


func SendFilCoin1(c *gin.Context) {

	addressfrom:= util.Decrypt(c.Param("from"))
	addressTo:=util.Decrypt(c.Param("to"))
	amountDeccrip:= util.Decrypt(c.Param("amount"))
	fmt.Println("******from******"+addressfrom)
	fmt.Println("******to******"+addressTo)
	fmt.Println("******amount******"+amountDeccrip)

	amount := com.StrTo(amountDeccrip).MustFloat64()
	code := e.ERROR
	msg := ""
	data := make(map[string]interface{})
	cid, err := api.SendFilCoin1(addressfrom, addressTo, amount)
	if err != nil {
		code = e.ERROR
		log.Println("------SendToAddress err------>", err.Error())
		data["result"] = nil
		msg = err.Error()
	} else {
		log.Printf("cid result: %s", cid)
		code = e.SUCCESS
		data["result"] = cid
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Fil_ValidateAddress(c *gin.Context) {
	address := util.Decrypt(c.Param("address"))
	code := e.ERROR
	data := make(map[string]interface{})
	validate, ok, err := api.VerifyAccount(address)
	if err != nil {
		data["isvalid"] = ok
		data["address"] = validate
		data["result"] = err.Error()

	} else {
		if !ok {
			code = e.ERROR
		} else {
			code = e.SUCCESS
		}
		data["isvalid"] = ok
		data["address"] = validate
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func Fil_wallethas(c *gin.Context) {
	address := util.Decrypt(c.Param("address"))
	code := e.ERROR
	data := make(map[string]interface{})
	ok, err := api.HasAddress(address)
	if err != nil {
		data["result"] = err.Error()

	} else {
		if ok {
			code = e.SUCCESS
		}
		data["result"] = ok

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func Fil_NewAccount(c *gin.Context) {
	code := e.ERROR
	data := make(map[string]interface{})
	account,err := api.NewAccount()
	if err != nil {
		data["result"] = err.Error()
	} else {
		code = e.SUCCESS
		data["result"] = account
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Fil_GetlastHeight(c *gin.Context) {
	code := e.ERROR
	data := make(map[string]interface{})
	height, err := api.GetLastHeight()
	if err != nil {
		data["result"] = err.Error()
	} else {
		code = e.SUCCESS
		data["result"] = height
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Fil_TransConfirms(c *gin.Context) {

	code := e.ERROR
	txid:= util.Decrypt(c.Param("txid"))
	data := make(map[string]interface{})
	number, err := api.GetTransConfirms(txid)
	if err != nil {
		data["result"] = err.Error()
	} else {
		code = e.SUCCESS
		data["result"] = number
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}



func Fil_ChangeIntegralState(c *gin.Context) {
	code := e.ERROR
	data := make(map[string]interface{})
	state:= util.Decrypt(c.Param("state"))
	txid:= util.Decrypt(c.Param("txid"))
	state64, err := strconv.ParseInt(state, 10, 64)
	err = api.ChangeIntegralState(txid, state64)
	if err != nil {
		data["result"] = err.Error()
	} else {
		code = e.SUCCESS
		data["result"] = txid
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}



func Fil_GetOneTransation(c *gin.Context) {
	code := e.ERROR
	data := make(map[string]interface{})
	state:= util.Decrypt(c.Param("state"))
	trans, err := api.GetOneTransation(state)
	if err != nil {
		data["result"] = err.Error()
	} else {
		code = e.SUCCESS
		data["result"] = trans
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}


