package schedule_generator

import (
	"database/sql"
	"fmt"

	"github.com/jorensjongers/scheduler/backend/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type DbController struct {
	db *sql.DB
}

func newDBController() DbController {
	db := createDB()
	return DbController{db}
}

// GetModelParameters -- returns the model parameters a stored in the DB
func (c DbController) GetModelParameters() (model.ModelParameters, error) {

	balanceScoresQuery := `
		SELECT min_balance, min_balance_jaev
		FROM model_parameters
		WHERE id = 1
	`

	var (
		minBalance     int
		minBalanceJaev int
	)

	if err := c.db.QueryRow(balanceScoresQuery).Scan(&minBalance, &minBalanceJaev); err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting model parameters from db")
	}

	shiftTypeParams, err := c.getShiftTypeParams()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting shift type parameters from db")
	}

	result := model.ModelParameters{
		MinBalance:          int32(minBalance),
		MinBalanceJaev:      int32(minBalanceJaev),
		ShiftTypeParameters: shiftTypeParams,
	}

	return result, nil
}

// SetModelParameters -- updates the model parameters in the DB
func (c DbController) SetModelParameters(params model.ModelParameters) error {

	setBalanceScoresQuery := `
		UPDATE model_parameters
		SET min_balance = ?, min_balance_jaev = ?
		WHERE id = 1
	`

	if _, err := c.db.Exec(setBalanceScoresQuery, params.MinBalance, params.MinBalanceJaev); err != nil {
		return errors.Wrap(err, "failed updating balance scores")
	}

	if err := c.setShiftTypeParams(params.ShiftTypeParameters); err != nil {
		return errors.Wrap(err, "failed setting shift type parameters in database")
	}

	return nil
}

func (c DbController) getShiftTypeParams() ([]model.ShiftTypeModelParameters, error) {

	shiftTypeParamsQuery := `
		SELECT shift_type, shift_workload, max_buffer
		FROM shift_type_parameters
	`

	rows, err := c.db.Query(shiftTypeParamsQuery)
	if err != nil {
		return []model.ShiftTypeModelParameters{}, err
	}
	defer rows.Close()

	result := []model.ShiftTypeModelParameters{}
	for rows.Next() {
		var (
			shiftType     string
			shiftWorkload float32
			maxBuffer     int32
		)

		if err := rows.Scan(&shiftType, &shiftWorkload, &maxBuffer); err != nil {
			return []model.ShiftTypeModelParameters{}, err
		}

		stp := model.ShiftTypeModelParameters{
			ShiftType:     model.ShiftType(shiftType),
			ShiftWorkload: shiftWorkload,
			MaxBuffer:     maxBuffer,
		}

		result = append(result, stp)
	}

	if err := rows.Err(); err != nil {
		return []model.ShiftTypeModelParameters{}, err
	}

	return result, nil
}

func (c DbController) setShiftTypeParams(stps []model.ShiftTypeModelParameters) error {

	deleteRowsQuery := "DELETE FROM shift_type_parameters"

	if _, err := c.db.Exec(deleteRowsQuery); err != nil {
		return errors.Wrap(err, "failed truncating shift type parameters table")
	}

	setSTPsQuery := `
		INSERT INTO shift_type_parameters(shift_type, shift_workload, max_buffer)
		VALUES (?, ?, ?)
	`

	for _, stp := range stps {
		if _, err := c.db.Exec(setSTPsQuery, stp.ShiftType, stp.ShiftWorkload, stp.MaxBuffer); err != nil {
			return errors.Wrap(err, "failed inserting shift type paramers in database")
		}
	}

	return nil

}

// GetInstanceData -- retrieves the instance data from database
func (c DbController) GetInstanceData() (model.InstanceData, error) {

	assitants, err := c.getAssistants()
	if err != nil {
		return model.InstanceData{}, errors.Wrap(err, "failed retreiving assistants from database")
	}

	days, err := c.getDays()
	if err != nil {
		return model.InstanceData{}, errors.Wrap(err, "failed retrieving days from database")
	}

	result := model.InstanceData{
		Assistants: assitants,
		Days:       days,
	}

	return result, nil
}

// SetInstanceData -- updates the instance data in the database
func (c DbController) SetInstanceData(data model.InstanceData) error {

	if err := c.setAssistants(data.Assistants); err != nil {
		return errors.Wrap(err, "failed updating assistants in db")
	}

	if err := c.setDays(data.Days); err != nil {
		return errors.Wrap(err, "failed updating days in db")
	}

	return nil
}

func (c DbController) getAssistants() ([]model.Assistant, error) {

	getAssistantsQuery := `
		SELECT id, name, type, free_days
		FROM assistant
	`

	rows, err := c.db.Query(getAssistantsQuery)
	if err != nil {
		return []model.Assistant{}, errors.Wrap(err, "failed querying database for assistants")
	}
	defer rows.Close()

	assistants := []model.Assistant{}
	for rows.Next() {
		var (
			id          int32
			name        string
			rawType     string
			freeDaysRaw string
		)

		if err := rows.Scan(&id, &name, &rawType, &freeDaysRaw); err != nil {
			return []model.Assistant{}, errors.Wrap(err, "failed scanning assistant row")
		}

		freeDays, err := stringToIntArray(freeDaysRaw)
		if err != nil {
			return []model.Assistant{}, errors.Wrap(err, "failed converting free days string to array")
		}

		assistant := model.Assistant{
			Id:       id,
			Name:     name,
			Type:     model.AssistantType(rawType),
			FreeDays: freeDays,
		}

		assistants = append(assistants, assistant)
	}

	if err := rows.Err(); err != nil {
		return []model.Assistant{}, errors.Wrap(err, "rows error when retrieving assistants")
	}

	return assistants, nil
}

func (c DbController) getDays() ([]model.Day, error) {

	getDaysQuery := `
		SELECT id, date, is_holiday
		FROM day
	`

	rows, err := c.db.Query(getDaysQuery)
	if err != nil {
		return []model.Day{}, errors.Wrap(err, "failed querying database for days")
	}
	defer rows.Close()

	days := []model.Day{}
	for rows.Next() {
		var (
			id        int32
			rawDate   string
			isHoliday bool
		)

		if err := rows.Scan(&id, &rawDate, &isHoliday); err != nil {
			return []model.Day{}, errors.Wrap(err, "failed scanning day row")
		}

		date, err := parseDate(rawDate)
		if err != nil {
			return []model.Day{}, errors.Wrap(err, fmt.Sprintf("failed parsing day: %s", rawDate))
		}

		day := model.Day{
			Id:        id,
			Date:      date,
			IsHoliday: isHoliday,
		}

		days = append(days, day)
	}

	if err := rows.Err(); err != nil {
		return []model.Day{}, errors.Wrap(err, "rows error when retrieving days")
	}

	return days, nil
}

func (c DbController) setAssistants(assistants []model.Assistant) error {

	deleteRowsQuery := "DELETE FROM assistant"
	if _, err := c.db.Exec(deleteRowsQuery); err != nil {
		return errors.Wrap(err, "failed truncating assistant table")
	}

	insertAssistantQuery := `
		INSERT INTO assistant(id, name, type, free_days)
		VALUES (?, ?, ?, ?)
	`
	for _, assistant := range assistants {
		_, err := c.db.Exec(insertAssistantQuery,
			assistant.Id,
			assistant.Name,
			assistant.Type,
			integerArrayToString(assistant.FreeDays))
		if err != nil {
			return errors.Wrap(err, "failed inserting assistant in database")
		}
	}

	return nil
}

func (c DbController) setDays(days []model.Day) error {
	deleteRowsQuery := "DELETE FROM day"
	if _, err := c.db.Exec(deleteRowsQuery); err != nil {
		return errors.Wrap(err, "failed truncating day table")
	}

	insertDayQuery := `
		INSERT INTO day(id, date, is_holiday)
		VALUES (?, ?, ?)
	`
	for _, day := range days {
		if _, err := c.db.Exec(insertDayQuery, day.Id, dateToString(day.Date), day.IsHoliday); err != nil {
			return errors.Wrap(err, "failed inserting day in database")
		}
	}

	return nil
}

func (c DbController) getSchedule() (model.Schedule, error) {

	scheduleQuery := `
		SELECT fairness_score, balance_score, jaev_fairness_score, jaev_balance_score
		FROM schedule
		WHERE id = 1
	`

	var (
		fairnessScore     float32
		balanceScore      int
		jaevFairnessScore float32
		jaevBalanceScore  int
	)

	err := c.db.QueryRow(scheduleQuery).Scan(
		&fairnessScore,
		&balanceScore,
		&jaevFairnessScore,
		&jaevBalanceScore,
	)

	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed getting schedule scores from db")
	}

	individualSchedules, err := c.getIndividualSchedules(balanceScore)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed getting individual schedules from db")
	}

	return model.Schedule{
		FairnessScore:       fairnessScore,
		BalanceScore:        int32(balanceScore),
		JaevFairnessScore:   jaevFairnessScore,
		JaevBalanceScore:    int32(jaevBalanceScore),
		IndividualSchedules: tagIndividualSchedules(individualSchedules, int32(balanceScore)),
	}, nil

}

func (c DbController) getIndividualSchedules(balanceScore int) ([]untaggedIndividualSchedule, error) {
	isQuery := `
		SELECT assistant_id, workload 
		FROM individual_schedule
	`

	rows, err := c.db.Query(isQuery)
	if err != nil {
		return []untaggedIndividualSchedule{}, err
	}

	result := []untaggedIndividualSchedule{}
	for rows.Next() {
		var (
			assistantId int32
			workload    float32
		)

		if err := rows.Scan(&assistantId, &workload); err != nil {
			return []untaggedIndividualSchedule{}, err
		}

		assignments, err := c.getAssignments(assistantId)
		if err != nil {
			return []untaggedIndividualSchedule{},
				errors.Wrap(err, fmt.Sprintf("failed getting assignments for assistant with id %d", assistantId))
		}

		is := untaggedIndividualSchedule{
			assistantId: assistantId,
			workload:    workload,
			assignments: assignments,
		}
		result = append(result, is)
	}

	if rows.Err() != nil {
		return []untaggedIndividualSchedule{}, err
	}

	return result, nil
}

func (c DbController) getAssignments(assistantId int32) ([]untaggedAssignment, error) {

	assignmentsQuery := `
		SELECT day_nb, shift_type
		FROM assignment
		WHERE assistant_id = ?
		ORDER BY day_nb ASC
	`

	rows, err := c.db.Query(assignmentsQuery, assistantId)
	if err != nil {
		return []untaggedAssignment{}, err
	}

	result := []untaggedAssignment{}
	for rows.Next() {
		var (
			dayNb        int32
			rawShiftType string
		)

		if err := rows.Scan(&dayNb, &rawShiftType); err != nil {
			return []untaggedAssignment{}, errors.Wrap(err, "failed scanning assignment")
		}

		shiftType, err := parseShiftType(rawShiftType)
		if err != nil {
			return []untaggedAssignment{}, errors.Wrap(err, fmt.Sprintf("failed parsing shift type: %s", rawShiftType))
		}

		assignment := untaggedAssignment{
			dayNb:     dayNb,
			shiftType: shiftType,
		}

		result = append(result, assignment)
	}

	if rows.Err() != nil {
		return []untaggedAssignment{}, err
	}

	return result, nil

}
