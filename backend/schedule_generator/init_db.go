package schedule_generator

import (
	"database/sql"
	"log"
	"os"
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

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	createTables(sqliteDatabase)
	initializeData(sqliteDatabase)

	log.Printf("Database initialized")
	return sqliteDatabase
}

func createTables(db *sql.DB) {

	queries := []string{
		`
			CREATE TABLE IF NOT EXISTS min_balance_score (
				id INTEGER PRIMARY KEY,
				score INTEGER NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS shift_type_params (
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

	initializeMinBalanceScoreSQL := `
		INSERT or IGNORE INTO min_balance_score(id, score)
		VALUES (1, ?)
	`

	statement, err := db.Prepare(initializeMinBalanceScoreSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := statement.Exec(initialModelParameters.BalanceMinimum); err != nil {
		log.Fatal(err.Error())
	}

	for _, stp := range initialModelParameters.ShiftTypeParams {
		initializeShiftTypeParamsSQL := `
			INSERT or IGNORE INTO shift_type_params(shift_type, fairness_weight, included_in_balance)
			VALUES (?, ?, ?)
		`

		statement, err := db.Prepare(initializeShiftTypeParamsSQL)
		if err != nil {
			log.Fatal(err.Error())
		}

		if _, err := statement.Exec(stp.ShiftType, stp.FairnessWeight, stp.IncludedInBalance); err != nil {
			log.Fatal(err.Error())
		}
	}
}
