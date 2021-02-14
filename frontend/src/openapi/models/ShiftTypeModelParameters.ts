/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftType } from './ShiftType';

/**
 * Holds the shift type specific model parameters.
 */
export type ShiftTypeModelParameters = {
    shift_type?: ShiftType;
    /**
     * The weight of this shift type in the fairness score.
     */
    fairness_weight?: number;
    /**
     * Whether or not this shift type is included when computing the balaance score.
     */
    included_in_balance?: boolean;
}
