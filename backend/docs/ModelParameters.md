# ModelParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinBalance** | Pointer to **int32** | The minimal balance score for an acceptable solution. | [optional] 
**MinBalanceJaev** | Pointer to **int32** | The mininmal balance score for JAEV shifts for an acceptable solution. | [optional] 
**ShiftTypeParameters** | Pointer to [**[]ShiftTypeModelParameters**](ShiftTypeModelParameters.md) |  | [optional] 

## Methods

### NewModelParameters

`func NewModelParameters() *ModelParameters`

NewModelParameters instantiates a new ModelParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelParametersWithDefaults

`func NewModelParametersWithDefaults() *ModelParameters`

NewModelParametersWithDefaults instantiates a new ModelParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinBalance

`func (o *ModelParameters) GetMinBalance() int32`

GetMinBalance returns the MinBalance field if non-nil, zero value otherwise.

### GetMinBalanceOk

`func (o *ModelParameters) GetMinBalanceOk() (*int32, bool)`

GetMinBalanceOk returns a tuple with the MinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalance

`func (o *ModelParameters) SetMinBalance(v int32)`

SetMinBalance sets MinBalance field to given value.

### HasMinBalance

`func (o *ModelParameters) HasMinBalance() bool`

HasMinBalance returns a boolean if a field has been set.

### GetMinBalanceJaev

`func (o *ModelParameters) GetMinBalanceJaev() int32`

GetMinBalanceJaev returns the MinBalanceJaev field if non-nil, zero value otherwise.

### GetMinBalanceJaevOk

`func (o *ModelParameters) GetMinBalanceJaevOk() (*int32, bool)`

GetMinBalanceJaevOk returns a tuple with the MinBalanceJaev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBalanceJaev

`func (o *ModelParameters) SetMinBalanceJaev(v int32)`

SetMinBalanceJaev sets MinBalanceJaev field to given value.

### HasMinBalanceJaev

`func (o *ModelParameters) HasMinBalanceJaev() bool`

HasMinBalanceJaev returns a boolean if a field has been set.

### GetShiftTypeParameters

`func (o *ModelParameters) GetShiftTypeParameters() []ShiftTypeModelParameters`

GetShiftTypeParameters returns the ShiftTypeParameters field if non-nil, zero value otherwise.

### GetShiftTypeParametersOk

`func (o *ModelParameters) GetShiftTypeParametersOk() (*[]ShiftTypeModelParameters, bool)`

GetShiftTypeParametersOk returns a tuple with the ShiftTypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftTypeParameters

`func (o *ModelParameters) SetShiftTypeParameters(v []ShiftTypeModelParameters)`

SetShiftTypeParameters sets ShiftTypeParameters field to given value.

### HasShiftTypeParameters

`func (o *ModelParameters) HasShiftTypeParameters() bool`

HasShiftTypeParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


