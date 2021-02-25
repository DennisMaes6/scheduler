package schedule_generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
	dbc DbController
}

func NewScheduleGenerator() ScheduleGenerator {
	return ScheduleGenerator{newDBController()}
}

func (s ScheduleGenerator) UpdateModelParameters(params model.ModelParameters) error {
	if err := s.dbc.SetModelParameters(params); err != nil {
		return errors.Wrap(err, "failed updating model parameters")
	}
	return nil
}

func (s ScheduleGenerator) GetModelParameters() (model.ModelParameters, error) {
	return s.dbc.GetModelParameters()
}

func (s ScheduleGenerator) UpdateInstanceData(data model.InstanceData) error {
	// TODO
	return nil
}

func (s ScheduleGenerator) GetInstanceData() (model.InstanceData, error) {
	return s.dbc.getInstanceData()
}

func (s ScheduleGenerator) GenerateSchedule() (model.Schedule, error) {

	if !cached {

		if err := s.generateDataFile(); err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed generating data file")
		}

		resStr, err := executeGenerateScheduleCmd()
		if err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed generating schedule")
		}

		if strings.Contains(resStr, "UNSATISFIABLE") {
			return model.Schedule{}, errors.New("model unsatisfiable")
		}

		res, err := parseSchedule(resStr)
		if err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed parsing schedule")
		}

		schedule = res
		//cached = true
	}

	return schedule, nil
}

func (s ScheduleGenerator) generateDataFile() error {

	file, err := os.Create("minizinc/data.dzn")
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed creating file %s", "minizinc/data.dzn"))
	}

	if err := writeInstanceSpecificData(file); err != nil {
		return errors.Wrap(err, "failed writning instance specific data")
	}

	modelParameters, err := s.dbc.GetModelParameters()
	if err != nil {
		return errors.Wrap(err, "databse controller error")
	}

	if err := writeModelParameters(file, modelParameters); err != nil {
		return errors.Wrap(err, "failed writing model parameters")
	}

	return nil

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
