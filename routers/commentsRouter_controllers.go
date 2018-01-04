package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ReviewService/controllers:QuestionController"] = append(beego.GlobalControllerRouter["ReviewService/controllers:QuestionController"],
		beego.ControllerComments{
			Method: "AddQuestion",
			Router: `/addQuestion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ReviewService/controllers:QuestionController"] = append(beego.GlobalControllerRouter["ReviewService/controllers:QuestionController"],
		beego.ControllerComments{
			Method: "GetAllQuestion",
			Router: `/getAllQuestion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ReviewService/controllers:UserController"] = append(beego.GlobalControllerRouter["ReviewService/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ReviewService/controllers:UserController"] = append(beego.GlobalControllerRouter["ReviewService/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ReviewService/controllers:UserController"] = append(beego.GlobalControllerRouter["ReviewService/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
