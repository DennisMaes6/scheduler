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
     * Represents the date of this day.
     */
    date: {
        /**
         * The day of month of this day.
         */
        day: number,
        /**
         * The month of this day. 1 = January, 12 = December
         */
        month: number,
        /**
         * The year of this day.
         */
        year: number,
    };
    /**
     * Indicates whether or not this day should be considered a holiday in the produced schedule.
     */
    is_holiday: boolean;
}
