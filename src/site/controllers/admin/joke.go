package admin

/**
 * 文章控制
 */

import (
 "site/models"
 "strings"
)


type JokeController struct{
	baseController
}


//joke列表
func (this *JokeController) List(){
	var list []*models.Joke 		//数组对象
	var joke models.Joke            //
	count, _ := joke.Query().Count()		//统计总数
	if count > 0{
		joke.Query().OrderBy("-id").All(&list)		//如果总数不为空,就全部取出
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.display()
}

//添加分类
func (this *JokeController) Add(){
	if this.Ctx.Request.Method == "POST"{
		_title := strings.TrimSpace(this.GetString("title"))
		_content := this.GetString("content")
		var joke models.Joke
		joke.Title = _title
		joke.Content = _content

		if err := joke.Insert(); err !=nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/joke/list",302)
	}
	this.display()
}


//删除
func (this *JokeController) Delete(){
	id,_ := this.GetInt("id")
	joke := models.Joke{Id:id}
	if joke.Read() == nil{
		joke.Delete()
	}
	this.Redirect("/admin/joke/list",302)
}


func (this *JokeController) Edit(){
	id,_:= this.GetInt("id")
	joke := models.Joke{Id:id}
	if joke.Read() != nil{
		this.Abort("404")
	}

	//如果提交修改
	if this.Ctx.Request.Method == "POST"{
		_title := strings.TrimSpace(this.GetString("title"))
		_content := this.GetString("content")
		joke.Title = _title
		joke.Content = _content
		joke.Update()
		this.Redirect("/admin/joke/list",302)
	}

	this.Data["joke"] = joke
	this.display()
}