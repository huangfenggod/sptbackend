package sql

import (
	"log"
	"time"
)

//插入订单
func InsertOrder(address string,duration int,amount int64,) {
	var rate float32
	switch duration {
	case 360:
		rate = 0.01388888889
	case 180:
		rate = 0.01277777778
	default:
		rate = 0.01111111111
	}
	log.Printf("insert order: %s %d %d", address, duration, amount)
	DB.Exec("insert into oder(address,duration,effective,amount,rate,create_time) values(?,?,?,?,?,?)",address,duration,duration,amount,rate,time.Now())
}
func InsertOrder1(address string,duration int,amount int64,times time.Time)  {
	var rate float32
	switch duration {
	case 360:
		rate =0.01388888889
	case 180:
		rate =0.01277777778
	default:
		rate =0.01111111111
	}
	log.Printf("insert order: %s %d %d",address,duration,amount)
	//DB.Exec("insert into oder(address,duration,effective,amount,rate,create_time) values(?,?,?,?,?,?)",address,duration,duration,amount,rate,time.Now())
	DB.Exec("insert into oder(address,duration,effective,amount,rate,create_time) values(?,?,?,?,?,?)",address,duration,duration,amount,rate,times)

}

//获取目前的质押总量
type result struct {
	Total int64 `db:"total"`
}
func GetTotalOfPledgeNow() int64 {
	var res result
	DB.Raw("select sum(amount) as total from oder where effective >0").Scan(&res)
	return res.Total
}

//获取某个地址的所有订单
func GetAllOrdersByAddress(address string)[]Order  {
	var orders []Order
	DB.Raw("select address,duration,effective,amount,gettoday,create_time from oder where address =?",address).Scan(&orders)
	return orders
}


