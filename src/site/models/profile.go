package models

import (
	"github.com/astaxie/beego/orm"
)


type Profile struct {
    Id          int
    Age         int16
    User        *User   `orm:"reverse(one)"` // 设置反向关系(可选)
}


//重构，重命名表,调用 base 重命名表名函数
func (m *Profile) TableName() string {
	return TableName("profile")
}


//新建用户
func (m *Profile) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return nil
}

func (m *Profile) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...);err != nil{
		return err
	}
	return nil 
}



func (m *Profile) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Profile) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Profile) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
