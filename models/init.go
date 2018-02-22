package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/cache"
)

var(
	Rc cache.Cache
	errRedis error
)

func Init() {

	//数据库初始化
	maxIdle := 30
	maxConn := 30

	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")

	dns := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8&loc=Local"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dns)

	//设置最大空闲连接数和最大数据库连接
	orm.SetMaxIdleConns("default", maxIdle)
	orm.SetMaxOpenConns("default", maxConn)

	//将用到的所有结构体映射到数据库中
	orm.RegisterModel(new(User),new(AllQuestion),new(Option))

	//自动建表
	orm.RunSyncdb("default",false,true)

	if beego.AppConfig.String("runmode") == beego.DEV {
		orm.Debug = true
	}

	//redis初始化
	if Rc,errRedis= cache.NewCache("redis",`{"key":"cacheRedis","conn":"127.0.0.1:6039","dbNum":"0","password":"thePassWord"}`); errRedis!=nil{
		panic(errRedis)
	}
}

//该方法可以直接得到操作的哪张表名，因为统一处理，所以表名暂时定为service_user和service_question
func TableName(suffixName string) string {
	return beego.AppConfig.String("db.prefix") + suffixName
}