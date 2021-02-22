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
    assistant_id?: number;
    /**
     * An array as long as the number of days of this schedule. Contains all the assignments that have been scheduled to this assistant.
     */
    assignments?: Array<Assignment>;
}
