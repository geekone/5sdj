package blog


type IndexController struct {
	baseController
}

func (this *IndexController) Index(){
	this.display()
}