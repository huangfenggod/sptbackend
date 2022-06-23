package router

func ApiGroup()  {
	group := GIN.Group("v1/api")
//登记领取lp奖励
	group.POST("/getlpreward", getLper)//
	//获取lp质押
	group.POST("/getlpstatus",getlpstatus)//
	//lp提现
	group.POST("/lpwithdraw",lpwithdraw)//
//获取通知
	group.GET("/notice",notice) //
//发布通知
	group.POST("/announce",announce) //
//关闭通知
	group.POST("/stopc",stopNotice) //
//参与质押
	group.POST("/pledge",pledge) //
	group.POST("/pledgefor",pledgefor)
	//查询当前地址是否为新用户，推荐人地址是否正确
	//group.GET("/getrecommender",getRecommender)
	group.GET("/isexists",isexists)

//质押利率返回
	group.GET("/pledgerate",pledgerate) //
	//获取用户社区推广相关信息
	group.GET("/userinfo",userInfo) //
//质押提现
	group.POST("/pledgewithdraw",pledgeWithDraw) //

//可提现和玩家等级
	group.GET("/simpleinfo",simpleInfo) //
	//订单的列表
	group.GET("/getorder",getorder)
	group.GET("/withdrawrecords",withDrawRecords)
	//group.GET("/zhiya",zhiya)
	//group.GET("/fenhong",fenhong)
	group.POST("/ming",mining)
	group.POST("/bound",bound)
	group.POST("/lpming",lpMing)
	group.GET("/getminginfo",getMinerInfo)
	group.POST("/mingwithdraw",mingWithDraw)

}
