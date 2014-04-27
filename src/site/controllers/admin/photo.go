package admin

/**
 * 图片的控制
 */

import (

 	"site/models"
 	"strings"
)


type PhotoController struct {
	baseController
}


//图片列表
func (this *PhotoController) List(){
	var list []*models.Photo
	var photo models.Photo

	count,_ := photo.Query().Count()
	if count > 0 {
		photo.Query().OrderBy("-id").All(&list)
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.display()
}


func(this *PhotoController) Add(){
	if this.Ctx.Request.Method == "POST" {
		_title := strings.TrimSpace(this.GetString("title"))
		_intro := this.GetString("intro")
		_img := strings.TrimSpace(this.GetString("img"))
		_url := strings.TrimSpace(this.GetString("url"))
		var photo  models.Photo
		photo.Title = _title
		photo.Intro = _intro
		photo.Img = _img
		photo.Url = _url
		if err := photo.Insert(); err != nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/photo/list",302)

	}

	this.display()
}


func (this *PhotoController) Delete(){
	id,_ := this.GetInt("id")
	photo := models.Photo{Id:id}
	if photo.Read() == nil{
		photo.Delete()
	}
	this.Redirect("/admin/photo/list",302)
}



func(this *PhotoController) Edit(){

	id,_:= this.GetInt("id")
	photo := models.Photo{Id:id}
	if photo.Read() != nil {
		this.Abort("404")
	}
	//TODO 都没有验证字段
	if this.Ctx.Request.Method == "POST"{
		_title := strings.TrimSpace(this.GetString("title"))
		_intro := this.GetString("intro")
		_img := strings.TrimSpace(this.GetString("img"))
		_url := strings.TrimSpace(this.GetString("url"))
		photo.Title = _title
		photo.Intro = _intro
		photo.Img = _img
		photo.Url = _url
		photo.Update()
		this.Redirect("/admin/photo/list",302)
	}

	this.Data["photo"] = photo
	this.display()
}
