package models

/**
 * models 的用户模型
 */


import (

	"github.com/astaxie/beego/orm"
)


type Post struct {
	Id         int64
    Title string
    User  *User  `orm:"rel(fk)"`
}

//重构，重命名表,调用 base 重命名表名函数
func (m *Post) TableName() string {
	return TableName("post")
}

//新建用户
func (m *Post) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return nil
}

func (m *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...);err != nil{
		return err
	}
	return nil 
}



func (m *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Post) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
