package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sptbackend/language"
	"sptbackend/myrsa"
	"sptbackend/service"
)
type withDraw struct {
	Address string `json:"address"`
	Password string`json:"password"`
}
func pledgeWithDraw(context *gin.Context) {
	var withdraw withDraw
	header := context.GetHeader("accept-language")
	context.BindJSON(&withdraw)

	switch header {
	case "zh-CN":
		if len(withdraw.Address) == 0 || len(withdraw.Password) < 22 {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		decode, err := myrsa.RSADecode(withdraw.Password)
		if err != nil {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		hash := service.IdentifyTransactionHash(decode)
		if !hash {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_PARAMS_WRONG})
			return
		} else {
			log.Printf("wihtdraw address :%s ,hash:%s", withdraw.Address, decode)
			success := service.PledgeWithDrwa(withdraw.Address)
			if success {
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: language.CHINESE_SUCCESS})
				return
			} else {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_FAIL})
			}
		}

	case "fr-FR":

		if len(withdraw.Address) == 0 || len(withdraw.Password) < 22 {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.ENGLISH_PARAMS_WRONG})
			return
		}
		decode, err := myrsa.RSADecode(withdraw.Password)
		if err != nil {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		hash := service.IdentifyTransactionHash(decode)
		if !hash {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.FRENCH_PARANS_WRONG})
			return
		} else {
			log.Printf("wihtdraw address :%s ,hash:%s", withdraw.Address, decode)

			success := service.PledgeWithDrwa(withdraw.Address)
			if success {
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: language.FRENCH_SUCCESS})
				return
			} else {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.FRENCH_FAIL})
				return
			}
		}

	default:
		if len(withdraw.Address) == 0 || len(withdraw.Password) < 22 {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.ENGLISH_PARAMS_WRONG})
			return
		}
		decode, err := myrsa.RSADecode(withdraw.Password)
		if err != nil {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		hash := service.IdentifyTransactionHash(decode)
		if !hash {
			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.ENGLISH_PARAMS_WRONG})
			return
		} else {
			log.Printf("wihtdraw address :%s ,hash:%s", withdraw.Address, decode)

			success := service.PledgeWithDrwa(withdraw.Address)
			if success {
				context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: language.ENGLISH_SUCCESS})
				return
			} else {
				context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.ENGLISH_FAIL})
				return
			}
		}
	}
}

func withDrawRecords(context *gin.Context)  {
	header :=context.GetHeader("accept-language")
	address := context.Query("address")
	typ := context.Query("type")
	switch header {
	case "zh-CN":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "沒有提現記錄"})
			return
		}

		if len(typ)==0 {
			draw := service.GetWtihDrawRecords(address,"zh-CN")
			if draw.Len()==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "沒有提現記錄"})
				return
			}else {
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Data: draw.Val()[:draw.Len()]})
				return
			}
		}
	case "fr-FR":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Pas d'enregistrement de retrait"})
			return
		}

		if len(typ)==0 {
			draw := service.GetWtihDrawRecords(address,"fr-FR")
			if draw.Len()==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Pas d'enregistrement de retrait"})
				return
			}else {
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Data: draw.Val()[:draw.Len()]})
				return
			}
		}
	default:
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "No withdrawal record"})
			return
		}
		if len(typ)==0 {
			draw := service.GetWtihDrawRecords(address,"fr-FR")
			if draw.Len()==0 {
				context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "No withdrawal record"})
				return
			}else {
				context.JSON(http.StatusOK,ResponseUtil{Status: true,Data: draw.Val()[:draw.Len()]})
				return
			}
		}
	}
}

