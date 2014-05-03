package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//Blog post 与 tag 的关系表
type BlogTagPost struct {
	Id         int64
	Tagid      int64 `orm:"index"`		//blog tag的id
	Postid     int64					//blog post文章的id
	Poststatus int8						//blog 发布的时间
	Posttime   time.Time
}

func(m *BlogTagPost) TableName() string{
	return TableName("blog_tag_post")
}

func (m *BlogTagPost) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *BlogTagPost) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *BlogTagPost) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *BlogTagPost) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *BlogTagPost) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
