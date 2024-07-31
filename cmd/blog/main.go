package main

import (
	"blog/config"
	"blog/internal/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadConfig()
	srv := http.NewServer()
	srv.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			if err := srv.Stop(ctx); err != nil {
				panic(err)
			}
			cancel()
			return
		case syscall.SIGHUP:
			config.LoadConfig()
		default:
			return
		}
	}
}
