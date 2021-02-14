/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Assistant } from './Assistant';
import type { IndividualSchedule } from './IndividualSchedule';
import type { ShiftType } from './ShiftType';

/**
 * Holds all data of a generated schedule.
 */
export type Schedule = {
    /**
     * The number of days for which this schedule is generated.
     */
    nb_days?: number;
    /**
     * The assistants involved in this schedule.
     */
    assistants?: Array<Assistant>;
    /**
     * The shift_types that were considered when generating this schedule.
     */
    shift_types?: Array<ShiftType>;
    individual_schedules?: Array<IndividualSchedule>;
}
