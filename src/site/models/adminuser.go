package models


import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Adminuser struct {
	Id int64
	Account		string `orm:"unique;size(15)"`
	Password	string `orm:"size(32"`
	Lastlogin	time.Time `orm:"auto_now_add;type(datetime)"`
	Lastip		string `orm:"size(32)"`	
}

func (m *Adminuser) TableName() string {
	return TableName("adminuser")
}


func (m *Adminuser) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err!=nil{
		return err
	}
	return nil
}


func (m *Adminuser) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...);err != nil{
		return err
	}
	return nil 
}



func (m *Adminuser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Adminuser) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Adminuser) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}