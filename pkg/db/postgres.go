package db

import (
	"database/sql"
	"fmt"
	"gocommerce/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.DBHost,
		config.Cfg.DBPort,
		config.Cfg.DBUser,
		config.Cfg.DBPassword,
		config.Cfg.DBName,
		config.Cfg.DBSSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully")

}

func GetDB() *sql.DB {
	return DB
}
