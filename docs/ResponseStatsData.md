# ResponseStatsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | Pointer to [**ResponseStatsTest**](ResponseStatsTest.md) |  | [optional] 
**Live** | Pointer to [**ResponseStatsTest**](ResponseStatsTest.md) |  | [optional] 

## Methods

### NewResponseStatsData

`func NewResponseStatsData() *ResponseStatsData`

NewResponseStatsData instantiates a new ResponseStatsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStatsDataWithDefaults

`func NewResponseStatsDataWithDefaults() *ResponseStatsData`

NewResponseStatsDataWithDefaults instantiates a new ResponseStatsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *ResponseStatsData) GetTest() ResponseStatsTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *ResponseStatsData) GetTestOk() (*ResponseStatsTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *ResponseStatsData) SetTest(v ResponseStatsTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *ResponseStatsData) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetLive

`func (o *ResponseStatsData) GetLive() ResponseStatsTest`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *ResponseStatsData) GetLiveOk() (*ResponseStatsTest, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *ResponseStatsData) SetLive(v ResponseStatsTest)`

SetLive sets Live field to given value.

### HasLive

`func (o *ResponseStatsData) HasLive() bool`

HasLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


