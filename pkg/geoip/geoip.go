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

type CountryRecord struct {
	Country struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
}

func Init(dbPath string) error {
	var err error
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock()
		db, err = maxminddb.Open(dbPath)
	})
	return err
}

func Close() error {
	mu.Lock()
	defer mu.Unlock()
	if db != nil {
		return db.Close()
	}
	return nil
}

func Lookup(ipStr string) string {
	mu.RLock()
	defer mu.RUnlock()

	if db == nil {
		return "Unknown"
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "Invalid IP"
	}

	if isLocalIP(ip) {
		return "Local"
	}

	var record CityRecord
	err := db.Lookup(ip, &record)
	if err != nil {
		return "Unknown"
	}

	return formatLocation(&record)
}

func formatLocation(record *CityRecord) string {
	city := getLocalizedName(record.City.Names)
	country := getLocalizedName(record.Country.Names)

	if city != "" && country != "" {
		return fmt.Sprintf("%s, %s", city, country)
	}
	if country != "" {
		return country
	}
	if city != "" {
		return city
	}
	return "Unknown"
}

func getLocalizedName(names map[string]string) string {
	if name, ok := names["zh-CN"]; ok && name != "" {
		return name
	}
	if name, ok := names["en"]; ok && name != "" {
		return name
	}
	for _, name := range names {
		if name != "" {
			return name
		}
	}
	return ""
}

func isLocalIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsPrivate() {
		return true
	}
	if ip.To4() == nil && ip.IsLinkLocalUnicast() {
		return true
	}
	return false
}

func IsInitialized() bool {
	mu.RLock()
	defer mu.RUnlock()
	return db != nil
}
