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
import { ShiftType } from './shiftType';


/**
 * Represents a single shift assignment.
 */
export interface Assignment { 
    /**
     * The day of the scheduling period for which this is an assignment.
     */
    day_nb: number;
    shift_type: ShiftType;
    /**
     * Indicates whether or not this assignment is part of a streak of free days as long as the min_balance score for this schedule.
     */
    part_of_min_balance: boolean;
}

