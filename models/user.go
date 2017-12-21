package models

import (
	"errors"
	_"strconv"
	_"time"
	"github.com/astaxie/beego/orm"
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
	Username string
	Password string
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

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

//查询用户是否在数据库中
//true表示未注册,false表示已注册
func GetUserByName(u *User) bool {
	err := orm.NewOrm().Read(u,"username")
	if err!=nil {
		return true
	}
	return false
}

func GetUserInfoByName(loginName string) (*User,error) {
	var user User
	err := orm.NewOrm().QueryTable(TableName(tabName)).Filter("username",loginName).One(&user)
	if err!=nil {
		return nil,err
	}
	return &user,err
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
