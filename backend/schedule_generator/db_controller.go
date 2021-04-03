package schedule_generator

import (
	"database/sql"
	"strconv"
	"strings"

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

	balanceMinimum, err := c.getMinBalanceScore()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting min balance score from db")
	}

	balanceMinimumJaev, err := c.getMinBalanceScoreJaev()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting min balance score jaev from db")
	}

	shiftTypeParams, err := c.getShiftTypeParams()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting shift type parameters from db")
	}

	result := model.ModelParameters{
		BalanceMinimum:     int32(balanceMinimum),
		BalanceMinimunJaev: int32(balanceMinimumJaev),
		ShiftTypeParams:    shiftTypeParams,
	}

	return result, nil
}

// SetModelParameters -- updates the model parameters in teh DB
func (c DbController) SetModelParameters(params model.ModelParameters) error {

	if err := c.setMinBalanceScore(params.BalanceMinimum); err != nil {
		return errors.Wrap(err, "failed setting min balance score in db")
	}

	if err := c.setMinBalanceScoreJaev(params.BalanceMinimunJaev); err != nil {
		return errors.Wrap(err, "failed setting min balance score in db")
	}

	for _, stps := range params.ShiftTypeParams {
		if err := c.setShiftTypeParams(stps); err != nil {
			return errors.Wrap(err, "failed setting shift type parameters in db")
		}
	}

	return nil
}

func (c DbController) getMinBalanceScore() (int, error) {

	minBalanceScoreQuery := `
		SELECT score
		FROM min_balance_score
		WHERE id = 1
	`

	var score int
	if err := c.db.QueryRow(minBalanceScoreQuery).Scan(&score); err != nil {
		return 0, err
	}

	return score, nil
}

func (c DbController) getMinBalanceScoreJaev() (int, error) {

	minBalanceScoreJaevQuery := `
		SELECT score
		FROM min_balance_score_jaev
		WHERE id = 1
	`

	var score int
	if err := c.db.QueryRow(minBalanceScoreJaevQuery).Scan(&score); err != nil {
		return 0, err
	}

	return score, nil
}

func (c DbController) getShiftTypeParams() ([]model.ShiftTypeModelParameters, error) {

	shiftTypeParamsQuery := `
		SELECT shift_type, shift_workload, max_buffer
		FROM shift_type_params
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

func (c DbController) setMinBalanceScore(newScore int32) error {

	setMBSQuery := `
		UPDATE min_balance_score
		SET score = ?
		WHERE id = 1
	`

	if _, err := c.db.Exec(setMBSQuery, newScore); err != nil {
		return err
	}

	return nil
}

func (c DbController) setMinBalanceScoreJaev(newScore int32) error {

	setMBSJaevQuery := `
		UPDATE min_balance_score_jaev
		SET score = ?
		WHERE id = 1
	`

	if _, err := c.db.Exec(setMBSJaevQuery, newScore); err != nil {
		return err
	}

	return nil
}

func (c DbController) setShiftTypeParams(stp model.ShiftTypeModelParameters) error {

	setSTPsQuery := `
		UPDATE shift_type_params
		SET shift_workload = ?,
		    max_buffer = ?
		WHERE shift_type = ?
	`

	if _, err := c.db.Exec(setSTPsQuery, stp.ShiftWorkload, stp.MaxBuffer, stp.ShiftType); err != nil {
		return err
	}

	return nil

}

func (c DbController) getInstanceData() (model.InstanceData, error) {
	nbWeeksQuery := `
		SELECT nb_weeks
		FROM nb_weeks
		WHERE id = 1
	`

	var nbWeeks int
	if err := c.db.QueryRow(nbWeeksQuery).Scan(&nbWeeks); err != nil {
		return model.InstanceData{}, errors.Wrap(err, "failed getting nb weeks from db")
	}

	holidaysQuery := `
		SELECT days
		FROM holidays
		WHERE id = 1
	`

	var rawHolidays string
	if err := c.db.QueryRow(holidaysQuery).Scan(&rawHolidays); err != nil {
		return model.InstanceData{}, errors.Wrap(err, "failed getting holidays from db")
	}

	holidays, err := stringToIntArray(rawHolidays)
	if err != nil {
		return model.InstanceData{}, errors.Wrap(err, "failed converting holiday string to array")
	}

	assistantInstanceQuery := `
		SELECT id, name, type, free_days
		FROM assistant_instance
	`

	rows, err := c.db.Query(assistantInstanceQuery)
	if err != nil {
		return model.InstanceData{}, errors.Wrap(err, "query error")
	}
	defer rows.Close()

	ais := []model.AssistantInstance{}
	for rows.Next() {
		var (
			id          int32
			name        string
			rawType     string
			freeDaysRaw string
		)

		if err := rows.Scan(&id, &name, &rawType, &freeDaysRaw); err != nil {
			return model.InstanceData{}, errors.Wrap(err, "scan error")
		}

		freeDays, err := stringToIntArray(freeDaysRaw)
		if err != nil {
			return model.InstanceData{}, errors.Wrap(err, "failed converting free days string to array")
		}

		ai := model.AssistantInstance{
			Id:       id,
			Name:     name,
			Type:     model.AssistantType(rawType),
			FreeDays: freeDays,
		}

		ais = append(ais, ai)
	}

	if err := rows.Err(); err != nil {
		return model.InstanceData{}, errors.Wrap(err, "rows error")
	}

	result := model.InstanceData{
		NbWeeks:    int32(nbWeeks),
		Holidays:   holidays,
		Assistants: ais,
	}

	return result, nil

}

func (c DbController) SetInstanceData(data model.InstanceData) error {

	if err := c.setNbWeeks(data.NbWeeks); err != nil {
		return errors.Wrap(err, "failed setting nb weeks in db")
	}

	if err := c.setHolidays(data.Holidays); err != nil {
		return errors.Wrap(err, "failed setting holidays in db")
	}

	if err := c.setAssistantInstances(data.Assistants); err != nil {
		return errors.Wrap(err, "failed setting assistant instances in db")
	}

	return nil
}

func (c DbController) setNbWeeks(nbWeeks int32) error {

	setMBSQuery := `
		UPDATE nb_weeks
		SET nb_weeks = ?
		WHERE id = 1
	`

	if _, err := c.db.Exec(setMBSQuery, nbWeeks); err != nil {
		return err
	}

	return nil
}

func (c DbController) setHolidays(holidays []int32) error {

	setHolidaysQuery := `
		UPDATE holidays
		SET days = ?
		WHERE id = 1
	`

	if _, err := c.db.Exec(setHolidaysQuery, integerArrayToString(holidays)); err != nil {
		return err
	}

	return nil
}

func (c DbController) setAssistantInstances(ais []model.AssistantInstance) error {
	deleteRowsQuery := "DELETE FROM assistant_instance"

	if _, err := c.db.Exec(deleteRowsQuery); err != nil {
		return errors.Wrap(err, "failed truncating assistant instance table")
	}

	setAIQuery := `
		INSERT INTO assistant_instance(id, name, type, free_days)
		VALUES (?, ?, ?, ?)
	`
	for _, ai := range ais {
		if _, err := c.db.Exec(setAIQuery, ai.Id, ai.Name, ai.Type, integerArrayToString(ai.FreeDays)); err != nil {
			return errors.Wrap(err, "failed updating assistant instance")
		}
	}

	return nil
}

func stringToIntArray(str string) ([]int32, error) {
	if str == "" {
		return []int32{}, nil
	}
	splitStr := strings.Split(strings.TrimSuffix(str, ","), ",")
	result := []int32{}
	for _, s := range splitStr {
		i, err := strconv.Atoi(s)
		if err != nil {
			return []int32{}, err
		}
		result = append(result, int32(i))
	}

	return result, nil
}
