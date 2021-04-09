/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Assignment } from './Assignment';

/**
 * Represents the work schedule of an individual assistant.
 */
export type IndividualSchedule = {
    /**
     * The identification number of the assistant for which this is an individual schedule.
     */
    assistant_id: number;
    /**
     * The workload of this inidividual schedule. Used when calculating the fairness score.
     */
    workload: number;
    /**
     * Contains all the individual assignments of this individual schedule.
     */
    assignments: Array<Assignment>;
}
