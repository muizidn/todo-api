package app

import (
	"database/sql"
	"fmt"

	"time"
)

type varchar = string
type text = string
type date = *time.Time

var _CurrentDB *sql.DB

func _connectMySQL() (*sql.DB, error) {
	dbUser := env.DbUser
	dbPass := env.DbPass
	dbHost := env.DbHost
	dbPort := env.DbPort
	dbName := env.DbName

	config := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?parseTime=true`, dbUser, dbPass, dbHost, dbPort, dbName)
	return sql.Open("mysql", config)
}

func _connectSQLite() (*sql.DB, error) {
	dbFile := env.DbFile
	config := fmt.Sprintf(`file:%s?cache=shared`, dbFile)
	return sql.Open("sqlite3", config)
}

func setupDatabase() *sql.DB {
	if _CurrentDB == nil {
		db, err := _connectSQLite()
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
		_CurrentDB = db
		return db
	}
	return _CurrentDB
}
