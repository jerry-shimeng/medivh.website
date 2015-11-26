package common
import (
	"fmt"
	"github.com/astaxie/beego"
)

var (
	StatusInfoUrl =fmt.Sprintf("%s/v1/status/all",beego.AppConfig.String("apiurl"))
	SystemInfoUrl =fmt.Sprintf("%s/v1/s/info",beego.AppConfig.String("apiurl"))

	OnceDataInfoUrl = fmt.Sprintf("%s/v1/c/counter",beego.AppConfig.String("apiurl"))
)