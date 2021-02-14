/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftType } from './ShiftType';

/**
 * Holds all model parameters to generate a schedule.
 */
export type ModelParameters = {
    /**
     * The minimal balance score for an acceptable solution
     */
    balance_minimum?: number;
    /**
     * The shifts to be included when calculating the balance score.
     */
    balance_shifts?: Array<ShiftType>;
    fairness_shift_weights?: Array<{
        shift_type?: ShiftType,
        /**
         * The weight of this shift type
         */
        weight?: number,
    }>;
}
