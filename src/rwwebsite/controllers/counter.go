package controllers
import "github.com/astaxie/beego"
type CounterController struct {
	beego.Controller
}

//@router /counter/ex/:id [get]
func (this *CounterController)Exception(){
	this.Data["id"] = this.Ctx.Input.Param(":id")
	this.TplNames = "ex.html"
}

//@router /counter/biz/:id [get]
func (this *CounterController)Business(){
	this.Data["id"] = this.Ctx.Input.Param(":id")
	this.TplNames = "biz.html"
}

//@router /counter/custom/:id [get]
func (this *CounterController)Custom()  {
	this.Data["id"] = this.Ctx.Input.Param(":id")
	this.TplNames = "custom.html"
}