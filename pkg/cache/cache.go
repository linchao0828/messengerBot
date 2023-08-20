package cache

import (
	"encoding/json"
	"github/linchao0828/messengerBot/infra"
	"time"
)

func Set(k string, v interface{}, expiration time.Duration) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return infra.Redis.Set(k, j, expiration).Err()
}

func Get(k string, res interface{}) bool {
	r := infra.Redis.Get(k).Val()
	if len(r) == 0 {
		return false
	}
	err := json.Unmarshal([]byte(r), &res)
	return err == nil
}
