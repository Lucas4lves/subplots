package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Driver *sql.DB
}

func NewDatabaseConfig(driverHandler, connString string) *DatabaseConfig {
	db, err := sql.Open(driverHandler, connString)

	if err != nil {
		log.Fatal("ERROR: unable to open database ->", err.Error())
		return nil
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	return &DatabaseConfig{
		Driver: db,
	}
}
