package service

import (
	"log"
	"sptbackend/config"
	"sptbackend/sql"
	"sptbackend/toolWithDrawRecords"
)
//lp取现
func LpWithDraw(address string) bool {
	lpamount ,_:= sql.GetLpamount(address)
	_,err  := Transfer(address, lpamount,config.Cfg.Ethereum.KeyStorePathLP)
	if err !=nil {
		log.Printf("address :%s lp with draw fail :%s",address,err)
		return false
	}
	sql.ResetLpZero(address)
	sql.InsertWithDraw(address,lpamount,2)
	return true
}

//质押取现

func PledgeWithDrwa(address string)bool  {
	info, exists := GetUserInfo(address)
	if !exists {
		return false
	}else {
		if info.Cashable <= 0 {
			return false
		}
		_, err:= Transfer(address, info.Cashable, config.Cfg.Ethereum.KeyStore)
		if err !=nil {
			log.Printf("address :%s lp with draw fail :%s",address,err)
			return false
		}
		sql.ResetPledgeToZero(address,info.Cashable)
		sql.InsertWithDraw(address,info.Cashable,1)
		return true
	}
}

func GetWtihDrawRecords(address string,language string) *toolWithDrawRecords.Array  {
	array := toolWithDrawRecords.Make(0, 5)
	draws := sql.GetWithDraw(address)
	if  len(draws)==0{
		return array
	}
	switch language {
	case "zh-CN":
		for i:=len(draws);i>0;i-- {
			draw := draws[i-1]
			var  co  toolWithDrawRecords.WithDrawRecord
			co.Address = draw.Address
			switch draw.Type {
			case 1:
				co.Typename = "質押提現"
			default:
				co.Typename = "lp收益提現"
			}
			co.Amount = float32(draw.Amount)/100000000
			co.CrateTime = draw.CrateTime.Format("2006-01-02 15:04:05")
			array.Append(co)
		}
		return array
	case "fr-FR":
		for i:=len(draws);i>0;i-- {
			draw := draws[i-1]
			var  co  toolWithDrawRecords.WithDrawRecord
			co.Address = draw.Address
			switch draw.Type {
			case 1:
				co.Typename = "Retrait du gage\n\n"
			default:
				co.Typename = "LP retrait des revenus"
			}
			co.Amount = float32(draw.Amount)/100000000
			co.CrateTime = draw.CrateTime.Format("2006-01-02 15:04:05")
			array.Append(co)
		}
		return array
	default:
		for i:=len(draws);i>0;i-- {
			draw := draws[i-1]
			var  co  toolWithDrawRecords.WithDrawRecord
			co.Address = draw.Address
			switch draw.Type {
			case 1:
				co.Typename = "Pledge withdrawal"
			default:
				co.Typename = "Withdrawal of LP income"
			}
			co.Amount = float32(draw.Amount)/100000000
			co.CrateTime = draw.CrateTime.Format("2006-01-02 15:04:05")
			array.Append(co)
		}
		return array
	}
}
