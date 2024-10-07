package database

import (
	"github.com/pkg/errors"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var (
	conn *gorm.DB
	once sync.Once
)

func connectMysql(dsn string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDatabaseConnection() (db *gorm.DB, err error) {
	dsn := "store:store.123@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
	var initErr error
	once.Do(func() {
		conn, err = connectMysql(dsn)
		if err != nil {
			initErr = errors.WithMessagef(err, "failed to open database: %q", dsn)
			return
		}
	})
	if initErr != nil {
		return nil, initErr
	}
	return conn, nil
}
