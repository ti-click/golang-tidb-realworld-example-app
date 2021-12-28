package db

import (
	"fmt"

	"os"

	"github.com/xesina/golang-echo-realworld-example-app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database connection infos
var dbName = "realworld"
var dbUser = ""
var dbPass = ""
var dbHost = ""
var dbPort = ""

// load database connection information
func init() {
	if envDbName := os.Getenv("DB_NAME"); envDbName != "" {
		dbName = envDbName
	}
	if envDbUser := os.Getenv("DB_USER"); envDbUser != "" {
		dbUser = envDbUser
	}
	if envDbPass := os.Getenv("DB_PASS"); envDbPass != "" {
		dbPass = envDbPass
	}
	if envDbHost := os.Getenv("DB_HOST"); envDbHost != "" {
		dbHost = envDbHost
	}
	if envDbPort := os.Getenv("DB_PORT"); envDbPort != "" {
		dbPort = envDbPort
	}
}

func getDSN() string {
	//if dbUser == "" { }
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
}

func getTiDB(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableNestedTransaction:                 true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return db, err
}

func New() *gorm.DB {
	dsn := getDSN()
	db, err := getTiDB(dsn)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}

func TestDB() *gorm.DB {
	testDsn := getDSN()
	db, err := getTiDB(testDsn)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.SetMaxIdleConns(3)
	}
	return db
}

func DropTestDB() error {
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
		&model.Article{},
		&model.Comment{},
		&model.Tag{},
	)
}
