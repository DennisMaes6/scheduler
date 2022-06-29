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
 * Holds the shift type specific model parameters.
 */
export interface WeightParameters { 
    /**
     * The weight of coverage in the fitness function of schedule
     */
    coverage: number;
    /**
     * The weight of balance in the fitness function of schedule
     */
    balance: number;
    /**
     * The weight of fairness in the fitness function of schedule
     */
    fairness: number;
}

