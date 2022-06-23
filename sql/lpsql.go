package sql

import (
	"log"
)

type lpTable struct {
	Lid int `db:"lid"`
	Address string `db:"address"`
	Amount int64 `db:"amount"`
	Yestoday int64 `db:"yestoday"`
}

type lprate struct {
	Amount int64 `db:"amount"`

}

func GetLpYesterday() int64 {
	var lp lprate
	DB.Raw("select amount from divid where id =1").Scan(&lp)
	return lp.Amount
}



type LpAddress struct {
	Address string `db:"address"`
}
func GetLpAddress() []LpAddress  {
	var address []LpAddress
	DB.Raw("select address from lp").Scan(&address)
	return address
}

func InsertLpUser(address string)bool  {
	DB.Exec("insert ignore into lp(address,amount) values(?,0)",address)
return true

}



//更新每日分红账户
func UpdateDivid(amount int64)  {
	exec := DB.Exec("update divid set amount = ? where id =1", amount)
	if exec.Error !=nil  {
		log.Println(exec.Error)
	}
}

//lpuser存在就修改不存在就插入操作
func UpdateLpUser(address string,amount int64)  {
	//DB.Exec(" insert into lp(address,amount) values(?,?) ON DUPLICATE KEY UPDATE amount =amount+ ?",address,amount,amount)
	DB.Debug().Exec("update lp set amount = amount + ? ,yestoday = ? where address =?",amount,amount,address)
}

//lp分红用户提现数据修改
func UpdateByWithdraw(address string)  {
	DB.Exec("update lp set amount = 0 where address = ? and exists(select 1 from lp where address =? limit 1) ",address,address)
}
//lp分红有多少
func GetLpamount(address string)(int64,int64)  {
	var lptable lpTable
	 DB.Raw("select amount ,yestoday from lp where address =?", address).Scan(&lptable)
	return lptable.Amount,lptable.Yestoday
}
//把lp账户可提现设置为0
func ResetLpZero(address string)  {
	DB.Exec("update lp set amount = 0 where address =?",address)
}
