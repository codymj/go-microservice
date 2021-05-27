package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// Database for sharing db connection
type Database struct {
	Conn *sql.DB
}

// Init initializes a database connection
func Init() (Database, error) {
	log.Println("Initializing new database connection...")

	// Get environment variables
	USER := os.Getenv("POSTGRES_USER")
	PASS := os.Getenv("POSTGRES_PASS")
	HOST := os.Getenv("POSTGRES_HOST")
	PORT := os.Getenv("POSTGRES_PORT")
	DB := os.Getenv("POSTGRES_DB")

	db := Database{}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s "+
			"sslmode=disable connect_timeout=5",
		HOST, PORT, USER, PASS, DB,
	)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Error connecting to database")
		return db, err
	}

	// Ping connection for success
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")

	return db, nil
}
