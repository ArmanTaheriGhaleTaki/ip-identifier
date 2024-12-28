package crud

import (
	"errors"
	"fmt"
	"log"
	ip_struct "mmd/internal/ip"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseInfo struct {
	address  string
	table    string
	user     string
	password string
	port     int
	sslmode  string
	TimeZone string
}

func PostgreConnect() *gorm.DB {
	var db databaseInfo
	db.table = os.Getenv("POSTGRES_DB")
	db.user = os.Getenv("POSTGRES_USER")
	db.password = os.Getenv("POSTGRES_PASSWORD")
	db.address = "db"
	db.port = 5432
	db.sslmode = "disable"
	db.TimeZone = "Asia/Tehran"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		db.address,
		db.user,
		db.password,
		db.table,
		db.port,
		db.sslmode,
		db.TimeZone,
	)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Printf("conneced to database")
	return dbConn
}

func DatabaseRetrieve(ip *ip_struct.IpInfo) {
	dbConn := PostgreConnect()
	err := dbConn.First(&ip, "ip = ?", ip.Ip).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("record not found for ip: %v", ip.Ip)
	} else {
		log.Printf("record found for ip: %v", ip.Ip)
	}
}

func DatabaseInsert(ip *ip_struct.IpInfo) {
	dbConn := PostgreConnect()
	dbConn.Create(&ip)
	log.Printf("inserted to database for ip: %v", ip.Ip)
}

func DatabaseMigrate(dbConn *gorm.DB) {
	if err := dbConn.AutoMigrate(ip_struct.IpInfo{}); err != nil {
		log.Fatalf("fialed to migrate: %v ", err)
	}
	println("migrait is successfully")
}
