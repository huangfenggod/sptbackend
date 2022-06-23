package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sptbackend/language"
	"sptbackend/myrsa"
	"sptbackend/nosql"
	"sptbackend/service"
	"sptbackend/sql"
	"strconv"
)

type pledges struct {
	Amount float32 `json:"amount"`
	Days int `json:"days"`
	Address string `json:"address"`
	Password string `json:"password"`
	Paddress string `json:"paddress"`
}
//获取质押比例而产生的的利率
func pledgerate(context *gin.Context)  {
	rate := nosql.GetRateFromRedis()
	context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success",Data: gin.H{"90days":float32(rate*100),"180days":float32(rate*230),"360days":float32(rate*500)}})
}



func pledge(context *gin.Context)  {
	header := context.GetHeader("accept-language")
	var ple pledges
	err := context.BindJSON(&ple)

	switch header {
	case "zh-CN":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		if len(ple.Password)<22||ple.Amount<100||len(ple.Address)==0|| !(ple.Days==90||ple.Days==180||ple.Days==360) {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		hash, err1:= myrsa.RSADecode(ple.Password)
		if err1!=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PASSWORD_WRONG})
			return
		}
		log.Printf("pledge into address:%s, amount: %s ,days:%s,hash:%s",ple.Address,strconv.FormatFloat(float64(ple.Amount),'e',5,32),strconv.Itoa(ple.Days),hash)
		//transactionHash := service.IdentifyTransactionHash(hash)
		//if !transactionHash {
		//	context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PASSWORD_WRONG})
		//	return
		//}

		//处理该账户存在的情况
		uid := sql.ExitsAddressUid(ple.Address)
		if uid >0 {
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			userInfo := sql.GetUserInfoByUid(uid)
			service.UpdateUserByNewOrder(uid,userInfo.Pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS})
			return
		}else
		{
			//如果该用户不存在创建新用户
			if len(ple.Paddress)==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_NO_RECOMMENDER})
				return
			}
			pid :=sql.ExitsAddressUid(ple.Paddress)
			if pid==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_NO_RECOMMENDER})
				return
			}
			//创建新用户
			sql.CreateUser(ple.Address,pid)
			newUid := sql.ExitsAddressUid(ple.Address)
			//插入订单
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			service.UpdateUserByNewOrder(newUid,pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS})
			return
		}
	case "fr-FR":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG})
			return
		}
		if len(ple.Password)<22||ple.Amount<100||len(ple.Address)==0|| !(ple.Days==90||ple.Days==180||ple.Days==360) {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG})
			fmt.Println(ple)
			return
		}
		hash, err1:= myrsa.RSADecode(ple.Password)
		if err1!=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PASSWORD_WRONG})
			return
		}
		log.Printf("pledge into address:%s, amount: %s ,days:%s,hash:%s",ple.Address,strconv.FormatFloat(float64(ple.Amount),'e',5,32),strconv.Itoa(ple.Days),hash)
		//transactionHash := service.IdentifyTransactionHash(hash)
		//if !transactionHash {
		//	context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PASSWORD_WRONG})
		//	return
		//}
		//处理该账户存在的情况
		uid := sql.ExitsAddressUid(ple.Address)
		if uid >0 {
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			userInfo := sql.GetUserInfoByUid(uid)
			service.UpdateUserByNewOrder(uid,userInfo.Pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS})
			return
		}else
		{
			//如果该用户不存在创建新用户
			if len(ple.Paddress)==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_NO_RECOMMENDER})
				return
			}
			pid :=sql.ExitsAddressUid(ple.Paddress)
			if pid==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_NO_RECOMMENDER})
				return
			}
			//创建新用户
			sql.CreateUser(ple.Address,pid)
			newUid := sql.ExitsAddressUid(ple.Address)
			//插入订单
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			service.UpdateUserByNewOrder(newUid,pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS})
			return
		}
	default:
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "fail,params wrong"})
			return
		}
		if len(ple.Password)<22||ple.Amount<100||len(ple.Address)==0|| !(ple.Days==90||ple.Days==180||ple.Days==360) {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "fail, some params wrong"})
			fmt.Println(ple)
			return
		}
		hash, err1:= myrsa.RSADecode(ple.Password)
		if err1!=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PASSWORD_WRONG})
			return
		}
		log.Printf("pledge into address:%s, amount: %s ,days:%s,hash:%s",ple.Address,strconv.FormatFloat(float64(ple.Amount),'e',5,32),strconv.Itoa(ple.Days),hash)

		//transactionHash := service.IdentifyTransactionHash(hash)
		//if !transactionHash {
		//	context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PASSWORD_WRONG})
		//	return
		//}
		//处理该账户存在的情况
		uid := sql.ExitsAddressUid(ple.Address)
		if uid >0 {
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			userInfo := sql.GetUserInfoByUid(uid)
			service.UpdateUserByNewOrder(uid,userInfo.Pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success"})
			return
		}else
		{
			//如果该用户不存在创建新用户
			if len(ple.Paddress)==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "No recommender"})
				return
			}
			pid :=sql.ExitsAddressUid(ple.Paddress)
			if pid==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "No recommender"})
				return
			}
			//创建新用户
			sql.CreateUser(ple.Address,pid)
			newUid := sql.ExitsAddressUid(ple.Address)
			//插入订单
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			service.UpdateUserByNewOrder(newUid,pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "pledge success"})
			return
		}

	}


}
func pledgefor(context *gin.Context){
	var ple pledges
	err := context.BindJSON(&ple)
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		if len(ple.Password)<22||ple.Amount<100|| !(ple.Days==90||ple.Days==180||ple.Days==360) {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		//处理该账户存在的情况
		uid := sql.ExitsAddressUid(ple.Address)
		if uid >0 {
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			userInfo := sql.GetUserInfoByUid(uid)
			service.UpdateUserByNewOrder(uid,userInfo.Pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS})
			return
		}else
		{
			//如果该用户不存在创建新用户
			if len(ple.Paddress)==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_NO_RECOMMENDER})
				return
			}
			pid :=sql.ExitsAddressUid(ple.Paddress)
			if pid==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_NO_RECOMMENDER})
				return
			}
			//创建新用户
			sql.CreateUser(ple.Address,pid)
			newUid := sql.ExitsAddressUid(ple.Address)
			//插入订单
			sql.InsertOrder(ple.Address,ple.Days,int64(ple.Amount*100000000))
			//新订单改变用户状态
			service.UpdateUserByNewOrder(newUid,pid,int64(ple.Amount*100000000))
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS})
			return
		}
}
