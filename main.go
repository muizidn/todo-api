package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/muizidn/todo-api/app"
)

const (
	port = 50051
)

func main() {
	app.Start(port)
}
