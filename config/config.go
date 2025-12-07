package config

import (
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(config)

type config struct {
	Mode            string
	LogLevel        string `mapstructure:"log_level"`
	LogPath         string `mapstructure:"log_path"`
	LogLevelAddr    string `mapstructure:"log_level_addr"`
	LogLevelPattern string `mapstructure:"log_level_pattern"`

	App      AppConfig
	Http     HttpConfig
	Postgres PostgresConfig
	Mysql    MysqlConfig
}

type AppConfig struct {
	JwtSecret          string `mapstructure:"jwt_secret"`           // JWT signing secret key
	JwtAccessDuration  int    `mapstructure:"jwt_access_duration"`  // Access token duration in minutes (default: 15)
	JwtRefreshDuration int    `mapstructure:"jwt_refresh_duration"` // Refresh token duration in days (default: 7)
	UploadPath         string `mapstructure:"upload_path"`
	GeoIPDBPath        string `mapstructure:"geoip_db_path"`
}

type HttpConfig struct {
	Addr string
}

type PostgresConfig struct {
	Host            string
	Port            int
	Dbname          string
	Username        string
	Password        string
	SSLMode         string `mapstructure:"ssl_mode"`
	Migrate         bool
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifeTime time.Duration `mapstructure:"conn_max_life_time"`
}

type MysqlConfig struct {
	Host            string
	Port            int
	Dbname          string
	Username        string
	Password        string
	SSLMode         string `mapstructure:"ssl_mode"`
	Migrate         bool
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifeTime time.Duration `mapstructure:"conn_max_life_time"`
}

func LoadConfig(paths ...string) {
	if len(paths) == 0 {
		viper.AddConfigPath(".")
		viper.AddConfigPath("config")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("../../config")
	} else {
		for _, path := range paths {
			viper.AddConfigPath(path)
		}
	}

	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("blog")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("mode", "debug")
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_path", "log")
	viper.SetDefault("log_level_pattern", "/log/level")
	viper.SetDefault("atomic_level_addr", "4240")

	viper.SetDefault("http.addr", ":8090")

	// JWT defaults
	viper.SetDefault("app.jwt_secret", "change-this-secret-in-production")
	viper.SetDefault("app.jwt_access_duration", 15)  // 15 minutes
	viper.SetDefault("app.jwt_refresh_duration", 7)  // 7 days

	// Read config.yaml (required)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("read config error: %v", err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		log.Panicf("unmarshal config err: %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config change: %s, %s, %s\n", e.Op.String(), e.Name, e.String())
		if err := viper.Unmarshal(Conf); err != nil {
			log.Printf("config change unmarshal err: %v", err)
		}
	})
	log.Println("load config successfully")
}
