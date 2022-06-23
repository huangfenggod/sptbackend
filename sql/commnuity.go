package sql

type mannualRate struct {
	Id int `db:"id"`
	Amount int `db:"amount"`
	Lock int `db:"lock"`
}
//查看手动比例
func GetRateMannual() float32 {
	var manRate mannualRate
	DB.Raw("select * from divid where id =2 ").Scan(&manRate)
	if manRate.Lock==1 {
		return 1
	}else {
		return float32(manRate.Amount)/100
	}
}

//所有order里面effective=1的把本金放到入到cashable里面去
func OverOrderToCashable()  {
	DB.Exec("update user u,(select address,sum(amount) s from oder where effective =1 group by address) t set u.cashable = cashable+ t.s where u.address =t.address")
}

//更新每日effective大于0的，每日订单分红
func UpdateOrderEveryday(rate float32)  {
	DB.Exec("update oder set effective = effective - 1 , gettoday = amount * rate *? ,distribution = distribution + gettoday where effective >0 ",rate)
}

//查询如果effective =1 整个团队质押总量分别减少，自身质押总量减少
func UpdateTotalPledgeOverOder()  {
	DB.Exec("update user u ,(select sum(amount) s ,address from oder  where effective =1 group by address) t set u.total_pledge = u.total_pledge - t.s where u.address = t.address")
}
//查询出订单effective=1的所有订单分别计算

type AddressAmount struct {
	Address string `db:"address"`
	Amount int64 `db:"amount"`
}
//获取到期时间需要减少的质押总量
func GetOrderEffectiveOver() []AddressAmount {
	var addressamouts []AddressAmount
	DB.Raw("select address,sum(amount) amount from oder where effective =1 group by address").Scan(&addressamouts)
	return addressamouts
}

//更新user表的每日质押获得,从订单表更新到用户表数据总和
func UpdateUserPledgeEveryday()  {
	DB.Exec("update user u , ( SELECT address ,sum(gettoday) s FROM oder  GROUP BY address) t set u.today_pget =t.s , u.cashable = u.cashable + t.s where t.address =u.address")
}


//查询没有子代的数据行
//func GetNoRecommender()[]CUser {
//	var cusers []CUser
//
//}






