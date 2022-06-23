package nosql

import (
	"context"
)

func SetInAddressForLp(address string)  {
	ctx := context.Background()
	if IsAlreadyHas(address) {
		return
	}
	Cli.SAdd(ctx, "lp", address)
}

func GetAddressForLp() []string {
	ctx := context.Background()
	members := Cli.SMembers(ctx, "lp")
	return members.Val()
}
func IsAlreadyHas(address string) bool {
	ctx := context.Background()
	is := Cli.SIsMember(ctx, "lp", address).Val()
	return is
}

func RemoveLp()  {
	ctx :=context.Background()
	Cli.SPopN(ctx, "lp", int64(len(GetAddressForLp())))
}



