/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { AssistantInstance } from './AssistantInstance';

/**
 * Holds the instance data for scheduling.
 */
export type InstanceData = {
    assistants: Array<AssistantInstance>;
    /**
     * The number of weeks to be scheduled in this instance
     */
    nb_weeks: number;
    holidays: Array<number>;
}
