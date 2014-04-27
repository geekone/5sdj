package blog

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
	this.moduleName = "blog"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
}



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