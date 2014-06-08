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


	//One To Many	测试数据

	// var profile models.Profile
	// profile.Age = 10
	// var user models.User
	// user.Username = "abc"
	// user.Password = "abc"
	// user.Email = "abc"
	// user.Profile = &profile
	// profile.Insert()
	// user.Insert()

	// var post models.Post
	// post.Title = "title"
	// post.User = &user
	// post.Insert()

	// var post1 models.Post
	// post1.Title = "title1"
	// post1.User = &user
	// post1.Insert()

	// var profile1 models.Profile
	// profile1.Age = 20
	// var user1 models.User
	// user1.Username = "abc2"
	// user1.Password = "abc2"
	// user1.Email = "abc2"
	// user1.Profile = &profile1
	// profile1.Insert()
	// user1.Insert()

	// var post2 models.Post
	// post2.Title = "title"
	// post2.User = &user1
	// post2.Insert()

	// var post3 models.Post
	// post3.Title = "title1"
	// post3.User = &user1
	// post3.Insert()



	// var posts []*models.Post
	// var post models.Post
	// post.Query().Filter("User",2).RelatedSel().All(&posts)
	// for _,post := range posts{
	// 	fmt.Printf("Id: %d, UserName: %s, Title: %s\n", post.Id, post.User.Username, post.Title)
	// }

	var user  models.User
	user.Query().Filter("Id",1).RelatedSel().One(&user)
	user.Delete()

	//TEST
	this.display()
}