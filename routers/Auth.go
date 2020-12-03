package api

import (

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/myxtype/filecoin-client/manager"
	"github.com/myxtype/filecoin-client/models"
	"github.com/myxtype/filecoin-client/pkg/e"
	"github.com/myxtype/filecoin-client/util"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {

	//获取参数
	username := util.Decrypt(c.Query("username"))
	password := util.Decrypt(c.Query("password"))
	//检查参数合法性
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	if ok {
		isExist := models.CheckAuth(manager.GetDbInstance(),a.Username, a.Password)
		if isExist {

			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {

			log.Printf("err.key %s,err.message %s",err.Key,err.Message)
		}

	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,

	})

}
