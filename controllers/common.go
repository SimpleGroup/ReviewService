package controllers

import (
	"github.com/astaxie/beego"
	"ReviewService/models"
)

const (
	MSG_Suc = 200
	MSG_ERR = 400
)

type BaseController struct {
	beego.Controller
}

type ResultData struct {
	Code int
	Msg string
	Data interface{}
}

func (base *BaseController)Prepare(){
	base.GetString("")
	base.CheckToken()
}

func (c *BaseController)CheckToken() {
	authString := c.Ctx.Input.Header("Authorization")
	beego.Debug("authString",authString)
	if isExist := models.Rc.IsExist(c.GetString("id")); isExist{
		//该用户已经登录，redis里面有token信息，下一步是验证token信息
	}else {
		//该用户未登录，必须先登录
	}
}

