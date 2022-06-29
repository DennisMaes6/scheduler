# ShiftTypeModelParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShiftType** | [**ShiftType**](ShiftType.md) |  | 
**ShiftWorkload** | **float32** | The weight of this shift type in the fairness score. | 
**ShiftCoverage** | **float32** | The amount of personnel needed for this particular shift. | 
**MaxBuffer** | **int32** | The number of assignments per assisant allowed above the minimun for this shift type. | 

## Methods

### NewShiftTypeModelParameters

`func NewShiftTypeModelParameters(shiftType ShiftType, shiftWorkload float32, shiftCoverage float32, maxBuffer int32, ) *ShiftTypeModelParameters`

NewShiftTypeModelParameters instantiates a new ShiftTypeModelParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftTypeModelParametersWithDefaults

`func NewShiftTypeModelParametersWithDefaults() *ShiftTypeModelParameters`

NewShiftTypeModelParametersWithDefaults instantiates a new ShiftTypeModelParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShiftType

`func (o *ShiftTypeModelParameters) GetShiftType() ShiftType`

GetShiftType returns the ShiftType field if non-nil, zero value otherwise.

### GetShiftTypeOk

`func (o *ShiftTypeModelParameters) GetShiftTypeOk() (*ShiftType, bool)`

GetShiftTypeOk returns a tuple with the ShiftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftType

`func (o *ShiftTypeModelParameters) SetShiftType(v ShiftType)`

SetShiftType sets ShiftType field to given value.


### GetShiftWorkload

`func (o *ShiftTypeModelParameters) GetShiftWorkload() float32`

GetShiftWorkload returns the ShiftWorkload field if non-nil, zero value otherwise.

### GetShiftWorkloadOk

`func (o *ShiftTypeModelParameters) GetShiftWorkloadOk() (*float32, bool)`

GetShiftWorkloadOk returns a tuple with the ShiftWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftWorkload

`func (o *ShiftTypeModelParameters) SetShiftWorkload(v float32)`

SetShiftWorkload sets ShiftWorkload field to given value.


### GetShiftCoverage

`func (o *ShiftTypeModelParameters) GetShiftCoverage() float32`

GetShiftCoverage returns the ShiftCoverage field if non-nil, zero value otherwise.

### GetShiftCoverageOk

`func (o *ShiftTypeModelParameters) GetShiftCoverageOk() (*float32, bool)`

GetShiftCoverageOk returns a tuple with the ShiftCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftCoverage

`func (o *ShiftTypeModelParameters) SetShiftCoverage(v float32)`

SetShiftCoverage sets ShiftCoverage field to given value.


### GetMaxBuffer

`func (o *ShiftTypeModelParameters) GetMaxBuffer() int32`

GetMaxBuffer returns the MaxBuffer field if non-nil, zero value otherwise.

### GetMaxBufferOk

`func (o *ShiftTypeModelParameters) GetMaxBufferOk() (*int32, bool)`

GetMaxBufferOk returns a tuple with the MaxBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuffer

`func (o *ShiftTypeModelParameters) SetMaxBuffer(v int32)`

SetMaxBuffer sets MaxBuffer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


