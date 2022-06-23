package sql

import "time"

func InsertWithDraw(address string,amount int64,typ int)  {
	DB.Exec("insert into withdraw(address,amount,type,create_time) values(?,?,?,?) ",address,amount,typ,time.Now())
}

type WithDrawOrder struct {
	Address string `db:"address"`
	Amount int64 `db:"amount"`
	Type  int `db:"type"`
	CrateTime time.Time `json:"crate_time"`
}
func GetWithDraw(address string)[]WithDrawOrder  {
	var withdraw []WithDrawOrder
	DB.Raw("select address, amount , type,create_time from withdraw where address = ? and exists(select 1 from withdraw where address =? limit 1)",address,address).Scan(&withdraw)
	return withdraw
}

func GetWithDrawType(address string ,typ int)[]WithDrawOrder  {
	var withdraw []WithDrawOrder
	DB.Raw("select address, amount , type,create_time from withdraw where address = ? and type =? and exists(select 1 from withdraw where address =? and type =? limit 1)",address,typ,address,typ).Scan(&withdraw)
	return withdraw
}


