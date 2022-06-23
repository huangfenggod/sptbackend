package sql

import (
	"errors"
)

type User struct {
	Uid int `db:"uid"`
	Address string `db:"address"`
	Pid int `db:"pid"`
	Cashable int64 `db:"cashable"`
	TotalPledge int64 `db:"total_pledge"`
	TodayPget int64 `db:"today_pget"`
	TodayReward int64 `db:"today_reward"`
	Grade int `db:"grade"`
	TeamNumber int `db:"team_number"`
	Comet int `db:"comet"`
	Planet int `db:"planet"`
	Fixstar int `db:"fixstar"`
	Outed int64 `db:"outed"`
	TotalReward int64 `db:"total_reward"`
	Recommenders int64 `db:"recommenders"`
	TotalTeamPledge int64 `db:"total_team_pledge"`
	Upgrade int `db:"upgrade"`
}
type CUser struct {
	Uid int `db:"uid"`
	Pid int `db:"pid"`
	TotalPledge int64 `db:"totalpledge"`
	Grade int `db:"grade"`
	TotalTeamPledge int64 `db:"total_team_pledge"`
	TodayReward int64 `db:"today_reward"`
}

func GetCUser() []CUser {
	var cuser []CUser
	DB.Raw("select uid , pid , total_pledge, grade,total_team_pledge,today_reward from user").Scan(&cuser)
	return cuser
}



func ExitsAddressUid(address string) int  {
	var  user User
	DB.Raw("select uid from user where address =? and exists(select 1 from user where address =? limit 1)",address,address).Scan(&user)
	return user.Uid

}
//查询所有用户信息
func GetUserInfoByUid(uid int) User {
	var user User
	DB.Raw("select * from user where uid =?",uid).Scan(&user)
	return user
}

//把所有今日奖励置为0
func UpdateUserTodaygetZero(){
	DB.Exec("update user set today_reward =0")
}


//改变质押可提现部分为0
func ResetPledgeToZero(address string,amount int64)  {
	DB.Debug().Exec("update user set cashable = 0 ,outed = outed+ ? where address =?",amount,address)
}


func CreateUser(address string, pid int )  {
	DB.Exec("insert into user(address,pid,isnew) values(?,?,0)",address,pid)
}

func UpdateUserSelf(uid int,amount int64)  {
	DB.Exec("update user set total_pledge=total_pledge+? , total_team_pledge =total_team_pledge + ? where uid =?",amount,amount,uid)
}

//查询uid是否存在
type xuid struct {
	Uid string `db:"uid"`
}

//查询当前uid账户的isnew如果为0，添加订单改变为1，老用户则不用


type isn struct{
	Isnew int `db:"isnew"`
}

func isnewAccount(uid int)bool  {
	var is []isn
	DB.Raw("select isnew from user where uid =?",uid).Scan(&is)
	if len(is)>0 &&is[0].Isnew==0 {
		return true
	}
	return false
}


//插入订单改变第一推荐人状态
func UpdateFirstRecommender(uid int,pid int,amount int64) bool {
	isnew := isnewAccount(uid)
	if isnew {
		DB.Exec("update user set cashable = cashable + ?,today_reward=today_reward+?,total_reward=total_reward+?,team_number =team_number +1, recommenders=recommenders+1,total_team_pledge=total_team_pledge+? where uid =?",amount/10,amount/10,amount/10,amount,pid)
		//DB.Exec("update user set total_reward=total_reward+?,team_number =team_number +1, recommenders=recommenders+1,total_team_pledge=total_team_pledge+? where uid =?",amount/10,amount,pid)
		DB.Debug().Exec("update user set isnew =1 where uid =?",uid)
	}else {
		DB.Exec("update user set cashable = cashable + ?,today_reward=today_reward+?,total_reward=total_reward+?,total_team_pledge=total_team_pledge+? where uid =?",amount/10,amount/10,amount/10,amount,pid)
		//DB.Exec("update user set total_reward=total_reward+?,total_team_pledge=total_team_pledge+? where uid =?",amount/10,amount,pid)
	}

	DB.Debug().Exec("update user set grade = 2,upgrade =1 where uid =? and recommenders >4 and grade =1 and total_team_pledge - total_pledge>=200000000000",pid)
	DB.Debug().Exec("update user set grade = 3,upgrade =1 where uid =? and grade =2 and comet >2 and total_team_pledge -total_pledge >1200000000000",pid)
	DB.Debug().Exec("update user set grade = 4,upgrade =1 where uid =? and grade =3 and planet>2 and total_team_pledge -total_pledge >5000000000000",pid)
	return isnew
}
//改变第一推荐人以上所有状态
func UpdateAboveAll(user User,amount int64,isnew bool)  {
	if user.Pid==0 {
		return
	}
	pid :=user.Pid

	for ;; {
		x := AboveAll(pid, amount,isnew)
		if  x==0 {
			break
		}
		pid =x
	}
}

type caluateUser struct{
	Co int `db:"co"`
	Pl int `db:"pl"`
	Fi int `db:"fi"`
}

type sumPledge struct {
	Recom int `db:"recom"`
	Total int64 `db:"total"`
}

func AboveAll(pid int ,amount int64,isnew bool ) int {
	if pid ==0 {
		return 0
	}
	var us []User
	var su []sumPledge
	var comet []caluateUser
	var planet []caluateUser
	var fixstar []caluateUser

	DB.Raw("select grade,pid from user where uid =?",pid).Scan(&us)
	DB.Raw("select count(uid) recom ,sum(total_team_pledge) total from user where pid =? ",pid).Scan(&su)
	DB.Raw("select count(uid) co from user where pid =? and grade =2",pid).Scan(&comet)
	DB.Raw("select count(uid) pl from user where  pid =? and grade =3",pid).Scan(&planet)
	DB.Raw("select count(uid) fi from user where pid =? and grade =4",pid).Scan(&fixstar)
	if isnew{
		switch us[0].Grade {
		case 1:
			if su[0].Recom>5&&su[0].Total>=200000000000 {
				DB.Debug().Exec("update user set team_number = team_number +1 ,total_team_pledge = total_team_pledge +?,grade =2 ,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}else {
				DB.Debug().Exec("update user set team_number =team_number +1 ,total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}
		case 2:
			if comet[0].Co>2&& su[0].Total>=1200000000000{
				DB.Debug().Exec("update user set team_number = team_number +1 ,total_team_pledge = total_team_pledge +?,grade =3 ,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}else {
				DB.Debug().Exec("update user set team_number =team_number +1 ,total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}
		case 3:
			if planet[0].Pl>2&&su[0].Total>=5000000000000 {
				DB.Debug().Exec("update user set team_number = team_number +1 ,total_team_pledge = total_team_pledge +?,grade =4 ,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}else {
				DB.Debug().Exec("update user set team_number =team_number +1 ,total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}
		default:
			DB.Debug().Exec("update user set team_number =team_number +1 ,total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
		}
	}else {
		switch us[0].Grade {
		case 1:
			if su[0].Recom>5&&su[0].Total>=200000000000 {
				DB.Debug().Exec("update user set total_team_pledge = total_team_pledge +?,grade =2 ,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}else {
				DB.Debug().Exec("update user set total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}
		case 2:
			if comet[0].Co>2&& su[0].Total>=1200000000000{
				DB.Debug().Exec("update user set  total_team_pledge = total_team_pledge +?,grade =3 ,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}else {
				DB.Debug().Exec("update user set  total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}
		case 3:
			if planet[0].Pl>2&&su[0].Total>=5000000000000 {
				DB.Debug().Exec("update user set total_team_pledge = total_team_pledge +?,grade =4 ,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}else {
				DB.Debug().Exec("update user set total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
			}
		default:
			DB.Debug().Exec("update user set total_team_pledge = total_team_pledge +?,comet =?,planet=?,fixstar=? where uid = ?",amount,comet[0].Co,planet[0].Pl,fixstar[0].Fi,pid)
		}
	}
	return us[0].Pid
}














func cuser(user User,amount int64,comet int,upgrade int) User{
	if comet ==1 {
	user.Comet = user.Comet+1
	}
	if upgrade !=0 {
		user.Upgrade=1
		user.Grade +=1
	}else {
		user.Upgrade=0
	}
	user.TeamNumber = user.TeamNumber+1
	user.TotalTeamPledge = user.TotalTeamPledge + amount
	return user
}

func cuserPlanet(user User,amount int64,planet int,upgrade int) User{
	if planet ==1 {
		user.Planet = user.Planet+1
	}
	if upgrade !=0 {
		user.Upgrade=1
		user.Grade +=1
	}else {
		user.Upgrade=0
	}
	user.TeamNumber = user.TeamNumber+1
	user.TotalTeamPledge = user.TotalTeamPledge + amount
	return user
}
func cuserFixstar(user User,amount int64,fixstar int,upgrade int) User {
	if fixstar ==1 {
		user.Fixstar = user.Fixstar+1
	}
	if upgrade !=0 {
		user.Upgrade=1
		user.Grade +=1
	}else {
		user.Upgrade=0
	}
	user.TeamNumber = user.TeamNumber+1
	user.TotalTeamPledge = user.TotalTeamPledge + amount
	return user
}


//更新当前状态根据自己团队以下的,并返回自己的pid
func updateRecommendSelf(sUser User,amount int64,isnew bool) (User,error) {
	var user User
	if sUser.Pid==0 {
		return user,errors.New("have no pid")
	}
	//DB.Exec("update user set team_number = team_number +1 ,total_team_pledge = total_pledge + ? where uid =?",amount,sUser.Pid)
	DB.Raw("select * from user where uid =?",sUser.Pid).Scan(&user)
	//传进来的是否为新账户如果为账户team_number+1 否不用
	if isnew {
		switch sUser.Upgrade {
		//0上级未升级的情况
		case 0:
			switch user.Grade {
			//当前等级为1
			case 1:
				if user.Recommenders>4 && user.TotalTeamPledge+amount-user.TotalPledge>200000000000 {
					DB.Exec("update user set grade =2 ,upgrade=1 ,team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",amount,user.Uid)
					return cuser(user,amount,1,1),nil
				}else {
					DB.Exec("update user set team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",amount,user.Uid)
					return cuser(user,amount,0,0),nil
				}
			case 2:
				if user.Comet>2 && user.TotalTeamPledge+amount-user.TotalPledge>1200000000000 {
					DB.Exec("update user set grade =3 ,upgrade=1 ,team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",amount,user.Uid)
					return cuserPlanet(user,amount,1,1),nil
				}else {
					DB.Exec("update user set team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",amount,user.Uid)
					return cuserPlanet(user,amount,0,0),nil
				}
			case 3:
				if user.Comet>2 && user.TotalTeamPledge+amount-user.TotalPledge>5000000000000 {
					DB.Exec("update user set grade =3 ,upgrade=1 ,team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",user.Uid)
					return cuserPlanet(user,amount,1,1),nil
				}else {
					DB.Exec("update user set team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",user.Uid)
					return cuserPlanet(user,amount,0,0),nil
				}
			default:
				DB.Exec("update user set team_number= team_number +1,total_team_pledge = total_team_pledge +? where  uid =? ",amount,user.Uid)
				return cuser(user,amount,0,0),nil
			}
		default:
			//上一级升级了的情况
			switch sUser.Grade {
			//升级为彗星
			case 2:
				//判断自己当前等级如果不为2，只添加团队人数+1和增加总质押+amount,且彗星+1
				if user.Grade==1 ||user.Grade==3 ||user.Grade==4{
					DB.Exec("update user set upgrade =0 ,total_team_pledge = total_team_pledge + ?,team_number = team_number+1 ,comet = comet +1 where uid =?",amount,sUser.Pid)
					return cuser(user,amount,1,0),nil
				}else {
					//如果为2就需要判断是否升级
					if user.Comet==2 {
						//小于2000spt
						if user.TotalTeamPledge -user.TotalPledge<1200000000000 {
							DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? ,team_number = team_number+1,comet = comet +1 where uid =?",amount,sUser.Pid)
							return cuser(user,amount,1,0),nil
						}else {
							//彗星推荐等于3了。团队总质押大于12000了可以升级了
							DB.Exec("update user set grade = 3 , upgrade=1,total_team_pledge = total_team_pledge + ? ,team_number = team_number+1,comet = comet +1 where uid =?",amount,sUser.Pid)
							return cuser(user,amount,1,1),nil
						}
					}else {
						//如果没拥有2颗彗星不能升级，只添加团队人数+1和增加总质押+amount,且彗星+1
						DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? ,team_number = team_number+1,comet = comet +1 where uid =?",amount,sUser.Pid)
						return cuser(user,amount,1,0),nil
					}
				}
			//	当前为行星判断是否升级为恒星
			case 3:
				//判断自己当前等级如果不为3，只添加团队人数+1和增加总质押+amount,且彗星+1
				if user.Grade==1 ||user.Grade==2 ||user.Grade==4{
					DB.Exec("update user set upgrade =0,comet=comet -1 ,total_team_pledge = total_team_pledge + ? ,team_number = team_number+1,planet = planet +1 where uid =?",amount,sUser.Pid)
					return cuserPlanet(user,amount,1,0),nil
				}else {
					//如果为2就需要判断是否升级
					if user.Planet==2 {
						//小于2000spt
						if user.TotalTeamPledge -user.TotalPledge<5000000000000 {
							DB.Exec("update user set upgrade =0,comet=comet -1,total_team_pledge = total_team_pledge + ? ,team_number = team_number+1,planet = planet +1 where uid =?",amount,sUser.Pid)
							return cuserPlanet(user,amount,1,0),nil
						}else {
							//彗星推荐等于3了。团队总质押大于50000了可以升级了
							DB.Exec("update user set grade = 3 , upgrade=1,comet=comet -1,total_team_pledge = total_team_pledge + ? ,team_number = team_number+1,planet = planet +1 where uid =?",amount,sUser.Pid)
							return cuserPlanet(user,amount,1,1),nil
						}
					}else {
						//如果没拥有2颗彗星不能升级，只添加团队人数+1和增加总质押+amount,且彗星+1
						DB.Exec("update user set upgrade =0, total_team_pledge = total_team_pledge + ?,comet=comet -1 ,team_number = team_number+1,planet = planet +1 where uid =?",amount,sUser.Pid)
						return cuserPlanet(user,amount,1,0),nil
					}
				}
			case 4:
				DB.Exec("update user set upgrade =0,planet=planet -1,team_number = team_number+1,fixstar = fixstar +1,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
				return cuser(user,amount,0,0),nil
				//如果为1或者4都不会导致升级，因此直接返回
			//为1.4星火会员和恒星不会升级直接加总额和推荐会员人数
			default:
				DB.Exec("update user set upgrade =0,team_number = team_number+1,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
				return cuser(user,amount,0,0),nil
			}
		}
	}else {
		switch sUser.Upgrade {
		case 0:
			//因为上一级未升级，所以当前用户只需改变团队人数和团队总质押,查看等级数量是否团队总质押量达到要求，如果达到就升级
			switch user.Grade{
			case 1:
				if user.Recommenders>4 && user.TotalTeamPledge - user.TotalPledge + amount>200000000000{
					DB.Exec("update user set upgrade =1,grade =2,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
					return cuser(user,amount,1,1),nil
				}else {
					DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
					return cuser(user,amount,0,0),nil
				}
			case 2:
				if user.Comet>2 &&user.TotalTeamPledge - user.TotalPledge + amount>1200000000000 {
					DB.Exec("update user set upgrade =1,grade =3,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
					return cuserPlanet(user,amount,1,1),nil
				}else {
					DB.Exec("update user set upgrade =0 ,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
					return cuser(user,amount,1,0),nil
				}
			case 3:
				if user.Planet>2 &&user.TotalTeamPledge - user.TotalPledge + amount>5000000000000 {
					DB.Exec("update user set upgrade =1,grade =3 ,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
					return cuserPlanet(user,amount,1,1),nil
				}else {
					DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
					return cuserPlanet(user,amount,1,0),nil
				}
			default:
				DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
				return cuser(user,amount,0,0),nil
			}
		default:
			//上一级升级了，分别查看升级的等级，彗星，行星，恒星
			switch sUser.Grade {
			//升级为彗星
			case 2:
				//判断自己当前等级如果不为2，只添加团队人数+1和增加总质押+amount,且彗星+1
				if user.Grade==1 ||user.Grade==3 ||user.Grade==4{
					DB.Exec("update user set upgrade =0 ,total_team_pledge = total_team_pledge + ? ,comet = comet +1 where uid =?",amount,sUser.Pid)
					return cuser(user,amount,1,0),nil
				}else {
					//如果为2就需要判断是否升级
					if user.Comet==2 {
						//小于2000spt
						if user.TotalTeamPledge -user.TotalPledge+amount<1200000000000 {
							DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? ,comet = comet +1 where uid =?",amount,sUser.Pid)
							return cuser(user,amount,1,0),nil
						}else {
							//彗星推荐等于3了。团队总质押大于12000了可以升级了
							DB.Exec("update user set grade = 3 , upgrade=1,total_team_pledge = total_team_pledge + ? ,comet = comet +1 where uid =?",amount,sUser.Pid)
							return cuser(user,amount,1,1),nil
						}
					}else {
						//如果没拥有2颗彗星不能升级，只添加团队人数+1和增加总质押+amount,且彗星+1
						DB.Exec("update user set upgrade =0,total_team_pledge = total_team_pledge + ? ,comet = comet +1 where uid =?",amount,sUser.Pid)
						return cuser(user,amount,1,0),nil
					}
				}
			//	当前为行星判断是否升级为恒星
			case 3:
				//判断自己当前等级如果不为3，只添加团队人数+1和增加总质押+amount,且彗星+1
				if user.Grade==1 ||user.Grade==2 ||user.Grade==4{
					DB.Exec("update user set upgrade =0,comet=comet -1 ,total_team_pledge = total_team_pledge + ? ,planet = planet +1 where uid =?",amount,sUser.Pid)
					return cuserPlanet(user,amount,1,0),nil
				}else {
					//如果为2就需要判断是否升级
					if user.Planet==2 {
						//小于2000spt
						if user.TotalTeamPledge -user.TotalPledge+amount<5000000000000 {
							DB.Exec("update user set upgrade =0,comet=comet -1,total_team_pledge = total_team_pledge + ? ,planet = planet +1 where uid =?",amount,sUser.Pid)
							return cuserPlanet(user,amount,1,0),nil
						}else {
							//彗星推荐等于3了。团队总质押大于50000了可以升级了
							DB.Exec("update user set grade = 4 , upgrade=1,comet=comet -1,total_team_pledge = total_team_pledge + ? ,planet = planet +1 where uid =?",amount,sUser.Pid)
							return cuserFixstar(user,amount,1,1),nil
						}
					}else {
						//如果没拥有2颗彗星不能升级，只添加团队人数+1和增加总质押+amount,且彗星+1
						DB.Exec("update user set upgrade =0, total_team_pledge = total_team_pledge + ?,comet=comet -1 ,planet = planet +1 where uid =?",amount,sUser.Pid)
						return cuserPlanet(user,amount,1,0),nil
					}
				}
			case 4:
				DB.Exec("update user set upgrade =0,fixstar = fixstar +1,planet = planet -1,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
				return cuser(user,amount,0,0),nil
				//如果为1或者4都不会导致升级，因此直接返回
			//为1.4星火会员和恒星不会升级直接加总额和推荐会员人数
			default:
				DB.Exec("update user set upgrade =0,planet = planet +1 ,total_team_pledge = total_team_pledge + ? where uid =?",amount,sUser.Pid)
				return cuser(user,amount,0,0),nil
			}
		}
	}
}


