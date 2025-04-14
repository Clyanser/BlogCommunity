package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1
	PermissionUser        Role = 2
	PermissionGuest       Role = 3
	PermissionDisableUser Role = 4
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionGuest:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁言"
	default:
		str = "其他"
	}
	return str
}
