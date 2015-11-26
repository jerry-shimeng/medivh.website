package controllers

import (
	"github.com/astaxie/beego"
"rwwebsite/common"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

//@router /home [get]
func (c *MainController) Get() {
	//获取所有连接信息
	model := new(common.StatusInfo)
	s,err := model.GetAll()
	if s!=nil && s.Code == 0 && s.Data!=nil {
		c.Data["list"] = s.Data
		temp,_:=json.Marshal(s.Data)
		c.Data["data"]= string(temp)
	}else if s != nil {
		c.Data["err"] = s.Message
	}else {
		c.Data["err"] = err
	}
	c.TplNames = "home.html"
}

//@router /home/:id [get]
func (this *MainController)Index(){
	id := this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.TplNames ="index.html"
}
