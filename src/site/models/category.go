package models

/**
 *  分类 模型
 */

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id int64
	Name string `orm:"size(45)"`		//分类名
	Cname string `orm:"size(45)"`		//中文分类名
	Catestr  string	`orm:size(45)"`		//再分类 joke 指笑话 pic 图片
	Cateid	int64								//笑话对应的分类ID
	Jokes []*Joke `orm:"reverse(many)"`	 //fk
}


func (m *Category) TableName() string {
	return TableName("category")
}

//新建分类
func (m *Category) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err !=nil{
		return err
	}
	return nil
}

func (m *Category) Read(fields ...string) error {
	if err :=orm.NewOrm().Read(m,fields...);err!=nil{
		return err
	}
	return nil
}


func (m *Category) Update(fields ...string) error{
	if _,err := orm.NewOrm().Update(m,fields...);err!=nil{
		return err
	}
	return nil
}

func (m *Category) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err != nil{
		return err
	}
	return nil
}

func (m *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}