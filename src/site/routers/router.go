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
    beego.Router("/admin/user/list",&admin.UserController{},"*:List")
    beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
    beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")
    beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")

    //admin - catgory
    beego.Router("/admin/category/list", &admin.CategoryController{}, "*:List")
    beego.Router("/admin/category/add", &admin.CategoryController{}, "*:Add")
    beego.Router("/admin/category/delete", &admin.CategoryController{}, "*:Delete")
    beego.Router("/admin/category/edit", &admin.CategoryController{}, "*:Edit")


    // admin - movie 
    beego.Router("/admin/movie/list",&admin.MovieController{},"*:List");
    beego.Router("/admin/movie/add",&admin.MovieController{},"*:Add")
    beego.Router("/admin/movie/delete",&admin.MovieController{},"*:Delete")
    beego.Router("/admin/movie/edit",&admin.MovieController{},"*:Edit")

    //admin - photo
    beego.Router("/admin/photo/list",&admin.PhotoController{},"*:List");
    beego.Router("/admin/photo/add",&admin.PhotoController{},"*:Add")
    beego.Router("/admin/photo/delete",&admin.PhotoController{},"*:Delete")
    beego.Router("/admin/photo/edit",&admin.PhotoController{},"*:Edit")



    beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
    beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
    //TODO Article 没补
    

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
