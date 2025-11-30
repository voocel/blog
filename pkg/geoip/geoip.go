package geoip

import (
	"fmt"
	"net"
	"sync"

	"github.com/oschwald/maxminddb-golang"
)

var (
	db   *maxminddb.Reader
	once sync.Once
	mu   sync.RWMutex
)

// CityRecord MaxMind City 数据库记录结构（也兼容 Country 数据库）
type CityRecord struct {
	City struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Country struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
		TimeZone  string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
}

// CountryRecord MaxMind Country 数据库记录结构
type CountryRecord struct {
	Country struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
}

// Init 初始化 GeoIP 数据库
func Init(dbPath string) error {
	var err error
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock()
		db, err = maxminddb.Open(dbPath)
	})
	return err
}

// Close 关闭数据库
func Close() error {
	mu.Lock()
	defer mu.Unlock()
	if db != nil {
		return db.Close()
	}
	return nil
}

// Lookup 查询 IP 地理位置
func Lookup(ipStr string) string {
	mu.RLock()
	defer mu.RUnlock()

	if db == nil {
		return "Unknown"
	}

	// 解析 IP
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "Invalid IP"
	}

	// 本地 IP 特殊处理
	if isLocalIP(ip) {
		return "Local"
	}

	// 查询数据库
	var record CityRecord
	err := db.Lookup(ip, &record)
	if err != nil {
		return "Unknown"
	}

	// 构建位置字符串
	return formatLocation(&record)
}

// formatLocation 格式化位置信息（兼容 City 和 Country 数据库）
func formatLocation(record *CityRecord) string {
	// 优先使用中文，其次英文
	city := getLocalizedName(record.City.Names)
	country := getLocalizedName(record.Country.Names)

	// City 和 Country 数据库都有国家信息
	if city != "" && country != "" {
		return fmt.Sprintf("%s, %s", city, country)
	}
	// Country 数据库只有国家
	if country != "" {
		return country
	}
	// City 数据库可能只有城市
	if city != "" {
		return city
	}
	return "Unknown"
}

// getLocalizedName 获取本地化名称（优先中文）
func getLocalizedName(names map[string]string) string {
	// 优先返回中文
	if name, ok := names["zh-CN"]; ok && name != "" {
		return name
	}
	// 其次英文
	if name, ok := names["en"]; ok && name != "" {
		return name
	}
	// 返回任意可用的名称
	for _, name := range names {
		if name != "" {
			return name
		}
	}
	return ""
}

// isLocalIP 判断是否为本地 IP
func isLocalIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsPrivate() {
		return true
	}
	// 检查是否为 IPv6 本地地址
	if ip.To4() == nil && ip.IsLinkLocalUnicast() {
		return true
	}
	return false
}

// IsInitialized 检查是否已初始化
func IsInitialized() bool {
	mu.RLock()
	defer mu.RUnlock()
	return db != nil
}
