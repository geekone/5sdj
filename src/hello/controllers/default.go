package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	//内置的log显示方法
	beego.SetLogger("file",`{"filename":"logs/hello.log"}`)
	beego.Trace("this is trace")
	beego.Debug("this is debug")
	beego.Info("this is info")
	beego.Warn("this is warn")
	beego.Error("this is error")
	//end


	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
