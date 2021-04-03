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
			Name:     "Test A",
			Type:     model.JA,
			FreeDays: []int32{1, 2, 3, 4, 5, 6, 7},
		},
		{
			Id:       2,
			Name:     "Test B",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       3,
			Name:     "Test C",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       4,
			Name:     "Test D",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       5,
			Name:     "Test E",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       6,
			Name:     "Test F",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       7,
			Name:     "Test G",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       8,
			Name:     "Test H",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       9,
			Name:     "Test I",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       10,
			Name:     "Test J",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       11,
			Name:     "Test K",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       12,
			Name:     "Test L",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       13,
			Name:     "Test M",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       14,
			Name:     "Test N",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       15,
			Name:     "Test O",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       16,
			Name:     "Test P",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       17,
			Name:     "Test Q",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       18,
			Name:     "Test R",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
	},
}
