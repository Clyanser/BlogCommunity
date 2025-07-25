package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

//作为分发

type Option struct {
	DB   bool
	User string // -u admin / -u user
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//	解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigretion()
	}
	if option.User == "admin" || option.User == "user" {
		//	创建用户
		CreateUser(option.User)
		return
	}
}
