package main

import (
	_ "ReviewService/routers"

	"github.com/astaxie/beego"
	"ReviewService/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		models.Init()
	}
	beego.Run()
}
