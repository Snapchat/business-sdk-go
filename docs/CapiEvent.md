# CapiEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PixelId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**SnapAppId** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventConversionType** | Pointer to **string** |  | [optional] 
**EventTag** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**HashedEmail** | Pointer to **string** |  | [optional] 
**HashedMobileAdId** | Pointer to **string** |  | [optional] 
**UuidC1** | Pointer to **string** |  | [optional] 
**HashedIdfv** | Pointer to **string** |  | [optional] 
**HashedPhoneNumber** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**HashedIpAddress** | Pointer to **string** |  | [optional] 
**ItemCategory** | Pointer to **string** |  | [optional] 
**ItemIds** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NumberItems** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**ClientDedupId** | Pointer to **string** |  | [optional] 
**DataUse** | Pointer to **string** |  | [optional] 
**SearchString** | Pointer to **string** |  | [optional] 
**PageUrl** | Pointer to **string** |  | [optional] 
**SignUpMethod** | Pointer to **string** |  | [optional] 
**HashedFirstNameSha** | Pointer to **string** |  | [optional] 
**HashedFirstNameSdx** | Pointer to **string** |  | [optional] 
**HashedMiddleNameSha** | Pointer to **string** |  | [optional] 
**HashedMiddleNameSdx** | Pointer to **string** |  | [optional] 
**HashedLastNameSha** | Pointer to **string** |  | [optional] 
**HashedLastNameSdx** | Pointer to **string** |  | [optional] 
**HashedCitySha** | Pointer to **string** |  | [optional] 
**HashedCitySdx** | Pointer to **string** |  | [optional] 
**HashedStateSha** | Pointer to **string** |  | [optional] 
**HashedStateSdx** | Pointer to **string** |  | [optional] 
**HashedZip** | Pointer to **string** |  | [optional] 
**HashedDobMonth** | Pointer to **string** |  | [optional] 
**HashedDobDay** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to **string** |  | [optional] 
**ClickId** | Pointer to **string** |  | [optional] 

## Methods

### NewCapiEvent

`func NewCapiEvent() *CapiEvent`

NewCapiEvent instantiates a new CapiEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapiEventWithDefaults

`func NewCapiEventWithDefaults() *CapiEvent`

NewCapiEventWithDefaults instantiates a new CapiEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPixelId

`func (o *CapiEvent) GetPixelId() string`

GetPixelId returns the PixelId field if non-nil, zero value otherwise.

### GetPixelIdOk

`func (o *CapiEvent) GetPixelIdOk() (*string, bool)`

GetPixelIdOk returns a tuple with the PixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelId

`func (o *CapiEvent) SetPixelId(v string)`

SetPixelId sets PixelId field to given value.

### HasPixelId

`func (o *CapiEvent) HasPixelId() bool`

HasPixelId returns a boolean if a field has been set.

### GetAppId

`func (o *CapiEvent) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CapiEvent) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CapiEvent) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CapiEvent) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetSnapAppId

`func (o *CapiEvent) GetSnapAppId() string`

GetSnapAppId returns the SnapAppId field if non-nil, zero value otherwise.

### GetSnapAppIdOk

`func (o *CapiEvent) GetSnapAppIdOk() (*string, bool)`

GetSnapAppIdOk returns a tuple with the SnapAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapAppId

`func (o *CapiEvent) SetSnapAppId(v string)`

SetSnapAppId sets SnapAppId field to given value.

### HasSnapAppId

`func (o *CapiEvent) HasSnapAppId() bool`

HasSnapAppId returns a boolean if a field has been set.

### GetEventType

`func (o *CapiEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CapiEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CapiEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CapiEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventConversionType

`func (o *CapiEvent) GetEventConversionType() string`

GetEventConversionType returns the EventConversionType field if non-nil, zero value otherwise.

### GetEventConversionTypeOk

`func (o *CapiEvent) GetEventConversionTypeOk() (*string, bool)`

GetEventConversionTypeOk returns a tuple with the EventConversionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventConversionType

`func (o *CapiEvent) SetEventConversionType(v string)`

SetEventConversionType sets EventConversionType field to given value.

### HasEventConversionType

`func (o *CapiEvent) HasEventConversionType() bool`

HasEventConversionType returns a boolean if a field has been set.

### GetEventTag

`func (o *CapiEvent) GetEventTag() string`

GetEventTag returns the EventTag field if non-nil, zero value otherwise.

### GetEventTagOk

`func (o *CapiEvent) GetEventTagOk() (*string, bool)`

GetEventTagOk returns a tuple with the EventTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTag

`func (o *CapiEvent) SetEventTag(v string)`

SetEventTag sets EventTag field to given value.

### HasEventTag

`func (o *CapiEvent) HasEventTag() bool`

HasEventTag returns a boolean if a field has been set.

### GetTimestamp

`func (o *CapiEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CapiEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CapiEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CapiEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHashedEmail

`func (o *CapiEvent) GetHashedEmail() string`

GetHashedEmail returns the HashedEmail field if non-nil, zero value otherwise.

### GetHashedEmailOk

`func (o *CapiEvent) GetHashedEmailOk() (*string, bool)`

GetHashedEmailOk returns a tuple with the HashedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedEmail

`func (o *CapiEvent) SetHashedEmail(v string)`

SetHashedEmail sets HashedEmail field to given value.

### HasHashedEmail

`func (o *CapiEvent) HasHashedEmail() bool`

HasHashedEmail returns a boolean if a field has been set.

### GetHashedMobileAdId

`func (o *CapiEvent) GetHashedMobileAdId() string`

GetHashedMobileAdId returns the HashedMobileAdId field if non-nil, zero value otherwise.

### GetHashedMobileAdIdOk

`func (o *CapiEvent) GetHashedMobileAdIdOk() (*string, bool)`

GetHashedMobileAdIdOk returns a tuple with the HashedMobileAdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedMobileAdId

`func (o *CapiEvent) SetHashedMobileAdId(v string)`

SetHashedMobileAdId sets HashedMobileAdId field to given value.

### HasHashedMobileAdId

`func (o *CapiEvent) HasHashedMobileAdId() bool`

HasHashedMobileAdId returns a boolean if a field has been set.

### GetUuidC1

`func (o *CapiEvent) GetUuidC1() string`

GetUuidC1 returns the UuidC1 field if non-nil, zero value otherwise.

### GetUuidC1Ok

`func (o *CapiEvent) GetUuidC1Ok() (*string, bool)`

GetUuidC1Ok returns a tuple with the UuidC1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidC1

`func (o *CapiEvent) SetUuidC1(v string)`

SetUuidC1 sets UuidC1 field to given value.

### HasUuidC1

`func (o *CapiEvent) HasUuidC1() bool`

HasUuidC1 returns a boolean if a field has been set.

### GetHashedIdfv

`func (o *CapiEvent) GetHashedIdfv() string`

GetHashedIdfv returns the HashedIdfv field if non-nil, zero value otherwise.

### GetHashedIdfvOk

`func (o *CapiEvent) GetHashedIdfvOk() (*string, bool)`

GetHashedIdfvOk returns a tuple with the HashedIdfv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedIdfv

`func (o *CapiEvent) SetHashedIdfv(v string)`

SetHashedIdfv sets HashedIdfv field to given value.

### HasHashedIdfv

`func (o *CapiEvent) HasHashedIdfv() bool`

HasHashedIdfv returns a boolean if a field has been set.

### GetHashedPhoneNumber

`func (o *CapiEvent) GetHashedPhoneNumber() string`

GetHashedPhoneNumber returns the HashedPhoneNumber field if non-nil, zero value otherwise.

### GetHashedPhoneNumberOk

`func (o *CapiEvent) GetHashedPhoneNumberOk() (*string, bool)`

GetHashedPhoneNumberOk returns a tuple with the HashedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedPhoneNumber

`func (o *CapiEvent) SetHashedPhoneNumber(v string)`

SetHashedPhoneNumber sets HashedPhoneNumber field to given value.

### HasHashedPhoneNumber

`func (o *CapiEvent) HasHashedPhoneNumber() bool`

HasHashedPhoneNumber returns a boolean if a field has been set.

### GetUserAgent

`func (o *CapiEvent) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CapiEvent) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CapiEvent) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CapiEvent) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetHashedIpAddress

`func (o *CapiEvent) GetHashedIpAddress() string`

GetHashedIpAddress returns the HashedIpAddress field if non-nil, zero value otherwise.

### GetHashedIpAddressOk

`func (o *CapiEvent) GetHashedIpAddressOk() (*string, bool)`

GetHashedIpAddressOk returns a tuple with the HashedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedIpAddress

`func (o *CapiEvent) SetHashedIpAddress(v string)`

SetHashedIpAddress sets HashedIpAddress field to given value.

### HasHashedIpAddress

`func (o *CapiEvent) HasHashedIpAddress() bool`

HasHashedIpAddress returns a boolean if a field has been set.

### GetItemCategory

`func (o *CapiEvent) GetItemCategory() string`

GetItemCategory returns the ItemCategory field if non-nil, zero value otherwise.

### GetItemCategoryOk

`func (o *CapiEvent) GetItemCategoryOk() (*string, bool)`

GetItemCategoryOk returns a tuple with the ItemCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCategory

`func (o *CapiEvent) SetItemCategory(v string)`

SetItemCategory sets ItemCategory field to given value.

### HasItemCategory

`func (o *CapiEvent) HasItemCategory() bool`

HasItemCategory returns a boolean if a field has been set.

### GetItemIds

`func (o *CapiEvent) GetItemIds() string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *CapiEvent) GetItemIdsOk() (*string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *CapiEvent) SetItemIds(v string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *CapiEvent) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetDescription

`func (o *CapiEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapiEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapiEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapiEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNumberItems

`func (o *CapiEvent) GetNumberItems() string`

GetNumberItems returns the NumberItems field if non-nil, zero value otherwise.

### GetNumberItemsOk

`func (o *CapiEvent) GetNumberItemsOk() (*string, bool)`

GetNumberItemsOk returns a tuple with the NumberItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberItems

`func (o *CapiEvent) SetNumberItems(v string)`

SetNumberItems sets NumberItems field to given value.

### HasNumberItems

`func (o *CapiEvent) HasNumberItems() bool`

HasNumberItems returns a boolean if a field has been set.

### GetPrice

`func (o *CapiEvent) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CapiEvent) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CapiEvent) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CapiEvent) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *CapiEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CapiEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CapiEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CapiEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTransactionId

`func (o *CapiEvent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CapiEvent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CapiEvent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CapiEvent) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetLevel

`func (o *CapiEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CapiEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CapiEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CapiEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetClientDedupId

`func (o *CapiEvent) GetClientDedupId() string`

GetClientDedupId returns the ClientDedupId field if non-nil, zero value otherwise.

### GetClientDedupIdOk

`func (o *CapiEvent) GetClientDedupIdOk() (*string, bool)`

GetClientDedupIdOk returns a tuple with the ClientDedupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDedupId

`func (o *CapiEvent) SetClientDedupId(v string)`

SetClientDedupId sets ClientDedupId field to given value.

### HasClientDedupId

`func (o *CapiEvent) HasClientDedupId() bool`

HasClientDedupId returns a boolean if a field has been set.

### GetDataUse

`func (o *CapiEvent) GetDataUse() string`

GetDataUse returns the DataUse field if non-nil, zero value otherwise.

### GetDataUseOk

`func (o *CapiEvent) GetDataUseOk() (*string, bool)`

GetDataUseOk returns a tuple with the DataUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUse

`func (o *CapiEvent) SetDataUse(v string)`

SetDataUse sets DataUse field to given value.

### HasDataUse

`func (o *CapiEvent) HasDataUse() bool`

HasDataUse returns a boolean if a field has been set.

### GetSearchString

`func (o *CapiEvent) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *CapiEvent) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *CapiEvent) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.

### HasSearchString

`func (o *CapiEvent) HasSearchString() bool`

HasSearchString returns a boolean if a field has been set.

### GetPageUrl

`func (o *CapiEvent) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *CapiEvent) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *CapiEvent) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *CapiEvent) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### GetSignUpMethod

`func (o *CapiEvent) GetSignUpMethod() string`

GetSignUpMethod returns the SignUpMethod field if non-nil, zero value otherwise.

### GetSignUpMethodOk

`func (o *CapiEvent) GetSignUpMethodOk() (*string, bool)`

GetSignUpMethodOk returns a tuple with the SignUpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpMethod

`func (o *CapiEvent) SetSignUpMethod(v string)`

SetSignUpMethod sets SignUpMethod field to given value.

### HasSignUpMethod

`func (o *CapiEvent) HasSignUpMethod() bool`

HasSignUpMethod returns a boolean if a field has been set.

### GetHashedFirstNameSha

`func (o *CapiEvent) GetHashedFirstNameSha() string`

GetHashedFirstNameSha returns the HashedFirstNameSha field if non-nil, zero value otherwise.

### GetHashedFirstNameShaOk

`func (o *CapiEvent) GetHashedFirstNameShaOk() (*string, bool)`

GetHashedFirstNameShaOk returns a tuple with the HashedFirstNameSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedFirstNameSha

`func (o *CapiEvent) SetHashedFirstNameSha(v string)`

SetHashedFirstNameSha sets HashedFirstNameSha field to given value.

### HasHashedFirstNameSha

`func (o *CapiEvent) HasHashedFirstNameSha() bool`

HasHashedFirstNameSha returns a boolean if a field has been set.

### GetHashedFirstNameSdx

`func (o *CapiEvent) GetHashedFirstNameSdx() string`

GetHashedFirstNameSdx returns the HashedFirstNameSdx field if non-nil, zero value otherwise.

### GetHashedFirstNameSdxOk

`func (o *CapiEvent) GetHashedFirstNameSdxOk() (*string, bool)`

GetHashedFirstNameSdxOk returns a tuple with the HashedFirstNameSdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedFirstNameSdx

`func (o *CapiEvent) SetHashedFirstNameSdx(v string)`

SetHashedFirstNameSdx sets HashedFirstNameSdx field to given value.

### HasHashedFirstNameSdx

`func (o *CapiEvent) HasHashedFirstNameSdx() bool`

HasHashedFirstNameSdx returns a boolean if a field has been set.

### GetHashedMiddleNameSha

`func (o *CapiEvent) GetHashedMiddleNameSha() string`

GetHashedMiddleNameSha returns the HashedMiddleNameSha field if non-nil, zero value otherwise.

### GetHashedMiddleNameShaOk

`func (o *CapiEvent) GetHashedMiddleNameShaOk() (*string, bool)`

GetHashedMiddleNameShaOk returns a tuple with the HashedMiddleNameSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedMiddleNameSha

`func (o *CapiEvent) SetHashedMiddleNameSha(v string)`

SetHashedMiddleNameSha sets HashedMiddleNameSha field to given value.

### HasHashedMiddleNameSha

`func (o *CapiEvent) HasHashedMiddleNameSha() bool`

HasHashedMiddleNameSha returns a boolean if a field has been set.

### GetHashedMiddleNameSdx

`func (o *CapiEvent) GetHashedMiddleNameSdx() string`

GetHashedMiddleNameSdx returns the HashedMiddleNameSdx field if non-nil, zero value otherwise.

### GetHashedMiddleNameSdxOk

`func (o *CapiEvent) GetHashedMiddleNameSdxOk() (*string, bool)`

GetHashedMiddleNameSdxOk returns a tuple with the HashedMiddleNameSdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedMiddleNameSdx

`func (o *CapiEvent) SetHashedMiddleNameSdx(v string)`

SetHashedMiddleNameSdx sets HashedMiddleNameSdx field to given value.

### HasHashedMiddleNameSdx

`func (o *CapiEvent) HasHashedMiddleNameSdx() bool`

HasHashedMiddleNameSdx returns a boolean if a field has been set.

### GetHashedLastNameSha

`func (o *CapiEvent) GetHashedLastNameSha() string`

GetHashedLastNameSha returns the HashedLastNameSha field if non-nil, zero value otherwise.

### GetHashedLastNameShaOk

`func (o *CapiEvent) GetHashedLastNameShaOk() (*string, bool)`

GetHashedLastNameShaOk returns a tuple with the HashedLastNameSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedLastNameSha

`func (o *CapiEvent) SetHashedLastNameSha(v string)`

SetHashedLastNameSha sets HashedLastNameSha field to given value.

### HasHashedLastNameSha

`func (o *CapiEvent) HasHashedLastNameSha() bool`

HasHashedLastNameSha returns a boolean if a field has been set.

### GetHashedLastNameSdx

`func (o *CapiEvent) GetHashedLastNameSdx() string`

GetHashedLastNameSdx returns the HashedLastNameSdx field if non-nil, zero value otherwise.

### GetHashedLastNameSdxOk

`func (o *CapiEvent) GetHashedLastNameSdxOk() (*string, bool)`

GetHashedLastNameSdxOk returns a tuple with the HashedLastNameSdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedLastNameSdx

`func (o *CapiEvent) SetHashedLastNameSdx(v string)`

SetHashedLastNameSdx sets HashedLastNameSdx field to given value.

### HasHashedLastNameSdx

`func (o *CapiEvent) HasHashedLastNameSdx() bool`

HasHashedLastNameSdx returns a boolean if a field has been set.

### GetHashedCitySha

`func (o *CapiEvent) GetHashedCitySha() string`

GetHashedCitySha returns the HashedCitySha field if non-nil, zero value otherwise.

### GetHashedCityShaOk

`func (o *CapiEvent) GetHashedCityShaOk() (*string, bool)`

GetHashedCityShaOk returns a tuple with the HashedCitySha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedCitySha

`func (o *CapiEvent) SetHashedCitySha(v string)`

SetHashedCitySha sets HashedCitySha field to given value.

### HasHashedCitySha

`func (o *CapiEvent) HasHashedCitySha() bool`

HasHashedCitySha returns a boolean if a field has been set.

### GetHashedCitySdx

`func (o *CapiEvent) GetHashedCitySdx() string`

GetHashedCitySdx returns the HashedCitySdx field if non-nil, zero value otherwise.

### GetHashedCitySdxOk

`func (o *CapiEvent) GetHashedCitySdxOk() (*string, bool)`

GetHashedCitySdxOk returns a tuple with the HashedCitySdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedCitySdx

`func (o *CapiEvent) SetHashedCitySdx(v string)`

SetHashedCitySdx sets HashedCitySdx field to given value.

### HasHashedCitySdx

`func (o *CapiEvent) HasHashedCitySdx() bool`

HasHashedCitySdx returns a boolean if a field has been set.

### GetHashedStateSha

`func (o *CapiEvent) GetHashedStateSha() string`

GetHashedStateSha returns the HashedStateSha field if non-nil, zero value otherwise.

### GetHashedStateShaOk

`func (o *CapiEvent) GetHashedStateShaOk() (*string, bool)`

GetHashedStateShaOk returns a tuple with the HashedStateSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedStateSha

`func (o *CapiEvent) SetHashedStateSha(v string)`

SetHashedStateSha sets HashedStateSha field to given value.

### HasHashedStateSha

`func (o *CapiEvent) HasHashedStateSha() bool`

HasHashedStateSha returns a boolean if a field has been set.

### GetHashedStateSdx

`func (o *CapiEvent) GetHashedStateSdx() string`

GetHashedStateSdx returns the HashedStateSdx field if non-nil, zero value otherwise.

### GetHashedStateSdxOk

`func (o *CapiEvent) GetHashedStateSdxOk() (*string, bool)`

GetHashedStateSdxOk returns a tuple with the HashedStateSdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedStateSdx

`func (o *CapiEvent) SetHashedStateSdx(v string)`

SetHashedStateSdx sets HashedStateSdx field to given value.

### HasHashedStateSdx

`func (o *CapiEvent) HasHashedStateSdx() bool`

HasHashedStateSdx returns a boolean if a field has been set.

### GetHashedZip

`func (o *CapiEvent) GetHashedZip() string`

GetHashedZip returns the HashedZip field if non-nil, zero value otherwise.

### GetHashedZipOk

`func (o *CapiEvent) GetHashedZipOk() (*string, bool)`

GetHashedZipOk returns a tuple with the HashedZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedZip

`func (o *CapiEvent) SetHashedZip(v string)`

SetHashedZip sets HashedZip field to given value.

### HasHashedZip

`func (o *CapiEvent) HasHashedZip() bool`

HasHashedZip returns a boolean if a field has been set.

### GetHashedDobMonth

`func (o *CapiEvent) GetHashedDobMonth() string`

GetHashedDobMonth returns the HashedDobMonth field if non-nil, zero value otherwise.

### GetHashedDobMonthOk

`func (o *CapiEvent) GetHashedDobMonthOk() (*string, bool)`

GetHashedDobMonthOk returns a tuple with the HashedDobMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedDobMonth

`func (o *CapiEvent) SetHashedDobMonth(v string)`

SetHashedDobMonth sets HashedDobMonth field to given value.

### HasHashedDobMonth

`func (o *CapiEvent) HasHashedDobMonth() bool`

HasHashedDobMonth returns a boolean if a field has been set.

### GetHashedDobDay

`func (o *CapiEvent) GetHashedDobDay() string`

GetHashedDobDay returns the HashedDobDay field if non-nil, zero value otherwise.

### GetHashedDobDayOk

`func (o *CapiEvent) GetHashedDobDayOk() (*string, bool)`

GetHashedDobDayOk returns a tuple with the HashedDobDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedDobDay

`func (o *CapiEvent) SetHashedDobDay(v string)`

SetHashedDobDay sets HashedDobDay field to given value.

### HasHashedDobDay

`func (o *CapiEvent) HasHashedDobDay() bool`

HasHashedDobDay returns a boolean if a field has been set.

### GetIntegration

`func (o *CapiEvent) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CapiEvent) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CapiEvent) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CapiEvent) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetClickId

`func (o *CapiEvent) GetClickId() string`

GetClickId returns the ClickId field if non-nil, zero value otherwise.

### GetClickIdOk

`func (o *CapiEvent) GetClickIdOk() (*string, bool)`

GetClickIdOk returns a tuple with the ClickId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickId

`func (o *CapiEvent) SetClickId(v string)`

SetClickId sets ClickId field to given value.

### HasClickId

`func (o *CapiEvent) HasClickId() bool`

HasClickId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


