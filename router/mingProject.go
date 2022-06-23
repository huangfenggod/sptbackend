package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	language2 "sptbackend/language"
	"sptbackend/myrsa"
	"sptbackend/service"
)

type miner struct {
	Address string `json:"address"`
	Amount int `json:"amount"`
	Password string `json:"password"`
}
func mining(c *gin.Context)  {
	language := c.GetHeader("accept-language")
	var m miner
	err := c.BindJSON(&m)
	address := m.Address
	password := m.Password
	amount := m.Amount
	power :=float32(0)
	switch language {
	case "zh-CN":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.CHINESE_PARAMS_WRONG})
			return
		}
		if len(password)<22||amount<400||len(address)==0 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.CHINESE_PARAMS_WRONG})
			return
		}
		hash, err1:= myrsa.RSADecode(password)
		if err1!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.CHINESE_PASSWORD_WRONG})
			return
		}
		switch amount {
		case 2000:
			power=55
		case 1000:
			power=26
		default:
			power=10
		}
		service.Mining(hash,address,power,amount)
		log.Printf("insert mining address:%s , hash :%s , amount :%d , power :%d",address,hash,amount,int(power))
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.CHINESE_SUCCESS})

	case "fr-FR":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.FRENCH_PARANS_WRONG})
			return
		}
		if len(password)<22||amount<400||len(address)==0 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.FRENCH_PARANS_WRONG})
			return
		}
		hash, err1:= myrsa.RSADecode(password)
		if err1!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.FRENCH_PASSWORD_WRONG})
			return
		}
		switch amount {
		case 2000:
			power=55
		case 1000:
			power=26
		default:
			power=10
		}
		service.Mining(hash,address,power,amount)
		log.Printf("insert mining address:%s , hash :%s , amount :%d , power :%d",address,hash,amount,int(power))
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.FRENCH_SUCCESS})

	default:
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_PARAMS_WRONG})
			return
		}
		if len(password)<22||amount<400||len(address)==0 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.ENGLISH_PARAMS_WRONG})
			return
		}
		hash, err1:= myrsa.RSADecode(password)
		if err1!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.ENGLISH_PASSWORD_WRONG})
			return
		}
		switch amount {
		case 2000:
			power=55
		case 1000:
			power=26
		default:
			power=10
		}
		service.Mining(hash,address,power,amount)
		log.Printf("insert mining address:%s , hash :%s , amount :%d , power :%d",address,hash,amount,int(power))
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_SUCCESS})
	}
}

type relation struct {
	Address string `json:"address"`
	Paddress string `json:"paddress"`
}

func bound(c *gin.Context){
	language := c.GetHeader("accept-language")
	var re relation
	err := c.BindJSON(&re)
	paddress := re.Paddress
	address := re.Address
	switch language {
	case "zh-CN":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.CHINESE_PARAMS_WRONG})
			return
		}
		if len(address)==0||len(paddress)==0{
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.CHINESE_PARAMS_WRONG})
			return
		}
		bind := service.MinerBind(address, paddress)
		if !bind {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "推薦人已綁定"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.CHINESE_SUCCESS})
	case "fr-FR":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.FRENCH_PARANS_WRONG})
			return
		}
		if len(address)==0||len(paddress)==0{
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.FRENCH_PARANS_WRONG})
			return
		}
		bind := service.MinerBind(address, paddress)
		if !bind {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Recommandation liée"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.FRENCH_SUCCESS})

	default:
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.ENGLISH_PARAMS_WRONG})
			return
		}
		if len(address)==0||len(paddress)==0{
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: language2.ENGLISH_PARAMS_WRONG})
			return
		}
		bind := service.MinerBind(address, paddress)
		if !bind {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Recommender bound"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_SUCCESS})
	}
}

func getMinerInfo(c *gin.Context)  {
	address := c.Query("address")
	info := service.GetMinerInfo(address)
	if len(info.Address)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "miner not exists"})
		return
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_SUCCESS,Data: gin.H{"amount":info.Withdraw,"address":info.Address,"recommender":info.Paddress,"buyPower":info.BuyPower,"rewardPower":info.RewardPower,"lpPower":info.LpPower}})
}

func lpMing(c *gin.Context)  {
	language := c.GetHeader("accept-language")
	var re relation
	err := c.BindJSON(&re)
	switch language {
	case "zh-CN":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.CHINESE_PARAMS_WRONG})
			return
		}
		address := re.Address
		lp, _ := service.MinerLp(address)
		if !lp {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "添加流動性lp值不足"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.CHINESE_SUCCESS})

	case "fr-FR":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.FRENCH_PARANS_WRONG})
			return
		}
		address := re.Address
		lp, _ := service.MinerLp(address)
		if !lp {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Ajouter une valeur LP de liquidité insuffisante"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.FRENCH_SUCCESS})

	default:
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_PARAMS_WRONG})
			return
		}
		address := re.Address
		lp, _ := service.MinerLp(address)
		if !lp {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Insufficient liquidity LP value added"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_SUCCESS})
	}
}
func mingWithDraw(c *gin.Context)  {
	language := c.GetHeader("accept-language")
	var re relation
	err := c.BindJSON(&re)
	switch language {
	case "zh-CN":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.CHINESE_PARAMS_WRONG})
			return
		}
		address := re.Address
		info := service.GetMinerInfo(address)
		if len(info.Address)==0 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "此帳戶沒有礦機算力"})
			return
		}
		if info.Withdraw<0.03 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "可提不足0.03eth"})
			return
		}
		eth := service.WithdrawEth(address)
		if !eth {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "可提不足0.03eth"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "提現成功等待撥幣"})
	case "fr-FR":
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.FRENCH_PARANS_WRONG})
			return
		}
		address := re.Address
		info := service.GetMinerInfo(address)
		if len(info.Address)==0 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Ce compte n'a pas de puissance de calcul"})
			return
		}
		if info.Withdraw<0.03 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Moins de 0,03 ETH"})
			return
		}
		eth := service.WithdrawEth(address)
		if !eth {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Moins de 0,03 ETH"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "Retrait réussi en attendant la composition de la monnaie"})
	default:
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: language2.ENGLISH_PARAMS_WRONG})
			return
		}
		address := re.Address
		info := service.GetMinerInfo(address)
		if len(info.Address)==0 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "This account has no mining machine capacity"})
			return
		}
		if info.Withdraw<0.03 {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Less than 0.03eth"})
			return
		}
		eth := service.WithdrawEth(address)
		if !eth {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "Less than 0.03eth"})
			return
		}
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "Withdrawal succeeded, waiting for currency allocation"})
	}
}
