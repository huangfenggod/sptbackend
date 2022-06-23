package crypto

import (
	"encoding/base64"
	"errors"
)

func BaseDecode(str string)(string,error ) {
	var  decodeStr string
	decodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return decodeStr,err
	}
	return string(decodeString),nil
}

func PasswordDecode(str string)(string,error ) {
	var  s string
	i := len(str)
	if i <10 {
		return s,errors.New("hash wrong")
	}
	s = str[:10]+str[20:]
	return s,nil
}
