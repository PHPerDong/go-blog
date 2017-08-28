package controllers

import (
	"github.com/astaxie/beego"
	"myblog/filters"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//fmt.Println(filters.IsLogin(c.Controller.Ctx))
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "首页"
	c.Layout = "layout/layout.tpl"
	c.TplName = "home/index.html"
}
