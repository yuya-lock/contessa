package lib

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Connect() (*gorm.DB, *sql.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Fatal("Error while loading .env file.\n%s", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, DBPORT, DBNAME,
	)
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 3 * time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
		Logger:          gormLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("Error when connecting to database.\n%s", err)
	}
	logrus.Info("Database connection succeeded.")
	logrus.Info(fmt.Sprintf("User: %s, Host: %s, Database: %s", user, host, DBNAME))
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	return db, sqlDB, nil
}

func InitDB() error {
	logrus.Info("Start initializing DB.")
	db, sqlDB, err := Connect()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when initializing DB")
	}
	defer sqlDB.Close()
	err = db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("error occured when migrating database\n%s", err)
	}
	return nil
}
