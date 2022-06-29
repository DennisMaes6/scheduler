# InstanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistants** | [**[]Assistant**](Assistant.md) | The assistant instances representing the assistants for which to produce a schedule. | 
**Days** | [**[]Day**](Day.md) | The days for which to produce a schedule. | 

## Methods

### NewInstanceData

`func NewInstanceData(assistants []Assistant, days []Day, ) *InstanceData`

NewInstanceData instantiates a new InstanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDataWithDefaults

`func NewInstanceDataWithDefaults() *InstanceData`

NewInstanceDataWithDefaults instantiates a new InstanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistants

`func (o *InstanceData) GetAssistants() []Assistant`

GetAssistants returns the Assistants field if non-nil, zero value otherwise.

### GetAssistantsOk

`func (o *InstanceData) GetAssistantsOk() (*[]Assistant, bool)`

GetAssistantsOk returns a tuple with the Assistants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistants

`func (o *InstanceData) SetAssistants(v []Assistant)`

SetAssistants sets Assistants field to given value.


### GetDays

`func (o *InstanceData) GetDays() []Day`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *InstanceData) GetDaysOk() (*[]Day, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *InstanceData) SetDays(v []Day)`

SetDays sets Days field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


