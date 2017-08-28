package models

import (
	"time"
)

type Admin struct {
	Id        int    `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Token     string `orm:"unique"`
	Avatar    string
	Email     string    `orm:"null"`
	Url       string    `orm:"null"`
	Signature string    `orm:"null;size(1000)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
	Roles     []*Role   `orm:"rel(m2m)"`
}
