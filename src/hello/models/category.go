package models

import (
	
)

type Category struct {
	Id int64
	Name string `orm:"size(45)"`	//分类名
	Cname string `orm:"size(45)"`	//中文分类名称
}

//重命令分类的表名
func (m *Category) TableName() string {
	return TableName("category")	//调用重命名表的公用函数
}