package main

import (
	sdk "github.com/Snapchat/business-sdk-go"
	"strconv"
	"time"
)

func runTestEventsSample(authToken, pixelId string) {
	capiClient := sdk.NewCapiClient(authToken)
	//capiClient := sdk.NewCapiClient(authToken, sdk.WithLaunchPadUrl("url_here"))
	capiClient.SetDebugging(true)

	event := sdk.NewCapiEvent()
	event.SetPixelId(pixelId)
	event.SetEventType("PURCHASE")
	event.SetEventConversionType("WEB")
	event.SetEventTag("event_tag_example")
	event.SetTimestamp(strconv.FormatInt(time.Now().UnixMilli(), 10))
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

	capiClient.SendTestEvent(*event)
	// Example to get and handle response objects
	// responseBody, httpResponse, err := capiClient.SendTestEvent(*event)

	capiClient.GetTestEventStats(pixelId)
	capiClient.GetTestEventLogs(pixelId)
}
