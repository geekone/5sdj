package admin

/**
 * 文章控制
 */

type ArticleController struct{
	baseController
}


func (this *ArticleController) List(){

	this.display()

}

//添加分类
func (this *ArticleController) Add(){
	this.display()
}