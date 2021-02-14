/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftTypeModelParameters } from './ShiftTypeModelParameters';

/**
 * Holds all model parameters to generate a schedule.
 */
export type ModelParameters = {
    /**
     * The minimal balance score for an acceptable solution
     */
    balance_minimum?: number;
    shift_type_params?: Array<ShiftTypeModelParameters>;
}
