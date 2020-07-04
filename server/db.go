package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbCredentials struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

func dbConnect(dbCred dbCredentials) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbCred.host, dbCred.port, dbCred.user, dbCred.password, dbCred.dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
