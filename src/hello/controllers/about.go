package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"hello/models"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Index(){

	var category models.Category
	beego.Debug(category)

	//提交POST的方法
	if this.Ctx.Request.Method == "POST"{
		 username :=  this.GetString("username")
		 this.Ctx.WriteString(username)
	}

	//读取默认的config文件
	appname := beego.AppConfig.String("appname")		
	beego.Trace(appname)

	//读取定义的conf文件
	myini,err := config.NewConfig("ini","conf/my.conf")
	if err != nil{
		beego.Debug(err)
	}
	beego.Trace(myini.String("appname"))
	
	//直接打印文字
	// this.Ctx.WriteString("hello world")
	this.TplNames = "about.tpl"
}