package idgen

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	uidPrefix    int64 = 1888000000000000
	postIDPrefix int64 = 2888000000000000
	last         int64
	mu           sync.Mutex
)

func Next() int64 {
	mu.Lock()
	defer mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if last == now {
		time.Sleep(1 * time.Millisecond)
		now = time.Now().UnixNano() / 1e6
	}
	last = now
	res, err := strconv.ParseInt(strconv.FormatInt(now, 10)[1:], 10, 64)
	if err != nil {
		fmt.Println(fmt.Sprintf("generate id fail! err: %T", err))
	}
	return res
}

func Uid() int64 {
	return uidPrefix + Next()
}

func PostID() int64 {
	return postIDPrefix + Next()
}
