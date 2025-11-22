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

func (dc *DatabaseConfig) CreatePlotsTable() error {

	sql := `
	    create table if not exists plots (
			id serial primary key,
			title varchar(255) unique not null,
			story varchar(512) not null,
			status boolean default true,
			created_at timestamp with time zone default now(),
			updated_at timestamp with time zone default now()
	);
	`

	_, err := dc.Driver.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}
