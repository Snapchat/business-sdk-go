package businesssdk

const (
	SDK_LANGUAGE             = "Go"
	SDK_VERSION              = "1.0.0"
	API_VERSION              = "v2"
	PROD_URL                 = "https://tr.snapchat.com/" + API_VERSION
	STAGING_URL              = "https://tr-shadow.snapchat.com/" + API_VERSION
	USER_AGENT               = "BusinessSDK/" + SDK_VERSION + "/Go"
	USER_AGENT_WITH_PAD      = USER_AGENT + " (LaunchPAD)"
	SDK_VERSION_HEADER       = "X-CAPI-BusinessSDK"
	INTEGRATION_SDK          = "business-sdk"
	SDK_VERSION_HEADER_VALUE = SDK_LANGUAGE + "/" + SDK_VERSION

	// LaunchPAD Support
	TEST_MODE_HEADER = "X-CAPI-Test-Mode"
	CAPI_PATH_HEADER = "X-CAPI-Path"

	// Specify the CAPI endpoint path where the request is forwarded
	CAPI_PATH      = "/" + API_VERSION + "/conversion"
	CAPI_PATH_TEST = CAPI_PATH + "/validate"
)
