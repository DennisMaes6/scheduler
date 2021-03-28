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
    fairness_score: number;
    /**
     * The balance score of this schedule
     */
    balance_score?: number;
    /**
     * The fairness score for teh JAEV schedule
     */
    jaev_fairness_score: number;
    /**
     * The balance for the JAEV schedule
     */
    jaev_balance?: number;
    /**
     * The number of days for which this schedule is generated.
     */
    nb_days?: number;
    holidays: Array<number>;
    /**
     * The assistants involved in this schedule.
     */
    assistants?: Array<Assistant>;
    /**
     * The shift types that were considered when generating this schedule.
     */
    shift_types?: Array<ShiftType>;
    individual_schedules?: Array<IndividualSchedule>;
}
