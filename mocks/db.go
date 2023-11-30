package mocks

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	_db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      _db,
		SkipInitializeWithVersion: true,
	}))
	if err != nil {
		panic(err)
	}
	return db, mock
}
