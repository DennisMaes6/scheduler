package schedule_generator

import (
	"fmt"
	"os"

	"github.com/jorensjongers/scheduler/backend/model"
)

func writeInstanceSpecificData(file *os.File) error {

	content := `
		nb_weeks = 4;
		nb_personnel = 18;
		T = [JA, JA, JA, JA, JA, JA, JA, JA_F, JA_F, SA, SA, SA, SA_F, SA_F, SA_NEO, SA_NEO, SA_F_NEO, SA_F_NEO];
		F = [{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}];
		S = {JANW, SAEW, JAEV, JAWH, SAWH, TSPT, CALL};
	`
	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil

}

func writeModelParameters(file *os.File, params model.ModelParameters) error {

	skeleton := `
		fairness_weight = %s;
		S_balance =  %s;
		min_balance = %d;
	`
	content := fmt.Sprintf(skeleton,
		buildFairnessWeightsString(params.ShiftTypeParams),
		buildBalanceShiftsString(params.ShiftTypeParams),
		params.BalanceMinimum)

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil

}

func buildFairnessWeightsString(stps []model.ShiftTypeModelParameters) string {
	result := "["

	for _, stp := range stps {
		result += fmt.Sprintf("%f,", stp.FairnessWeight)
	}

	result += "]"

	return result

}

func buildBalanceShiftsString(stps []model.ShiftTypeModelParameters) string {
	result := "{"

	for _, stp := range stps {
		if stp.IncludedInBalance {
			result += fmt.Sprintf("%s,", (stp.ShiftType))
		}
	}

	result += "}"

	return result
}
