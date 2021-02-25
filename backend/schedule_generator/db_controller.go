package schedule_generator

import (
	"database/sql"

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

	shiftTypeParams, err := c.getShiftTypeParams()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting shift type parameters from db")
	}

	result := model.ModelParameters{
		BalanceMinimum:  int32(balanceMinimum),
		ShiftTypeParams: shiftTypeParams,
	}

	return result, nil
}

// SetModelParameters -- updates the model parameters in teh DB
func (c DbController) SetModelParameters(params model.ModelParameters) error {

	if err := c.setMinBalanceScore(params.BalanceMinimum); err != nil {
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

func (c DbController) getShiftTypeParams() ([]model.ShiftTypeModelParameters, error) {

	shiftTypeParamsQuery := `
		SELECT shift_type, fairness_weight, included_in_balance
		FROM shift_type_params
	`

	rows, err := c.db.Query(shiftTypeParamsQuery)
	defer rows.Close()
	if err != nil {
		return []model.ShiftTypeModelParameters{}, err
	}

	result := []model.ShiftTypeModelParameters{}
	for rows.Next() {
		var (
			shiftType         string
			fairnessWeight    float32
			includedInBalance bool
		)

		if err := rows.Scan(&shiftType, &fairnessWeight, &includedInBalance); err != nil {
			return []model.ShiftTypeModelParameters{}, err
		}

		stp := model.ShiftTypeModelParameters{
			ShiftType:         model.ShiftType(shiftType),
			FairnessWeight:    fairnessWeight,
			IncludedInBalance: includedInBalance,
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

func (c DbController) setShiftTypeParams(stp model.ShiftTypeModelParameters) error {

	setSTPsQuery := `
		UPDATE shift_type_params
		SET fairness_weight = ?,
		    included_in_balance = ?
		WHERE shift_type = ?
	`

	if _, err := c.db.Exec(setSTPsQuery, stp.FairnessWeight, stp.IncludedInBalance, stp.ShiftType); err != nil {
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

	assistantInstanceQuery := `
		SELECT id, type
		FROM assistant_instance
	`

	rows, err := c.db.Query(assistantInstanceQuery)
	defer rows.Close()
	if err != nil {
		return model.InstanceData{}, errors.Wrap(err, "query error")
	}

	ais := []model.AssistantInstance{}
	for rows.Next() {
		var (
			id      int32
			rawType string
		)

		if err := rows.Scan(&id, &rawType); err != nil {
			return model.InstanceData{}, errors.Wrap(err, "scan error")
		}

		ai := model.AssistantInstance{
			Id:   id,
			Type: model.AssistantType(rawType),
		}

		ais = append(ais, ai)
	}

	if err := rows.Err(); err != nil {
		return model.InstanceData{}, errors.Wrap(err, "rows error")
	}

	result := model.InstanceData{
		NbWeeks:    int32(nbWeeks),
		Assistants: ais,
	}

	return result, nil

}
