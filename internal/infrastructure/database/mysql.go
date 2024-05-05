package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dsn string) (db *gorm.DB, err error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
