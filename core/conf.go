package core

import (
	"GoBlog/config"
	"GoBlog/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConfig, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("config Init Unmarsshall: %v", err)
	}
	log.Println("config yamlFile get success!")
	fmt.Println(c)
	//需要一个全局变量用于保存配置文件，存放在golbal 目录下
	global.Config = c

}
