package cron

import (
	"github.com/robfig/cron/v3"
)

func InitCron()  {
	crontab := cron.New(cron.WithSeconds())
	//点差分红
	spec :="0 0 21 * * ?"
	crontab.AddFunc(spec,LpDivid)
	spec1 :="0 0 3  * * ?"
	crontab.AddFunc(spec1,CommunityDividend)
	sepc2 := "0 0 1 * * ?"
	crontab.AddFunc(sepc2,TodayRewardZero)
	spec3 :="0 0 0,8,16 * * ?"
	crontab.AddFunc(spec3,MiningCal)
	crontab.Start()
}


