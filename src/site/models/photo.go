package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/**
 * 图片
 */
type Photo struct {
	Id int64
	Title string		`orm:"size(100)"`			//标题
	Intro string 		`orm:"size(255)"` 			//简介
	Img	string			//图片本地地址
	Url	string			//图片远程地址
	Status	int8							//状态
	Created	time.Time   `orm:"type(datetime);index"` //建立时间 `orm:"auto_now_add;type(datetime)"`
	Category *Category `orm:"rel(fk)"`	
}

//重命令表
func (m *Photo) TableName() string {
	return TableName("photo")
}


//插入图片的方法
func (m *Photo) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err !=nil{
		return err
	}
	return nil
}

func (m *Photo) Read(fields ...string) error{
	if err := orm.NewOrm().Read(m,fields...); err!= nil{
		return err
	}
	return nil
}

func (m *Photo) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...);err != nil{
		return err
	}
	return nil
}


func (m *Photo) Delete() error{
	if _,err:=orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}


func (m *Photo) Query() orm.QuerySeter{
	return orm.NewOrm().QueryTable(m)
}