# TestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ErrorRecords** | Pointer to [**[]ResponseErrorRecords**](ResponseErrorRecords.md) |  | [optional] 
**WarningRecords** | Pointer to [**[]ResponseErrorRecords**](ResponseErrorRecords.md) |  | [optional] 
**ValidatedFields** | Pointer to [**[]ValidatedFields**](ValidatedFields.md) |  | [optional] 

## Methods

### NewTestResponse

`func NewTestResponse() *TestResponse`

NewTestResponse instantiates a new TestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResponseWithDefaults

`func NewTestResponseWithDefaults() *TestResponse`

NewTestResponseWithDefaults instantiates a new TestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *TestResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TestResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TestResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TestResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetErrorRecords

`func (o *TestResponse) GetErrorRecords() []ResponseErrorRecords`

GetErrorRecords returns the ErrorRecords field if non-nil, zero value otherwise.

### GetErrorRecordsOk

`func (o *TestResponse) GetErrorRecordsOk() (*[]ResponseErrorRecords, bool)`

GetErrorRecordsOk returns a tuple with the ErrorRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecords

`func (o *TestResponse) SetErrorRecords(v []ResponseErrorRecords)`

SetErrorRecords sets ErrorRecords field to given value.

### HasErrorRecords

`func (o *TestResponse) HasErrorRecords() bool`

HasErrorRecords returns a boolean if a field has been set.

### GetWarningRecords

`func (o *TestResponse) GetWarningRecords() []ResponseErrorRecords`

GetWarningRecords returns the WarningRecords field if non-nil, zero value otherwise.

### GetWarningRecordsOk

`func (o *TestResponse) GetWarningRecordsOk() (*[]ResponseErrorRecords, bool)`

GetWarningRecordsOk returns a tuple with the WarningRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRecords

`func (o *TestResponse) SetWarningRecords(v []ResponseErrorRecords)`

SetWarningRecords sets WarningRecords field to given value.

### HasWarningRecords

`func (o *TestResponse) HasWarningRecords() bool`

HasWarningRecords returns a boolean if a field has been set.

### GetValidatedFields

`func (o *TestResponse) GetValidatedFields() []ValidatedFields`

GetValidatedFields returns the ValidatedFields field if non-nil, zero value otherwise.

### GetValidatedFieldsOk

`func (o *TestResponse) GetValidatedFieldsOk() (*[]ValidatedFields, bool)`

GetValidatedFieldsOk returns a tuple with the ValidatedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedFields

`func (o *TestResponse) SetValidatedFields(v []ValidatedFields)`

SetValidatedFields sets ValidatedFields field to given value.

### HasValidatedFields

`func (o *TestResponse) HasValidatedFields() bool`

HasValidatedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


