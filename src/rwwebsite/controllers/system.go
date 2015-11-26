package controllers
import (
	"github.com/astaxie/beego"
	"rwwebsite/common"
)

type SystemInfoController struct {
	beego.Controller
}

//@router /sys/:id [get]
func (this *SystemInfoController)GetAll(){
	id := this.Ctx.Input.Param(":id")
	sys := new(common.SystemInfo)
	s,err := sys.GetSysInfo(id)
	if s != nil {
		this.Data["json"] = s
	}
	if err != nil {
		this.Data["json"] = err
	}
	this.ServeJson()
}

//@router /system/:id [get]
func (this *SystemInfoController)Index(){
	id := this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.TplNames = "sys.html"
}