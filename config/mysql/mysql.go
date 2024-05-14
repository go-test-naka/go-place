package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-test-naka/go-place/config/env"
	"github.com/go-test-naka/go-place/log"
)

func CreateConnection() *sql.DB {
	log.Info("Connecting to database...")
	var err error
	//root:my-secret-pw@tcp(localhost:3306)/sys
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", env.DB_USER, env.DB_PASS, env.DB_HOST, env.DB_PORT, env.DB_SCHEMA)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Error("Failed to connect to database", err)
		panic(err.Error())
	}
	//defer db.Close()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		log.Error("Failed to connect to database", err)
		panic(err.Error())
	}

	return db
}
