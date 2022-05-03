package schedule_generator

import "github.com/DennisMaes6/scheduler/backend/model"

var initialModelParameters model.ModelParameters = model.ModelParameters{
	MinBalance:     7,
	MinBalanceJaev: 4,
	ShiftTypeParameters: []model.ShiftTypeModelParameters{
		{
			ShiftType:     model.SANW,
			ShiftCoverage: 2,
			ShiftWorkload: 3,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.JAWE,
			ShiftCoverage: 2,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.JAEV,
			ShiftCoverage: 2,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.JAHO,
			ShiftCoverage: 2,
			ShiftWorkload: 1,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.SAWE,
			ShiftCoverage: 2,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.SAHO,
			ShiftCoverage: 2,
			ShiftWorkload: 1,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.SAEV1,
			ShiftCoverage: 2,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.SAEV2,
			ShiftCoverage: 2,
			ShiftWorkload: 1,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.TPWE,
			ShiftCoverage: 2,
			ShiftWorkload: 3,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.TPHO,
			ShiftCoverage: 2,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.TPNF,
			ShiftCoverage: 2,
			ShiftWorkload: 2,
			MaxBuffer:     1,
		},
		{
			ShiftType:     model.CALL,
			ShiftCoverage: 2,
			ShiftWorkload: 0,
			MaxBuffer:     1,
		},
	},
}

var initialInstanceData model.InstanceData = model.InstanceData{
	Days: []model.Day{
		{
			Id: 1,
			Date: model.DayDate{
				Day:   1,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: true,
		},
		{
			Id: 2,
			Date: model.DayDate{
				Day:   2,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 3,
			Date: model.DayDate{
				Day:   3,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 4,
			Date: model.DayDate{
				Day:   4,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 5,
			Date: model.DayDate{
				Day:   5,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 6,
			Date: model.DayDate{
				Day:   6,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 7,
			Date: model.DayDate{
				Day:   7,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 8,
			Date: model.DayDate{
				Day:   8,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 9,
			Date: model.DayDate{
				Day:   9,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 10,
			Date: model.DayDate{
				Day:   10,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 11,
			Date: model.DayDate{
				Day:   11,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: true,
		},
		{
			Id: 12,
			Date: model.DayDate{
				Day:   12,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 13,
			Date: model.DayDate{
				Day:   13,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 14,
			Date: model.DayDate{
				Day:   14,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 15,
			Date: model.DayDate{
				Day:   15,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 16,
			Date: model.DayDate{
				Day:   16,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 17,
			Date: model.DayDate{
				Day:   17,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 18,
			Date: model.DayDate{
				Day:   18,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 19,
			Date: model.DayDate{
				Day:   19,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 20,
			Date: model.DayDate{
				Day:   20,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 21,
			Date: model.DayDate{
				Day:   21,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 22,
			Date: model.DayDate{
				Day:   22,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 23,
			Date: model.DayDate{
				Day:   23,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 24,
			Date: model.DayDate{
				Day:   24,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 25,
			Date: model.DayDate{
				Day:   25,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 26,
			Date: model.DayDate{
				Day:   26,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 27,
			Date: model.DayDate{
				Day:   27,
				Month: 1,
				Year:  2021,
			},
			IsHoliday: false,
		},
		{
			Id: 28,
			Date: model.DayDate{
				Day:   28,
				Month: 1,
				Year:  2021,
			},
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
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       9,
			Name:     "Test I",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       10,
			Name:     "Test J",
			Type:     model.JA,
			FreeDays: []int32{},
		},
		{
			Id:       11,
			Name:     "Test K",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       12,
			Name:     "Test L",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       13,
			Name:     "Test M",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       14,
			Name:     "Test N",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       15,
			Name:     "Test O",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       16,
			Name:     "Test P",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       17,
			Name:     "Test Q",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       18,
			Name:     "Test R",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       19,
			Name:     "Test S",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       20,
			Name:     "Test T",
			Type:     model.JA_F,
			FreeDays: []int32{},
		},
		{
			Id:       21,
			Name:     "Test U",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       22,
			Name:     "Test V",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       23,
			Name:     "Test W",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       24,
			Name:     "Test X",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       25,
			Name:     "Test Y",
			Type:     model.SA,
			FreeDays: []int32{},
		},
		{
			Id:       26,
			Name:     "Test Z",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       27,
			Name:     "Test AA",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       28,
			Name:     "Test AB",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       29,
			Name:     "Test AC",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       30,
			Name:     "Test AD",
			Type:     model.SA_F,
			FreeDays: []int32{},
		},
		{
			Id:       31,
			Name:     "Test AE",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       32,
			Name:     "Test AF",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       33,
			Name:     "Test AG",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       34,
			Name:     "Test AH",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       35,
			Name:     "Test AI",
			Type:     model.SA_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       36,
			Name:     "Test AJ",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       37,
			Name:     "Test AK",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       38,
			Name:     "Test AL",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       39,
			Name:     "Test AM",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
		{
			Id:       40,
			Name:     "Test AN",
			Type:     model.SA_F_NEO,
			FreeDays: []int32{},
		},
	},
}
