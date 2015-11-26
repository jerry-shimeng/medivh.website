package controllers
import "github.com/astaxie/beego"
type DetailController struct {
	beego.Controller
}

//@router /detail/:id [get]
func (this *DetailController)Get(){
	id := this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.TplNames="detail.html"
}