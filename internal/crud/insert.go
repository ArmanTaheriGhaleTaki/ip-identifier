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

type database_info struct {
	address  string
	table    string
	user     string
	password string
	port     int
	sslmode  string
	TimeZone string
}

func Postgre_connect() *gorm.DB {
	var db database_info
	db.table = os.Getenv("POSTGRES_DB")
	db.user = os.Getenv("POSTGRES_USER")
	db.password = os.Getenv("POSTGRES_PASSWORD")
	db.address = "db"
	db.port = 5432
	db.sslmode = "disable"
	db.TimeZone = "Asia/Tehran"
	// Construct the DSN (Data Source Name)
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
	// Open a connection to the database
	db_conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Printf("conneced to database")
	return db_conn

}
func Database_retrieve(ip *ip_struct.IP_info) {
	db_conn := Postgre_connect()

	err := db_conn.First(&ip, "ip = ?", ip.Ip).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("record not found for ip: %v", ip.Ip)
	} else {
		log.Printf("record found for ip: %v", ip.Ip)
	}

}
func Database_insert(ip *ip_struct.IP_info) {
	db_conn := Postgre_connect()
	db_conn.Create(&ip)
	log.Printf("inserted to database for ip: %v", ip.Ip)
}
func Database_migrate(db_conn *gorm.DB) {
	// Migrate the schema
	if err := db_conn.AutoMigrate(ip_struct.IP_info{}); err != nil {
		log.Fatalf("fialed to migrate: %v ", err)
	}
	println("migrait is successfully")
}
