package schedule_generator

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createDB() *sql.DB {
	if _, err := os.Stat("sqlite-database.db"); os.IsNotExist(err) {

		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()
		log.Println("sqlite-database.db file created")
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File                   // Defer Closing the database
	createTables(sqliteDatabase)

	return sqliteDatabase // Create Database Tables
}

type DbController struct {
	db *sql.DB
}

func newDBController() DbController {
	db := createDB()
	return DbController{db}
}

func createTables(db *sql.DB) {
	print("here create tables \n")
}
