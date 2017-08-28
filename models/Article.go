package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id            int       `orm:"pk;auto"`
	Title         string    `orm:"unique"`
	Content       string    `orm:"type(text);null"`
	InTime        time.Time `orm:"auto_now_add;type(datetime)"`
	User          *User     `orm:"rel(fk)"`
	Category      *Category `orm:"rel(fk)"`
	View          int       `orm:"default(0)"`
	ReplyCount    int       `orm:"default(0)"`
	LastReplyUser *User     `orm:"rel(fk);null"`
	LastReplyTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func AddArticle(article *Article) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(article)
	return id
}
