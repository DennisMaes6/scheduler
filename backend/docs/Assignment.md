# Assignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayNb** | **int32** | The day of the scheduling period for which this is an assignment. | 
**ShiftType** | [**ShiftType**](ShiftType.md) |  | 
**PartOfMinBalance** | **bool** | Indicates whether or not this assignment is part of a streak of free days as long as the min_balance score for this schedule. | 

## Methods

### NewAssignment

`func NewAssignment(dayNb int32, shiftType ShiftType, partOfMinBalance bool, ) *Assignment`

NewAssignment instantiates a new Assignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentWithDefaults

`func NewAssignmentWithDefaults() *Assignment`

NewAssignmentWithDefaults instantiates a new Assignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayNb

`func (o *Assignment) GetDayNb() int32`

GetDayNb returns the DayNb field if non-nil, zero value otherwise.

### GetDayNbOk

`func (o *Assignment) GetDayNbOk() (*int32, bool)`

GetDayNbOk returns a tuple with the DayNb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayNb

`func (o *Assignment) SetDayNb(v int32)`

SetDayNb sets DayNb field to given value.


### GetShiftType

`func (o *Assignment) GetShiftType() ShiftType`

GetShiftType returns the ShiftType field if non-nil, zero value otherwise.

### GetShiftTypeOk

`func (o *Assignment) GetShiftTypeOk() (*ShiftType, bool)`

GetShiftTypeOk returns a tuple with the ShiftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftType

`func (o *Assignment) SetShiftType(v ShiftType)`

SetShiftType sets ShiftType field to given value.


### GetPartOfMinBalance

`func (o *Assignment) GetPartOfMinBalance() bool`

GetPartOfMinBalance returns the PartOfMinBalance field if non-nil, zero value otherwise.

### GetPartOfMinBalanceOk

`func (o *Assignment) GetPartOfMinBalanceOk() (*bool, bool)`

GetPartOfMinBalanceOk returns a tuple with the PartOfMinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfMinBalance

`func (o *Assignment) SetPartOfMinBalance(v bool)`

SetPartOfMinBalance sets PartOfMinBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


