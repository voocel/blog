package lock

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	ttl              = time.Second * 30
	resetTTLInterval = ttl / 3
	tryLockInterval  = time.Second
	unlockScript     = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end`
)

var (
	ErrLockFailed = errors.New("lock failed")
	ErrTimeout    = errors.New("timeout")
)

type RedLock struct {
	clients        []*redis.Client
	successClients []*redis.Client
	script         *redis.Script
	resource       string
	randomValue    string
	watchDog       chan struct{}
}

func NewRedLock(clients []*redis.Client, resource string) *RedLock {
	return &RedLock{
		clients:  clients,
		script:   redis.NewScript(unlockScript),
		resource: resource,
	}
}

func (l *RedLock) TryLock(ctx context.Context) error {
	randomValue := uuid.New().String()
	var wg sync.WaitGroup
	wg.Add(len(l.clients))
	successClients := make(chan *redis.Client, len(l.clients))
	for _, client := range l.clients {
		go func(client *redis.Client) {
			defer wg.Done()
			success, err := client.SetNX(ctx, l.resource, randomValue, ttl).Result()
			if err != nil {
				return
			}
			if !success {
				return
			}

			go l.startWatchDog()
			successClients <- client
		}(client)
	}

	wg.Wait()
	close(successClients)
	if len(successClients) < len(l.clients)/2+1 {
		for client := range successClients {
			go func(client *redis.Client) {
				ctx, cancel := context.WithTimeout(context.Background(), ttl)
				l.script.Run(ctx, client, []string{l.resource}, randomValue)
				cancel()
			}(client)
		}
		return ErrLockFailed
	}

	l.randomValue = randomValue
	l.successClients = nil
	for successClient := range successClients {
		l.successClients = append(l.successClients, successClient)
	}

	return nil
}

func (l *RedLock) startWatchDog() {
	l.watchDog = make(chan struct{})
	ticker := time.NewTicker(resetTTLInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for _, client := range l.successClients {
				go func(client *redis.Client) {
					ctx, cancel := context.WithTimeout(context.Background(), ttl-resetTTLInterval)
					client.Expire(ctx, l.resource, ttl)
					cancel()
				}(client)
			}
		case <-l.watchDog:
			return
		}
	}
}

func (l *RedLock) Unlock(ctx context.Context) error {
	for _, client := range l.successClients {
		go func(client *redis.Client) {
			l.script.Run(ctx, client, []string{l.resource}, l.randomValue)
		}(client)
	}
	close(l.watchDog)
	return nil
}
