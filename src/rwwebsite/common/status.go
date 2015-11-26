package common
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"rwwebsite/models"
)

type StatusInfo struct {

}

func (this *StatusInfo)GetAll()(*models.ReturnModel,error){
	url:= StatusInfoUrl
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
	//fmt.Println(string(body))
	model := new(models.ReturnModel)
	json.Unmarshal(body,model)
	return model,nil
}

