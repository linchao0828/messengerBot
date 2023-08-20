package cache

import "fmt"

func AvatarCache(address string) string {
	return fmt.Sprintf("avatar:%s", address)
}
