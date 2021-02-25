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
    id?: number;
    type?: AssistantType;
}
