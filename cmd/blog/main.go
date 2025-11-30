package main

import (
	"blog/config"
	"blog/internal/entity"
	"blog/internal/http"
	"blog/internal/repository/postgres"
	"blog/pkg/geoip"
	"blog/pkg/log"
	"blog/pkg/util"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	createAdmin = flag.Bool("create-admin", false, "创建管理员账号")
)

func main() {
	config.LoadConfig()
	flag.BoolVar(&config.Conf.Postgres.Migrate, "migrate", true, "是否自动迁移数据库")
	flag.Parse()

	log.Init("http", "debug")

	if config.Conf.App.GeoIPDBPath != "" {
		if err := geoip.Init(config.Conf.App.GeoIPDBPath); err != nil {
			log.Warnf("GeoIP 数据库初始化失败（将使用 'Unknown' 作为位置）: %v", err)
		} else {
			log.Info("GeoIP 数据库初始化成功")
		}
	}

	// 如果设置了创建管理员账号参数
	if *createAdmin {
		if err := createAdminAccount(); err != nil {
			log.Errorf("创建管理员账号失败: %v", err)
			os.Exit(1)
		}
		log.Info("管理员账号创建成功!")
		os.Exit(0)
	}

	srv := http.NewServer()
	srv.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	for {
		s := <-ch
		log.Infof("[%v] Shutting down...", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			if err := srv.Stop(ctx); err != nil {
				panic(err)
			}
			geoip.Close()
			log.Sync()
			cancel()
			return
		case syscall.SIGHUP:
			config.LoadConfig()
		default:
			return
		}
	}
}

// createAdminAccount 创建管理员账号
func createAdminAccount() error {
	dbRepo, err := postgres.New()
	if err != nil {
		return fmt.Errorf("数据库连接失败: %w", err)
	}
	defer func() {
		dbRepo.DbWClose()
		dbRepo.DbRClose()
	}()

	db := dbRepo.GetDbW()

	// 检查是否已存在管理员账号
	var existingUser entity.User
	err = db.Where("email = ? OR role = ?", "admin@example.com", "admin").First(&existingUser).Error
	if err == nil {
		return fmt.Errorf("管理员账号已存在: %s", existingUser.Email)
	}

	// 加密密码
	hashedPassword, err := util.HashPassword("admin123")
	if err != nil {
		return fmt.Errorf("密码加密失败: %w", err)
	}

	// 创建管理员用户
	admin := &entity.User{
		Username: "admin",
		Email:    "admin@example.com",
		Password: hashedPassword,
		Role:     "admin",
		Bio:      "Administrator",
	}

	// 保存到数据库
	if err := db.Create(admin).Error; err != nil {
		return fmt.Errorf("保存管理员账号失败: %w", err)
	}

	fmt.Printf("\n=== 管理员账号创建成功 ===\n")
	fmt.Printf("邮箱: %s\n", admin.Email)
	fmt.Printf("密码: admin123\n")
	fmt.Printf("角色: %s\n", admin.Role)
	fmt.Printf("========================\n\n")

	return nil
}
