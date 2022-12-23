package config

import (
	"fmt"
	"go-best-practice/internal/entity"
	"go-best-practice/internal/utilities"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDb(config *AppConfig) *gorm.DB {
	loggerLevel := logger.Silent
	debugLogEnv := []string{"debug", "development"}
	if utilities.Contains(debugLogEnv, config.Env) {
		loggerLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  loggerLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + config.Database.Host + " port=" + config.Database.Port + " user=" + config.Database.Username + " password=" + config.Database.Port + " dbname=" + config.Database.Database + " sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	// sqlDB, _ := db.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)

	// sqlDB.SetConnMaxLifetime(time.Hour)

	// log.Println("Configuration : Database successfully connected")

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	dberr := db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
	)

	if dberr != nil && dberr.Error() != "" {
		fmt.Println(dberr.Error())
		log.Fatal(dberr)
	}

	log.Println("Configuration : Database successfully migrated")
}
