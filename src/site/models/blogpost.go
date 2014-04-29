package models

/***
 * 博客文章
 */

 import (
 	"time"
 	"github.com/astaxie/beego/orm"
 )

type BlogPost struct {
	Id int64
	Userid 	 int64	`orm:"index"`
	Author	 string  `orm:"size(15)"`
	Title    string `orm:"size(100)"`
	Color	 string	 `orm:"size(7)"`
	Urlname  string	 `orm:"size(100);index"`
	Urltype	 int8
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"type(datetime);index"`
	Views    int64
	Status   int8
	Updated  time.Time `orm:"type(datetime)"`
	Istop    int8
}


func (m *BlogPost) TableName() string {
	return TableName("blog_post")
}

//新建
func (m *BlogPost) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err !=nil{
		return err
	}
	return nil
}

//读取	根据要显示字段 返回回填的model
func (m *BlogPost) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err!=nil{
		return err
	}
	return nil
}

//更新
func (m *BlogPost) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil{
		return err
	}
	return nil
}

//TODO 删除
func (m *BlogPost) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil 
}


func (m *BlogPost) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}