package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04v", rand.Intn(10000))
}
