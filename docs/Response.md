# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ErrorRecords** | Pointer to [**[]ResponseErrorRecords**](ResponseErrorRecords.md) |  | [optional] 

## Methods

### NewResponse

`func NewResponse() *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetErrorRecords

`func (o *Response) GetErrorRecords() []ResponseErrorRecords`

GetErrorRecords returns the ErrorRecords field if non-nil, zero value otherwise.

### GetErrorRecordsOk

`func (o *Response) GetErrorRecordsOk() (*[]ResponseErrorRecords, bool)`

GetErrorRecordsOk returns a tuple with the ErrorRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecords

`func (o *Response) SetErrorRecords(v []ResponseErrorRecords)`

SetErrorRecords sets ErrorRecords field to given value.

### HasErrorRecords

`func (o *Response) HasErrorRecords() bool`

HasErrorRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


