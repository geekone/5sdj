package admin

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"site/models"
	
)

type baseController struct {
	beego.Controller
	adminid         int64
	account       	string
	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare(){
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()
}


//登录状态验证
func (this *baseController) auth(){
	if this.controllerName == "adminuser" && (this.actionName=="login" || this.actionName == "logout"){
		//如果是adminuser login 和 logout 跳过验证
	} else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, password := arr[0], arr[1]
			adminid, _ := strconv.ParseInt(idstr, 10, 0)
			if adminid > 0 {
				var adminuser models.Adminuser
				adminuser.Id = adminid
				if adminuser.Read() == nil && password == adminuser.Password {
					this.adminid = adminuser.Id
					this.account = adminuser.Account
				}
			}
		}

		if this.adminid == 0 {
			this.Redirect("/admin/adminuser/login", 302)
		}
	}
}


//渲染模版
func (this *baseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
	}
	//一些layout 需要的变量
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.adminid
	this.Data["adminname"] = this.account
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = tplname
}



//显示错误提示
func (this *baseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	//TODO this.Data["adminid"] = this.userid
	// this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = this.moduleName + "/" + "showmsg.html"
	this.Render()
	this.StopRun()
}


//获取用户IP地址
func (this *baseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr,":")
	return s[0]
}


// func (this *baseController) getTime() time.Time {
// 	options := models.GetOptions()
// 	timezone := float64(0)
// 	if v, ok := options["timezone"]; ok {
// 		timezone, _ = strconv.ParseFloat(v, 64)
// 	}
// 	add := timezone * float64(time.Hour)
// 	return time.Now().UTC().Add(time.Duration(add))
// }