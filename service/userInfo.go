package service

import "sptbackend/sql"

func GetUserInfo(address string) (sql.User,bool)  {
	var user sql.User
	uid := sql.ExitsAddressUid(address)
	if uid >0 {
		info := sql.GetUserInfoByUid(uid)
		return info,true
	}else {
		return user,false
	}

}

