package mp

import (
	"github.com/maciekmm/messenger-platform-go-sdk"
	"sync"
)

var (
	once sync.Once
	Cli  *messenger.Messenger
)

func Init(accessToken, verifyToken string) {
	once.Do(func() {
		Cli = &messenger.Messenger{
			AccessToken: accessToken,
			VerifyToken: verifyToken,
			Debug:       messenger.DebugAll,
		}
	})
}
