/*
Snap Conversions API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package businesssdk

import (
	"encoding/json"
)

// ResponseLogs struct for ResponseLogs
type ResponseLogs struct {
	Status *string `json:"status,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Logs []ResponseLogsLog `json:"logs,omitempty"`
}

// NewResponseLogs instantiates a new ResponseLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLogs() *ResponseLogs {
	this := ResponseLogs{}
	return &this
}

// NewResponseLogsWithDefaults instantiates a new ResponseLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLogsWithDefaults() *ResponseLogs {
	this := ResponseLogs{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseLogs) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLogs) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseLogs) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResponseLogs) SetStatus(v string) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ResponseLogs) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLogs) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ResponseLogs) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ResponseLogs) SetReason(v string) {
	o.Reason = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ResponseLogs) GetLogs() []ResponseLogsLog {
	if o == nil || o.Logs == nil {
		var ret []ResponseLogsLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLogs) GetLogsOk() ([]ResponseLogsLog, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ResponseLogs) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []ResponseLogsLog and assigns it to the Logs field.
func (o *ResponseLogs) SetLogs(v []ResponseLogsLog) {
	o.Logs = v
}

func (o ResponseLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

type NullableResponseLogs struct {
	value *ResponseLogs
	isSet bool
}

func (v NullableResponseLogs) Get() *ResponseLogs {
	return v.value
}

func (v *NullableResponseLogs) Set(val *ResponseLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLogs(val *ResponseLogs) *NullableResponseLogs {
	return &NullableResponseLogs{value: val, isSet: true}
}

func (v NullableResponseLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

