package schedule_generator

import "github.com/jorensjongers/scheduler/backend/model"

var initialModelParameters model.ModelParameters = model.ModelParameters{
	BalanceMinimum:     7,
	BalanceMinimunJaev: 4,
	ShiftTypeParams: []model.ShiftTypeModelParameters{
		{
			ShiftType:     model.JANW,
			ShiftWorkload: 3,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.JAWH,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.SAEW,
			ShiftWorkload: 3,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.SAWH,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.TSPT,
			ShiftWorkload: 3,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.CALL,
			ShiftWorkload: 0,
			MaxBuffer:     1,
		},
	},
}

var initialInstanceData model.InstanceData = model.InstanceData{
	NbWeeks:  4,
	Holidays: []int32{4, 20},
	Assistants: []model.AssistantInstance{
		{
			Id:       1,
			Type:     model.JA,
			FreeDays: []int32{1, 2, 3, 4, 5, 6, 7},
		},
		{
			Id:       2,
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       3,
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       4,
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       5,
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       6,
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       7,
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       8,
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       9,
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       10,
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       11,
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       12,
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       13,
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       14,
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       15,
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       16,
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       17,
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       18,
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
	},
}
