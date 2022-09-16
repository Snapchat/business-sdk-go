package businesssdk

import "net/http"

type AsyncResponse struct {
	Response     *Response
	HttpResponse *http.Response
	Err          error
}
