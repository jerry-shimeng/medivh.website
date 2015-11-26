package common
import (
	"net/http"
	"io/ioutil"
	"errors" 
"rwwebsite/models"
	"encoding/json"
)

type SystemInfo struct  {


}
func (this *SystemInfo)GetSysInfo(id string)(*models.ReturnModel,error){
	res,err := http.Get(SystemInfoUrl+"/"+id)
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

func (this *SystemInfo)ApiGet(id,query string)(*models.ReturnModel,error){
	url :=SystemInfoUrl+"/"+id+"?"+query
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

//func (this *SystemInfo)ApiPost(id string,body *[]byte) (*models.ReturnModel,error) {
//
//}