# Assistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identification number of this assistant. | 
**Name** | **string** | The name of this assistant. | 
**Type** | [**AssistantType**](AssistantType.md) |  | 
**FreeDays** | **[]int32** | The day numbers corresponding to the free days granted to this assistant for the current scheduling period. | 

## Methods

### NewAssistant

`func NewAssistant(id int32, name string, type_ AssistantType, freeDays []int32, ) *Assistant`

NewAssistant instantiates a new Assistant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantWithDefaults

`func NewAssistantWithDefaults() *Assistant`

NewAssistantWithDefaults instantiates a new Assistant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Assistant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Assistant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Assistant) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Assistant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Assistant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Assistant) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Assistant) GetType() AssistantType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Assistant) GetTypeOk() (*AssistantType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Assistant) SetType(v AssistantType)`

SetType sets Type field to given value.


### GetFreeDays

`func (o *Assistant) GetFreeDays() []int32`

GetFreeDays returns the FreeDays field if non-nil, zero value otherwise.

### GetFreeDaysOk

`func (o *Assistant) GetFreeDaysOk() (*[]int32, bool)`

GetFreeDaysOk returns a tuple with the FreeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeDays

`func (o *Assistant) SetFreeDays(v []int32)`

SetFreeDays sets FreeDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


