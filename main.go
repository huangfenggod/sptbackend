package main

import (
	"fmt"
	"sptbackend/config"
	"sptbackend/cron"
	"sptbackend/log"
	"sptbackend/myrsa"
	"sptbackend/nosql"
	"sptbackend/router"
	"sptbackend/service"
	"sptbackend/sql"
)

//优化查找团队人数的地方
//一段使用cte递归查询伞下总业绩的sql
//with recursive  cte as(select uid,address,pid from user where pid = 8  union all select u.uid ,u.address,u.pid from cte c inner join user u on c.uid = u.pid)select uid,address,pid from cte order by uid;
func main() {
	config.Config()
	log.InitLog()
	fmt.Println("system start:")
	err := myrsa.InitRsa()
	if err !=nil {
		fmt.Println("private key wrong")
		return
	}
	service.InitDial()
	nosql.InitRedis()
	sql.InitDatabase()
	router.InitGin()
	cron.InitCron()
	router.ApiGroup()

	router.GIN.Run(":8009")
	defer nosql.Cli.Close()
	defer service.Conn.Close()
	defer sql.DB.Close()
}
