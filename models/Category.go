package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"unique"`
}

func FindAllCategoty() []*Category {
	o := orm.NewOrm()
	var category Category
	var categorys []*Category
	o.QueryTable(category).All(&categorys)
	return categorys
}
