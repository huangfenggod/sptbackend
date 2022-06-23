package nosql

import (
	"context"
	"strconv"
)

func SetRateIntoRedis(rate float32)  {
	ctx := context.Background()
	Cli.SAdd(ctx,"rate",rate)
}

func RemoveRateRedis()  {
	ctx := context.Background()
	Cli.SPop(ctx,"rate")
}

func GetRateFromRedis()float32  {
	ctx := context.Background()
	member := Cli.SRandMember(ctx, "rate")
	float, err := strconv.ParseFloat(member.Val(), 10)
	if  err !=nil{
		return 1
	}else {
	return float32(float)
	}
}
