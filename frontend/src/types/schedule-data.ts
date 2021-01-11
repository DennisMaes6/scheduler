// interface that represents a schedule
export type ScheduleData = {
    nb_days: number;
    employees: Employee[];
    shift_types: ShiftType[];
    individual_schedules: IndividualSchedule[];
}

export type Employee = {
    number: number;
    type: EmployeeType;
}

export enum EmployeeType {
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


