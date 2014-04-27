package models


import (
	"time"
)

type Adminuser struct {
	Id int64
	Adminuser string `orm:"unique;size(15)"`
	Adminpass string `orm:"size(32"`
	Lastlogin	time.Time `orm:"auto_now_add;type(datetime)"`
}

