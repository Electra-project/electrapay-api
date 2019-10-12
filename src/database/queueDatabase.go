package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var queuedb *sql.DB

func init() {

	var err error
	if os.Getenv("GO_ENV") != "production" {
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	queuedbName := os.Getenv("QUEUEDB")
	queuedbUser := os.Getenv("QUEUEUSER")
	queuedbPass := os.Getenv("QUEUEPASSWORD")
	queuedbHost := os.Getenv("QUEUEHOST")
	queuedbPort := os.Getenv("QUEUEPORT")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", queuedbUser, queuedbPass, queuedbHost, queuedbPort, queuedbName)

	queuedb, err = sql.Open("mysql", dbSource)

	if err != nil {
		panic(err)
	}
}

// Get a database instance.
func GetQueueDatabase() *sql.DB {
	return queuedb
}
