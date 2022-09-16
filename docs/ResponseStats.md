# ResponseStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ResponseStatsData**](ResponseStatsData.md) |  | [optional] 

## Methods

### NewResponseStats

`func NewResponseStats() *ResponseStats`

NewResponseStats instantiates a new ResponseStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStatsWithDefaults

`func NewResponseStatsWithDefaults() *ResponseStats`

NewResponseStatsWithDefaults instantiates a new ResponseStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponseStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ResponseStats) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseStats) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseStats) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseStats) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStats

`func (o *ResponseStats) GetStats() ResponseStatsData`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ResponseStats) GetStatsOk() (*ResponseStatsData, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ResponseStats) SetStats(v ResponseStatsData)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ResponseStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


