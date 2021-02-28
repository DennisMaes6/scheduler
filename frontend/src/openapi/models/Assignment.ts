/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftType } from './ShiftType';

/**
 * Represents a single shift assignment.
 */
export type Assignment = {
    shift_type: ShiftType;
    /**
     * whether or not this assignment is part of a min balance free block.
     */
    part_of_min_balance: boolean;
}
