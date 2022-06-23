package myrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
)

var rsaPrivateKey *rsa.PrivateKey
//需要对rsa文件进行初始化
func InitRsa() error {
	file, err := readFile()
	if err!=nil {

		return err
	}
	key, err1 := getPrivateKey(file)
	if  err1 !=nil{
		return err1
	}
	rsaPrivateKey =key
	return nil
}

//RSA解密方法
func RSADecode(data string) (string,error) {
	var str string
	decodeString, err2 := base64.StdEncoding.DecodeString(data)
	if err2 !=nil{
		return str ,err2
	}
	v15, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodeString)
	if err!=nil {
		return str ,err
	}
	return string(v15),nil
}

func readFile() ([]byte,error){
	fileInfo, err := ioutil.ReadFile("./private")
	if err !=nil {
		log.Printf("read private error :%s",err)
	}
	return fileInfo,err
}
func getPrivateKey(private []byte) (*rsa.PrivateKey,error) {
	decode, _ := pem.Decode(private)
	key, err := x509.ParsePKCS1PrivateKey(decode.Bytes)
	if err!=nil {
		log.Printf("read private error :%s",err)
	}
	return key,err
}




