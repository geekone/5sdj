package admin

import (
	"site/models"
	"strings"
)


type BlogpostController struct{
	baseController
}


func (this *BlogpostController) List(){
	var list []*models.BlogPost
	var blogPost models.BlogPost
	count,_ := blogPost.Query().Count()
	if count > 0 {
		blogPost.Query().OrderBy("-id").All(&list)
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.display()

}

//添加分类
func (this *BlogpostController) Add(){
	
	if this.Ctx.Request.Method == "POST"{
		title:= strings.TrimSpace(this.GetString("title"))
		status,_:= this.GetInt("status")
		istop,_:= this.GetInt("istop")
		content := this.GetString("content")

		var blogPost models.BlogPost
		blogPost.Title = title 
		blogPost.Status = int8(status)
		blogPost.Istop = int8(istop)
		blogPost.Content = content

		if err :=blogPost.Insert();err != nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/blogpost/list",302)

	}

	this.display()
}



func (this *BlogpostController) Delete(){
	id,_ := this.GetInt("id")
	blogPost := models.BlogPost{Id:id}
	if blogPost.Read() == nil{
		blogPost.Delete()
	}
	this.Redirect("/admin/blogpost/list",302)
}


func (this *BlogpostController) Edit(){
	id,_:=this.GetInt("id")
	blogPost := models.BlogPost{Id:id}
	if blogPost.Read() != nil{
		this.Abort("404")
	}
	//TODO 没有完成修改
	this.Data["blogPost"] = blogPost
	this.display()
}