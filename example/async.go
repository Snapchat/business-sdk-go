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
