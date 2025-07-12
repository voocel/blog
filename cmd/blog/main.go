package main

import (
	"blog/config"
	"blog/internal/entity"
	"blog/internal/http"
	"blog/internal/repository/postgres"
	"blog/internal/repository/redis"
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
	flag.BoolVar(&config.Conf.Postgres.Migrate, "migrate", true, "是否自动创建表")
	flag.Parse()

	log.Init("http", "debug")
	redis.Init()

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
		log.Infof("[%v]Shutting down...", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			if err := srv.Stop(ctx); err != nil {
				panic(err)
			}
			redis.Close()
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
	// 连接数据库
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
	err = db.Where("email = ? OR role = ?", "admin@gmail.com", "admin").First(&existingUser).Error
	if err == nil {
		return fmt.Errorf("管理员账号已存在: %s", existingUser.Email)
	}

	// 加密密码
	hashedPassword, err := util.HashPassword("123456")
	if err != nil {
		return fmt.Errorf("密码加密失败: %w", err)
	}

	// 创建管理员用户
	admin := &entity.User{
		Username:      "admin",
		Email:         "admin@gmail.com",
		Password:      hashedPassword,
		Nickname:      "管理员",
		Role:          "admin",
		Status:        "active",
		Scores:        0,
		Sex:           0,
		Source:        0,
		Birthday:      time.Now(),
		LastLoginTime: time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 保存到数据库
	if err := db.Create(admin).Error; err != nil {
		return fmt.Errorf("保存管理员账号失败: %w", err)
	}

	fmt.Printf("\n=== 管理员账号创建成功 ===\n")
	fmt.Printf("邮箱: %s\n", admin.Email)
	fmt.Printf("密码: 123456\n")
	fmt.Printf("角色: %s\n", admin.Role)
	fmt.Printf("状态: %s\n", admin.Status)
	fmt.Printf("========================\n\n")

	return nil
}
