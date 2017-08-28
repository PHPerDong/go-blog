package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sluu99/uuid"
	"myblog/models"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {
	this.Data["PageTitle"] = "注册"
	this.Layout = "layout/layout.tpl"
	this.TplName = "home/reg.html"
}

func (this *RegController) Post() {
	//var user models.User
	inputs := this.Input()
	username := inputs.Get("username")
	password := inputs.Get("pass")
	email := inputs.Get("email")
	token := uuid.Rand().Hex()
	user := models.User{Username: username, Password: password, Email: email, Token: token}
	models.SaveUser(&user)
	/*if err == nil {
		this.Redirect("/", 302)
	} else {
		fmt.Println(err)
		this.Redirect("/", 302)
	}*/
	this.Redirect("/", 302)

}
