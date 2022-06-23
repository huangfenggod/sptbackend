package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sptbackend/language"
	"sptbackend/service"
)

func getorder(context *gin.Context)  {
	header := context.GetHeader("accept-language")
	address := context.Query("address")
	switch header {
	case "zh-CN":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		orders := service.GetAllOrderByAddress(address,header)
		if orders.Len()==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "沒有質押訂單"})
			return
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: orders.Val()[:orders.Len()]})
		return
	case "fr-FR":
		if len(address)==0{
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_WRONG_ADDRESS})
			return
		}
		orders := service.GetAllOrderByAddress(address,header)
		if orders.Len()==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Aucune commande de gage"})
			return
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS,Data: orders.Val()[:orders.Len()]})
		return

	default:

		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_WRONG_ADDRESS})
			return
		}
		orders := service.GetAllOrderByAddress(address,header)
		if orders.Len()==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "No pledge order"})
			return
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_SUCCESS,Data: orders.Val()[:orders.Len()]})
		return
	}


}
