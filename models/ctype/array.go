package ctype

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Array []string

func (t *Array) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	if string(v) == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(string(v), "\n")
	return nil
}

func (t Array) Value() (driver.Value, error) {
	// 将数组转换为值
	return strings.Join(t, "\n"), nil
}
