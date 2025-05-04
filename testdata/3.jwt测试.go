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
	//claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrbmFtZSI6IjEwMTkiLCJyb2xlIjoxLCJ1c2VyaWQiOjEsImV4cCI6MTc0NjUwMDE3Ny45NTA0NTEsImlzcyI6Inh4In0.4T501-3Wy4ajHBffDb5dtVa6OMXHT0F8ZhiQ9XN_34o")
	//fmt.Println(claims, err)
}
