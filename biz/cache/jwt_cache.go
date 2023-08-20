package cache

import (
	"context"
	"fmt"
	"github/linchao0828/messengerBot/infra"
	"time"
)

var Jwt = &jwtCache{}

type jwtCache struct{}

func (*jwtCache) blockKey(k string) string {
	return fmt.Sprintf("jwt:block:%s", k)
}

func (j *jwtCache) Block(ctx context.Context, k string) error {
	return infra.Redis.WithContext(ctx).Set(j.blockKey(k), "true", 60*time.Minute).Err()
}

func (j *jwtCache) IsBlock(ctx context.Context, k string) bool {
	return infra.Redis.WithContext(ctx).Get(j.blockKey(k)).Val() == "true"
}
