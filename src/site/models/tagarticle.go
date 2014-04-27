package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//标签内容关系表
type TagArticle struct {
	Id         int64
	Tagid      int64 `orm:"index"`
	Articleid     int64
	Articlestatus int8
	Posttime   time.Time
}

func (m *TagArticle) TableName() string {
	return TableName("tag_article")
}

func (m *TagArticle) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *TagArticle) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *TagArticle) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *TagArticle) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *TagArticle) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
