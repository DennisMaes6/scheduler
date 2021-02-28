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
		shift_workload = %s;
		max_buffer = %s;
		min_balance = %d;
	`
	content := fmt.Sprintf(skeleton,
		buildShiftWorkloadString(params.ShiftTypeParams),
		buildMaxBufferString(params.ShiftTypeParams),
		params.BalanceMinimum)

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil

}

func buildShiftWorkloadString(stps []model.ShiftTypeModelParameters) string {
	result := "["

	for _, stp := range stps {
		result += fmt.Sprintf("%f,", stp.ShiftWorkload)
	}

	result += "0.0]"

	return result

}

func buildMaxBufferString(stps []model.ShiftTypeModelParameters) string {
	result := "["

	for _, stp := range stps {
		result += fmt.Sprintf("%d,", stp.MaxBuffer)
	}

	result += "0]"

	return result
}
