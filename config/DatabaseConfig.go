package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	host       = "localhost"
	port       = "3399"
	user       = "user"
	password   = "password"
	database   = "database"
	driverName = "mysql"
)

func DatabaseConnection() *sql.DB {

	db, err := sql.Open(driverName, user+":"+password+"@tcp("+host+":"+port+")/"+database+"?parseTime=true")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}

	fmt.Println("database connected!")

	return db
}

func DatabaseMigration(db *sql.DB) {

	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migration",
	}

	// Execution all migration files
	n, err := migrate.Exec(db, driverName, migrations, migrate.Up)
	if err != nil {
		log.Error().Msg("Errror Migration... " + err.Error())
		return
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
