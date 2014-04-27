package models

/**
 * Blog 分类 模型
 */

import (
	"github.com/astaxie/beego/orm"
)

type BlogCategory struct {
	Id int64
	Name string `orm:"unique;size(20)"`		//分类名
}

//重命名
func (m *BlogCategory) TableName() string {
	return TableName("blog_category")
}

//新建
func (m *BlogCategory) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err !=nil{
		return err
	}
	return nil
}

/**
 * 读取再回传值给model
 */
func (m *BlogCategory) Read(fields ...string) error{
	if err := orm.NewOrm().Read(m,fields...); err!=nil{
		return err
	}
	return nil
}

/**
 * 更新再回传model
 */
func (m *BlogCategory) Update(fields ...string) error {
	if _,err :=  orm.NewOrm().Update(m,fields...); err != nil{
		return err
	}
	return nil
}


/**
 *	删除 出错返回 m
 */
func (m *BlogCategory) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}


func (m *BlogCategory) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}