package cron

import (
	"fmt"
	"log"
	"sptbackend/nosql"
	"sptbackend/service"
	"sptbackend/sql"
)

func TodayRewardZero()  {
	sql.UpdateUserTodaygetZero()
}



func CommunityDividend() {
	//每日分配比例

	log.Println("start divid")
	var rate float32
	manRate := sql.GetRateMannual()
	if manRate != 1 {
		rate = calculateRate(sql.GetTotalOfPledgeNow())
	} else {
		rate = manRate
	}
	//把这个rate值存储在redis里面，随时可传给前端使用
	fmt.Println(111)
	nosql.RemoveRateRedis()
	nosql.SetRateIntoRedis(rate)

	//订单每日更新
	sql.OverOrderToCashable()//把即结束的本金传入可提现
	service.OverOrder()

	sql.UpdateTotalPledgeOverOder()//把质押总量减少本金
	sql.UpdateOrderEveryday(rate)//
	sql.UpdateUserPledgeEveryday()//更新每日质押获得
//分别计算团队质押量
	CommuityDivid()


}
func calculateRate(pledgeNumber int64)float32  {
	rate := 1-float32(pledgeNumber/50000)*0.05
	if rate< 0.3{
		rate =0.3
	}
	return rate
}
