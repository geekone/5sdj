package models

/**
 * 视频模型
 */

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Movie struct {
	Id int64
	Title string `orm:"size(100)"`			//标题
	Intro string `orm:"size(255)"` 			//简介
	Status	int8							//状态
	Created	 time.Time 	`orm:"type(datetime);index"`		//建立时间 `orm:"auto_now_add;type(datetime)"`
	Category *Category `orm:"rel(fk)"`	
}


//重命名表名
func (m *Movie) TableName() string {
	return TableName("movie")
}

/**
 * 新建
 */
func (m *Movie) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err!=nil{
		return err
	}
	return nil
}


/**
 * 读取
 */
func (m *Movie) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...);err!=nil{
		return err
	}
	return nil
}

/**
 * 更新
 */
func (m *Movie) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...);err != nil{
		return err
	}
	return nil
}


/**
 * 删除
 */
func(m *Movie) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err!=nil{
		return err 
	}
	return nil
}


func (m *Movie) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}