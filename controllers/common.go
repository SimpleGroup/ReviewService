package controllers

import (
	"github.com/astaxie/beego"
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

