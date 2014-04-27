package admin

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"site/models"
)

type UserController struct {
	baseController
}


func (this *UserController) List(){
	var user models.User
	var list []*models.User
	count,_:=user.Query().Count()
	if count > 0{
		user.Query().OrderBy("-id").All(&list)
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.display()
}


//添加用户
func (this *UserController) Add(){

	errmsg := make(map[string] string)

	if this.Ctx.Request.Method == "POST"{
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		password2 := strings.TrimSpace(this.GetString("password2"))
		email := strings.TrimSpace(this.GetString("email"))
	
		//验证
		valid := validation.Validation{}

		if v := valid.Required(username, "username"); !v.Ok {
			errmsg["username"] = "请输入用户名"
		} else if v := valid.MaxSize(username, 15, "username"); !v.Ok {
			errmsg["username"] = "用户名长度不能大于15个字符"
		}

		if v := valid.Required(password, "password"); !v.Ok {
			errmsg["password"] = "请输入密码"
		}

		if v := valid.Required(password2, "password2"); !v.Ok {
			errmsg["password2"] = "请再次输入密码"
		} else if password != password2 {
			errmsg["password2"] = "两次输入的密码不一致"
		}

		if v := valid.Required(email, "email"); !v.Ok {
			errmsg["email"] = "请输入email地址"
		} else if v := valid.Email(email, "email"); !v.Ok {
			errmsg["email"] = "Email无效"
		}

		//如果都验证通过了，就保存并跳转
		if len(errmsg) == 0{
			var user models.User
			user.Username = username
			user.Password = models.Md5([]byte(password))
			user.Email = email
			if err := user.Insert(); err != nil {
				this.showmsg(err.Error())
			}
			this.Redirect("/admin/user/list",302)
		}
		
	}

	this.Data["errmsg"] = errmsg
	this.display()
}