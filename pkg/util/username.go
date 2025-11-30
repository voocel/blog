package util

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomUsername 生成随机用户名
// 格式: user_ + 时间戳后6位 + 随机4位数字
func GenerateRandomUsername() string {
	// 使用当前时间作为随机种子
	rand.Seed(time.Now().UnixNano())
	
	// 获取时间戳后6位
	timestamp := time.Now().Unix() % 1000000
	
	// 生成4位随机数字
	randomNum := rand.Intn(10000)
	
	return fmt.Sprintf("user_%06d%04d", timestamp, randomNum)
}
