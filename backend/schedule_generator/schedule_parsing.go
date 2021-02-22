package schedule_generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jorensjongers/scheduler/backend/model"
	"github.com/pkg/errors"
)

func parseSchedule(scheduleStr string) (model.Schedule, error) {

	nbDays, err := extractNbDays(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting nb days")
	}

	assistants, err := extractAssistants(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting assistants")
	}

	shiftTypes, err := extractShiftTypes(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting shift types")
	}

	individualSchedules, err := extractIndividualSchedules(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting individual schedules")
	}

	result := model.Schedule{
		NbDays:              nbDays,
		Assistants:          assistants,
		ShiftTypes:          shiftTypes,
		IndividualSchedules: individualSchedules,
	}

	return result, nil
}

func extractNbDays(scheduleStr string) (int32, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "nb_days")
	nbDays, err := strconv.Atoi(strings.Split(lines[0], ":")[1])
	if err != nil {
		return 0, err
	}

	return int32(nbDays), nil
}

func extractAssistants(scheduleStr string) ([]model.Assistant, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "assistant")

	assistants := []model.Assistant{}
	for _, line := range lines {
		split_line := strings.Split(line, " ")

		id, err := strconv.Atoi(strings.Split(split_line[0], ":")[1])
		if err != nil {
			return []model.Assistant{}, err
		}

		assistantType, err := parseAssistantType(strings.Split(split_line[1], ":")[1])
		if err != nil {
			return []model.Assistant{}, err
		}

		assistant := model.Assistant{
			Id:   int32(id),
			Type: assistantType,
		}

		assistants = append(assistants, assistant)
	}

	return assistants, nil
}

func parseAssistantType(typeStr string) (model.AssistantType, error) {
	switch typeStr {
	case "JA":
		return model.JA, nil
	case "JA_F":
		return model.JA_F, nil
	case "SA":
		return model.SA, nil
	case "SA_F":
		return model.SA_F, nil
	case "SA_NEO":
		return model.SA_NEO, nil
	case "SA_F_NEO":
		return model.SA_F_NEO, nil
	default:
		return "", fmt.Errorf("invalid assistant type string: %s", typeStr)
	}
}

func extractShiftTypes(scheduleStr string) ([]model.ShiftType, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "shift_types")
	shiftTypesLine := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")

	shiftTypes := []model.ShiftType{}
	for i := 0; i < len(shiftTypesLine); i++ {
		shiftType, err := parseShiftType(shiftTypesLine[i])
		if err != nil {
			return []model.ShiftType{}, err
		}
		shiftTypes = append(shiftTypes, shiftType)
	}

	return shiftTypes, nil
}

func parseShiftType(shiftTypeStr string) (model.ShiftType, error) {

	switch shiftTypeStr {
	case "JAEV":
		return model.JAEV, nil
	case "JANW":
		return model.JANW, nil
	case "JAWH":
		return model.JAWH, nil
	case "SAEW":
		return model.SAEW, nil
	case "SAWH":
		return model.SAWH, nil
	case "TSPT":
		return model.TSPT, nil
	case "CALL":
		return model.CALL, nil
	case "FREE":
		return model.FREE, nil
	default:
		return "", fmt.Errorf("invalid shift type string: %s", shiftTypeStr)
	}
}

func extractIndividualSchedules(scheduleStr string) ([]model.IndividualSchedule, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "assistant")

	result := []model.IndividualSchedule{}
	for _, line := range lines {

		entries := strings.Split(strings.TrimSpace(line), " ")
		assistantId, err := strconv.Atoi(strings.Split(entries[0], ":")[1])
		if err != nil {
			return []model.IndividualSchedule{}, err
		}

		assignments := []model.ShiftType{}
		for i := 2; i < len(entries); i++ {
			assignment, err := parseShiftType(entries[i])
			if err != nil {
				return []model.IndividualSchedule{}, err
			}
			assignments = append(assignments, assignment)
		}

		minBalance, err := extractMinBalance(scheduleStr)
		if err != nil {
			return []model.IndividualSchedule{}, err
		}

		is := model.IndividualSchedule{
			AssistantId: int32(assistantId),
			Assignments: tagAssignments(assignments, minBalance),
		}

		result = append(result, is)
	}

	return result, nil
}

func filterLines(lines []string, substr string) []string {
	filtered_lines := []string{}
	for _, line := range lines {
		if strings.Contains(line, substr) {
			filtered_lines = append(filtered_lines, line)
		}
	}

	return filtered_lines
}

func extractMinBalance(scheduleStr string) (int, error) {
	minBalanceLines := filterLines(strings.Split(scheduleStr, "\n"), "min_balance")
	return strconv.Atoi(strings.Split(minBalanceLines[0], ":")[1])
}

func tagAssignments(assignments []model.ShiftType, minBalance int) []model.Assignment {

	result := []model.Assignment{}

	firstShiftDone := false
	lastShiftDone := false
	firstDay := 0
	count := 0

	for i, shiftType := range assignments {
		result = append(result, model.Assignment{
			ShiftType:        shiftType,
			PartOfMinBalance: false,
		})

		if firstShiftDone && !lastShiftDone {
			if shiftType == model.FREE {
				if count == 0 {
					firstDay = i
				}
				count++
			} else if assignments[i-1] == model.FREE && shiftType != model.FREE {
				if count == minBalance {
					for j := firstDay; j < i; j++ {
						result[j] = model.Assignment{
							ShiftType:        assignments[j],
							PartOfMinBalance: true,
						}
					}
				}
				count = 0
			}
		}

		if !firstShiftDone && shiftType != model.FREE {
			firstShiftDone = true
		}

		if !lastShiftDone && allEqual(assignments[i:], model.FREE) {
			lastShiftDone = true
		}
	}
	return result
}

func allEqual(shiftTypes []model.ShiftType, condition model.ShiftType) bool {
	for _, shiftType := range shiftTypes {
		if shiftType != condition {
			return false
		}
	}
	return true
}
