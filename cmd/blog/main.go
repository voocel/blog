package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"blog/config"
	"blog/internal/entity"
	"blog/internal/http"
	"blog/internal/repository/postgres"
	"blog/pkg/geoip"
	"blog/pkg/log"
	"blog/pkg/util"
)

var (
	createAdmin = flag.Bool("create-admin", false, "Create admin account")
)

func main() {
	config.LoadConfig()
	flag.BoolVar(&config.Conf.Postgres.Migrate, "migrate", true, "Auto migrate database")
	flag.Parse()

	log.Init(config.Conf.LogLevel, config.Conf.LogPath)

	if config.Conf.App.GeoIPDBPath != "" {
		if err := geoip.Init(config.Conf.App.GeoIPDBPath); err != nil {
			log.Errorf("GeoIP database initialization failed (will use 'Unknown' as location): %v", err)
		} else {
			log.Info("GeoIP database initialized successfully")
		}
	}

	// If create admin account parameter is set
	if *createAdmin {
		if err := createAdminAccount(); err != nil {
			log.Errorf("Failed to create admin account: %v", err)
			os.Exit(1)
		}
		log.Info("Admin account created successfully!")
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

// createAdminAccount creates admin account
func createAdminAccount() error {
	reader := bufio.NewReader(os.Stdin)

	username := promptInput(reader, "Enter admin username (default: admin): ", "admin")
	email := promptInput(reader, "Enter admin email (required): ", "")
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	password := promptInput(reader, "Enter admin password (required): ", "")
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	dbRepo, err := postgres.New()
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}
	defer func() {
		dbRepo.DbWClose()
		dbRepo.DbRClose()
	}()

	db := dbRepo.GetDbW()

	var existingUser entity.User
	err = db.Where("email = ?", email).First(&existingUser).Error
	if err == nil {
		return fmt.Errorf("admin account already exists: %s", existingUser.Email)
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return fmt.Errorf("password encryption failed: %w", err)
	}

	admin := &entity.User{
		Username:   username,
		Email:      email,
		Password:   hashedPassword,
		Provider:   "email",
		ProviderID: email,
		Role:       "admin",
		Bio:        "Administrator",
	}

	if err := db.Create(admin).Error; err != nil {
		return fmt.Errorf("failed to save admin account: %w", err)
	}

	fmt.Printf("\n=== Admin Account Created Successfully ===\n")
	fmt.Printf("Email: %s\n", admin.Email)
	fmt.Printf("Username: %s\n", admin.Username)
	fmt.Printf("Role: %s\n", admin.Role)
	fmt.Printf("==========================================\n\n")

	return nil
}

func promptInput(reader *bufio.Reader, label, def string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return def
	}
	return text
}
