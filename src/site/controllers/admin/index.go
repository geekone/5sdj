package admin

import (
	"site/models"
	// "fmt"
)


type IndexController struct {
	baseController
}


func (this *IndexController) Index(){

	//One To One 

	// var profile models.Profile
	// profile.Age = 10
	// var user models.User
	// user.Username = "abc1"
	// user.Password = "abc1"
	// user.Email = "abc1"
	// user.Profile = &profile

	// profile.Insert()
	// if err := user.Insert(); err !=nil{
	// 			this.showmsg(err.Error());
	// }

	// var user  models.User
	// user.Query().Filter("Id",1).RelatedSel().One(&user)
	// fmt.Println(user.Username)
	// fmt.Println(user.Profile.Age)

	//通过 User 反向查询 Profile：

	// var profile models.Profile
	// profile.Query().Filter("User__Id", 1).One(&profile)
	// fmt.Println(profile.Age)


	//One To Many

	var profile models.Profile
	profile.Age = 10
	var user models.User
	user.Username = "abc"
	user.Password = "abc"
	user.Email = "abc"
	user.Profile = &profile
	profile.Insert()
	user.Insert()

	var post models.Post
	post.Title = "title"
	post.User = &user
	post.Insert()

	post.Title = "title1"
	post.User = &user
	post.Insert()

	// var posts []*models.Post
	// var post models.Post
	// post.Query().Filter("User",15).RelatedSel().All(&posts)


	//TEST
	this.display()
}