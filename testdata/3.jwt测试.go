package main

import (
	"GoBlog/core"
	"GoBlog/global"
	"GoBlog/utils/jwts"
	"fmt"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GetToken(jwts.JwtPayload{
		UserID: 1,
		Role:   1,
		//Username: "admin",
		Nickname: "1019",
	})
	fmt.Println(token, err)
	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja25hbWUiOiIxMDE5Iiwicm9sZSI6MSwidXNlcmlkIjoxLCJleHAiOjE3NDU0MDAwMjMuMjk0NTQ0LCJpc3MiOiJ4eCJ9.acGG6KJbNlz92M4bJc5PgwWrcSsJXcodN9F4f1tXPMo")
	fmt.Println(claims, err)
}
