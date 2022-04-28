# IndividualSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | **int32** | The identification number of the assistant for which this is an individual schedule. | 
**Workload** | **float32** | The workload of this inidividual schedule. Used when calculating the fairness score. | 
**Assignments** | [**[]Assignment**](Assignment.md) | Contains all the individual assignments of this individual schedule. | 

## Methods

### NewIndividualSchedule

`func NewIndividualSchedule(assistantId int32, workload float32, assignments []Assignment, ) *IndividualSchedule`

NewIndividualSchedule instantiates a new IndividualSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualScheduleWithDefaults

`func NewIndividualScheduleWithDefaults() *IndividualSchedule`

NewIndividualScheduleWithDefaults instantiates a new IndividualSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantId

`func (o *IndividualSchedule) GetAssistantId() int32`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *IndividualSchedule) GetAssistantIdOk() (*int32, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *IndividualSchedule) SetAssistantId(v int32)`

SetAssistantId sets AssistantId field to given value.


### GetWorkload

`func (o *IndividualSchedule) GetWorkload() float32`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *IndividualSchedule) GetWorkloadOk() (*float32, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *IndividualSchedule) SetWorkload(v float32)`

SetWorkload sets Workload field to given value.


### GetAssignments

`func (o *IndividualSchedule) GetAssignments() []Assignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *IndividualSchedule) GetAssignmentsOk() (*[]Assignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *IndividualSchedule) SetAssignments(v []Assignment)`

SetAssignments sets Assignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


