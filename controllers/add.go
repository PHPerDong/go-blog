package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/filters"
	"myblog/models"
	"strconv"
)

type AddArticleController struct {
	beego.Controller
}

func (this *AddArticleController) Get() {
	this.Data["IsLogin"], this.Data["UserInfo"] = filters.IsLogin(this.Controller.Ctx)
	this.Data["category"] = models.FindAllCategoty()
	this.Data["PageTitle"] = "提问"
	this.Layout = "layout/layout.tpl"
	this.TplName = "home/add.html"
}

func (this *AddArticleController) Post() {

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	cid := this.Input().Get("cid")
	s, _ := strconv.Atoi(cid)
	category := models.Category{Id: s}
	fmt.Println(category)
	_, user := filters.IsLogin(this.Ctx)
	article := models.Article{Title: title, Content: content, Category: &category, User: &user}
	id := models.AddArticle(&article)
	this.Redirect("/article/"+strconv.FormatInt(id, 10), 302)
}
