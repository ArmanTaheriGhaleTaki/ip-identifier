package main

import (
	database "mmd/internal/crud"
)

func main() {
	db_conn := database.Postgre_connect()
	database.Database_migrate(db_conn)
}
