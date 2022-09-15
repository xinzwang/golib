package utils

import (
	"fmt"
	"sync"
	"time"

	"e.coding.net/itdesk/weixin/golib/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	mysqlConfig config.MysqlConfig
	mysqlLock   sync.Mutex
	mysqlIns    *gorm.DB
)

func SetMysqlConfig(cfg *config.MysqlConfig) {
	mysqlConfig = *cfg
}

func NewMysqlByConfig() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Addr,
		mysqlConfig.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlDB := db.DB()
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn) //设置连接池 0为不限制
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn) //设置最大链接数
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	// db.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
	// 	if !scope.HasError() {
	// 		timeNow := time.Now().Unix()
	// 		if field, ok := scope.FieldByName("CreateDate"); ok {
	// 			if field.IsBlank {
	// 				field.Set(timeNow)
	// 			}
	// 		}
	// 		if field, ok := scope.FieldByName("UpdateDate"); ok {
	// 			if field.IsBlank {
	// 				field.Set(timeNow)
	// 			}
	// 		}
	// 	}
	// })

	return db, nil
}

func MysqlIns() *gorm.DB {
	if mysqlIns == nil {
		mysqlLock.Lock()
		defer mysqlLock.Unlock()
		if mysqlIns == nil {
			var err error
			mysqlIns, err = NewMysqlByConfig() // db_test
			if err != nil {
				panic(err)
			}
		}
	}
	return mysqlIns
}
