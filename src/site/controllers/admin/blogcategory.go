package admin

import (
	"github.com/astaxie/beego/validation"
	"site/models"
	"strings"
)


type BlogcategoryController struct{
	baseController
}


func (this *BlogcategoryController) List(){
	var list []*models.BlogCategory  //准备list 指针装载查询的所有分类
	var blogCategory models.BlogCategory 
	count,_:=blogCategory.Query().Count()
	if  count > 0{   
		blogCategory.Query().OrderBy("-id").All(&list)
	}
	this.Data["count"] = count
	this.Data["list"] = list 
	this.display()

}

//添加分类
func (this *BlogcategoryController) Add(){

	if this.Ctx.Request.Method == "POST"{
		name := strings.TrimSpace(this.GetString("name"))

		var blogCategory models.BlogCategory
		blogCategory.Name = name
		if err:=blogCategory.Insert(); err != nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/blogcategory/list",302)
	}

	this.display()
}


//删除,通过id
func (this *BlogcategoryController) Delete(){
	id,_:=this.GetInt("id")
	blogCategory := models.BlogCategory{Id:id}
	if blogCategory.Read() == nil{
		blogCategory.Delete()
	}
	this.Redirect("/admin/blogcategory/list",302)
}

//编辑
func (this *BlogcategoryController) Edit() {
	id, _ := this.GetInt("id")
	blogCategory := models.BlogCategory{Id: id}
	if blogCategory.Read() != nil {
		this.Abort("404")
	}


	errmsg :=make(map[string] string)
	if this.Ctx.Request.Method == "POST"{
		name := strings.TrimSpace(this.GetString("name"))
		valid := validation.Validation{}
		if v:=valid.Required(name,"name");!v.Ok{
			errmsg["name"] = "name is Required"
		}else{
			blogCategory.Name = name
		}

		if len(errmsg) == 0{
			blogCategory.Update()
			this.Redirect("/admin/blogcategory/list",302)
		}
	}

	this.Data["errmsg"] = errmsg
	this.Data["blogCategory"] = blogCategory
	
	this.display()
}