# Day

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The day number of this day in the current scheduling period. | 
**Date** | [**DayDate**](DayDate.md) |  | 
**IsHoliday** | **bool** | Indicates whether or not this day should be considered a holiday in the produced schedule. | 

## Methods

### NewDay

`func NewDay(id int32, date DayDate, isHoliday bool, ) *Day`

NewDay instantiates a new Day object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayWithDefaults

`func NewDayWithDefaults() *Day`

NewDayWithDefaults instantiates a new Day object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Day) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Day) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Day) SetId(v int32)`

SetId sets Id field to given value.


### GetDate

`func (o *Day) GetDate() DayDate`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Day) GetDateOk() (*DayDate, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Day) SetDate(v DayDate)`

SetDate sets Date field to given value.


### GetIsHoliday

`func (o *Day) GetIsHoliday() bool`

GetIsHoliday returns the IsHoliday field if non-nil, zero value otherwise.

### GetIsHolidayOk

`func (o *Day) GetIsHolidayOk() (*bool, bool)`

GetIsHolidayOk returns a tuple with the IsHoliday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHoliday

`func (o *Day) SetIsHoliday(v bool)`

SetIsHoliday sets IsHoliday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


