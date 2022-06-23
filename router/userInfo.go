package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sptbackend/language"
	"sptbackend/service"
	"sptbackend/sql"
)

func isexists(context *gin.Context)  {
	//header := context.GetHeader("accept-language")
	address := context.Query("address")
	paddress := context.Query("paddress")
	if len(address)>0 {
		uid := sql.ExitsAddressUid(address)
		if uid>0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: true})
			return
		}else {
			if len(paddress)>0 {
				addressUid := sql.ExitsAddressUid(paddress)
				if addressUid>0 {
					context.JSON(http.StatusOK,ResponseUtil{Status: true})
					return
				}
			}
		}
	}
	context.JSON(http.StatusOK,ResponseUtil{Status: false})
	return
}

func userInfo(context *gin.Context)  {
	header := context.GetHeader("accept-language")
	address := context.Query("address")
	switch header {
	case "zh-CN":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
		}
		info, exists := service.GetUserInfo(address)
		pidInfo := sql.GetUserInfoByUid(info.Pid)
		if !exists {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_WRONG_ADDRESS})
			return
		}else {
			var grade string
			switch info.Grade {
			case 4:
				grade = "恒星玩家"
			case 2:
				grade = "彗星玩家"
			case 3:
				grade = "行星玩家"
			default:
				grade = "星火玩家"
			}
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: gin.H{"address":info.Address,"recommender":pidInfo.Address,"grade":grade,"teamnumber":info.TeamNumber+1,"teampledge":float32(info.TotalTeamPledge)/100000000,"recommenders":info.Recommenders,"comet":info.Comet,"planet":info.Planet,"fixedstart":info.Fixstar,"todayreward":float32(info.TodayReward)/100000000}})
			return
		}

	case "fr-FR":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG})
		}
		info, exists := service.GetUserInfo(address)
		pidInfo := sql.GetUserInfoByUid(info.Pid)
		if !exists {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_WRONG_ADDRESS})
			return
		}else {
			var grade string
			switch info.Grade {
			case 4:
				grade = "Star player"
			case 2:
				grade = "Comète player"
			case 3:
				grade = "Joueur planétaire"
			default:
				grade = "Starfire player"
			}
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS,Data: gin.H{"address":info.Address,"recommender":pidInfo.Address,"grade":grade,"teamnumber":info.TeamNumber+1,"teampledge":float32(info.TotalTeamPledge)/100000000,"recommenders":info.Recommenders,"comet":info.Comet,"planet":info.Planet,"fixedstart":info.Fixstar,"todayreward":float32(info.TodayReward)/100000000}})
			return
		}
	default:

		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PARAMS_WRONG})
		}
		info, exists := service.GetUserInfo(address)
		pidInfo := sql.GetUserInfoByUid(info.Pid)
		if !exists {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_WRONG_ADDRESS})
			return
		}else {
			var grade string
			switch info.Grade {
			case 4:
				grade = "Fixedstar player"
			case 2:
				grade = "Comet player"
			case 3:
				grade = "Planet player"
			default:
				grade = "Spark player"
			}
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_SUCCESS,Data: gin.H{"address":info.Address,"recommender":pidInfo.Address,"grade":grade,"teamnumber":info.TeamNumber+1,"teampledge":float32(info.TotalTeamPledge)/100000000,"recommenders":info.Recommenders,"comet":info.Comet,"planet":info.Planet,"fixedstart":info.Fixstar,"todayreward":float32(info.TodayReward)/100000000 }})
			return
		}

	}




}

func simpleInfo(context *gin.Context)  {
	header := context.GetHeader("accept-language")
	address := context.Query("address")
	switch header {
	case "zh-CN":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		info, b := service.GetUserInfo(address)
		if !b {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_WRONG_ADDRESS,Data: gin.H{"grade":"尚未成為星火玩家","cashable":0}})
			return
		}
		var grade string
		switch info.Grade {
		case 2:
			grade = "彗星玩家"
		case 3:
			grade = "行星玩家"
		case 4:
			grade = "恒星玩家"
		default:
			grade = "星火玩家"
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: gin.H{"grade":grade,"cashable":float32(info.Cashable)/100000000}})
		return
	case "fr-FR":
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG})
			return
		}
		info, b := service.GetUserInfo(address)
		if !b {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_WRONG_ADDRESS,Data: gin.H{"grade":"N'est pas encore devenu un joueur Spark","cashable":0}})
			return
		}
		var grade string
		switch info.Grade {
		case 2:
			grade = "Comète player"
		case 3:
			grade = "Joueur planétaire"
		case 4:
			grade = "Star player"
		default:
			grade = "Starfire player"
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS,Data: gin.H{"grade":grade,"cashable":float32(info.Cashable)/100000000}})
		return
	default:
		if len(address)==0 {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PARAMS_WRONG})
			return
		}
		info, b := service.GetUserInfo(address)
		if !b {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_WRONG_ADDRESS,Data: gin.H{"grade":"Not yet a spark player","cashable":0}})
			return
		}
		var grade string
		switch info.Grade {
		case 2:
			grade = "Comet player"
		case 3:
			grade = "Planet player"
		case 4:
			grade = "Fixed start"
		default:
			grade = "Spark player"
		}
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: gin.H{"grade":grade,"cashable":float32(info.Cashable)/100000000}})
		return
	}
}

type Recommender struct {
	Address string `json:"address"`
	Paddress string `json:"paddress"`
}

func getRecommender(context *gin.Context)  {
	var rd Recommender
	header := context.GetHeader("accept-language")
	err := context.BindJSON(&rd)
	switch header {
	case "zh-CN":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
		}
	}
}
