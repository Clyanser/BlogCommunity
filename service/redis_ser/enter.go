package redis_ser

import (
	"GoBlog/global"
	"GoBlog/utils"
	"fmt"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("%s%s", prefix, token), "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	//判断是否在redis中
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.IsInList(prefix+"*"+token, keys) {
		return true
	}
	return false
}
