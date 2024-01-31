package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	operatorName := "go-postgresql"
	logTime := time.Now()
	operatorLog := "we have created a log for you to store and review later."

	if err := insertDataIntoPostgres(operatorName, operatorLog, logTime); err != nil {
		fmt.Printf("Error inserting data: %s", err)
	}

}

func insertDataIntoPostgres(operatorName, operatorLog string, logTime time.Time) error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	tableName := os.Getenv("DB_TABLE")

	// Connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Insert data into the table
	insertStmt := fmt.Sprintf("INSERT INTO %s(operator_name, log_time, log_file) VALUES ($1, $2, $3)", tableName)
	_, err = db.Exec(insertStmt, operatorName, logTime, operatorLog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")

	return nil
}
