package schedule_generator

import "github.com/jorensjongers/scheduler/backend/model"

var initialModelParameters model.ModelParameters = model.ModelParameters{
	BalanceMinimum: 3,
	ShiftTypeParams: []model.ShiftTypeModelParameters{
		{
			ShiftType:         model.JAEV,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
		{
			ShiftType:         model.JANW,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
		{
			ShiftType:         model.JAWH,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
		{
			ShiftType:         model.SAEW,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
		{
			ShiftType:         model.SAWH,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
		{
			ShiftType:         model.TSPT,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
		{
			ShiftType:         model.CALL,
			FairnessWeight:    1,
			IncludedInBalance: true,
		},
	},
}
