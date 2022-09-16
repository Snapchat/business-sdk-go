package businesssdk

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)

// TODO: Add more documentation
type ConversionApi interface {
	SetDebugging(bool)
	SendEventSync(CapiEvent) (*Response, *http.Response, error)
	SendEventsSync([]CapiEvent) (*Response, *http.Response, error)
	// Todo implement callback structure
	SendEventAsync(CapiEvent)
	SendEventsAsync([]CapiEvent)
}

// CapiClient DefaultApi service
type CapiClient struct {
	apiClient      *APIClient
	cfg            *Configuration
	ctx            context.Context
	isDebugEnabled bool
}

type ContextOption func(*CapiClient)

func NewCapiClient(longLivedToken string, opts ...ContextOption) *CapiClient {
	c := CapiClient{}

	c.cfg = NewConfiguration()
	c.ctx = context.WithValue(context.Background(), ContextAccessToken, longLivedToken)

	// apply context options
	for _, o := range opts {
		o(&c)
	}
	c.cfg.AddDefaultHeader(SDK_VERSION_HEADER, SDK_VERSION_HEADER_VALUE)

	c.apiClient = NewAPIClient(c.cfg)
	return &c
}

// WithLaunchPadUrl sets configuration based on launchPadUrl
func WithLaunchPadUrl(launchPadUrl string) ContextOption { // HL
	return func(c *CapiClient) {
		c.cfg.UserAgent = USER_AGENT_WITH_PAD
		c.cfg.Host = launchPadUrl
	}
}

func (c *CapiClient) SetDebugging(isEnabled bool) {
	c.isDebugEnabled = isEnabled
	c.log(log.InfoLevel, "Debug mode is enabled")
}

func (c *CapiClient) SendTestEvent(event CapiEvent) (*TestResponse, *http.Response, error) {
	return c.SendTestEvents([]CapiEvent{event})
}

func (c *CapiClient) SendTestEvents(events []CapiEvent) (*TestResponse, *http.Response, error) {
	c.setEventIntegration(&events)
	c.logOutgoingEvents(events)

	request := c.apiClient.DefaultApi.SendTestData(c.ctx).Body(events)
	resp, r, err := request.Execute()
	c.logResponse(r, err)
	return resp, r, err
}

func (c *CapiClient) GetTestEventLogs(assetId string) (*ResponseLogs, *http.Response, error) {
	request := c.apiClient.DefaultApi.ConversionValidateLogs(c.ctx).AssetId(assetId)
	resp, r, err := request.Execute()
	c.logResponse(r, err)
	return resp, r, err
}

func (c *CapiClient) GetTestEventStats(assetId string) (*ResponseStats, *http.Response, error) {
	request := c.apiClient.DefaultApi.ConversionValidateStats(c.ctx).AssetId(assetId)
	resp, r, err := request.Execute()
	c.logResponse(r, err)
	return resp, r, err
}

func (c *CapiClient) SendEventSync(event CapiEvent) (*Response, *http.Response, error) {
	return c.SendEventsSync([]CapiEvent{event})
}

func (c *CapiClient) SendEventsSync(events []CapiEvent) (*Response, *http.Response, error) {
	c.setEventIntegration(&events)
	c.logOutgoingEvents(events)

	request := c.apiClient.DefaultApi.SendData(c.ctx).Body(events)
	resp, r, err := request.Execute()
	c.logResponse(r, err)
	return resp, r, err
}

func (c *CapiClient) SendEvent(event CapiEvent, rc chan *AsyncResponse) {
	c.SendEvents([]CapiEvent{event}, rc)
}

func (c *CapiClient) SendEvents(events []CapiEvent, rc chan *AsyncResponse) {
	go c.sendAsyncHelper(events, rc)
}

func (c *CapiClient) sendAsyncHelper(events []CapiEvent, rc chan *AsyncResponse) {
	c.setEventIntegration(&events)
	c.logOutgoingEvents(events)

	request := c.apiClient.DefaultApi.SendData(c.ctx).Body(events)
	resp, r, err := request.Execute()
	c.logResponse(r, err)
	rc <- &AsyncResponse{Response: resp, HttpResponse: r, Err: err}
}

func (c *CapiClient) setEventIntegration(events *[]CapiEvent) {
	for i := range *events {
		(*events)[i].SetIntegration(INTEGRATION_SDK)
	}
}

func (c *CapiClient) logOutgoingEvents(events []CapiEvent) {
	js, err := json.Marshal(events)
	if err != nil {
		c.log(log.ErrorLevel, "Failed to JSON Marshal and log events")
	} else {
		c.log(log.InfoLevel, "Sending events: "+string(js))
	}
}

func (c *CapiClient) logResponse(r *http.Response, err error) {
	fields := log.Fields{}

	if err != nil {
		fields["error"] = err
		if r != nil {
			if dump, errDump := httputil.DumpResponse(r, true); errDump == nil {
				fields["response"] = string(dump)
			}
		}
		c.logWithFields(log.ErrorLevel, "Failed to send event with error", fields)
	} else {
		if dump, errDump := httputil.DumpResponse(r, true); errDump == nil {
			fields["response"] = string(dump)
		}
		c.logWithFields(log.InfoLevel, "Request Succeeded", fields)
	}
}

func (c *CapiClient) log(level log.Level, message string) {
	if !c.isDebugEnabled {
		return
	}

	c.logWithFields(level, message, log.Fields{})
}

func (c *CapiClient) logWithFields(level log.Level, message string, fields log.Fields) {
	if !c.isDebugEnabled {
		return
	}

	msg := "[Snap Business SDK] " + message
	switch level {
	case log.InfoLevel:
		log.WithFields(fields).Info(msg)
	case log.WarnLevel:
		log.WithFields(fields).Warn(msg)
	case log.ErrorLevel:
		log.WithFields(fields).Error(msg)
	default:
		log.WithFields(fields).Debug(msg)
	}
}
