package schedule_generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jorensjongers/scheduler/backend/model"
	"github.com/pkg/errors"
)

type unfairnessPerShiftType struct {
	shiftType       model.ShiftType
	fairnessScore   float32
	highestWorkload []int32 // assistant IDs
	lowestWorkload  []int32 // assistant IDs
}

func parseSchedule(scheduleStr string) (model.Schedule, error) {

	nbDays, err := extractNbDays(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting nb days")
	}

	shiftTypes, err := extractShiftTypes(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting shift types")
	}

	fairnessScore := getFairnessScore(shiftTypes)

	balanceScore, err := extractMinBalance(scheduleStr)
	if err != nil {
		return model.Schedule{}, err
	}

	individualSchedules, err := extractIndividualSchedules(scheduleStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting individual schedules")
	}

	assistants, err := extractAssistants(scheduleStr, getUnfairness(individualSchedules, shiftTypes, int(nbDays)), fairnessScore)
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

func extractAssistants(scheduleStr string, unfairness []unfairnessPerShiftType, fairnessScore float32) ([]model.Assistant, error) {
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

		_, err = strconv.Atoi(strings.Split(split_line[2], ":")[1])
		if err != nil {
			return []model.Assistant{}, err
		}

		huw := getHighestUnfairWorkload(unfairness, fairnessScore, int32(id))
		luw := filterQualifications(getLowestUnfairWorkload(unfairness, fairnessScore, int32(id)), assistantType)

		assistant := model.Assistant{
			Id:                    int32(id),
			Type:                  assistantType,
			HighestUnfairWorkload: filter(huw, luw),
			LowestUnfairWorkload:  filter(luw, huw),
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

func extractShiftTypes(scheduleStr string) ([]model.ScheduleShiftTypes, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "fairness_per_shift_type")
	shiftTypesLine := strings.TrimSpace(strings.Split(lines[0], ":")[1])

	result := []model.ScheduleShiftTypes{}
	for _, shiftTypeStr := range strings.Split(shiftTypesLine, " ") {
		shiftType, err := parseShiftType(strings.Split(shiftTypeStr, "=")[0])
		if err != nil {
			return []model.ScheduleShiftTypes{}, err
		}

		fairnessScore, err := strconv.ParseFloat(strings.Split(shiftTypeStr, "=")[1], 32)
		if err != nil {
			return []model.ScheduleShiftTypes{}, err
		}

		result = append(result, model.ScheduleShiftTypes{
			FairnessScore: float32(fairnessScore),
			ShiftType:     shiftType,
		})
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
		for i := 3; i < len(entries); i++ {
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

		balanceShifts, err := extractBalanceShifts(scheduleStr)
		if err != nil {
			return []model.IndividualSchedule{}, err
		}

		is := model.IndividualSchedule{
			AssistantId: int32(assistantId),
			Assignments: tagAssignments(assignments, minBalance, balanceShifts),
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

func extractBalanceShifts(scheduleStr string) ([]model.ShiftType, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "balance_shifts")
	return parseShiftTypes(strings.Split(lines[0], ":")[1])
}

func parseShiftTypes(line string) ([]model.ShiftType, error) {
	shiftTypesLine := strings.Split(strings.TrimSpace(line), " ")

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

func oneOf(test model.ShiftType, sts []model.ShiftType) bool {
	for _, st := range sts {
		if st == test {
			return true
		}
	}
	return false
}

func getFairnessScore(shiftTypes []model.ScheduleShiftTypes) float32 {
	result := float32(0.0)
	for _, st := range shiftTypes {
		if st.FairnessScore > result {
			result = float32(st.FairnessScore)
		}
	}
	return result
}

func getUnfairness(schedules []model.IndividualSchedule,
	shiftTypes []model.ScheduleShiftTypes, nbDays int) []unfairnessPerShiftType {

	result := []unfairnessPerShiftType{}
	for _, shiftType := range shiftTypes {
		highestWorkload := 0
		assistantsHighestWorkload := []int32{}
		lowestWorkload := nbDays
		assistantsLowestWorkload := []int32{}

		for _, schedule := range schedules {
			count := countOccurences(shiftType.ShiftType, schedule)
			if count > highestWorkload {
				highestWorkload = count
				assistantsHighestWorkload = []int32{schedule.AssistantId}
			} else if count == highestWorkload {
				assistantsHighestWorkload = append(assistantsHighestWorkload, schedule.AssistantId)
			}

			if count < lowestWorkload {
				lowestWorkload = count
				assistantsLowestWorkload = []int32{schedule.AssistantId}
			} else if count == lowestWorkload {
				assistantsLowestWorkload = append(assistantsLowestWorkload, schedule.AssistantId)
			}
		}

		unfairness := unfairnessPerShiftType{
			shiftType:       shiftType.ShiftType,
			fairnessScore:   shiftType.FairnessScore,
			highestWorkload: assistantsHighestWorkload,
			lowestWorkload:  assistantsLowestWorkload,
		}
		result = append(result, unfairness)
	}

	return result
}

func countOccurences(shiftType model.ShiftType, schedule model.IndividualSchedule) int {
	count := 0
	for _, assignment := range schedule.Assignments {
		if assignment.ShiftType == shiftType {
			count++
		}
	}
	return count
}

func getHighestUnfairWorkload(unfairness []unfairnessPerShiftType, fairnessScore float32, id int32) []model.ShiftType {

	result := []model.ShiftType{}
	for _, unfairnessShiftType := range unfairness {
		if contains(unfairnessShiftType.highestWorkload, id) && unfairnessShiftType.fairnessScore == fairnessScore {
			result = append(result, unfairnessShiftType.shiftType)
		}
	}
	return result
}

func getLowestUnfairWorkload(unfairness []unfairnessPerShiftType, fairnessScore float32, id int32) []model.ShiftType {

	result := []model.ShiftType{}
	for _, unfairnessShiftType := range unfairness {
		if contains(unfairnessShiftType.lowestWorkload, id) && unfairnessShiftType.fairnessScore == fairnessScore {
			result = append(result, unfairnessShiftType.shiftType)
		}
	}
	return result
}

func contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func filter(first []model.ShiftType, second []model.ShiftType) []model.ShiftType {

	result := []model.ShiftType{}
	for _, firstEl := range first {
		appears := false
		for _, secondEl := range second {
			if firstEl == secondEl {
				appears = true
			}
		}
		if !appears {
			result = append(result, firstEl)
		}
	}

	return result
}

func filterQualifications(sts []model.ShiftType, at model.AssistantType) []model.ShiftType {
	result := []model.ShiftType{}
	for _, st := range sts {
		if oneOf(st, getQualifications(at)) {
			result = append(result, st)
		}
	}
	return result
}

func getQualifications(at model.AssistantType) []model.ShiftType {
	switch at {
	case model.JA:
		return []model.ShiftType{model.JAEV, model.JAWH, model.JANW}
	case model.JA_F:
		return []model.ShiftType{model.JAWH}
	case model.SA:
		return []model.ShiftType{model.SAEW, model.CALL, model.SAWH}
	case model.SA_F:
		return []model.ShiftType{model.CALL, model.SAWH}
	case model.SA_NEO:
		return []model.ShiftType{model.SAEW, model.CALL, model.SAWH, model.TSPT}
	case model.SA_F_NEO:
		return []model.ShiftType{model.CALL, model.SAWH, model.TSPT}
	default:
		return []model.ShiftType{}
	}
}
