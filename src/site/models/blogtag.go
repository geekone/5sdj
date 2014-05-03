package models

import (

	"github.com/astaxie/beego/orm"
)

//Blog标签表
type BlogTag struct {
	Id int64
	Name string `orm:"size(20);index"`
	Count int64
}


func(m *BlogTag) TableName() string {
	return TableName("blog_tag")
}

func (m *BlogTag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *BlogTag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *BlogTag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}


//表查询
func (m *BlogTag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}