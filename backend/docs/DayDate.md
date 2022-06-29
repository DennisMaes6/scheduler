# DayDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | **int32** | The day of month of this day. | 
**Month** | **int32** | The month of this day. 1 &#x3D; January, 12 &#x3D; December | 
**Year** | **int32** | The year of this day. | 

## Methods

### NewDayDate

`func NewDayDate(day int32, month int32, year int32, ) *DayDate`

NewDayDate instantiates a new DayDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayDateWithDefaults

`func NewDayDateWithDefaults() *DayDate`

NewDayDateWithDefaults instantiates a new DayDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *DayDate) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DayDate) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DayDate) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMonth

`func (o *DayDate) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *DayDate) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *DayDate) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetYear

`func (o *DayDate) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *DayDate) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *DayDate) SetYear(v int32)`

SetYear sets Year field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


