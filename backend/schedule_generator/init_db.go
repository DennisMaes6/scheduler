package schedule_generator

import (
	"database/sql"
	"log"
	"os"

	"github.com/pkg/errors"
)

func createDB() *sql.DB {
	if _, err := os.Stat("sqlite-database.db"); os.IsNotExist(err) {

		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()
		log.Printf("New db file created")
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")

	if err := createTables(sqliteDatabase); err != nil {
		log.Fatal(errors.Wrap(err, "failed initializing tables is db"))
	}

	if err := initializeData(sqliteDatabase); err != nil {
		log.Fatal(errors.Wrap(err, "failed initializing data in db"))
	}

	log.Printf("Database initialized succesfully")
	return sqliteDatabase
}

func createTables(db *sql.DB) error {

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
		`
			CREATE TABLE IF NOT EXISTS nb_weeks (
				id INTEGER PRIMARY KEY,
				nb_weeks INTEGER NOT NULL
			)
		`,
		`
		CREATE TABLE IF NOT EXISTS assistant_instance (
			id INTEGER PRIMARY KEY,
			type TEXT NOT NULL
		)
	`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return errors.Wrap(err, "failed creating creating table")
		}
	}

	return nil
}

func initializeData(db *sql.DB) error {

	initMBSQuery := `
		INSERT or IGNORE INTO min_balance_score(id, score)
		VALUES (1, ?)
	`

	if _, err := db.Exec(initMBSQuery, initialModelParameters.BalanceMinimum); err != nil {
		return errors.Wrap(err, "failed initailizing minimin balance score")
	}

	initSTPsQuery := `
		INSERT or IGNORE INTO shift_type_params(shift_type, fairness_weight, included_in_balance)
		VALUES (?, ?, ?)
    `
	for _, stp := range initialModelParameters.ShiftTypeParams {
		if _, err := db.Exec(initSTPsQuery, stp.ShiftType, stp.FairnessWeight, stp.IncludedInBalance); err != nil {
			return errors.Wrap(err, "failed initailizing shift type parameters")
		}
	}

	initNbWeeksQuery := `
		INSERT or IGNORE INTO nb_weeks(id, nb_weeks)
		VALUES (1, ?)
	`

	if _, err := db.Exec(initNbWeeksQuery, initialInstanceData.NbWeeks); err != nil {
		return errors.Wrap(err, "failed initializing number of weeks")
	}

	initAIQuery := `
		INSERT or IGNORE INTO assistant_instance(id, type)
		VALUES (?, ?)
	`
	for _, ai := range initialInstanceData.Assistants {
		if _, err := db.Exec(initAIQuery, ai.Id, ai.Type); err != nil {
			return errors.Wrap(err, "failed initializing assistant instance")
		}
	}
	return nil
}
