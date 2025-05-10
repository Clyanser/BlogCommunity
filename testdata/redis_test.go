package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	err := rdb.Set("xxx2", "value2", 10*time.Second).Err()
	fmt.Println(err)
	cmd := rdb.Keys("*")
	keys, err := cmd.Result()
	fmt.Println(keys, err)
}
