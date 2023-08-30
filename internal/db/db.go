package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

func Get(connStr string) (*DB, error) {

	// Gets db instance with the given connection string
	db, err := get(connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error{
	return d.Client.Close()
}

func get(connStr string) (*sql.DB, error) {
	// Opens connection to Postgres
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ping to check if connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}