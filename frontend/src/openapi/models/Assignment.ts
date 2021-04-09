/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftType } from './ShiftType';

/**
 * Represents a single shift assignment.
 */
export type Assignment = {
    /**
     * The day of the scheduling period for which this is an assignment.
     */
    day_nb: number;
    shift_type: ShiftType;
    /**
     * Indicates whether or not this assignment is part of a streak of free days as long as the min_balance score for this schedule.
     */
    part_of_min_balance: boolean;
}
