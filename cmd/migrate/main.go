package main

import (
	database "mmd/internal/crud"
)

func main() {
	dbConn := database.PostgreConnect()
	database.DatabaseMigrate(dbConn)
}
