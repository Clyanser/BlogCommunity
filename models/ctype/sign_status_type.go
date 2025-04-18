package ctype

import "encoding/json"

type StatusType int

const (
	SignQQ     StatusType = 1 //qq
	SignWechat            = 2 //微信
	SignEmail             = 3 //github
)

func (s StatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s StatusType) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "qq"
	case SignWechat:
		str = "微信"
	case SignEmail:
		str = "github"
	default:
		str = "其他"
	}
	return str
}
