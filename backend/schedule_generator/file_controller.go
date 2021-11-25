package schedule_generator

import (
	"fmt"
	"os"
	"sort"

	"github.com/DennisMaes6/scheduler/backend/model"
)

func writeData(file *os.File, params model.ModelParameters, data model.InstanceData) error {

	sort.Sort(ByShiftType(params.ShiftTypeParameters))

	skeleton := `
		nb_weeks = %d;
		H = %s;
		nb_holidays = %d;
		nb_personnel = %d;
		personnel_id = %s;
		T = %s;
		F = %s;
		
		shift_workload = %s;
		max_buffer = %s;
		min_balance = %d;
	`
	content := fmt.Sprintf(skeleton,
		len(data.Days)/7,
		buildHolidaysString(data.Days),
		nbHolidays(data.Days),
		len(data.Assistants),
		buildIdsString(data.Assistants),
		buildAssistantTypesString(data.Assistants),
		buildFreeDaysString(data.Assistants, data.Days),
		buildShiftWorkloadString(params.ShiftTypeParameters),
		buildMaxBufferString(params.ShiftTypeParameters),
		params.MinBalance,
	)

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil

}

func nbHolidays(days []model.Day) int {
	result := 0
	for _, day := range days {
		if day.IsHoliday {
			result += 1
		}
	}
	return result
}

func buildHolidaysString(days []model.Day) string {
	result := "{"

	for _, day := range days {
		if day.IsHoliday {
			result += fmt.Sprintf("%d,", day.Id)
		}
	}

	result += "}"

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

func buildFreeDaysString(assistants []model.Assistant, days []model.Day) string {
	result := "["

	for _, assistant := range assistants {
		freeDays := union(assistant.FreeDays, days)
		result += fmt.Sprintf("{%s},", integerArrayToString(freeDays))
	}

	result += "]"

	return result
}

func union(daysIds []int32, days []model.Day) []int32 {
	result := []int32{}

	for _, dayId := range daysIds {
		if intInArray(dayId, daysToIds(days)) {
			result = append(result, dayId)
		}
	}
	return result
}

func intInArray(i int32, array []int32) bool {
	for _, b := range array {
		if b == i {
			return true
		}
	}
	return false
}

func daysToIds(days []model.Day) []int32 {
	result := []int32{}
	for _, day := range days {
		result = append(result, day.Id)
	}
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

	skeleton := `
		nb_weeks = %d;
		H = %s;
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
		buildFreeDaysString(jas, data.Days),
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
			for _, assignment := range is.assignments {
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
