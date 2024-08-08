package main

import (
	"blog/config"
	"blog/internal/http"
	"blog/internal/repository/redis"
	"blog/pkg/log"
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadConfig()
	flag.BoolVar(&config.Conf.Mysql.Migrate, "migrate", true, "是否自动创建表")
	flag.Parse()
	log.Init("http", "debug")
	redis.Init()
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
