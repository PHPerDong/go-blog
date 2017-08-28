package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myblog/models"
	_ "myblog/routers"
)

func init() {
	//orm.Debug = true
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/myblog-go?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(
		new(models.Admin),
		new(models.Role),
		new(models.Permission),
		new(models.Article),
		new(models.User),
		new(models.Category),
		new(models.Tag),
	)
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetStaticPath("/img", "images")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "js")
	beego.SetStaticPath("/layui", "layui")
}
