package main

import (
	database "IpIdentifier/internal/crud"
)

func main() {
	dbConn := database.PostgreConnect()
	database.DatabaseMigrate(dbConn)
}
