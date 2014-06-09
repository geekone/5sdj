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
	var page int64				//当前页号
	var pagesize int64 = 10 		//当前每页显示的数量


	var list []*models.Joke 		//数组对象
	var joke models.Joke            

	if page,_ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page -1) * pagesize;

	count, _ := joke.Query().Count()		//统计总数
	if count > 0{
		joke.Query().OrderBy("-id").Limit(pagesize,offset).All(&list)		//如果总数不为空,就全部取出
	}
	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page,count,pagesize,"/admin/joke/list",true).ToString()
	this.display()
}

//添加分类
func (this *JokeController) Add(){
	if this.Ctx.Request.Method == "POST"{
		_title := strings.TrimSpace(this.GetString("title"))
		_content := this.GetString("content")
		// _cid,_:= this.GetInt("cid")
		var joke models.Joke
		joke.Title = _title
		joke.Content = _content
		// joke.CategoryId = _cid

		if err := joke.Insert(); err !=nil{
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/joke/list",302)
	}
	var categories []*models.Category
	var category models.Category
	//category.Query().Filter("catestr", "joke").OrderBy("-id").All(&categories)			//取出所有笑话类的分类
	category.Query().OrderBy("-id").All(&categories)
	this.Data["categories"] = categories
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