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

var cachedSchedule model.Schedule = model.Schedule{}
var cached bool = false

type ScheduleGenerator struct {
	dbc DbController
}

func NewScheduleGenerator() ScheduleGenerator {
	return ScheduleGenerator{newDBController()}
}

func (s ScheduleGenerator) GetInstanceData() (model.InstanceData, error) {
	return s.dbc.GetInstanceData()
}

func (s ScheduleGenerator) UpdateInstanceData(data model.InstanceData) error {
	return s.dbc.SetInstanceData(data)
}

func (s ScheduleGenerator) GetModelParameters() (model.ModelParameters, error) {
	return s.dbc.GetModelParameters()
}

func (s ScheduleGenerator) UpdateModelParameters(params model.ModelParameters) error {
	return s.dbc.SetModelParameters(params)
}

func (s ScheduleGenerator) GetScheduleFromDb() (model.Schedule, error) {
	cmd := exec.Command("java", "-cp", "/Users/jorensjongers/thesis/out/artifacts/scheduler_jar/scheduler.jar", "Main")

	if err := cmd.Run(); err != nil {
		return model.Schedule{}, errors.Wrap(err, "failed starting schedule generation")
	}

	return s.dbc.getSchedule()
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

		res, err := parseFirstStageSchedule(resStr)
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

		cachedSchedule = combinedRes
		//cached = true
	}

	return cachedSchedule, nil
}

func (s ScheduleGenerator) generateDataFile() error {

	file, err := os.Create("minizinc/data.dzn")
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed creating file %s", "minizinc/data.dzn"))
	}

	instanceData, err := s.dbc.GetInstanceData()
	if err != nil {
		return errors.Wrap(err, "database controller error")
	}

	fmt.Printf("nb weeks: %d\n", len(instanceData.Days)/7)

	modelParameters, err := s.dbc.GetModelParameters()
	if err != nil {
		return errors.Wrap(err, "databse controller error")
	}

	fmt.Printf("min balance: %d\n", modelParameters.MinBalance)
	fmt.Printf("min balance JAEV: %d\n", modelParameters.MinBalanceJaev)

	if err := writeData(file, modelParameters, instanceData); err != nil {
		return errors.Wrap(err, "failed writing instance specific data")
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

func (s ScheduleGenerator) generateJaevDataFile(res firstStageSchedule) error {

	file, err := os.Create("minizinc/data_jaev.dzn")
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed creating jaev file %s", "minizinc/data.dzn"))
	}

	instanceData, err := s.dbc.GetInstanceData()
	if err != nil {
		return errors.Wrap(err, "database controller error")
	}

	modelParameters, err := s.dbc.GetModelParameters()
	if err != nil {
		return errors.Wrap(err, "database controller error")
	}

	if err := writeJaevData(file, res, modelParameters, instanceData); err != nil {
		return errors.Wrap(err, "failed writing model parameters")
	}

	return nil
}
