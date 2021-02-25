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

var initialInstanceData model.InstanceData = model.InstanceData{
	NbWeeks: 4,
	Assistants: []model.AssistantInstance{
		{
			Id:   1,
			Type: model.JA,
		},
		{
			Id:   2,
			Type: model.JA,
		},
		{
			Id:   3,
			Type: model.JA,
		},
		{
			Id:   4,
			Type: model.JA,
		},
		{
			Id:   5,
			Type: model.JA,
		},
		{
			Id:   6,
			Type: model.JA,
		},
		{
			Id:   7,
			Type: model.JA,
		},
		{
			Id:   8,
			Type: model.JA_F,
		},
		{
			Id:   9,
			Type: model.JA_F,
		},
		{
			Id:   10,
			Type: model.SA,
		},
		{
			Id:   11,
			Type: model.SA,
		},
		{
			Id:   12,
			Type: model.SA,
		},
		{
			Id:   13,
			Type: model.SA_F,
		},
		{
			Id:   14,
			Type: model.SA_F,
		},
		{
			Id:   15,
			Type: model.SA_NEO,
		},
		{
			Id:   16,
			Type: model.SA_NEO,
		},
		{
			Id:   17,
			Type: model.SA_F_NEO,
		},
		{
			Id:   18,
			Type: model.SA_F_NEO,
		},
	},
}
