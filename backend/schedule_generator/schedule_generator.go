package schedule_generator

import "github.com/jorensjongers/scheduler/backend/model"

type ScheduleGenerator struct {
}

func (scheduleGenerator ScheduleGenerator) GenerateSchedule() model.Schedule {
	return model.Schedule{
		NbDays:              7,
		Assistants:          []model.Assistant{},
		ShiftTypes:          "",
		IndividualSchedules: []model.IndividualSchedule{},
	}
}
