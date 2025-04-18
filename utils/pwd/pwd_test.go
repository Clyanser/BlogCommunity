package pwd

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}
func TestVerifyPassword(t *testing.T) {
	fmt.Println(VerifyPassword("$2a$04$sWvCItz2LBdPGoqqSACmjOhhPL5Z1n9rdSDLxKgVYful5Mls0HJlm", "123452"))
}
