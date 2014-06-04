package models


/**
 *  分类 模型
 */

import (
 	"time"
	"github.com/astaxie/beego/orm"
)

type Joke struct {
	Id int64
	Title    string    `orm:"size(100)"`
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Created  time.Time `orm:"type(datetime);index"`
	Status   int8
	Url		 string    `orm:"size(100)"`
	CategoryId int64	
	// Category *Category `orm:"rel(fk)"`	
}


func (m *Joke) TableName() string {
	return TableName("joke")
}

//新建分类
func (m *Joke) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err !=nil{
		return err
	}
	return nil
}

func (m *Joke) Read(fields ...string) error {
	if err :=orm.NewOrm().Read(m,fields...);err!=nil{
		return err
	}
	return nil
}


func (m *Joke) Update(fields ...string) error{
	if _,err := orm.NewOrm().Update(m,fields...);err!=nil{
		return err
	}
	return nil
}

func (m *Joke) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err != nil{
		return err
	}
	return nil
}

func (m *Joke) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}