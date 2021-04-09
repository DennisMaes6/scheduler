/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { AssistantType } from './AssistantType';

/**
 * Holds all information for one assistant.
 */
export type Assistant = {
    /**
     * The identification number of this assistant.
     */
    id: number;
    /**
     * The name of this assistant.
     */
    name: string;
    type: AssistantType;
    /**
     * The day numbers corresponding to the free days granted to this assistant for the current scheduling period.
     */
    free_days: Array<number>;
}
