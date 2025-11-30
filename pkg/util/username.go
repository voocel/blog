package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomUsername() string {
	rand.Seed(time.Now().UnixNano())

	timestamp := time.Now().Unix() % 1000000
	randomNum := rand.Intn(10000)

	return fmt.Sprintf("user_%06d%04d", timestamp, randomNum)
}
