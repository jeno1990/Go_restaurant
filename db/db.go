package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlStorage(cfg mysql.Config) (*sql.DB, error) {
	dsn := cfg.FormatDSN() // this is the driver to sql
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
