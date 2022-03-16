package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Start ...
func Start() {
	var err error

	db, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_PROTOCOL")+"("+os.Getenv("DB_HOST")+")"+"/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatalf("[FATAL] Could not open connection to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("[FATAL] Could not open connect to database: %s", err)
	}

	db.SetConnMaxLifetime(time.Hour * 1)
}
