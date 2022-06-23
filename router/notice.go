package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sptbackend/language"
	"sptbackend/sql"
)

func notice(context  *gin.Context)  {
	header := context.GetHeader("accept-language")
	if header =="zh-CN" {
		noti := sql.GetNotice(1)
		content :=sql.GetContent(4)
		if noti.Open !=1 {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: gin.H{"text":gin.H{"title":content.Title,"content":content.Content},"notice":gin.H{"isopen":false}}})
			return
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: gin.H{"text":gin.H{"title":content.Title,"content":content.Content},"notice":gin.H{"isopen":true,"content":noti.Content}}})
		return
	}else if  header=="fr-FR"{
		noti := sql.GetNotice(3)
		content :=sql.GetContent(6)
		if noti.Open !=1 {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS,Data: gin.H{"text":gin.H{"title":content.Title,"content":content.Content},"notice":gin.H{"isopen":false}}})
			return
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS,Data: gin.H{"text":gin.H{"title":content.Title,"content":content.Content},"notice":gin.H{"isopen":true,"content":noti.Content}}})
		return
	}else {
		noti := sql.GetNotice(2)
		content :=sql.GetContent(5)
		if noti.Open !=1 {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_SUCCESS,Data: gin.H{"text":gin.H{"title":content.Title,"content":content.Content},"notice":gin.H{"isopen":false}}})
			return
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_SUCCESS,Data: gin.H{"text":gin.H{"title":content.Title,"content":content.Content},"notice":gin.H{"isopen":true,"content":noti.Content}}})
		return
	}
}

type announceContent struct {
	Chinese string `json:"chinese"`
	English string `json:"english"`
	French string `json:"french"`
}
func announce(context  *gin.Context)  {

	var  aounceText announceContent
	err := context.BindJSON(&aounceText)
	if err !=nil {
		context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "announce param wrong"})
		return
	}
	if len(aounceText.Chinese)==0||len(aounceText.French)==0||len(aounceText.English)==0 {
		context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "language param lack"})
		return
	}
	sql.UpdateNotice(aounceText.Chinese,1)
	sql.UpdateNotice(aounceText.English,2)
	sql.UpdateNotice(aounceText.French,3)
	context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "announce success"})
}

func stopNotice(context  *gin.Context) {
	sql.StopAnnounce()
	context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "stop success"})
}
