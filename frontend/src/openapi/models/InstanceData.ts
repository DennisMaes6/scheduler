/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Assistant } from './Assistant';
import type { Day } from './Day';

/**
 * Holds the instance data for scheduling.
 */
export type InstanceData = {
    /**
     * The assistant instances representing the assistants for which to produce a schedule.
     */
    assistants: Array<Assistant>;
    /**
     * The days for which to produce a schedule.
     */
    days: Array<Day>;
}
