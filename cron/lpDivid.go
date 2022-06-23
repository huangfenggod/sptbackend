package cron

import (
	"log"
	"sptbackend/service"
	"sptbackend/sql"
	"sptbackend/tool"
)



func LpDivid()  {
	log.Printf("begin lpdivd")
	//这里把数据库的地址导入进来
//这里实现把持久化数据数据导入到redis中然后redis进行晒选处理
	yesterday := sql.GetLpYesterday()
	//for i:=0;i<len(lpaddress);i++ {
	//	nosql.SetInAddressForLp(lpaddress[i].Address)
	//}

	//lp := nosql.GetAddressForLp()
	//if len(lp)==0 {
	//	return
	//}
	//今日之前总量

	livid, _ := service.GetLivid()

	distrbute := livid - yesterday //今日可分配数量
	log.Printf("today distribute :%d",distrbute)
	lpAllAddress := sql.GetLpAddress()
	//检查是否是会员
	lpaddress := tool.Make(0,10)
	for j:=0;j<len(lpAllAddress);j++ {
		uid := sql.ExitsAddressUid(lpAllAddress[j].Address)
		if uid !=0 {
			address :=tool.LpAddressAmount{Address: lpAllAddress[j].Address}
			lpaddress.Append(address)
		}
	}
	lpaa := tool.Make(0, 10)
	sum :=int64(0)
	for i :=0;i<lpaddress.Len();i++ {
		usdtLp, _ := service.GetAccountLpBalanceOf(lpaddress.Get(i).Address)
		bnbLp ,_:=	service.GetBnbSptLp(lpaddress.Get(i).Address)
		lpTotal := usdtLp + bnbLp
		lpaddress1 :=  tool.LpAddressAmount{Amount: lpTotal,Address: lpaddress.Get(i).Address}
		lpaa.Append(lpaddress1)
		sum += lpTotal
	}

	if distrbute<=0||sum<=0{
		log.Println("lp divid over")
		return
	}
	for i :=0;i<lpaa.Len();i++ {
		if lpaa.Get(i).Amount<=0{
			continue
		}
		sql.UpdateLpUser(lpaa.Get(i).Address,(lpaa.Get(i).Amount/1000000)*distrbute/(sum/1000000))
	}
	//nosql.RemoveLp()
	sql.UpdateDivid(livid)
	log.Println("lp divid over")
}

