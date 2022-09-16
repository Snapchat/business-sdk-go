# ResponseLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Logs** | Pointer to [**[]ResponseLogsLog**](ResponseLogsLog.md) |  | [optional] 

## Methods

### NewResponseLogs

`func NewResponseLogs() *ResponseLogs`

NewResponseLogs instantiates a new ResponseLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLogsWithDefaults

`func NewResponseLogsWithDefaults() *ResponseLogs`

NewResponseLogsWithDefaults instantiates a new ResponseLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponseLogs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseLogs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseLogs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseLogs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ResponseLogs) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseLogs) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseLogs) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseLogs) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLogs

`func (o *ResponseLogs) GetLogs() []ResponseLogsLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ResponseLogs) GetLogsOk() (*[]ResponseLogsLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ResponseLogs) SetLogs(v []ResponseLogsLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ResponseLogs) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


