package schedule_generator

import "github.com/jorensjongers/scheduler/backend/model"

// Sort interface methods for shift type parameters
type ByShiftType []model.ShiftTypeModelParameters

var shiftTypeOrder []model.ShiftType = []model.ShiftType{
	model.JANW,
	model.JAWH,
	model.SAEW,
	model.SAWH,
	model.TSPT,
	model.CALL,
}

func (a ByShiftType) Len() int {
	return len(a)
}

func (a ByShiftType) Less(i, j int) bool {
	return index(shiftTypeOrder, a[i].ShiftType) < index(shiftTypeOrder, a[j].ShiftType)
}

func (a ByShiftType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func index(slice []model.ShiftType, item model.ShiftType) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}
	return -1
}
