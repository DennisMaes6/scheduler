package schedule_generator

import (
	"fmt"
	"os"

	"github.com/jorensjongers/scheduler/backend/model"
)

func writeInstanceSpecificData(file *os.File, data model.InstanceData) error {

	skeleton := `
		nb_weeks = %d;
		H = {%s};
		nb_personnel = %d;
		T = %s;
		F = %s;
		S = {JANW, SAEW, JAWH, SAWH, TSPT, CALL};
	`
	content := fmt.Sprintf(skeleton,
		data.NbWeeks,
		integerArrayToString(data.Holidays),
		len(data.Assistants),
		buildAssistantTypesString(data.Assistants),
		buildFreeDaysString(data.Assistants),
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

func buildFreeDaysString(assistants []model.AssistantInstance) string {
	result := "["

	for _, assistant := range assistants {
		result += fmt.Sprintf("{%s},", integerArrayToString(assistant.FreeDays))
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

func writeInstanceSpecificDataJaev(file *os.File, schedule model.Schedule, data model.InstanceData) error {

	skeleton := `
		nb_weeks = %d;
		H = {%s};
		nb_personnel = %d;
		schedule = [|%s];
		F = %s;
	`

	content := fmt.Sprintf(skeleton,
		(schedule.NbDays / 7),
		integerArrayToString(schedule.Holidays),
		len(filterAssistant(model.JA, data.Assistants)),
		buildScheduleString(schedule, data.Assistants),
		buildFreeDaysString(filterAssistant(model.JA, data.Assistants)),
	)

	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}

func filterAssistant(at model.AssistantType, assistants []model.AssistantInstance) []model.AssistantInstance {
	result := []model.AssistantInstance{}

	for _, a := range assistants {
		if a.Type == at {
			result = append(result, a)
		}
	}

	return result
}

func buildScheduleString(schedule model.Schedule, ais []model.AssistantInstance) string {

	result := ""
	assistants := filterAssistant(model.JA, ais)

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

func idMatches(id int32, assistants []model.AssistantInstance) bool {

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
