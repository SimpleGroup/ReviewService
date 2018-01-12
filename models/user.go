package models

import (
	_"strconv"
	_"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

var (
	UserList map[string]*User
	tabName ="user"
)

func init() {
	UserList = make(map[string]*User)
	u := User{0, "astaxie", "11111"}
	UserList["user_11111"] = &u
}

type User struct {
	Id       int
	UserName string `orm:"column(user_name)"`
	PassWord string `orm:"column(user_pwd)"`
	CreateTime
}

//自定义表名
func (u *User)TableName() string {
	return "service_user"
}

//注册用户时往数据库加用户
func AddUser(u *User) bool {
	var isSu bool
	if GetUserByName(u) {
		_, e := orm.NewOrm().Insert(u)
		if e!=nil {
			isSu=false
		}else {
			isSu=true
		}
	}else {
		isSu=false
	}
	return isSu
}

//查询用户是否在数据库中
//true表示未注册,false表示已注册
func GetUserByName(u *User) bool {
	err := orm.NewOrm().Read(u,"user_name")
	if err!=nil {
		return true
	}
	return false
}

func GetUserInfoByName(loginName string) (*User,error) {
	var user User
	err := orm.NewOrm().QueryTable(TableName(tabName)).Filter("user_name",loginName).One(&user)
	if err!=nil {
		return nil,err
	}
	return &user,err
}

