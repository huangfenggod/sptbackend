package service

import (
	"sptbackend/sql"
	"sptbackend/toolorder"
	"strconv"
)
//pid必须不为0
func UpdateUserByNewOrder(uid int,pid int,amount int64)  {
	//改变自身状态
	sql.UpdateUserSelf(uid,amount)
	//改变当前推荐人状态
	if pid==0 {
		return
	}
	isnew := sql.UpdateFirstRecommender(uid, pid, amount)
	userPid := sql.GetUserInfoByUid(pid)
//改变整个团队状态
	sql.UpdateAboveAll(userPid,amount,isnew)
}




func GetAllOrderByAddress(address string ,language string) *toolorder.Array  {
	array := toolorder.Make(0, 5)
	orders := sql.GetAllOrdersByAddress(address)
	if len(orders) ==0 {
	return array
	}
	switch language {
	case "zh-CN":
		for i:=len(orders);i>0;i-- {
			order := orders[i-1]
			var co toolorder.ChineseOrder
			co.Address = order.Address
			co.Name =strconv.Itoa(order.Duration)+"天質押"
			co.Effective = order.Effective
			co.Amount = float32(order.Amount)/100000000
			co.Gettoday = float32(order.Gettoday)/100000000
			co.CreateTime = order.CreateTime.Format("2006-01-02 15:04:05")
			array.Append(co)
		}
		return array
	case "fr-FR":
		for i:=len(orders);i>0;i--{
			order := orders[i-1]
			var co toolorder.ChineseOrder
			co.Address = order.Address
			co.Name = strconv.Itoa(order.Duration)+"Gage céleste"
			co.Effective = order.Effective
			co.Amount = float32(order.Amount)/100000000
			co.Gettoday = float32(order.Gettoday)/100000000
			co.CreateTime = order.CreateTime.Format("2006-01-02 15:04:05")

			array.Append(co)
		}
		return array
	default:
		for i:=len(orders);i>0;i--{
			order := orders[i-1]
			var co toolorder.ChineseOrder
			co.Address = order.Address
			co.Name = strconv.Itoa(order.Duration)+"days pledge"
			co.Effective = order.Effective
			co.Amount = float32(order.Amount)/100000000
			co.Gettoday = float32(order.Gettoday)/100000000
			co.CreateTime = order.CreateTime.Format("2006-01-02 15:04:05")
			array.Append(co)
		}
		return array
	}
}

