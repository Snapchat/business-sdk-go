# ResponseLogsLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventConversionType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to **string** |  | [optional] 
**EventMetadata** | Pointer to [**CapiEvent**](CapiEvent.md) |  | [optional] 
**ErrorRecords** | Pointer to **[]string** |  | [optional] 
**WarningRecords** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResponseLogsLog

`func NewResponseLogsLog() *ResponseLogsLog`

NewResponseLogsLog instantiates a new ResponseLogsLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLogsLogWithDefaults

`func NewResponseLogsLogWithDefaults() *ResponseLogsLog`

NewResponseLogsLogWithDefaults instantiates a new ResponseLogsLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ResponseLogsLog) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseLogsLog) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseLogsLog) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ResponseLogsLog) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEventType

`func (o *ResponseLogsLog) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ResponseLogsLog) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ResponseLogsLog) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ResponseLogsLog) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventConversionType

`func (o *ResponseLogsLog) GetEventConversionType() string`

GetEventConversionType returns the EventConversionType field if non-nil, zero value otherwise.

### GetEventConversionTypeOk

`func (o *ResponseLogsLog) GetEventConversionTypeOk() (*string, bool)`

GetEventConversionTypeOk returns a tuple with the EventConversionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventConversionType

`func (o *ResponseLogsLog) SetEventConversionType(v string)`

SetEventConversionType sets EventConversionType field to given value.

### HasEventConversionType

`func (o *ResponseLogsLog) HasEventConversionType() bool`

HasEventConversionType returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseLogsLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseLogsLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseLogsLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseLogsLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIntegration

`func (o *ResponseLogsLog) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ResponseLogsLog) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ResponseLogsLog) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ResponseLogsLog) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetEventMetadata

`func (o *ResponseLogsLog) GetEventMetadata() CapiEvent`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *ResponseLogsLog) GetEventMetadataOk() (*CapiEvent, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *ResponseLogsLog) SetEventMetadata(v CapiEvent)`

SetEventMetadata sets EventMetadata field to given value.

### HasEventMetadata

`func (o *ResponseLogsLog) HasEventMetadata() bool`

HasEventMetadata returns a boolean if a field has been set.

### GetErrorRecords

`func (o *ResponseLogsLog) GetErrorRecords() []string`

GetErrorRecords returns the ErrorRecords field if non-nil, zero value otherwise.

### GetErrorRecordsOk

`func (o *ResponseLogsLog) GetErrorRecordsOk() (*[]string, bool)`

GetErrorRecordsOk returns a tuple with the ErrorRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecords

`func (o *ResponseLogsLog) SetErrorRecords(v []string)`

SetErrorRecords sets ErrorRecords field to given value.

### HasErrorRecords

`func (o *ResponseLogsLog) HasErrorRecords() bool`

HasErrorRecords returns a boolean if a field has been set.

### GetWarningRecords

`func (o *ResponseLogsLog) GetWarningRecords() []string`

GetWarningRecords returns the WarningRecords field if non-nil, zero value otherwise.

### GetWarningRecordsOk

`func (o *ResponseLogsLog) GetWarningRecordsOk() (*[]string, bool)`

GetWarningRecordsOk returns a tuple with the WarningRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRecords

`func (o *ResponseLogsLog) SetWarningRecords(v []string)`

SetWarningRecords sets WarningRecords field to given value.

### HasWarningRecords

`func (o *ResponseLogsLog) HasWarningRecords() bool`

HasWarningRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


