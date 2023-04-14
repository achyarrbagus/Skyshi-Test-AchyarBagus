package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	var MYSQL_HOST = os.Getenv("MYSQL_HOST")
	var MYSQL_PORT = os.Getenv("MYSQL_PORT")
	var MYSQL_USER = os.Getenv("MYSQL_USER")
	var MYSQL_DBNAME = os.Getenv("MYSQL_DBNAME")
	var MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")

	// dsn := MYSQL_USER + ":" + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DBNAME)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to database")
}
