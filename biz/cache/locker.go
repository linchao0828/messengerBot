package cache

import (
	"context"
	"github/linchao0828/messengerBot/infra"
	"time"
)

var Locker = &locker{}

type locker struct{}

func (l *locker) TryLock(ctx context.Context, key string, ttl time.Duration) bool {
	success, err := infra.Redis.WithContext(ctx).SetNX(key, "lock", ttl).Result()
	if err != nil {
		return false
	}
	return success
}

func (l *locker) Unlock(ctx context.Context, key string) {
	infra.Redis.WithContext(ctx).Del(key)
}
