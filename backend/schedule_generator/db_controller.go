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

	initializeData(sqliteDatabase)

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

	queries := []string{
		`
			CREATE TABLE IF NOT EXISTS min_balance_score (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				score INTEGER NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS modelshift_type_params (
				shift_type TEXT PRIMARY KEY,
				fairness_weight REAL,
				included_in_balance BOOL
			)
		`,
	}

	for _, query := range queries {
		statement, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		statement.Exec()
	}
}

func initializeData(db *sql.DB) {

}
