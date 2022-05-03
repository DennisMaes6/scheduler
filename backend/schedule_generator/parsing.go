package schedule_generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DennisMaes6/scheduler/backend/model"
	"github.com/pkg/errors"
)

type firstStageSchedule struct {
	fairnessScore       float32
	balanceScore        int32
	individualSchedules []untaggedIndividualSchedule
}

type untaggedIndividualSchedule struct {
	assistantId int32
	workload    float32
	assignments []untaggedAssignment
}

type untaggedAssignment struct {
	dayNb     int32
	shiftType model.ShiftType
}

func parseFirstStageSchedule(scheduleStr string) (firstStageSchedule, error) {

	fairnessScore, err := extractFairnessScore(scheduleStr)
	if err != nil {
		return firstStageSchedule{}, errors.Wrap(err, "failed extracting fairness score")
	}

	balanceScore, err := extractBalanceScore(scheduleStr)
	if err != nil {
		return firstStageSchedule{}, errors.Wrap(err, "failed extracting balance score")
	}

	individualSchedules, err := extractIndividualSchedules(scheduleStr)
	if err != nil {
		return firstStageSchedule{}, errors.Wrap(err, "failed extracting individual schedules")
	}

	result := firstStageSchedule{
		fairnessScore:       float32(fairnessScore),
		balanceScore:        int32(balanceScore),
		individualSchedules: individualSchedules,
	}

	return result, nil
}

func extractFairnessScore(scheduleStr string) (float64, error) {
	fairnessLines := filterLines(strings.Split(scheduleStr, "\n"), "fairness_score")
	return strconv.ParseFloat(strings.Split(fairnessLines[0], ":")[1], 32)
}

func extractBalanceScore(scheduleStr string) (int, error) {
	minBalanceLines := filterLines(strings.Split(scheduleStr, "\n"), "balance_score")
	return strconv.Atoi(strings.Split(minBalanceLines[0], ":")[1])
}

func extractIndividualSchedules(scheduleStr string) ([]untaggedIndividualSchedule, error) {
	lines := filterLines(strings.Split(scheduleStr, "\n"), "assistant")

	result := []untaggedIndividualSchedule{}
	for _, line := range lines {

		entries := strings.Split(strings.TrimSpace(line), " ")

		assistantId, err := strconv.Atoi(strings.Split(entries[0], ":")[1])
		if err != nil {
			return []untaggedIndividualSchedule{}, errors.Wrap(err, "failed parsing assistant id")
		}

		workload, err := strconv.ParseFloat(strings.Split(entries[1], ":")[1], 32)
		if err != nil {
			return []untaggedIndividualSchedule{}, errors.Wrap(err, "failed parsing assistant workload")
		}

		assignments := []untaggedAssignment{}
		for i := 2; i < len(entries); i++ {
			shiftType, err := parseShiftType(entries[i])
			if err != nil {
				return []untaggedIndividualSchedule{}, errors.Wrap(err, "failed parsing assignment")
			}
			assignment := untaggedAssignment{
				dayNb:     int32(i - 1),
				shiftType: shiftType,
			}
			assignments = append(assignments, assignment)
		}

		is := untaggedIndividualSchedule{
			assistantId: int32(assistantId),
			workload:    float32(workload),
			assignments: assignments,
		}

		result = append(result, is)
	}

	return result, nil
}

func parseShiftType(shiftTypeStr string) (model.ShiftType, error) {

	switch shiftTypeStr {
	case "JAEV":
		return model.JAEV, nil
	case "SANW":
		return model.SANW, nil
	case "JAWE":
		return model.JAWE, nil
	case "JAHO":
		return model.JAHO, nil
	case "SAEW":
		return model.SAEW, nil
	case "SAWE":
		return model.SAWE, nil
	case "SAHO":
		return model.SAHO, nil
	case "SAEV1":
		return model.SAEV1, nil
	case "SAEV2":
		return model.SAEV2, nil
	case "TPWE":
		return model.TPWE, nil
	case "TPHO":
		return model.TPHO, nil
	case "TPNF":
		return model.TPNF, nil
	case "CALL":
		return model.CALL, nil
	case "FREE":
		return model.FREE, nil
	default:
		return "", fmt.Errorf("invalid shift type string: %s", shiftTypeStr)
	}
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

func parseAndCombineSchedule(jaevResStr string, res firstStageSchedule) (model.Schedule, error) {

	jaevFairnessScore, err := extractFairnessScore(jaevResStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting jaev fairness score")
	}

	jaevBalanceScore, err := extractBalanceScore(jaevResStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting jaev balance score")
	}

	individualSchedulesJaev, err := extractIndividualSchedules(jaevResStr)
	if err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed extracting jaev individual schedules")
	}

	individualSchedules := combineIndividualSchedules(res.individualSchedules, individualSchedulesJaev)

	return model.Schedule{
		FairnessScore:       res.fairnessScore,
		BalanceScore:        float32(res.balanceScore),
		JaevFairnessScore:   float32(jaevFairnessScore),
		JaevBalanceScore:    float32(jaevBalanceScore),
		IndividualSchedules: tagIndividualSchedules(individualSchedules, res.balanceScore),
	}, nil
}

func combineIndividualSchedules(
	firstStageISs []untaggedIndividualSchedule,
	jaevISs []untaggedIndividualSchedule) []untaggedIndividualSchedule {
	result := []untaggedIndividualSchedule{}

	for _, is := range firstStageISs {
		found, jaevIs := findISWithId(is.assistantId, jaevISs)
		if found {
			result = append(result, jaevIs)
		} else {
			result = append(result, is)
		}

	}

	return result
}

func findISWithId(id int32, iss []untaggedIndividualSchedule) (bool, untaggedIndividualSchedule) {
	for _, is := range iss {
		if is.assistantId == id {
			return true, is
		}
	}

	return false, untaggedIndividualSchedule{}
}

func tagIndividualSchedules(iss []untaggedIndividualSchedule, balanceScore int32) []model.IndividualSchedule {

	result := []model.IndividualSchedule{}

	for _, uis := range iss {
		streaks := streaksOfLength(findStreaks(uis), balanceScore)
		taggedAssignments := []model.Assignment{}
		for _, ua := range uis.assignments {
			taggedAssignment := model.Assignment{
				DayNb:            ua.dayNb,
				ShiftType:        ua.shiftType,
				PartOfMinBalance: inStreak(streaks, ua.dayNb),
			}
			taggedAssignments = append(taggedAssignments, taggedAssignment)
		}

		taggedIndividualSchedule := model.IndividualSchedule{
			AssistantId: uis.assistantId,
			Workload:    uis.workload,
			Assignments: taggedAssignments,
		}

		result = append(result, taggedIndividualSchedule)
	}

	return result
}

type streak struct {
	firstDay int32
	lastDay  int32
}

func findStreaks(uis untaggedIndividualSchedule) []streak {

	var (
		firstShiftDone bool = false
		streakActive   bool
		firstDay       int32
	)

	streaks := []streak{}

	for _, ua := range uis.assignments {
		if firstShiftDone {
			if !streakActive && !countBalance(ua) {
				firstDay = int32(ua.dayNb)
				streakActive = true
			}
			if streakActive && countBalance(ua) {
				streakActive = false
				streak := streak{
					firstDay: firstDay,
					lastDay:  int32(ua.dayNb),
				}
				streaks = append(streaks, streak)
			}
		}

		if !firstShiftDone && (ua.shiftType != model.FREE && ua.shiftType != model.JAEV) {
			firstShiftDone = true
		}
	}

	return streaks
}

func countBalance(ua untaggedAssignment) bool {
	if ua.dayNb%7 == 2 || ua.dayNb%7 == 3 {
		return ua.shiftType != model.FREE && ua.shiftType != model.JAEV
	}

	return ua.shiftType != model.FREE &&
		ua.shiftType != model.JAEV &&
		ua.shiftType != model.JAHO &&
		ua.shiftType != model.SAHO &&
		ua.shiftType != model.TPHO
}

func streaksOfLength(streaks []streak, length int32) []streak {
	result := []streak{}

	for _, streak := range streaks {
		if (streak.lastDay - streak.firstDay) == length {
			result = append(result, streak)
		}
	}
	return result
}

func inStreak(streaks []streak, dayNb int32) bool {
	for _, streak := range streaks {
		if dayNb >= streak.firstDay && dayNb < streak.lastDay {
			return true
		}
	}

	return false
}
