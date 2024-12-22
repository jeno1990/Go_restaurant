package main

import (
	"basic_go_backend/cmd/api"
	"basic_go_backend/config"
	"basic_go_backend/db"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	configs := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAdress,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, e := db.NewMysqlStorage(configs)
	if e != nil {
		log.Fatal(e)
	}
	// the first ping
	initStorage(db)

	server := api.NewAPIServer(":8081", db)
	err := server.Run()
	if err != nil {
		log.Fatal(err)

	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
}
