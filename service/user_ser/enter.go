package user_ser

import (
	"GoBlog/service/redis_ser"
	"GoBlog/utils/jwts"
	"time"
)

type UserService struct {
}

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Sub(now)
	redis_ser.Logout(token, diff)

	return redis_ser.Logout(token, diff)
}
