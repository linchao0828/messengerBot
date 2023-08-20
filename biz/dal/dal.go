package dal

import (
	"fmt"
	"strings"

	"github/linchao0828/messengerBot/biz/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dsn string) {
	var err error
	var db *gorm.DB

	if strings.HasSuffix(dsn, "sqlite.db") {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn))
	}

	if err != nil {
		panic(fmt.Errorf("connect db fail: %s", err))
	}

	query.SetDefault(db)
}
