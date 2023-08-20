package cache

import (
	"fmt"
	"github/linchao0828/messengerBot/infra"
	"time"
)

var Config = &configCache{}

type configCache struct{}

func (*configCache) key(key string) string {
	return fmt.Sprintf("config:%s", key)
}

func (c *configCache) Set(k, v string) error {
	return infra.Redis.Set(c.key(k), v, 3*time.Second).Err()
}

func (*configCache) Get(k string) (string, bool) {
	v := infra.Redis.Get(k).Val()
	return v, len(v) > 0
}
