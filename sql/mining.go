package sql

import (
	"time"
)

func InsertMiningOrder(address string,power float32,amount int,txhash string)  {
	DB.Exec("insert moder(address,power,amount,days,create_time,over_time,txhash) values(?,?,?,1080,?,?,?)",address,power,amount,time.Now(),time.Unix(time.Now().Unix()+1080*24*60*60, 0),txhash)
}

func UpdateMinerByBuy(address string, power float32)  {
DB.Exec("insert miner(address,buy_power) values(?,?) on duplicate key update buy_power = buy_power +?",address,power,power)
}

func UpdateMinerByLp(address string,power float32)  {
	DB.Exec("insert miner(address ,lp_power) values(?,?) on duplicate key update lp_power = ?",address,power,power)
}
func UpdateminerBYLpEveryTime(address string,power float32)  {
	DB.Exec("update miner set lp_power = ? where address =?",power,address)
}

func UpdateMinerByReward(address string,power float32)  {
DB.Exec("update user set reward_power = reward_power + ? where address =?",power,address)
}
//上2级人员奖励算力
func FindBindAddress(address string,power float32){
	exec := DB.Debug().Exec("update (select paddress from miner where address = ?) p ,miner m set m.reward_power = m.reward_power + ? where  m.address =  p.paddress ", address,power/10)
	if exec.RowsAffected>0{
	DB.Debug().Exec("update (select m1.paddress from miner m1,(select paddress from miner where address =?)p1 where m1.address = p1.paddress) p2 ,miner m set m.reward_power = m.reward_power + ? where  m.address =  p2.paddress",address,power*0.06)
	}
}


//paddress存在return true
func GetMinerPaddress(address string) bool {
	var miner MinerInfo
	DB.Raw("select paddress from miner where address =?",address).Scan(&miner)
	if len(miner.Paddress)>0 {
		return true
	}else {
	return 	false
	}

}

func MinerBind(address string,paddress string)  {
	DB.Exec("insert into miner(address,paddress) values(?,?) on duplicate key update paddress=?",address,paddress,paddress)
}
type MinerInfo struct {
	Address string `db:"address"`
	Paddress string `db:"paddress""`
	BuyPower float32 `db:"buy_power"`
	RewardPower float32 `db:"reward_power"`
	LpPower float32 `db:"lp_power"`
	Withdraw float32 `db:"withdraw"`
}

func GetMinerLPInfo()[]MinerInfo  {
	var minfo []MinerInfo
	DB.Raw("select address ,lp_power from miner where lp_power>0").Scan(&minfo)
	return minfo
}


func GetMinerInfo(address string) MinerInfo{
	var mi MinerInfo
	DB.Raw("select address,paddress,buy_power,reward_power,lp_power,withdraw from miner where address =?",address).Scan(&mi)
	return mi
}

func UpdateMinerWithdraw(total float32)  {
	DB.Debug().Exec("update miner set withdraw =withdraw + (buy_power+reward_power+lp_power)*? ",total)
}

type moderInfo struct {
	Address string `db:"address"`
	Power float32 `db:"power"`
}




func GetOverMiningAddressOrder() []moderInfo{
	var minfo []moderInfo
	DB.Raw("select address,power from moder where isover =0 and now()>over_time").Scan(&minfo)
	DB.Exec("update moder set isover =1 where now()>over_time and isover =0")
	return minfo
}

func UpdatePowerTimeUpOder(address string,power float32)  {
	DB.Exec("update miner set buy_power =buy_power -? where address =?",power,address)
	exec := DB.Debug().Exec("update (select paddress from miner where address = ?) p ,miner m set m.reward_power = m.reward_power - ? where  m.address =  p.paddress ", address,power/10)
	if exec.RowsAffected>0{
		DB.Debug().Exec("update (select m1.paddress from miner m1,(select paddress from miner where address =?)p1 where m1.address = p1.paddress) p2 ,miner m set m.reward_power = m.reward_power - ? where  m.address =  p2.paddress",address,power*0.06)
	}
}

func WithdrawEth(address string)bool  {
	exec := DB.Exec("update miner set call='Y' where address =? and withdraw >0.03", address)
	if exec.RowsAffected>0 {
		return true
	}else {
		return false
	}
}
