package main

import (
	"oruka/api"
	dbconnection "oruka/db-connection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client := dbconnection.Connect()

	api.Run(client)
}
