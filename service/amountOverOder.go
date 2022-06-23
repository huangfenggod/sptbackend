package service

import "sptbackend/sql"

func OverOrder()  {
	over := sql.GetOrderEffectiveOver()
	for i:=0;i<len(over);i++ {

		addressAmount := over[i]
		sql.DB.Exec("update user set total_team_pledge = total_team_pledge - ? where address = ?",addressAmount.Amount,addressAmount.Address)
		var uid int
		 uid = sql.ExitsAddressUid(addressAmount.Address)
		for ;; {
			userInfo := sql.GetUserInfoByUid(uid)

			if userInfo.Pid==0{
				break
			}
			sql.DB.Debug().Exec("update user set total_team_pledge = total_team_pledge - ? where uid =?",addressAmount.Amount,userInfo.Pid)
			uid = userInfo.Pid
		}
	}
}
