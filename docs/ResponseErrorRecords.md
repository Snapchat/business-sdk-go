# ResponseErrorRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**RecordIndexes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewResponseErrorRecords

`func NewResponseErrorRecords() *ResponseErrorRecords`

NewResponseErrorRecords instantiates a new ResponseErrorRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseErrorRecordsWithDefaults

`func NewResponseErrorRecordsWithDefaults() *ResponseErrorRecords`

NewResponseErrorRecordsWithDefaults instantiates a new ResponseErrorRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ResponseErrorRecords) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseErrorRecords) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseErrorRecords) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseErrorRecords) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRecordIndexes

`func (o *ResponseErrorRecords) GetRecordIndexes() []int32`

GetRecordIndexes returns the RecordIndexes field if non-nil, zero value otherwise.

### GetRecordIndexesOk

`func (o *ResponseErrorRecords) GetRecordIndexesOk() (*[]int32, bool)`

GetRecordIndexesOk returns a tuple with the RecordIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIndexes

`func (o *ResponseErrorRecords) SetRecordIndexes(v []int32)`

SetRecordIndexes sets RecordIndexes field to given value.

### HasRecordIndexes

`func (o *ResponseErrorRecords) HasRecordIndexes() bool`

HasRecordIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


