package main

import (
	"GoBlog/core"
	"testing"
)

func TestNewCode(t *testing.T) {
	core.InitConf()
	core.InitLogger()
	NewCode().send("2585925873@qq.com", "验证码是 1019")
}
