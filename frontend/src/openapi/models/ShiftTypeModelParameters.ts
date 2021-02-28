/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftType } from './ShiftType';

/**
 * Holds the shift type specific model parameters.
 */
export type ShiftTypeModelParameters = {
    shift_type: ShiftType;
    /**
     * The weight of this shift type in the fairness score.
     */
    shift_workload: number;
    /**
     * The number of assignments per assisant allowed above the minimun for this shift type.
     */
    max_buffer: number;
}
