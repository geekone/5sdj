package models

/***
 * 博客文章
 */

 import (
 	"time"
 	"github.com/astaxie/beego/orm"
 )

type Article struct {
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


func (m *Article) TableName() string {
	return TableName("article")
}

//新建文章
func (m *Article) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err !=nil{
		return err
	}
	return nil
}


func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}