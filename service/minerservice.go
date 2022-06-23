package service

import (
	"errors"
	"sptbackend/sql"
)

func Mining(txhash string,address string,power float32,sptAmount int){
	sql.InsertMiningOrder(address,power,sptAmount,txhash)
	sql.UpdateMinerByBuy(address,power)
	UpdateMinerTwo(address,power)
}
//改变推荐人上2级算力
func UpdateMinerTwo(address string,power float32)  {
	sql.FindBindAddress(address,power)
}

func MinerBind(address string,paddress string) bool {
	b := sql.GetMinerPaddress(address)
	if !b {
		sql.MinerBind(address,paddress)
		return true
	}
	return false
}

func GetMinerInfo(address string) sql.MinerInfo{
	info := sql.GetMinerInfo(address)
	return info
}

const Lpvalue = int64(600000000)

func MinerLp(address string) (bool,error) {
	lp, err := GetBnbSptLp(address)
	if err!=nil {
		return false,errors.New("bsn chain wrong try again")
	}
	lpPower := lp/Lpvalue
	if lpPower<1 {
		return false,errors.New("lp value not enough")
	}
	sql.UpdateMinerByLp(address,float32(lpPower))
	return true,nil
}

func WithdrawEth(address string)bool {
	eth := sql.WithdrawEth(address)
	return eth
}
