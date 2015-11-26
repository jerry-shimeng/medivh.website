package models
import "encoding/json"

type CmdModel struct {
	Mark,Alias string
	Level int
	Module int `json:"ModuleType"`
	Counter int `json:"CounterType"`
}

func (this *CmdModel)ToJson()string{
	msg,_ := json.Marshal(this)
	return string(msg)
}

