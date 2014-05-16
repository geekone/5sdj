package routers

import (
	"site/controllers/home"
	"site/controllers/blog"
	"site/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
    //home
	beego.Router("/", &home.IndexController{},"*:Index")

	//blog
	beego.Router("/blog", &blog.IndexController{},"*:Index")

	//admin
    beego.Router("/admin", &admin.IndexController{},"*:Index")
    
     //admin - adminuser
    beego.Router("/admin/adminuser/list", &admin.AdminuserController{}, "*:List")       
    beego.Router("/admin/adminuser/add", &admin.AdminuserController{}, "*:Add")
    beego.Router("/admin/adminuser/delete", &admin.AdminuserController{}, "*:Delete")
    beego.Router("/admin/adminuser/edit", &admin.AdminuserController{}, "*:Edit")
    beego.Router("/admin/adminuser/login",&admin.AdminuserController{},"*:Login")
    beego.Router("/admin/adminuser/logout",&admin.AdminuserController{},"*:Logout")

    //TODO User
    beego.Router("/admin/user/list",&admin.UserController{},"*:List")           //用户列表
    beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")           //用户添加
    beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")     //用户删除
    beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")         //用户编辑

    //admin - catgory
    beego.Router("/admin/category/list", &admin.CategoryController{}, "*:List")     //分类列表
    beego.Router("/admin/category/add", &admin.CategoryController{}, "*:Add")       //添加分类
    beego.Router("/admin/category/delete", &admin.CategoryController{}, "*:Delete") //删除分类
    beego.Router("/admin/category/edit", &admin.CategoryController{}, "*:Edit")     //编辑分类


    // admin - movie 
    beego.Router("/admin/movie/list",&admin.MovieController{},"*:List");            //视频列表
    beego.Router("/admin/movie/add",&admin.MovieController{},"*:Add")               //添加视频
    beego.Router("/admin/movie/delete",&admin.MovieController{},"*:Delete")         //删除视频
    beego.Router("/admin/movie/edit",&admin.MovieController{},"*:Edit")             //编辑视频

    //admin - photo
    beego.Router("/admin/photo/list",&admin.PhotoController{},"*:List");            //图片列表
    beego.Router("/admin/photo/add",&admin.PhotoController{},"*:Add")               //添加图片
    beego.Router("/admin/photo/delete",&admin.PhotoController{},"*:Delete")         //删除图片
    beego.Router("/admin/photo/edit",&admin.PhotoController{},"*:Edit")             //编辑图片
    beego.Router("/admin/photo/upload",&admin.PhotoController{},"*:Upload")         //上传图片


    beego.Router("/admin/joke/list", &admin.JokeController{}, "*:List")             //Joke列表
    beego.Router("/admin/joke/add", &admin.JokeController{}, "*:Add")               //添加joke
    beego.Router("/admin/joke/delete",&admin.JokeController{},"*:Delete")         //删除Joke
    beego.Router("/admin/joke/edit",&admin.JokeController{},"*:Edit")             //编辑Joke

    








    //admin - blog category
    beego.Router("/admin/blogcategory/list", &admin.BlogcategoryController{}, "*:List")
    beego.Router("/admin/blogcategory/add", &admin.BlogcategoryController{}, "*:Add")
    beego.Router("/admin/blogcategory/delete", &admin.BlogcategoryController{}, "*:Delete")
    beego.Router("/admin/blogcategory/edit", &admin.BlogcategoryController{}, "*:Edit")


    //admin - blog blogpost
    beego.Router("/admin/blogpost/list", &admin.BlogpostController{}, "*:List")
    beego.Router("/admin/blogpost/add", &admin.BlogpostController{}, "*:Add")
    beego.Router("/admin/blogpost/delete", &admin.BlogpostController{}, "*:Delete")
    beego.Router("/admin/blogpost/edit", &admin.BlogpostController{}, "*:Edit")
    beego.Router("/admin/blogpost/save", &admin.BlogpostController{}, "post:Save")


}
