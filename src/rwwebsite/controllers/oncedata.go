package controllers
import (
	"github.com/astaxie/beego"
	"strings"
	"rwwebsite/common"
)
type OnceDataController struct  {
	beego.Controller
}

//@router /counter/business/:id [get]
func (this *OnceDataController)GetInfo(){
	params := this.Ctx.Input.Params
	id := params[":id"]
	url := this.Ctx.Input.Uri()

	arr := strings.Split(url,"?")
	if len(arr) != 2 {
		this.Ctx.Output.Status = 502
	}else {
		d,err := new(common.CmdSender).Get(id,common.OnceDataInfoUrl,arr[1])
		if err != nil {
			this.Data["json"] = err
		}else {
			this.Data["json"] = d
		}
	}
	this.ServeJson()
}

//@router /counter/exception/:id [get]
func (this *OnceDataController)Exception(){
	params := this.Ctx.Input.Params
	id := params[":id"]
	url := this.Ctx.Input.Uri()

	arr := strings.Split(url,"?")
	if len(arr) != 2 {
		this.Ctx.Output.Status = 502
	}else {
		d,err := new(common.CmdSender).Get(id,common.OnceDataInfoUrl,arr[1])
		if err != nil {
			this.Data["json"] = err
		}else {
			this.Data["json"] = d
		}
	}

	this.ServeJson()
}

//@router /counter/customer/:id [get]
func (this *OnceDataController)Customer(){
	params := this.Ctx.Input.Params
	id := params[":id"]
	url := this.Ctx.Input.Uri()

	arr := strings.Split(url,"?")
	if len(arr) != 2 {
		this.Ctx.Output.Status = 502
	}else {
		d,err := new(common.CmdSender).Get(id,common.OnceDataInfoUrl,arr[1])
		if err != nil {
			this.Data["json"] = err
		}else {
			this.Data["json"] = d
		}
	}

	this.ServeJson()
}