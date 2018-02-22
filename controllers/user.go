package controllers

import (
	"ReviewService/models"
	"ReviewService/utils"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title Login
// @Description Logs user into the system
// @Param	username		formData 	string	true		"The username for login"
// @Param	password		formData 	string	true		"The password for login"
// @Success 200 {object} models.user.UserRe
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password:=u.GetString("password")
	user, e1 := models.GetUserInfoByName(username)
	resData:=new(ResultData)
	userRe:=new(models.UserRe)
	if e1!=nil {
		resData.Code=MSG_ERR
		resData.Msg="用户不存在"
	} else {
		if user.UserName==username&&user.PassWord==password {
			user.Token=utils.CreateToken(user)
			resData.Code=MSG_Suc
			resData.Msg="登录成功"
			userRe.Uid=user.Id
			userRe.Token=utils.CreateToken(user)
		}else {
			resData.Code=MSG_ERR
			resData.Msg="用户名或密码错误"
		}
	}
	resData.Data=userRe
	u.Data["json"] = resData
	u.ServeJSON()
}

//@Title Register
//@Description add user into server
//@Param	username		formData	string	true		"The username for register"
//@Param	password		formData	string	true		"The password for register"
//@Success 200 register success
//@Failure 403 user has Registered
//@router /register [post]
func (u *UserController) Register()  {
	username := u.GetString("username")
	password := u.GetString("password")
	user:=models.User{UserName:username,PassWord:password}
	resData:=new(ResultData)
	isSu := models.AddUser(&user)
	if isSu {
		resData.Code=MSG_Suc
		resData.Msg="注册成功"
	}else {
		resData.Code=MSG_ERR
		resData.Msg="注册失败,该用户已存在"
	}
	u.Data["json"]=resData
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [post]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

