package schedule_generator

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func createDB() *sql.DB {
	if _, err := os.Stat("sqlite-database-workload-score-model.db"); os.IsNotExist(err) {

		file, err := os.Create("sqlite-database-workload-score-model.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()
		log.Printf("New db file created")
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database-workload-score-model.db")

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
			CREATE TABLE IF NOT EXISTS min_balance_score_jaev (
				id INTEGER PRIMARY KEY,
				score INTEGER NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS shift_type_params (
				shift_type TEXT PRIMARY KEY,
				shift_workload REAL,
				max_buffer INTEGER
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS nb_weeks (
				id INTEGER PRIMARY KEY,
				nb_weeks INTEGER NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS holidays (
				id INTEGER PRIMARY KEY,
				days TEXT NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS assistant_instance (
				id INTEGER PRIMARY KEY,
				name TEXT NOT NULL,
				type TEXT NOT NULL,
				free_days TEXT NOT NULL
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

	initMBSJaevQuery := `
		INSERT or IGNORE INTO min_balance_score_jaev(id, score)
		VALUES (1, ?)
	`

	if _, err := db.Exec(initMBSJaevQuery, initialModelParameters.BalanceMinimunJaev); err != nil {
		return errors.Wrap(err, "failed initailizing minimin balance score")
	}

	initSTPsQuery := `
		INSERT or IGNORE INTO shift_type_params(shift_type, shift_workload, max_buffer)
		VALUES (?, ?, ?)
    `
	for _, stp := range initialModelParameters.ShiftTypeParams {
		if _, err := db.Exec(initSTPsQuery, stp.ShiftType, stp.ShiftWorkload, stp.MaxBuffer); err != nil {
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

	initHolidaysQuery := `
		INSERT or IGNORE INTO holidays(id, days)
		VALUES (1, ?)
	`

	if _, err := db.Exec(initHolidaysQuery, integerArrayToString(initialInstanceData.Holidays)); err != nil {
		return errors.Wrap(err, "failed initializing holidays")
	}

	initAIQuery := `
		INSERT or IGNORE INTO assistant_instance(id, name, type, free_days)
		VALUES (?, ?, ?, ?)
	`
	for _, ai := range initialInstanceData.Assistants {
		if _, err := db.Exec(initAIQuery, ai.Id, ai.Name, ai.Type, integerArrayToString(ai.FreeDays)); err != nil {
			return errors.Wrap(err, "failed initializing assistant instance")
		}
	}
	return nil
}

func integerArrayToString(array []int32) string {
	result := ""
	for _, h := range array {
		result += fmt.Sprintf("%s,", strconv.Itoa(int(h)))
	}
	return result
}
