package db

import "github.com/jinzhu/gorm"

type IConnection interface {
	MakeConnect() (*gorm.DB, error)
	GetConnection() *gorm.DB
	RunMigration(models ...interface{})
	DeleteTable(models ...interface{})
}
