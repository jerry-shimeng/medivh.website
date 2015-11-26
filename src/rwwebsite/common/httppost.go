package common
import (
	"net/http"
	"strings"
	"io/ioutil"
	"rwwebsite/models"
	"encoding/json"
)

func HttpPost(url string,postdata *string)(*[]byte,error){

	r := strings.NewReader(*postdata)

	res,err := http.Post(url,"application/json",r)
	if err != nil {
		return nil,err
	}
	defer  res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil,err
	}
	return &body,nil
}

func HttpPostBackReturnModel(url string,postdata *[]byte)(*models.ReturnModel,error){
	data := string(*postdata)
	r,err := HttpPost(url,&data)
	if err != nil {
		return nil,err
	}
	var v models.ReturnModel
	err = json.Unmarshal(*r,&v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}