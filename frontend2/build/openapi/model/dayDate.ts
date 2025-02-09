/**
 * Scheduler API
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Represents the date of this day.
 */
export interface DayDate { 
    /**
     * The day of month of this day.
     */
    day: number;
    /**
     * The month of this day. 1 = January, 12 = December
     */
    month: number;
    /**
     * The year of this day.
     */
    year: number;
}

