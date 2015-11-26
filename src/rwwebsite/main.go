package main

import (
	_ "rwwebsite/routers"
	"github.com/astaxie/beego"
	"time"
)

func main() {
	beego.AddFuncMap("GetTimeString",ConvertTimeString)
	beego.StaticDir["/static"] = "static"
	beego.Run()
}

func ConvertTimeString(i float64)string{
	return time.Unix(int64(i),0).Format("2006-01-02 15:04:05")
}

