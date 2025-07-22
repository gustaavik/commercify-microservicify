package db

import (
	"fmt"
	"log"
	"os"

	"github.com/zenfulcode/commercifyms/pkg/common"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbType := os.Getenv("DB_DRIVER")

	var db *gorm.DB
	var err error

	switch dbType {
	case "postgres":
		db, err = initPostgresDB()
	case "sqlite", "":
		db, err = initSQLiteDB()
	default:
		log.Fatalf("Unsupported DB_DRIVER: %s", dbType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

// initSQLiteDB initializes SQLite database connection
func initSQLiteDB() (*gorm.DB, error) {
	dbPath := common.GetEnv("DB_NAME", "./commercify.db")
	if dbPath == "" {
		dbPath = "./commercify.db"
	}
	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

func initPostgresDB() (*gorm.DB, error) {
	// Get database connection details from environment
	host := common.GetEnv("DB_HOST", "localhost")
	port := common.GetEnv("DB_PORT", "5432")
	user := common.GetEnv("DB_USER", "postgres")
	password := common.GetEnv("DB_PASSWORD", "postgres")
	dbname := common.GetEnv("DB_NAME", "commercify")
	sslmode := common.GetEnv("DB_SSL_MODE", "disable")

	// Build connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open database connection
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
