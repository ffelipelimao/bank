package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
