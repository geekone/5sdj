package admin

import (
	"github.com/astaxie/beego"
	"strings"
)

type baseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare(){
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
}

// func (this *baseController) auth(){
// 	if this.controllerName == "account" && (this.actionName=="login" || this.actionName == "logout"){

// 	}else{
// 		arr := strings.Split(this.Ctx.GetCookie("auth"),"|")
		
// 	}
// }


//渲染模版
func (this *baseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
	}
	
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = tplname
}



//显示错误提示
func (this *baseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	// this.Data["adminid"] = this.userid
	// this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = this.moduleName + "/" + "showmsg.html"
	this.Render()
	this.StopRun()
}