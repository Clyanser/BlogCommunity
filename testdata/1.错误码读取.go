package main

import (
	"GoBlog/models/res"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type ErrMap map[res.ErrorCode]string

const filepath = "models/res/err_code.json"

func main() {
	byteData, err := os.ReadFile(filepath)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
}
