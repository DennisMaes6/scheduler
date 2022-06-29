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
 * The skill category of an assistant.
 */
export type AssistantType = 'JA' | 'JA_F' | 'SA' | 'SA_F' | 'SA_NEO' | 'SA_F_NEO' | 'FELLOWS';

export const AssistantType = {
    Ja: 'JA' as AssistantType,
    JaF: 'JA_F' as AssistantType,
    Sa: 'SA' as AssistantType,
    SaF: 'SA_F' as AssistantType,
    SaNeo: 'SA_NEO' as AssistantType,
    SaFNeo: 'SA_F_NEO' as AssistantType,
    Fellows: 'FELLOWS' as AssistantType
};

