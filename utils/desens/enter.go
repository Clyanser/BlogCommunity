package desens

import "strings"

func DesensitizationEmail(email string) string {
	//2585925873@qq.com
	//25****5873@qq.com
	eslist := strings.Split(email, "@")
	if len(eslist) != 2 {
		return ""
	}
	return eslist[0][:1] + "****@" + eslist[1]
}
func DesensitizationTel(tel string) string {
	//15170806298
	//151****6298
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
