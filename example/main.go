package main

func main() {
	authToken := "TOKEN_WITHOUT_BEARER_PREFIX"
	pixelId := "PIXEL_ID"

	//runSyncSample(authToken, pixelId)
	//runAsyncSample(authToken, pixelId)
	runTestEventsSample(authToken, pixelId)
}
