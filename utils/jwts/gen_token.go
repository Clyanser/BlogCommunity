package jwts

import (
	"GoBlog/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

func GetToken(user JwtPayload) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), //默认两小时过期
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(MySecret)
}
