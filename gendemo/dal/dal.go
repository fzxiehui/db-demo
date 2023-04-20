package dal

import (
	"fmt"
	"sync"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"

	"github.com/ag9920/db-demo/gendemo/dal/model"
)

/*
 * 只是维护了内存中的数据库连接，完成初始化，和业务无关。
 */
var DB *gorm.DB
var once sync.Once

func init() {
	// 只会执行一次
	once.Do(func() {
		DB = ConnectDB().Debug()
		// 自动迁移
		_ = DB.AutoMigrate(&model.User{}, &model.Passport{})
	})
}

func ConnectDB() (conn *gorm.DB) {
	//  打开数据库连接, :memory: 表示在内存中
	conn, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
