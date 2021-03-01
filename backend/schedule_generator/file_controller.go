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

func writeInstanceSpecificDataJaev(file *os.File, schedule model.Schedule) error {

	skeleton := `
		nb_weeks = %d;
		nb_personnel = %d;
		schedule = [|%s];
		F = %s;
	`

	content := fmt.Sprintf(skeleton,
		(schedule.NbDays / 7),
		len(filterAssistant(model.JA, schedule.Assistants)),
		buildScheduleString(schedule),
		buildFreeDaysString(len(filterAssistant(model.JA, schedule.Assistants))),
	)

	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}

func filterAssistant(at model.AssistantType, assistants []model.Assistant) []model.Assistant {
	result := []model.Assistant{}

	for _, a := range assistants {
		if a.Type == at {
			result = append(result, a)
		}
	}

	return result
}

func buildScheduleString(schedule model.Schedule) string {

	result := ""
	assistants := filterAssistant(model.JA, schedule.Assistants)

	for _, is := range schedule.IndividualSchedules {
		if idMatches(is.AssistantId, assistants) {
			result += "\n"
			for _, assignment := range is.Assignments {
				result += fmt.Sprintf("%s, ", string(assignment.ShiftType))
			}
			result += " |"
		}
	}

	return result
}

func idMatches(id int32, assistants []model.Assistant) bool {

	for _, a := range assistants {
		if a.Id == id {
			return true
		}
	}

	return false
}

func writeModelParametersJaev(file *os.File, params model.ModelParameters) error {

	if _, err := file.WriteString(fmt.Sprintf("min_balance = %d;", params.BalanceMinimunJaev)); err != nil {
		return err
	}
	return nil
}
