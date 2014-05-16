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
	var page int64 				//当前页号
	var pagesize int64 = 10 	//当前每页显示的数量

	var list []*models.Category
	var category models.Category

	//如果当前页号为空
	if page,_ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page -1) * pagesize	//偏移量
	count,_ := category.Query().Count()		//总数

	if count > 0{
		category.Query().OrderBy("-id").Limit(pagesize,offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] =  models.NewPager(page, count, pagesize, "/admin/category/list", true).ToString()
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