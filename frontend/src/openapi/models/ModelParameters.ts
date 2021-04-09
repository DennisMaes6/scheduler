/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ShiftTypeModelParameters } from './ShiftTypeModelParameters';

/**
 * Holds all model parameters for generating schedules.
 */
export type ModelParameters = {
    /**
     * The minimal balance score for an acceptable solution.
     */
    min_balance?: number;
    /**
     * The mininmal balance score for JAEV shifts for an acceptable solution.
     */
    min_balance_jaev?: number;
    shift_type_parameters?: Array<ShiftTypeModelParameters>;
}
