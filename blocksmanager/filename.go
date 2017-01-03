package blocksmanager
import (
	"encoding/json"
	"encoding/base64"
)
func GetFileName(smallfile SmallFile)(string,error) {
	body, err := json.Marshal(smallfile)
	if err != nil {
		return "",err
	}
	return base64.URLEncoding.EncodeToString(body),nil
}

func GetSmallFile(s string)(*SmallFile , error) {
	body, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	var smallfile SmallFile
	err = json.Unmarshal(body,&smallfile)
	if err !=  nil{
		return nil,err
	}
	return &smallfile,nil
}
