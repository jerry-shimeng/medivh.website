package routers

import (
	"rwwebsite/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Include( &controllers.MainController{})
	beego.Include( &controllers.DetailController{})
	beego.Include( &controllers.SystemInfoController{})
	beego.Include( &controllers.CounterController{})
	beego.Include( &controllers.OnceDataController{})
}
