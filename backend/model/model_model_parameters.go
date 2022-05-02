/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// ModelParameters - Holds all model parameters for generating schedules.
type ModelParameters struct {

	// The minimal balance score for an acceptable solution.
	MinBalance int32 `json:"min_balance,omitempty"`

	// The mininmal balance score for JAEV shifts for an acceptable solution.
	MinBalanceJaev int32 `json:"min_balance_jaev,omitempty"`

	ShiftTypeParameters []ShiftTypeModelParameters `json:"shift_type_parameters,omitempty"`
}
