package schedule_generator

import (
	"fmt"
	"os"
	"sort"

	"github.com/jorensjongers/scheduler/backend/model"
)

func writeData(file *os.File, params model.ModelParameters, data model.InstanceData) error {

	sort.Sort(ById(data.Assistants))
	sort.Sort(ByShiftType(params.ShiftTypeParams))

	skeleton := `
		nb_weeks = %d;
		H = {%s};
		nb_personnel = %d;
		Ids= %s;
		T = %s;
		F = %s;
		
		shift_workload = %s;
		max_buffer = %s;
		min_balance = %d;
	`
	content := fmt.Sprintf(skeleton,
		len(data.Days)/7,
		buildHolidaysString(data.Days),
		len(data.Assistants),
		buildIdsString(data.Assistants),
		buildAssistantTypesString(data.Assistants),
		buildFreeDaysString(data.Assistants),
		buildShiftWorkloadString(params.ShiftTypeParams),
		buildMaxBufferString(params.ShiftTypeParams),
		params.MinBalance,
	)

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil

}

func buildHolidaysString(days []model.Day) string {
	result := "["

	for _, day := range days {
		if day.IsHoliday {
			result += fmt.Sprintf("%d,", day.Id)
		}
	}

	result += "]"

	return result
}

func buildIdsString(ais []model.Assistant) string {
	result := "["

	for _, assistant := range ais {
		result += fmt.Sprintf("%d,", assistant.Id)
	}

	result += "]"

	return result
}

func buildAssistantTypesString(ais []model.Assistant) string {
	result := "["

	for _, assistant := range ais {
		result += fmt.Sprintf("%s,", assistant.Type)
	}

	result += "]"

	return result
}

func buildFreeDaysString(assistants []model.Assistant) string {
	result := "["

	for _, assistant := range assistants {
		result += fmt.Sprintf("{%s},", integerArrayToString(assistant.FreeDays))
	}

	result += "]"

	return result
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

func writeJaevData(file *os.File,
	schedule firstStageSchedule,
	params model.ModelParameters,
	data model.InstanceData) error {

	jas := filterAssistant(model.JA, data.Assistants) // consider only assistants of type JA
	sort.Sort(ById(jas))

	skeleton := `
		nb_weeks = %d;
		H = {%s};
		nb_personnel = %d;
		personnel_id = %s;
		schedule = [|%s];
		F = %s;
		min_balance = %d;
	`

	content := fmt.Sprintf(skeleton,
		len(data.Days)/7,
		buildHolidaysString(data.Days),
		len(jas),
		buildIdsString(jas),
		buildScheduleString(schedule, jas),
		buildFreeDaysString(jas),
		params.MinBalanceJaev,
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

func buildScheduleString(schedule firstStageSchedule, jas []model.Assistant) string {
	result := ""

	for _, is := range schedule.individualSchedules {
		if idIn(is.assistantId, jas) {
			result += "\n\t\t\t"
			for _, assignment := range is.assigments {
				result += fmt.Sprintf("%s, ", string(assignment.shiftType))
			}
			result += " |"
		}
	}

	return result
}

func idIn(id int32, assistants []model.Assistant) bool {

	for _, a := range assistants {
		if a.Id == id {
			return true
		}
	}

	return false
}
