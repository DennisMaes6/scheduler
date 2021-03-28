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

var generateScheduleJaevArgs []string = []string{
	"--solver",
	"gurobi",
	"minizinc/scheduler_jaev.mzn",
	"minizinc/data_jaev.dzn",
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
	if err := s.dbc.SetInstanceData(data); err != nil {
		return errors.Wrap(err, "failed updating instance data")
	}
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

		resStr, err := executeGenerateScheduleCmd(generateScheduleCmd, generateScheduleArgs)
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

		if err := s.generateJaevDataFile(res); err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed generating jaev data file")
		}

		jaevResStr, err := executeGenerateScheduleCmd(generateScheduleCmd, generateScheduleJaevArgs)
		if err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed generating schedule")
		}

		if strings.Contains(jaevResStr, "UNSATISFIABLE") {
			return model.Schedule{}, errors.New("JAEV model unsatisfiable")
		}

		combinedRes, err := parseAndCombineSchedule(jaevResStr, res)
		if err != nil {
			return model.Schedule{}, errors.Wrap(err, "failed parsing schedule")
		}

		schedule = combinedRes
		//cached = true
	}

	return schedule, nil
}

func (s ScheduleGenerator) generateDataFile() error {

	file, err := os.Create("minizinc/data.dzn")
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed creating file %s", "minizinc/data.dzn"))
	}

	instanceData, err := s.dbc.getInstanceData()
	if err != nil {
		return errors.Wrap(err, "database controller error")
	}

	if err := writeInstanceSpecificData(file, instanceData); err != nil {
		return errors.Wrap(err, "failed writing instance specific data")
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

func executeGenerateScheduleCmd(cmdStr string, argsStr []string) (string, error) {
	cmd := exec.Command(cmdStr, argsStr...)

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

func (s ScheduleGenerator) generateJaevDataFile(res model.Schedule) error {

	file, err := os.Create("minizinc/data_jaev.dzn")
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed creating jaev file %s", "minizinc/data.dzn"))
	}

	instanceData, err := s.dbc.getInstanceData()
	if err != nil {
		return errors.Wrap(err, "database controller error")
	}

	if err := writeInstanceSpecificDataJaev(file, res, instanceData); err != nil {
		return errors.Wrap(err, "failed writing instance specific data")
	}

	modelParameters, err := s.dbc.GetModelParameters()
	if err != nil {
		return errors.Wrap(err, "database controller error")
	}

	if err := writeModelParametersJaev(file, modelParameters); err != nil {
		return errors.Wrap(err, "failed writing model parameters")
	}

	return nil
}
