package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id              int64
	Username        string
	Password        string
	Email           string
	Mobiel          string
	Last_login_time time.Time `orm:"auto_now_add;type(datetime)"`
	Last_login_ip   string
	Status          int16
	Token           string `orm:"unique"`
}

func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func Login(email string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Email", email).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}
