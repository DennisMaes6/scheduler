package schedule_generator

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

const fileName = "demo.db"

func createDB() *sql.DB {
	newDb := false  // TODO: Hier op true zetten om db te initialiseren

	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		file, err := os.Create(fileName) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()
		newDb = true
		log.Printf("New db file created")
	}

	sqliteDatabase, _ := sql.Open("sqlite3", fileName)

	if newDb {
		if err := createTables(sqliteDatabase); err != nil {
			log.Fatal(errors.Wrap(err, "failed initializing tables is db"))
		}
		
		if err := initializeData(sqliteDatabase); err != nil {
			log.Fatal(errors.Wrap(err, "failed initializing data in db"))
		} 
	}

	log.Printf("Database initialized succesfully")
	return sqliteDatabase
}

func createTables(db *sql.DB) error {

	queries := []string{
		`
			PRAGMA foreign_keys
		`,
		`
			CREATE TABLE IF NOT EXISTS model_parameters (
				id INTEGER PRIMARY KEY,
				min_balance INTEGER NOT NULL,
				min_balance_jaev INTEGER NOT NULL

			)
		`,
		`
			CREATE TABLE IF NOT EXISTS shift_type_parameters (
				shift_type TEXT PRIMARY KEY,
				shift_workload REAL,
				shift_coverage REAL,
				max_buffer INTEGER
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS day (
				id INTEGER PRIMARY KEY,
				date TEXT NOT NULL,
				is_holiday BOOL NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS assistant (
				id INTEGER PRIMARY KEY,
				name TEXT NOT NULL,
				type TEXT NOT NULL,
				free_days TEXT NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS schedule (
				id INTEGER PRIMARY_KEY,
				fairness_score REAL NOT NULL,
				balance_score REAL NOT NULL,
				jaev_fairness_score REAL NOT NULL,
				jaev_balance_score REAL NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS individual_schedule (
				assistant_id INTEGER PRIMARY_KEY,
				workload REAL NOT NULL,
				FOREIGN KEY (assistant_id) REFERENCES assistant_instance(id) ON DELETE CASCADE
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS assignment (
				assistant_id INTEGER NOT NULL,
				day_nb INTEGER NOT NULL,
				shift_type TEXT NOT NULL,
				PRIMARY KEY (assistant_id, day_nb),
				FOREIGN KEY (assistant_id) REFERENCES assistant(id) ON DELETE CASCADE,
				FOREIGN KEY (day_nb) REFERENCES day(id) ON DELETE CASCADE
			)
		`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed creating table with query: %s", query))
		}
	}
	
	return nil
}

func initializeData(db *sql.DB) error {

	initModelParametersQuery := `
		INSERT or IGNORE INTO model_parameters(id, min_balance, min_balance_jaev)
		VALUES (1, ?, ?)
	`

	_, err := db.Exec(
		initModelParametersQuery,
		initialModelParameters.MinBalance,
		initialModelParameters.MinBalanceJaev,
	)
	if err != nil {
		return errors.Wrap(err, "failed initailizing minimin balance score")
	}

	initSTPsQuery := `
		INSERT or IGNORE INTO shift_type_parameters(shift_type, shift_workload, shift_coverage, max_buffer)
		VALUES (?, ?, ?, ?)
    `
	for _, stp := range initialModelParameters.ShiftTypeParameters {
		if _, err := db.Exec(initSTPsQuery, stp.ShiftType, stp.ShiftWorkload, stp.ShiftCoverage, stp.MaxBuffer); err != nil {
			return errors.Wrap(err, "failed initailizing shift type parameters")
		}
	}

	initDaysQuery := `
		INSERT or IGNORE INTO day(id, date, is_holiday)
		VALUES (?, ?, ?)
	`

	for _, day := range initialInstanceData.Days {
		if _, err := db.Exec(initDaysQuery, day.Id, dateToString(day.Date), day.IsHoliday); err != nil {
			return errors.Wrap(err, "failed initializing days in database")
		}
	}

	initAIQuery := `
		INSERT or IGNORE INTO assistant(id, name, type, free_days)
		VALUES (?, ?, ?, ?)
	`
	for _, ai := range initialInstanceData.Assistants {
		if _, err := db.Exec(initAIQuery, ai.Id, ai.Name, ai.Type, integerArrayToString(ai.FreeDays)); err != nil {
			return errors.Wrap(err, "failed initializing assistant instance")
		}
	}
	return nil
}
