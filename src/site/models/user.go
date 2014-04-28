package models

/**
 * models 的用户模型
 */


import (
	"time"
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id         int64
	Username   string    `orm:"unique;size(15)"`	//登录名
	Password   string    `orm:"size(32)"`			//密码
	Email      string    `orm:"size(50)"`			//邮件
	Lastlogin  time.Time `orm:"auto_now_add;type(datetime)"`		//最后登录时间，自动添加

}

//重构，重命名表,调用 base 重命名表名函数
func (m *User) TableName() string {
	return TableName("user")
}

//新建用户
func (m *User) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...);err != nil{
		return err
	}
	return nil 
}



func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
