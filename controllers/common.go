package controllers

import "github.com/astaxie/beego"

const (
	MSG_Suc = 200
	MSG_ERR = 403
)

type BaseController struct {
	beego.Controller
}

type ResultData struct {
	Code int
	Msg string
	data interface{}
}

