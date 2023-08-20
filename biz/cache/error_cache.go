package cache

import (
	"context"
	"fmt"
	"github/linchao0828/messengerBot/infra"
	"strconv"
)

var ErrTimes = &errTimesCache{}

type errTimesCache struct{}

func (*errTimesCache) key(bucket, k string) string {
	return fmt.Sprintf("err:times:%s:%s", bucket, k)
}

func (e *errTimesCache) Add(ctx context.Context, bucket, k string) (int64, error) {
	return infra.Redis.WithContext(ctx).IncrBy(e.key(bucket, k), 1).Result()
}

func (e *errTimesCache) Get(ctx context.Context, bucket, k string) int64 {
	if r, err := infra.Redis.WithContext(ctx).Get(e.key(bucket, k)).Result(); err == nil {
		res, _ := strconv.ParseInt(r, 10, 64)
		return res
	}
	return 0
}
