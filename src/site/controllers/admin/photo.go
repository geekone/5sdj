package admin

/**
 * 图片的控制
 */

import (
 	"fmt"
 	"site/models"
 	"strings"
 	"os"
 	"time"
)


type PhotoController struct {
	baseController
}


//图片列表
func (this *PhotoController) List(){
	var page int64 				//当前页号
	var pagesize int64 = 2 	//当前每页显示的数量
	var list []*models.Photo
	var photo models.Photo
	//如果当前页号为空
	if page,_ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page -1) * pagesize	//偏移量
	count,_ := photo.Query().Count()
	if count > 0 {
		photo.Query().OrderBy("-id").Limit(pagesize,offset).RelatedSel().All(&list)
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] =  models.NewPager(page, count, pagesize, "/admin/photo/list", true).ToString()
	this.display()
}


func(this *PhotoController) Add(){
	if this.Ctx.Request.Method == "POST" {
		_title := strings.TrimSpace(this.GetString("title"))
		_intro := this.GetString("intro")
		_img := strings.TrimSpace(this.GetString("img"))
		_url := strings.TrimSpace(this.GetString("url"))
		_cid,_:= this.GetInt("cid")
		var photo  models.Photo
		photo.Title = _title
		photo.Intro = _intro
		photo.Img = _img
		photo.Url = _url
		_category := models.Category{Id:_cid}
		photo.Category = &_category
		if err := photo.Insert(); err != nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/photo/list",302)

	}
	var categories []*models.Category
	var category models.Category
	category.Query().Filter("cateid", 3).OrderBy("-id").All(&categories)
	this.Data["categories"] = categories
	this.display()
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


//上传图片
func (this *PhotoController) Upload(){
	if this.Ctx.Request.Method == "POST"{
		_,header,err := this.GetFile("upfile")
		ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename,"."):])
		out := make(map[string]string)
		out["url"] = ""
		out["fileType"] = ext
		out["original"] = header.Filename
		out["state"] = "SUCCESS"
		if err !=nil{
			this.showmsg(err.Error())
		}else{
			savepath := "./static/upload/" + time.Now().Format("20060102")
			if err := os.MkdirAll(savepath, os.ModePerm); err != nil {
				out["state"] = err.Error()
			} else {
				filename := fmt.Sprintf("%s/%d%s", savepath, time.Now().UnixNano(), ext)
				if err := this.SaveToFile("upfile", filename); err != nil {
					out["state"] = err.Error()
				} else {
					out["url"] = filename[1:]
				}
			}
		}
		
	}

	this.display()
}