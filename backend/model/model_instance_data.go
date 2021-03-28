/*
 * Scheduler API
 *
 * Basic API for retrieving schedules, based on a ScheduleInput object.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// InstanceData - Holds the instance data for scheduling.
type InstanceData struct {
	Assistants []AssistantInstance `json:"assistants"`

	// The number of weeks to be scheduled in this instance
	NbWeeks int32 `json:"nb_weeks"`

	Holidays []int32 `json:"holidays"`
}
