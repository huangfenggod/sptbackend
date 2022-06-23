package router

//
//func zhiya(context *gin.Context) {
//	address := context.Query("address")
//	paddress := context.Query("paddress")
//	time1 := context.Query("time")
//	location, err3 := time.ParseInLocation("2006-01-02 15:04:05", time1, time.Local)
//	if err3 !=nil {
//		context.JSON(http.StatusOK, ResponseUtil{Status: false,Msg: "fase"})
//		return
//	}
//
//	amount := context.Query("amount")
//	days := context.Query("days")
//	amount1, err2 := strconv.ParseFloat(amount, 32)
//	if err2 != nil {
//		context.JSON(http.StatusOK, ResponseUtil{Status: false,Msg: "fase"})
//		return
//	}
//	dayss, err := strconv.ParseInt(days, 10, 64)
//	if err != nil {
//		context.JSON(http.StatusOK, ResponseUtil{Status: false,Msg: "fase"})
//		return
//	}
//	uid := sql.ExitsAddressUid(address)
//	fmt.Println(uid)
//	if uid > 0 {
//		sql.InsertOrder1(address, int(dayss), int64(amount1*100000000),location)
//		//新订单改变用户状态
//		userInfo := sql.GetUserInfoByUid(uid)
//		service.UpdateUserByNewOrder(uid, userInfo.Pid, int64(amount1*100000000))
//		context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: language.CHINESE_SUCCESS})
//		return
//	} else
//	{
//		//如果该用户不存在创建新用户
//		if len(paddress) == 0 {
//			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_NO_RECOMMENDER})
//			return
//		}
//		pid := sql.ExitsAddressUid(paddress)
//		if pid == 0 {
//			context.JSON(http.StatusOK, ResponseUtil{Status: false, Msg: language.CHINESE_NO_RECOMMENDER})
//			return
//		}
//		//创建新用户
//		sql.CreateUser(address, pid)
//		newUid := sql.ExitsAddressUid(address)
//		//插入订单
//		sql.InsertOrder1(address, int(dayss), int64(amount1*100000000),location)
//		//新订单改变用户状态
//		service.UpdateUserByNewOrder(newUid, pid, int64(amount1*100000000))
//		context.JSON(http.StatusOK, ResponseUtil{Status: true, Msg: language.CHINESE_SUCCESS})
//		return
//	}
//}
//
//func fenhong(context *gin.Context)  {
//cron.CommunityDividend()
//context.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success"})
//}
