package models

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DefaulDB = &gormDb{}

type gormDb struct {
	db *gorm.DB
}

func (g *gormDb) DB() *gorm.DB {
	return g.db
}

func (g *gormDb) Init(dsn string, dst ...interface{}) error {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	g.db = db
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	db.AutoMigrate(dst...)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
