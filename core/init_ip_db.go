package core

import (
	ipUtils "GoBlog/utils/ip"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
)

var searcher *xdb.Searcher

func InitIpDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Fatalf("init ip2region.xdb err: %v\n", err)
		return
	}
	searcher = _searcher
}
func GetIpAddr(ip string) (addr string) {
	if ipUtils.HasLocalIPAddr(ip) {
		return "内网ip"
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("fatil to searchIP(%s): %s/n", ip, err)
		return "异常ip地址"
	}
	addrList := strings.Split(region, "|")
	if len(addrList) <= 5 {
		logrus.Warnf("ip(%s) not valid ip", ip)
		return "非法ip地址"
	}
	//addrList(国家 0 省份 城市 运营商) 5个部分
	country := addrList[0]
	province := addrList[2]
	city := addrList[3]

	if province != "0" && city != "0" {
		return fmt.Sprintf("%s|%s", province, city)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s|%s", country, province)
	}
	if country != "0" {
		return fmt.Sprintf("%s", country)
	}

	return region
}
