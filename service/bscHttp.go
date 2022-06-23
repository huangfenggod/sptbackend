package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sptbackend/config"
	"sptbackend/nosql"
)

type BscResponse struct {
	Status	string `json:"status"`
	Message string 	`json:"message"`
	Result map[string]interface{}	`json:"result"`
}


//方法1
func BscHttp(transaction string)(error , BscResponse){
	var bsc BscResponse
	params := url.Values{}
	Url, _:= url.Parse("https://api.bscscan.com/api/")
	params.Set("module","transaction")
	params.Set("action","gettxreceiptstatus")
	params.Set("txhash",transaction)
	params.Set("apikey",config.Cfg.ApiKey)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	response, err := http.Get(urlPath)
	if err !=nil {
		log.Fatalf("get bscan transaction:%s fail",transaction)
		return err, bsc
	}
	defer response.Body.Close()
	all,_ := ioutil.ReadAll(response.Body)
	toStruct := StringToStruct(all)
	return err, toStruct
}

func StringToStruct(byteInfo []byte ) BscResponse  {
	var bsc BscResponse
	err := json.Unmarshal(byteInfo, &bsc)
	if err !=nil {
		log.Fatal("err get from bsc ")
	}
	return bsc
}
func IdentifyTransactionHash2(transactionHash string) bool{
	has := nosql.IsHasTxhash(transactionHash)
	if has {
		return false
	}
	err, response := BscHttp(transactionHash)
	if err !=nil {
		log.Printf(" tradeHash:%s acquire treasure fail because check BNB fail: %s",transactionHash,err)
		return false
	}
	if response.Result["status"] != "1" {
		return false
	}
	nosql.InsertHashToRedis(transactionHash)
	return true
}
