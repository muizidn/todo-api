package app

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"time"
)

type varchar = string
type text = string
type date = *time.Time

var _CurrentDB *sql.DB

func _connectMySQL() (*sql.DB, error) {
	dbUser := viper.GetString("TODO_DBUSER")
	dbPass := viper.GetString("TODO_DBPASS")
	dbHost := viper.GetString("TODO_DBHOST")
	dbPort := viper.GetString("TODO_DBPORT")
	dbName := viper.GetString("TODO_DBNAME")

	config := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?parseTime=true`, dbUser, dbPass, dbHost, dbPort, dbName)
	return sql.Open("mysql", config)
}

func _connectSQLite() (*sql.DB, error) {
	dbFile := viper.GetString("TODO_DBFILE")
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
