package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["PageTitle"] = "登入"
	this.Layout = "layout/layout.tpl"
	this.TplName = "home/login.html"
}

func (this *LoginController) Post() {
	flash := beego.NewFlash()
	//sess := this.StartSession()
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//return

	email := this.Input().Get("email")
	password := this.Input().Get("pass")
	//user := models.User{Email:email,Password:password}
	if flag, user := models.Login(email, password); flag {

		this.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		//this.SetSession("uid", user.Id)

		this.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&this.Controller)
		this.Redirect("/add", 302)
	}

}

func (this *LoginController) Logout() {
	this.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	this.Redirect("/", 302)
}
