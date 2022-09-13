package database

import (
	"fmt"
	"time"

	//The mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbInstance *sqlx.DB
)

const (
	host     = "domain.in"
	password = "password"
	username = "username"
	database = "password"
)

// MakeDB returns the sqlx.DB singleton if it exists, else creates one
func MakeDB() (*sqlx.DB, error) {
	var err error
	if dbInstance == nil {
		dbInstance, err = sqlx.Open("mysql", makeDsn())
		if err != nil {
			return nil, err
		}
		dbInstance.SetConnMaxLifetime(10 * time.Second)
	}
	return dbInstance, err
}

func makeDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, database)
	return
}
