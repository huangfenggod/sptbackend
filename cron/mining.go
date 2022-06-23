package cron

import (
	"sptbackend/config"
	"sptbackend/service"
	"sptbackend/sql"
)
//每日更新总获得和算力值
func MiningCal()  {
	sql.UpdateMinerWithdraw(config.Config().Power)
	order := sql.GetOverMiningAddressOrder()
	for i:=0;i<len(order);i++ {
		sql.UpdatePowerTimeUpOder(order[i].Address,order[i].Power)
	}
	lpinfo := sql.GetMinerLPInfo()
	for j:=0;j<len(lpinfo);j++ {
		address := lpinfo[j].Address
		lp, _ := service.GetBnbSptLp(address)
		if  float32(lp/service.Lpvalue) !=lpinfo[j].LpPower {
			sql.UpdateminerBYLpEveryTime(address,float32(lp/service.Lpvalue))
		}
	}
}
