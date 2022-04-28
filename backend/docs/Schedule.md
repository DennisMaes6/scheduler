# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FairnessScore** | **float32** | The fairness score of this schedule. | 
**BalanceScore** | **float32** | The balance score of this schedule. | 
**JaevFairnessScore** | **float32** | The fairness score for the JAEV shifts of this schedule. | 
**JaevBalanceScore** | **float32** | The balance score for the JAEV shifts of this schedule. | 
**IndividualSchedules** | Pointer to [**[]IndividualSchedule**](IndividualSchedule.md) |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule(fairnessScore float32, balanceScore float32, jaevFairnessScore float32, jaevBalanceScore float32, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFairnessScore

`func (o *Schedule) GetFairnessScore() float32`

GetFairnessScore returns the FairnessScore field if non-nil, zero value otherwise.

### GetFairnessScoreOk

`func (o *Schedule) GetFairnessScoreOk() (*float32, bool)`

GetFairnessScoreOk returns a tuple with the FairnessScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairnessScore

`func (o *Schedule) SetFairnessScore(v float32)`

SetFairnessScore sets FairnessScore field to given value.


### GetBalanceScore

`func (o *Schedule) GetBalanceScore() float32`

GetBalanceScore returns the BalanceScore field if non-nil, zero value otherwise.

### GetBalanceScoreOk

`func (o *Schedule) GetBalanceScoreOk() (*float32, bool)`

GetBalanceScoreOk returns a tuple with the BalanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceScore

`func (o *Schedule) SetBalanceScore(v float32)`

SetBalanceScore sets BalanceScore field to given value.


### GetJaevFairnessScore

`func (o *Schedule) GetJaevFairnessScore() float32`

GetJaevFairnessScore returns the JaevFairnessScore field if non-nil, zero value otherwise.

### GetJaevFairnessScoreOk

`func (o *Schedule) GetJaevFairnessScoreOk() (*float32, bool)`

GetJaevFairnessScoreOk returns a tuple with the JaevFairnessScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJaevFairnessScore

`func (o *Schedule) SetJaevFairnessScore(v float32)`

SetJaevFairnessScore sets JaevFairnessScore field to given value.


### GetJaevBalanceScore

`func (o *Schedule) GetJaevBalanceScore() float32`

GetJaevBalanceScore returns the JaevBalanceScore field if non-nil, zero value otherwise.

### GetJaevBalanceScoreOk

`func (o *Schedule) GetJaevBalanceScoreOk() (*float32, bool)`

GetJaevBalanceScoreOk returns a tuple with the JaevBalanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJaevBalanceScore

`func (o *Schedule) SetJaevBalanceScore(v float32)`

SetJaevBalanceScore sets JaevBalanceScore field to given value.


### GetIndividualSchedules

`func (o *Schedule) GetIndividualSchedules() []IndividualSchedule`

GetIndividualSchedules returns the IndividualSchedules field if non-nil, zero value otherwise.

### GetIndividualSchedulesOk

`func (o *Schedule) GetIndividualSchedulesOk() (*[]IndividualSchedule, bool)`

GetIndividualSchedulesOk returns a tuple with the IndividualSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualSchedules

`func (o *Schedule) SetIndividualSchedules(v []IndividualSchedule)`

SetIndividualSchedules sets IndividualSchedules field to given value.

### HasIndividualSchedules

`func (o *Schedule) HasIndividualSchedules() bool`

HasIndividualSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


