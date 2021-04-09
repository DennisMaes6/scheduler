/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

/**
 * Represents a single day in the scheduling period.
 */
export type Day = {
    /**
     * The day number of this day in the current scheduling period.
     */
    id: number;
    /**
     * A string representation of the date that is represented by this day.
     */
    date: string;
    /**
     * Indicates whether or not this day should be considered a holiday in the produced schedule.
     */
    is_holiday: boolean;
}
