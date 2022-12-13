package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Con *gorm.DB

type MigrationHistory struct {
	Id       string    `gorm:"primary_key;not_null;size:50" json:"id"`
	RunnerAt time.Time `gorm:"autoCreateTime" json:"runnerAt"`
}

func ConnectToDatabase(withDatabase bool) error {
	mysqlString := GetStringDatabase(withDatabase)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	Con, err = gorm.Open(mysql.Open(mysqlString), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}

func CloseConnection() {
	dbInstance, _ := Con.DB()
	_ = dbInstance.Close()
}
