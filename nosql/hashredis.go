package nosql

import (
	"context"
	"time"
)

func InsertHashToRedis(hash string)  {
	ctx := context.Background()
	if IsHasTxhash(hash){
		return
	}
	Cli.SAdd(ctx, "tx", hash)
	Cli.Expire(ctx,"tx",9999*time.Hour)
}

func IsHasTxhash(hash string)bool  {
	ctx := context.Background()
	val := Cli.SIsMember(ctx, "tx", hash).Val()
	if val {
		return true
	}else {
		return false
	}
}
