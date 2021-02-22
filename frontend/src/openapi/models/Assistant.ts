/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { AssistantType } from './AssistantType';
import type { ShiftType } from './ShiftType';

/**
 * Represents an assistant.
 */
export type Assistant = {
    /**
     * The identification number of this assistant.
     */
    id?: number;
    type?: AssistantType;
    highest_unfair_workload?: Array<ShiftType>;
    lowest_unfair_workload?: Array<ShiftType>;
}
