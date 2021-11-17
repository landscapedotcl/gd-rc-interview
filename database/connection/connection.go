package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // it provides the drivers for postgres
)

type PostgreClient struct {
	*sql.DB
}

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "landscape-db"
)

// Fetches the database connection
func GetPostgreClient() *PostgreClient {
	// data source name
	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"

	// Open method does not create the connection, it simply check if the arguments work properly
	// That's why we must check with Ping() method if it's working!
	db, err := sql.Open("postgres", dsn)
	// Close DB after program exits
	defer db.Close()
	if err != nil {
		log.Fatalf("There is a problem with typed data source name: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("There is a problem with the connection to the database: %v", err)
	}

	fmt.Println("Connected to database successfully")

	// Return the database
	return &PostgreClient{db}
}
