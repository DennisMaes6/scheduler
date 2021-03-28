/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { AssistantType } from './AssistantType';

/**
 * Represents an assistant.
 */
export type Assistant = {
    /**
     * The identification number of this assistant.
     */
    id: number;
    type: AssistantType;
    /**
     * The workload on this assiastant in the schedule.
     */
    workload: number;
}
