package schedule_generator

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jorensjongers/scheduler/backend/model"
	"github.com/pkg/errors"
)

const generateScheduleCmd string = "minizinc"

var generateScheduleArgs []string = []string{
	"--solver",
	"gurobi",
	"minizinc/scheduler.mzn",
	"minizinc/data.dzn",
}

var schedule model.Schedule = model.Schedule{}
var cached bool = false

type ScheduleGenerator struct {
}

func (scheduleGenerator ScheduleGenerator) GenerateSchedule() (model.Schedule, error) {

	if !cached {
		resStr, err := executeGenerateScheduleCmd()
		if err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed generating schedule")
		}

		res, err := parseSchedule(resStr)
		if err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed parsing schedule")
		}

		schedule = res
		cached = true
	}

	return schedule, nil
}

func executeGenerateScheduleCmd() (string, error) {
	cmd := exec.Command(generateScheduleCmd, generateScheduleArgs...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", errors.Wrap(err, "failed connecting to std out")
	}
	if err := cmd.Start(); err != nil {
		return "", errors.Wrap(err, "failed starting schedule generation")
	}

	var resultRaw []byte
	if resultRaw, err = ioutil.ReadAll(stdout); err != nil {
		return "", errors.Wrap(err, "failed reading result from std out")
	}

	return string(resultRaw), nil
}

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
	lines := strings.Split(scheduleStr, "\n")
	nbDays, err := strconv.Atoi(strings.Split(lines[0], ":")[1])
	if err != nil {
		return 0, err
	}

	return int32(nbDays), nil
}

func extractAssistants(scheduleStr string) ([]model.Assistant, error) {
	lines := strings.Split(scheduleStr, "\n")

	assistants := []model.Assistant{}
	for i := 2; i < len(lines)-2; i++ {
		line := strings.Split(lines[i], " ")

		id, err := strconv.Atoi(strings.Split(line[0], ":")[1])
		if err != nil {
			return []model.Assistant{}, err
		}

		assistantType, err := parseAssistantType(strings.Split(line[1], ":")[1])
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
	lines := strings.Split(scheduleStr, "\n")
	shiftTypesLine := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")

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
	lines := strings.Split(scheduleStr, "\n")

	result := []model.IndividualSchedule{}
	for i := 2; i < len(lines)-2; i++ {

		entries := strings.Split(strings.TrimSpace(lines[i]), " ")
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

		is := model.IndividualSchedule{
			AssistantId: int32(assistantId),
			Assignments: assignments,
		}

		result = append(result, is)
	}

	return result, nil
}
