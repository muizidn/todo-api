package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/muizidn/todo-api/app"
)

const (
	port = 50051
)

func main() {
	app.Start(port)
}
