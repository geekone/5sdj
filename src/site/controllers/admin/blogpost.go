package admin

import (
	"site/models"
	"strings"
	"github.com/astaxie/beego/orm"
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


//保存和更新重构在一起
func (this *BlogpostController) Save(){
	var (
		id int64 =0
		title string = strings.TrimSpace(this.GetString("title"))
		content string = this.GetString("content")
		tags string = strings.TrimSpace(this.GetString("tags"))
		urlname string = strings.TrimSpace(this.GetString("urlname"))
		color string = strings.TrimSpace(this.GetString("color"))
		//TODO timestr string = strings.TrimSpace(this.GetString("posttime"))
		status int64 = 0
		istop int8 =0
		urltype int8 = 0
		blogPost models.BlogPost
	)

	if title == ""{
		this.showmsg("标题不能为空")
	}

	id,_ = this.GetInt("id")
	status,_ = this.GetInt("status")

	if this.GetString("istop") == "1"{
		istop = 1
	}

	if this.GetString("urltype") == "1"{
		urltype = 1
	}

	//TODO A

	//整理TAG,先分离tags成arr,
	addtags := make([]string,0)
	if tags != ""{
		tagarr :=  strings.Split(tags,",")
		for _,v := range tagarr{
			if tag := strings.TrimSpace(v);tag != ""{
					exists := false		//判断前面已经加入了
					for _,vv := range addtags {
						if vv == tag{
							exists = true 
							break
						}
					}

					if !exists {
						addtags = append(addtags,tag)
					}
			}
		}
	}

	//如果ID = 0 是插入
	if id < 1 {
		blogPost.Userid = this.adminid
		blogPost.Author = this.account
		blogPost.Insert()
	}else{
		// blogPost.Id = id
		//TODO ERR

		// if blogPost.Tags != ""{

		// }
	}


	if len(addtags) > 0 {
		for _,v:=range addtags {
			tag := models.Tag{Name:v}
			if tag.Read("Name") == orm.ErrNoRows {
				tag.Count = 1
				tag.Insert()		//添加tag表或更新tag次数
			}else{
				tag.Count += 1
				tag.Update("Count")
			}
			//TODO 添加到tag post 关联表里
			//tp := models.TagPost{Tagid:tag.Id,Postid:blogPost.Id,Poststatus:int8(status)}
			//tp.Insert()
		}
		blogPost.Tags = "," + strings.Join(addtags,",") + ","	//重返回post tags字段内容

	}



	//最后可以update一起

	blogPost.Status = int8(status)
	blogPost.Title = title
	blogPost.Color = color
	blogPost.Istop = istop
	blogPost.Content = content
	blogPost.Urlname = urlname
	blogPost.Urltype = urltype

	blogPost.Update()

	this.Redirect("/admin/blogpost/list",302)
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