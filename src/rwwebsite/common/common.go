package common
import "strings"

func StringConvertToArray(str string)*[]string{
	if len(str) == 0 {
		return nil
	}
	s := strings.Split(str,",")
	return &s
}