package admin

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"site/models"
	// "fmt"
	// "strconv"
)


/**
 * 分类控制
 */

type CategoryController struct{
	baseController
}


func (this *CategoryController) List(){
	var list []*models.Category
	var category models.Category
	category.Query().OrderBy("-id").All(&list)
	this.Data["list"] = list
	this.display()

}

//添加分类
func (this *CategoryController) Add(){
	errmsg :=make(map[string] string)
	if this.Ctx.Request.Method == "POST"{
		name := strings.TrimSpace(this.GetString("name"))
		valid := validation.Validation{}
		if v:=valid.Required(name,"name");!v.Ok{
			errmsg["name"] = "name is Required"
		}

		cname := strings.TrimSpace(this.GetString("cname"))
		cateid,_ := this.GetInt("cateid")
		catestr := strings.TrimSpace(this.GetString("catestr"))
		// cateid:= int64(_cateid)

		if len(errmsg) == 0{
			var category models.Category
			category.Name = name
			category.Cname = cname
			category.Catestr =catestr
			category.Cateid = cateid
			if err := category.Insert(); err !=nil{
				this.showmsg(err.Error());
			}
			this.Redirect("/admin/category/list",302)
		}
	}

	this.Data["errmsg"] = errmsg
	this.display()
}


func(this *CategoryController) Delete(){
	id,_:=this.GetInt("id")
	
	category := models.Category{Id:id}
	if category.Read() == nil{
		category.Delete()
	}

	this.Redirect("/admin/category/list",302)

}



//编辑
func (this *CategoryController) Edit() {
	id, _ := this.GetInt("id")
	category := models.Category{Id: id}
	if category.Read() != nil {
		this.Abort("404")
	}


	errmsg :=make(map[string] string)
	if this.Ctx.Request.Method == "POST"{
		name := strings.TrimSpace(this.GetString("name"))
		cname := strings.TrimSpace(this.GetString("cname"))
		cateid,_ := this.GetInt("cateid")
		catestr := strings.TrimSpace(this.GetString("catestr"))
		valid := validation.Validation{}
		if v:=valid.Required(name,"name");!v.Ok{
			errmsg["name"] = "name is Required"
		}else{
			category.Name = name
			category.Cname = cname
			category.Catestr =catestr
			category.Cateid = cateid
		}

		if len(errmsg) == 0{
			category.Update()
			this.Redirect("/admin/category/list",302)
		}
	}

	this.Data["errmsg"] = errmsg
	this.Data["category"] = category
	
	this.display()
}