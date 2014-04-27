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
	Title    string `orm:"size(100)"`
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"type(datetime);index"`
	Status   int8
	Views    int64
	Istop    int8
	Updated  time.Time `orm:"type(datetime)"`
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

//删除
func (m *BlogPost) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil 
}


func (m *BlogPost) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}