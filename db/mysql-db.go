package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	SQL *gorm.DB
}

func New() *Connection {
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", "root:2land@tcp(10.0.16.26:3306)/2land?charset=utf8mb4&parseTime=True&loc=Local")
	//defer db.Close()
	db.LogMode(true)

	if err != nil {
		panic(err.Error())
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err.Error())
	}

	return &Connection{
		SQL: db,
	}
}

func (conn *Connection) MakeConnect() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", "root:2land@tcp(10.0.16.26:3306)/2land?charset=utf8mb4&parseTime=True&loc=Local")
	//defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err.Error())
	}
	conn.SQL = db
	return db, err
}

func (conn *Connection) GetConnection() *gorm.DB {
	return conn.SQL
}

func (conn *Connection) RunMigration(models ...interface{}) {
	for i := 0; i < len(models); i++ {
		conn.SQL.AutoMigrate(models[i])
	}
}

func (conn *Connection) DeleteTable(models ...interface{}) {
	for i := 0; i < len(models); i++ {
		if conn.SQL.HasTable(models[i]) {
			conn.SQL.DropTable(models[i])
		}
	}
}
