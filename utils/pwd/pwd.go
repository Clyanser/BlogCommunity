package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd 加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}

// VerifyPassword 验证密码 hash之后的密码  输入的密码
func VerifyPassword(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
