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
     * The fairness score of this schedule
     */
    fairness_score?: number;
    /**
     * The balance score of this schedule
     */
    balance_score?: number;
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
    shift_types?: Array<{
        /**
         * The fairness score for this shift type
         */
        fairness_score?: number,
        shift_type?: ShiftType,
    }>;
    individual_schedules?: Array<IndividualSchedule>;
}
