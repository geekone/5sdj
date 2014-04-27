package admin

import (
	"site/models"
	"strings"
)

type MovieController struct {
	baseController
}

//视频列表
func (this *MovieController) List(){
	var list []*models.Movie
	var movie models.Movie
	count,_ := movie.Query().Count()		//统计视频数量
	if count > 0{
		movie.Query().OrderBy("-id").All(&list)
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.display()
}


//添加视频
func (this *MovieController) Add(){
	if this.Ctx.Request.Method == "POST"{
		_title := strings.TrimSpace(this.GetString("title"))
		_intro := this.GetString("intro")
		var movie models.Movie
		movie.Title = _title
		movie.Intro = _intro
		if err := movie.Insert(); err!=nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/movie/list",302)
	}

	this.display()
}

func (this *MovieController) Delete(){
	id,_ := this.GetInt("id")
	movie := models.Movie{Id:id}
	if movie.Read() == nil{
		movie.Delete()
	}
	this.Redirect("/admin/movie/list",302)
}


func (this *MovieController) Edit(){
	id,_:= this.GetInt("id")
	movie := models.Movie{Id:id}
	if movie.Read() != nil {
		this.Abort("404")
	}
	//TODO 都没有验证字段
	if this.Ctx.Request.Method == "POST"{
		_title := strings.TrimSpace(this.GetString("title"))
		_intro := this.GetString("intro")
		movie.Title = _title
		movie.Intro = _intro
		movie.Update()
		this.Redirect("/admin/movie/list",302)
	}

	this.Data["movie"] = movie
	this.display()
}