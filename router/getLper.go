package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sptbackend/language"
	"sptbackend/myrsa"
	"sptbackend/service"
	"sptbackend/sql"
)

type param struct {
	Address string `json:"address"`
	Password string `json:"password"`
}
func getLper(context *gin.Context)  {
	header := context.GetHeader("accept-language")
	var para param
	err := context.BindJSON(&para)
	switch header {
	case "zh-CN":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		//nosql.SetInAddressForLp(para.Address)
		sql.InsertLpUser(para.Address)
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "lp質押收益領取成功，收益每天定時發放"})
	case "fr-FR":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG} )
			return
		}

		sql.InsertLpUser(para.Address)
		//nosql.SetInAddressForLp(para.Address)
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "Les gains lp plaid sont collectés avec succès. Les gains sont distribués périodiquement tous les jours"})
	default:
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PARAMS_WRONG})
			return
		}
		//nosql.SetInAddressForLp(para.Address)
		sql.InsertLpUser(para.Address)
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "Lp pledge return get success, return every time"})
	}

}

func getlpstatus(context *gin.Context)  {
	header := context.GetHeader("accept-language")
	var para param
	err := context.BindJSON(&para)
	switch header {
	case "zh-CN":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		lpamount ,yestoday := sql.GetLpamount(para.Address)
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS,Data: gin.H{"lpamount":float32(lpamount)/100000000,"lpyesterday":float32(yestoday)/100000000}})
	case "fr-FR":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG} )
			return
		}
		lpamount ,yestoday := sql.GetLpamount(para.Address)
		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS,Data: gin.H{"lpamount":float32(lpamount)/100000000,"lpyesterday":float32(yestoday)/100000000}})
	default:
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PARAMS_WRONG})
			return
		}
		lpamount ,yestoday := sql.GetLpamount(para.Address)

		context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_SUCCESS,Data: gin.H{"lpamount":float32(lpamount)/100000000,"lpyesterday":float32(yestoday)/100000000}})
	}
}

func lpwithdraw(context *gin.Context){

	header := context.GetHeader("accept-language")
	var para param
	err := context.BindJSON(&para)
	switch header {
	case "zh-CN":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}
		if len(para.Address)==0 ||len(para.Password)<22{
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PARAMS_WRONG})
			return
		}

		//hash, err1 := crypto.PasswordDecode(para.Password)
		hash, err1:= myrsa.RSADecode(para.Password)
		if err1!=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PASSWORD_WRONG})
			return
		}
		transactionHash := service.IdentifyTransactionHash(hash)
		if !transactionHash {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.CHINESE_PASSWORD_WRONG})
			return
		}
		draw := service.LpWithDraw(para.Address)
		if draw {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_SUCCESS})
		}else {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.CHINESE_FAIL})
		}
	case "fr-FR":
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG} )
			return
		}
		if len(para.Address)==0 ||len(para.Password)<22{
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PARANS_WRONG})
			return
		}
		hash, err1:= myrsa.RSADecode(para.Password)
		if err1!=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PASSWORD_WRONG})
			return
		}
		transactionHash := service.IdentifyTransactionHash(hash)
		if !transactionHash {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.FRENCH_PASSWORD_WRONG})
			return
		}
		draw := service.LpWithDraw(para.Address)
		if draw {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_SUCCESS})
		}else {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.FRENCH_FAIL})
		}
	default:
		if err !=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PARAMS_WRONG})
			return
		}
		if len(para.Address)==0 ||len(para.Password)<22{
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PASSWORD_WRONG})
			return
		}
		hash, err1:= myrsa.RSADecode(para.Password)
		if err1!=nil {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PASSWORD_WRONG})
			return
		}
		transactionHash := service.IdentifyTransactionHash(hash)
		if !transactionHash {
			context.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language.ENGLISH_PASSWORD_WRONG})
			return
		}
		draw := service.LpWithDraw(para.Address)
		if draw {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_SUCCESS})
		}else {
			context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language.ENGLISH_FAIL})
		}
	}


}

