/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { IndividualSchedule } from './IndividualSchedule';

/**
 * Holds all data of a generated schedule.
 */
export type Schedule = {
    /**
     * The fairness score of this schedule.
     */
    fairness_score: number;
    /**
     * The balance score of this schedule.
     */
    balance_score: number;
    /**
     * The fairness score for the JAEV shifts of this schedule.
     */
    jaev_fairness_score: number;
    /**
     * The balance score for the JAEV shifts of this schedule.
     */
    jaev_balance_score: number;
    individual_schedules?: Array<IndividualSchedule>;
}
