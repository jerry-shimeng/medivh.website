package common
import (
	"rwwebsite/models"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
)

type CmdSender struct {

}

//发送请求：GET
func (this *CmdSender)Get(id,url,query string)(*models.ReturnModel,error) {
	url =url+"/"+id+"?"+query
	res,err := http.Get(url)
	if err != nil {
		return nil,err
	}
	if res == nil {
		return nil,errors.New("连接失败")
	}
	defer res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	if err!=nil {
		return nil,err
	}

	model := new(models.ReturnModel)
	json.Unmarshal(body,model)
	return model,nil
}