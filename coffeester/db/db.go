package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10 
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d,err := sql.Open("pgx",dsn)
	if err != nil {
		return nil,err
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(int(maxDbLifetime))
	
	err = testDb(d)
}

func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err 
	}
	fmt.Println("*** Pinged database syccessfully! ***")
	return nil 
}