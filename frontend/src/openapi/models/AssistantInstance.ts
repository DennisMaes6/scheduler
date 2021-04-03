/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { AssistantType } from './AssistantType';

/**
 * Holds all instance data for one assistant.
 */
export type AssistantInstance = {
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
     * The free days granted to this assistant for the current scheduling period.
     */
    free_days: Array<number>;
}
