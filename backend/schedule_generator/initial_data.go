package schedule_generator

import "github.com/jorensjongers/scheduler/backend/model"

var initialModelParameters model.ModelParameters = model.ModelParameters{
	MinBalance:     7,
	MinBalanceJaev: 4,
	ShiftTypeParameters: []model.ShiftTypeModelParameters{
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
	Days: []model.Day{
		{
			Id:        1,
			Date:      "2021-01-01",
			IsHoliday: true,
		},
		{
			Id:        2,
			Date:      "2021-01-02",
			IsHoliday: false,
		},
		{
			Id:        3,
			Date:      "2021-01-03",
			IsHoliday: false,
		},
		{
			Id:        4,
			Date:      "2021-01-01",
			IsHoliday: false,
		},
		{
			Id:        5,
			Date:      "2021-01-05",
			IsHoliday: false,
		},
		{
			Id:        6,
			Date:      "2021-01-06",
			IsHoliday: false,
		},
		{
			Id:        7,
			Date:      "2021-01-07",
			IsHoliday: false,
		},
		{
			Id:        8,
			Date:      "2021-01-08",
			IsHoliday: false,
		},
		{
			Id:        9,
			Date:      "2021-01-09",
			IsHoliday: true,
		},
		{
			Id:        10,
			Date:      "2021-01-10",
			IsHoliday: false,
		},
		{
			Id:        11,
			Date:      "2021-01-11",
			IsHoliday: false,
		},
		{
			Id:        12,
			Date:      "2021-01-12",
			IsHoliday: false,
		},
		{
			Id:        13,
			Date:      "2021-01-13",
			IsHoliday: false,
		},
		{
			Id:        14,
			Date:      "2021-01-14",
			IsHoliday: false,
		},
		{
			Id:        15,
			Date:      "2021-01-15",
			IsHoliday: false,
		},
		{
			Id:        16,
			Date:      "2021-01-16",
			IsHoliday: false,
		},
		{
			Id:        17,
			Date:      "2021-01-17",
			IsHoliday: false,
		},
		{
			Id:        18,
			Date:      "2021-01-18",
			IsHoliday: false,
		},
		{
			Id:        19,
			Date:      "2021-01-19",
			IsHoliday: false,
		},
		{
			Id:        20,
			Date:      "2021-01-20",
			IsHoliday: false,
		},
		{
			Id:        21,
			Date:      "2021-01-21",
			IsHoliday: false,
		},
		{
			Id:        22,
			Date:      "2021-01-22",
			IsHoliday: false,
		},
		{
			Id:        23,
			Date:      "2021-01-23",
			IsHoliday: false,
		},
		{
			Id:        24,
			Date:      "2021-01-24",
			IsHoliday: false,
		},
		{
			Id:        25,
			Date:      "2021-01-25",
			IsHoliday: false,
		},
		{
			Id:        26,
			Date:      "2021-01-26",
			IsHoliday: false,
		},
		{
			Id:        27,
			Date:      "2021-01-27",
			IsHoliday: false,
		},
		{
			Id:        28,
			Date:      "2021-01-29",
			IsHoliday: false,
		},
	},
	Assistants: []model.Assistant{
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
