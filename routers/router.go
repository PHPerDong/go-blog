package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
	"myblog/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "GET:Logout")
	beego.Router("/reg", &controllers.RegController{})
	beego.InsertFilter("/add", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/add", &controllers.AddArticleController{})
}
