package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	//内置的log显示方法
	// beego.SetLogger("file",`{"filename":"logs/hello.log"}`)		//log写到了文件里
	// beego.BeeLogger.DelLogger("console")						//log只写到文件，不打印到console
	//beego.SetLogFuncCall(true)								//输出文件名和行号
	//beego.SetLevel(beego.LevelInfo)							//设置级别
	// beego.Trace("this is trace")
	// beego.Debug("this is debug")
	// beego.Info("this is info")
	// beego.Warn("this is warn")
	// beego.Error("this is error")
	//end


	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
