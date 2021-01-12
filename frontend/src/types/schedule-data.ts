// interface that represents a schedule
export type ScheduleData = {
    nb_days: number;
    assistants: Assistant[];
    shift_types: ShiftType[];
    individual_schedules: IndividualSchedule[];
}

export type Assistant = {
    id: number;
    type: AssistantType;
}

export enum AssistantType {
    JA = "JA",
    JA_F = "JA_F",
    SA = "SA",
    SA_F = "SA_F",
    SA_NEO = "SA_NEO",
    SA_F_NEO = "SA_F_NEO"
}

export enum ShiftType {
    JA_E = "JA_E",
    JA_NW = "JA_NW",
    JA_WH = "JA_WH",
    SA_EW = "SA_EW",
    SA_WH = "SA_WH",
    TS = "TS",
    C = "C",
    O = "O",
}

export type IndividualSchedule = {
    assistant_id: number,
    assignments: ShiftType[],
}


