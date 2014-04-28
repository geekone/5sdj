package admin


import (
	"site/models"
	"strings"
	"strconv"
)

type AdminuserController struct {
	baseController
}

func (this *AdminuserController) List(){
	var list []*models.Adminuser
	var adminuser models.Adminuser
	count,_:= adminuser.Query().Count()
	if count > 0 {
		adminuser.Query().OrderBy("-id").All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list 
	this.display()
}


func (this *AdminuserController) Add(){

	errmsg :=make(map[string] string)

	if this.Ctx.Request.Method == "POST" {
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))

		


		if len(errmsg) == 0 {
			var adminuser models.Adminuser
			adminuser.Account = account
			adminuser.Password = password
			if err := adminuser.Insert();err != nil {
				this.showmsg(err.Error())
			}
			this.Redirect("/admin/adminuser/list",302)
		}
		
	}

	this.Data["errmsg"] = errmsg
	this.display()
}


func (this *AdminuserController) Delete(){
	id,_:=this.GetInt("id")
	adminuser := models.Adminuser{Id:id}
	if adminuser.Read() == nil{
		adminuser.Delete()
	}
	this.Redirect("/admin/adminuser/list",302)
}


func (this *AdminuserController) Edit(){
	id,_:=this.GetInt("id")
	adminuser := models.Adminuser{Id:id}
	if adminuser.Read() != nil{
		this.Abort("404")
	}

	errmsg := make(map[string] string)
	if this.Ctx.Request.Method == "POST"{
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))

		adminuser.Account = account
		adminuser.Password = password

		if len(errmsg) == 0 {
			adminuser.Update()
			this.Redirect("/admin/adminuser/list",302)
		}
	}

	this.Data["errmsg"] = errmsg
	this.Data["adminuser"] = adminuser
	this.display()
}



//登录
func (this *AdminuserController) Login(){
	if this.Ctx.Request.Method == "POST"{
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		// remember := this.GetString("remember")

		if account != "" && password != ""{
			var adminuser models.Adminuser
			adminuser.Account = account
			if adminuser.Read("account") != nil	|| adminuser.Password != password {
				this.Data["errmsg"] = "账号或密码错误"
			}else{

					adminuser.Lastip = this.getClientIp()
					adminuser.Update()
					authkey := adminuser.Password
					//TODO remember
					//TODO password  没加密
					this.Ctx.SetCookie("auth",strconv.FormatInt(adminuser.Id,10) + "|" + authkey)

			}
			this.Redirect("/admin",302)
		}

	}

	this.TplNames = "admin/admin_login.html"
}

//管理员退出
func (this *AdminuserController) Logout(){
	this.Ctx.SetCookie("auth","")
	this.Redirect("/admin/adminuser/login",302)
}