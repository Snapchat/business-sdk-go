# CAPI Business SDK in Go

## Introduction
The Snapchat Conversions API (CAPI) allows you to pass web, app, and offline events to Snap via a Server-to-Server (S2S) integration. Similar to our other Direct Data Integrations, Snap Pixel and App Ads Kit (SDK), by passing these events, you can access post-view and post-swipe campaign reporting to measure performance and incrementality. Depending on the data shared and timeliness of integration, it’s also possible to leverage events passed via Conversions API for solutions such as custom audience targeting, campaign optimization, Dynamic Ads, and more.

Business SDK is an API client that facilitates requests and authentication processes with CAPI as if they were local functions of the supported languages. There have been similar products (e.g. Facebook Business SDK for Conversion API), which largely simplified the integration for advertising platforms. To improve the developer experience and CAPI adoption, we bundle the dedicated CAPI client, hashing libraries, and code examples into one SDK that is available in multiple languages. In addition, our CAPI Business SDK paves the way for Privacy-Enhancing Technologies in CAPI v2, with seamless integration of the Launch Pad.
## Installation
The Go CAPI Business SDK module will be published as part of the public github repository.

Add the Business SDK dependency to your go project with the go get command

`go get github.com/Snapchat/business-sdk-go`


Import the package in your go file with the following declaration
Go:

`Import "github.com/Snapchat/business-sdk-go"`

## Code Examples
### Sending Production Events
Please use the code example below as a template on sending your conversion events to Snap Conversion API. More conversion parameters are expected to be provided in practice.

Example 1: Send a single CAPI event
```
package main

import (
sdk "github.com/Snapchat/business-sdk-go"
)

func main() {
    auth_token := "TOKEN_WITHOUT_BEARER_PREFIX"
    pixel_id := "TEST_PIXEL_ID"
    
    capiClient := sdk.NewCapiClient(longLivedToken)
    //capiClient := sdk.NewCapiClient(longLivedToken, sdk.WithLaunchPadUrl("LAUNCHPAD_URL_HERE"))
    capiClient.SetDebugging(true)
    
    event := sdk.NewCapiEvent()
    event.SetPixelId(pixel_id)
    event.SetEventType("PURCHASE")
    event.SetEventConversionType("WEB")
    event.SetEventTag("event_tag_example")
    event.SetTimestamp(1656022510346)
    event.SetUuidC1("34dd6077-e3a0-4b1c-9f91-a690ea0e335d")
    // we support pass in plaintext email (it will be hashed and set to hashed_email automatically)
    event.SetEmail("test@example.com")
    // you can also pass hashed email if preferred
    // event.SetHashedEmail("f660ab912ec121d1b1e928a0bb4bc61b15f5ad44d5efdc4e1c92a25e99b8e44a")
    // we support pass in plaintext phone number (it will be hashed and set to hashed_phone_number automatically)
    event.SetPhoneNumber("1234567890")
    // you can also pass hashed phone number if preferred
    //event.SetHashedPhoneNumber("a2b5e507dfb65941ff4be6e4fc089313cbbb640da5fd6fbc4e3d2e2f3abe92cc")
    // we support pass in plaintext ip address (it will be hashed and set to hashed_ip_address automatically)
    event.SetIpAddress("12.34.56.78")
    // you can also pass hashed ip address if preferred
    // event.SetHashedIpAddress("f1412386aa8db2579aff2636cb9511cacc5fd9880ecab60c048508fbe26ee4d9")
    event.SetItemCategory("item_category_example")
    event.SetItemIds("item_ids_example")
    event.SetDescription("description_example")
    event.SetNumberItems("number_items_example")
    event.SetPrice("price_example")
    event.SetCurrency("USD")
    event.SetTransactionId("transaction_id_example")
    event.SetLevel("level_example")
    event.SetClientDedupId("client_dedup_id_example")
    event.SetSearchString("search_string_example")
    event.SetPageUrl("page_url_example")
    event.SetSignUpMethod("sign_up_method_example")
    event.SetFirstName("test_first")
    //event.SetHashedFirstNameSha("d99156483b6a99eb5f5a1707f7330e1c979a648b47a379d56a0d6850a9a9c76c")
    //event.SetHashedFirstNameSdx("T231")
    event.SetMiddleName("")
    //event.SetHashedMiddleNameSha("")
    //event.SetHashedMiddleNameSdx("")
    event.SetLastName("test_last")
    //event.SetHashedLastNameSha("19fc3d9f9f6fad30ccbbebd51f67515dc95d8a5ef363fd35c34a2f47064d43bd")
    //event.SetHashedLastNameSdx("T234")
    event.SetCity("Los Angeles")
    //event.SetHashedCitySha("9f2608067816e38c85edfb0c3985feff32def8b5dc17bb522ffc2e877e9b386b")
    //event.SetHashedCitySdx("9f2608067816e38c85edfb0c3985feff32def8b5dc17bb522ffc2e877e9b386b")
    event.SetState("CA")
    //event.SetHashedStateSha("6959097001d10501ac7d54c0bdb8db61420f658f2922cc26e46d536119a31126")
    event.SetHashedStateSdx("C000") //(hashed_state_sdx="C000")
    event.SetZip("00000")
    //event.SetHashedZip("e7042ac7d09c7bc41c8cfa5749e41858f6980643bc0db1a83cc793d3e24d3f77")
    //event.SetClickId("click_id_example")
    
    asyncRespChan := make(chan *sdk.AsyncResponse, 1)
    // send single event asynchronously
    capiClient.SendEvent(*event, asyncRespChan)
    
	defer func() {
		for i := 0; i < 2; i++ {
			result := <-asyncRespChan

			if result.Err != nil {
				log.Error("Error when calling `ConversionApi.SendEvent``: %v\n", result.Err)
				log.Error("Full HTTP response: %v\n", result.HttpResponse)
			}
			fmt.Fprintf(os.Stdout, "Response from `ConversionApi.SendEventAsync`: %v\n", result.Response)
		}
	}()
}
```


Example 2: Sending a batch of CAPI events (up to 2000)
```
package main

import (
   "fmt"
   "os"
   "strconv"
   "time"

   sdk "github.com/Snapchat/business-sdk-go"
   log "github.com/sirupsen/logrus"
)

func runAsyncSample(authToken, pixelId string) {
   capiClient := sdk.NewCapiClient(authToken)
   //capiClient := sdk.NewCapiClient(auth_token, sdk.WithLaunchPadUrl("url_here"))
   capiClient.SetDebugging(true)

   event1 := sdk.NewCapiEvent()
   event1.SetPixelId(pixelId)
   event1.SetEventType("PURCHASE")
   event1.SetEventConversionType("WEB")
   event1.SetEventTag("event_tag_example")
   event1.SetTimestamp(strconv.FormatInt(time.Now().UnixMilli(), 10))
   event1.SetUuidC1("34dd6077-e3a0-4b1c-9f91-a690ea0e335d")
   event1.SetEmail("test@example.com")
   event1.SetPhoneNumber("1234567890")
   event1.SetIpAddress("12.34.56.78")
   event1.SetItemCategory("item_category_example")
   event1.SetItemIds("item_ids_example")
   event1.SetDescription("description_example")
   event1.SetNumberItems("number_items_example")
   event1.SetPrice("price_example")
   event1.SetCurrency("USD")
   event1.SetTransactionId("transaction_id_example")
   event1.SetLevel("level_example")
   event1.SetClientDedupId("client_dedup_id_example")
   event1.SetSearchString("search_string_example")
   event1.SetPageUrl("page_url_example")
   event1.SetSignUpMethod("sign_up_method_example")
   event1.SetFirstName("test_first")

   event2 := sdk.NewCapiEvent()
   event2.SetPixelId(pixelId)
   event2.SetEventType("PAGE_VIEW")
   event2.SetEventConversionType("WEB")
   event2.SetEventTag("event_tag_example")
   event2.SetTimestamp(strconv.FormatInt(time.Now().UnixMilli(), 10))
   event2.SetUuidC1("34dd6077-e3a0-4b1c-9f91-a690ea0e335d")
   event2.SetEmail("test@example.com")
   event2.SetPageUrl("page_url_example")

   asyncRespChan := make(chan *sdk.AsyncResponse, 2)
   capiClient.SendEvents([]sdk.CapiEvent{*event1, *event2}, asyncRespChan)

   defer func() {
       for i := 0; i < 2; i++ {
           result := <-asyncRespChan

           if result.Err != nil {
               log.Error("Error when calling `ConversionApi.SendEvent``: %v\n", result.Err)
               log.Error("Full HTTP response: %v\n", result.HttpResponse)
           }
           fmt.Fprintf(os.Stdout, "Response from `ConversionApi.SendEventAsync`: %v\n", result.Response)
       }
   }()
}
```

### Sending Test Events
Snap’s Conversion API provides the validate, log, and stats endpoints for advertisers to test their events and obtain more information on how each of the test event was processed.

Creating and setting up a test event is the same as setting up to send a production event. Clients must simply call the SendTestEvent function instead of the production functions.

The stats and logs should be called after sending and receiving a successful response from the test event endpoint.

```
capiClient.SendTestEvent(*event)
// Example to get and handle response objects
// responseBody, httpResponse, err := capiClient.SendTestEvent(*event)

responseStats, httpRespStats, err := capiClient.GetTestEventStats(pixelId)
responseLogs, httpRespLogs, err := capiClient.GetTestEventLogs(pixelId)
```

## Notes:
1) Initiate ConversionApi
```
sdk.NewCapiClient(longLivedToken) 
// with launchpad URL
// sdk.NewCapiClient(longLivedToken, sdk.WithLaunchPadUrl("url_here")) 
```
   * if the Launch Pad has been set up under your domain. Conversion events will be forwarded to Snap transparently. (Other MPC features will be introduced in later versions).
   * Otherwise, you can initiate the instance using `sdk.NewCapiClient(longLivedToken)` Conversion events are sent back to Snap from the business SDK directly.
   * It’s recommended to create a dedicated instance per thread to avoid any potential issues.
2) API Token
   * To use the Conversions API, you need to use the access token for auth. See [here](https://marketingapi.snapchat.com/docs/conversion.html#auth-requirements) to generate the token.
3) Build CapiEvent
   * Please check with the section [Conversion Parameters](https://marketingapi.snapchat.com/docs/conversion.html#additional-data-formatting-guidelines) and provide as much information as possible when creating the CapiEvent object.
   * Every CAPI attribute has a corresponding setter in the CapiEvent class following the camelcase naming convention.  
   * At least one of the following parameters is required in order to successfully send events via the Conversions API. When possible, we recommend passing all of the below parameters to improve performance:
     * hashed_email
     * hashed_phone
     * hashed_ip and user_agent
     * hashed_mobile_ad_id
   * Any setter starting with the prefix of “hashed” (e.g. `SetHashedEmail()`) accepts the hashing PII value only (see [Data Hygiene](https://marketingapi.snapchat.com/docs/conversion.html#data-hygiene)). Please use the unhashed setter (e.g. `SetEmail()`) if you want the business sdk to normalize and hash the PII field on your behalf.
   * We highly recommend passing cookie1 `uuid_c1`, if available, as this will increase your match rate. You can access a 1st party cookie by looking at the _scid value under your domain if you are using the Pixel SDK.
4) Send event(s) asynchronously
   * Conversion events can be sent individually via `SendEvent(event CapiEvent, rc chan *AsyncResponse)`.
   * Conversion events can be reported in batch using `SendEvents(events []CapiEvent, rc chan *AsyncResponse)` if they are buffered in your application Please check example/async.go for more details. We recommend a 1000 QPS limit for sending us requests. You may send up to 2000 events per batch request, and can thus send up to 2M events/sec. Sending more than 2000 events per batch will result in a 400 error.
   * Events are encapsulated in an asynchronous request in both solutions by which your application won’t be blocked. The response is logged by the asynchronous go routine (under debugging mode)
   * The asynchronous go routine will send back the response to your handler through a response channel that you can provide to the SendEvents function.
5) Test Events, Logs, and Stats
   * Conversion events can be sent for testing and validation via the `SendTestEvent(event CapiEvent)`.
   * Conversion API also provides logging endpoint. It provides a summary of test CAPI events sent to the test endpoint within the past day
   * Converion API’s stats endpoint provides basic stats and summary of the test events sent.
6) Debugging Mode
   * When debugging mode is turned on, we will log the events, api call response and exceptions using the logrus logger.
     * By default, the logger will log to the console
