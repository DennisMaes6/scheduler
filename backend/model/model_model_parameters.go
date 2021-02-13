/*
 * Scheduler API
 *
 * Basic API for retrieving schedules, based on a ScheduleInput object.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// ModelParameters - Holds all model parameters to generate a schedule.
type ModelParameters struct {

	// The minimal balance score for an acceptable solution
	BalanceMinimum int32 `json:"balance_minimum,omitempty"`

	// The shifts to be included when calculating the balance score.
	BalanceShifts []ShiftType `json:"balance_shifts,omitempty"`

	FairnessShiftWeights []ModelParametersFairnessShiftWeights `json:"fairness_shift_weights,omitempty"`
}
