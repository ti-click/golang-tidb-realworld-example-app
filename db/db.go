package db

import (
	"fmt"
	"strings"

	"os"

	"github.com/ti-click/golang-tidb-realworld-example-app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database connection infos
var dbName = "realworld"
var dbUser = "root"
var dbPass = ""
var dbHost = "127.0.0.1"
var dbPort = "4000"

// test Database
const TestDBName = "realworld_test"

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
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
}

func getDSNWithoutDBName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort)
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
	dbName = TestDBName
	testDsn := getDSN()
	db, err := getTiDB(testDsn)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1049: Unknown database") {
			// create database when not found
			tmpDsn := getDSNWithoutDBName()
			db, err := getTiDB(tmpDsn)
			if err != nil {
				panic("can not connect db without dbname: " + err.Error())
			}
			err = db.Exec("create database " + TestDBName).Error
			if err != nil {
				panic("can not create database " + err.Error())
			}
			sqlDB, _ := db.DB()
			sqlDB.Close()
			db, err = getTiDB(testDsn)
			if err != nil {
				panic("can not new test database " + err.Error())
			}
		} else {
			panic("storage err: " + err.Error())
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.SetMaxIdleConns(3)
	}
	return db
}

func DropTestDB(db *gorm.DB) error {
	err := db.Exec("drop database " + TestDBName).Error
	if err != nil {
		panic("can not drop test database")
	}
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
