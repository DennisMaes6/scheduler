/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ShiftTypeModelParameters - Holds the shift type specific model parameters.
type ShiftTypeModelParameters struct {

	ShiftType ShiftType `json:"shift_type"`

	// The weight of this shift type in the fairness score.
	ShiftWorkload float32 `json:"shift_workload"`

	// The amount of personnel needed for this particular shift.
	ShiftCoverage float32 `json:"shift_coverage"`

	// The number of assignments per assisant allowed above the minimun for this shift type.
	MaxBuffer int32 `json:"max_buffer"`
}

// AssertShiftTypeModelParametersRequired checks if the required fields are not zero-ed
func AssertShiftTypeModelParametersRequired(obj ShiftTypeModelParameters) error {
	elements := map[string]interface{}{
		"shift_type": obj.ShiftType,
		"shift_workload": obj.ShiftWorkload,
		"shift_coverage": obj.ShiftCoverage,
		"max_buffer": obj.MaxBuffer,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseShiftTypeModelParametersRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ShiftTypeModelParameters (e.g. [][]ShiftTypeModelParameters), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseShiftTypeModelParametersRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aShiftTypeModelParameters, ok := obj.(ShiftTypeModelParameters)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertShiftTypeModelParametersRequired(aShiftTypeModelParameters)
	})
}
