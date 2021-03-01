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

	fairnessScore, err := extractFairnessScore(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting fairness score")
	}

	shiftTypes, err := extractShiftTypes(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting shift types")
	}

	balanceScore, err := extractMinBalance(scheduleStr)
	if err != nil {
		return model.Schedule{}, err
	}

	individualSchedules, err := extractIndividualSchedules(scheduleStr, balanceScore, shiftTypes)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting individual schedules")
	}

	assistants, err := extractAssistants(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting assistants")
	}

	result := model.Schedule{
		FairnessScore:       fairnessScore,
		BalanceScore:        int32(balanceScore),
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

		workload, err := strconv.ParseFloat(strings.Split(split_line[2], ":")[1], 32)
		if err != nil {
			return []model.Assistant{}, err
		}

		assistant := model.Assistant{
			Id:       int32(id),
			Type:     assistantType,
			Workload: float32(workload),
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
	shiftTypesLine := strings.TrimSpace(strings.Split(lines[0], ":")[1])

	result := []model.ShiftType{}
	for _, shiftTypeStr := range strings.Split(shiftTypesLine, " ") {
		shiftType, err := parseShiftType(strings.Split(shiftTypeStr, "=")[0])
		if err != nil {
			return []model.ShiftType{}, err
		}

		result = append(result, shiftType)
	}
	return result, nil
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

func extractIndividualSchedules(scheduleStr string, balanceScore int, shiftTypes []model.ShiftType) ([]model.IndividualSchedule, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "assistant")

	result := []model.IndividualSchedule{}
	for _, line := range lines {

		entries := strings.Split(strings.TrimSpace(line), " ")
		assistantId, err := strconv.Atoi(strings.Split(entries[0], ":")[1])
		if err != nil {
			return []model.IndividualSchedule{}, err
		}

		assignments := []model.ShiftType{}
		for i := 3; i < len(entries); i++ {
			assignment, err := parseShiftType(entries[i])
			if err != nil {
				return []model.IndividualSchedule{}, err
			}
			assignments = append(assignments, assignment)
		}

		is := model.IndividualSchedule{
			AssistantId: int32(assistantId),
			Assignments: tagAssignments(assignments, balanceScore, shiftTypes),
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

func tagAssignments(assignments []model.ShiftType, minBalance int, balanceShifts []model.ShiftType) []model.Assignment {

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
			if !oneOf(shiftType, balanceShifts) {
				if count == 0 {
					firstDay = i
				}
				count++
			} else if assignments[i-1] == model.FREE && oneOf(shiftType, balanceShifts) {
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

		if !lastShiftDone && allNotOfOneOf(assignments[i:], balanceShifts) {
			lastShiftDone = true
		}
	}
	return result
}

func allNotOfOneOf(shiftTypes []model.ShiftType, test []model.ShiftType) bool {
	for _, shiftType := range shiftTypes {
		if oneOf(shiftType, test) {
			return false
		}
	}
	return true
}

func oneOf(test model.ShiftType, sts []model.ShiftType) bool {
	for _, st := range sts {
		if st == test {
			return true
		}
	}
	return false
}

func extractFairnessScore(scheduleStr string) (float32, error) {
	fairnessLines := filterLines(strings.Split(scheduleStr, "\n"), "fairness_score")
	rawFloat, err := strconv.ParseFloat(strings.Split(fairnessLines[0], ":")[1], 32)
	if err != nil {
		return 0.0, errors.Wrap(err, "failed parsing fairness score")
	}

	return float32(rawFloat), nil
}

func parseAndCombineSchedule(jaevResStr string, res model.Schedule) (model.Schedule, error) {

	jaevFairnessScore, err := extractFairnessScore(jaevResStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting jaev fairness")
	}

	jaevBalanceScore, err := extractMinBalance(jaevResStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting jaev fairness")
	}

	individualSchedulesJaev, err := extractIndividualSchedules(jaevResStr, int(res.BalanceScore), res.ShiftTypes)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting jaev individual schedules")
	}

	individualSchedules := []model.IndividualSchedule{}
	for _, is := range res.IndividualSchedules {
		if assistantIsOfType(res.Assistants, is.AssistantId, model.JA) {
			schedule, err := getSchedule(individualSchedulesJaev, is.AssistantId)
			if err != nil {
				return model.Schedule{}, err
			}
			individualSchedules = append(individualSchedules, schedule)
		} else {
			individualSchedules = append(individualSchedules, is)
		}
	}

	return model.Schedule{
		FairnessScore:       res.FairnessScore,
		BalanceScore:        res.BalanceScore,
		JaevFairnessScore:   jaevFairnessScore,
		JaevBalance:         int32(jaevBalanceScore),
		NbDays:              res.NbDays,
		Assistants:          res.Assistants,
		ShiftTypes:          res.ShiftTypes,
		IndividualSchedules: individualSchedules,
	}, nil
}

func assistantIsOfType(assistants []model.Assistant, id int32, at model.AssistantType) bool {
	for _, a := range assistants {
		if a.Id == id {
			return a.Type == at
		}
	}
	return false
}

func getSchedule(individualSchedules []model.IndividualSchedule, id int32) (model.IndividualSchedule, error) {
	for _, is := range individualSchedules {
		if is.AssistantId == id {
			return is, nil
		}
	}

	return model.IndividualSchedule{}, errors.New("schedule not found")
}
