package models

import ()

type Role struct {
	Id          int           `orm:"pk;auto"`
	Name        string        `orm:"unique"`
	Users       []*Admin      `orm:"reverse(many)"`
	Permissions []*Permission `orm:"rel(m2m)"`
}
