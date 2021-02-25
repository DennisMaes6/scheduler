package schedule_generator

import (
	"fmt"
	"os"

	"github.com/jorensjongers/scheduler/backend/model"
)

func writeInstanceSpecificData(file *os.File, data model.InstanceData) error {

	skeleton := `
		nb_weeks = %d;
		nb_personnel = %d;
		T = %s;
		F = %s;
		S = {JANW, SAEW, JAWH, SAWH, TSPT, CALL};
	`
	content := fmt.Sprintf(skeleton,
		data.NbWeeks,
		len(data.Assistants),
		buildAssistantTypesString(data.Assistants),
		buildFreeDaysString(len(data.Assistants)),
	)

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil

}

func buildAssistantTypesString(ais []model.AssistantInstance) string {
	result := "["

	for _, ai := range ais {
		result += fmt.Sprintf("%s,", (ai.Type))
	}

	result += "]"

	return result
}

func buildFreeDaysString(nbAssistants int) string {
	result := "["

	for i := 0; i < nbAssistants; i++ {
		result += "{},"
	}

	result += "]"

	return result
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

	result += "0.0]"

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
