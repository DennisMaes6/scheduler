package schedule_generator

import "github.com/DennisMaes6/scheduler/backend/model"

// Sort interface methods for shift type parameters
type ByShiftType []model.ShiftTypeModelParameters

var shiftTypeOrder []model.ShiftType = []model.ShiftType{
	model.SANW,
	model.JAWE,
	model.JAHO,
	model.SAEW,
	model.SAWE,
	model.SAHO,
	model.TPWE,
	model.TPHO,
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
