package database

import (
	"log"
	"test-go-developer/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=" + configs.DatabaseHost + " " + "port=" + configs.DatabasePort + " " + "user=" + configs.DatabaseUser + " " + "password=" + configs.DatabasePassword + " " +
		"dbname=" + configs.DatabaseName + " " + "sslmode=" + configs.DatabaseSSL
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger: func() logger.Interface {
			if configs.DatabaseLog == "disable" {
				return logger.Default.LogMode(logger.Silent)
			}
			return logger.Default.LogMode(logger.Info)
		}(),
	})

	if err != nil {
		panic("failed to connect database error: " + err.Error())
	}
	log.Println("Database connection established")
	DB = db
}
