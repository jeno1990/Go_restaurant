package main

/*

if the automation does not work use
migrate -path ./cmd/migrate/migrations -database "mysql://<db_user>:<db_pass>@tcp(127.0.0.1:3306)/<db_name>" up
migrate -path ./cmd/migrate/migrations -database "mysql://<db_user>:<db_pass>@tcp(127.0.0.1:3306)/<db_name>" down
*/
import (
	"log"
	"os"

	"basic_go_backend/config"
	"basic_go_backend/db"

	// _ "github.com/go-sql-driver/mysql" // mysql driver
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := mysqlDriver.Config{
		Passwd:               config.Envs.DBPassword,
		User:                 config.Envs.DBUser,
		Addr:                 config.Envs.DBAdress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMysqlStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	// if cmd == "up" {
	// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 		log.Fatal(err)
	// 	}
	// }
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
