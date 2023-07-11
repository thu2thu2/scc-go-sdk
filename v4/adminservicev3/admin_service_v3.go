/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.70.0-7df966bf-20230419-195904
 */

// Package adminservicev3 : Operations and models for the AdminServiceV3 service
package adminservicev3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/scc-go-sdk/v4/common"
)

// AdminServiceV3 : The Security and Compliance Center API reference.
//
// API Version: 3.0.0
type AdminServiceV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "admin_service"

const ParameterizedServiceURL = "https://{region}.cloud.ibm.com/instances/{instance_id}/v3"

var defaultUrlVariables = map[string]string{
	"region":      "us-south.compliance",
	"instance_id": "instance_id",
}

// AdminServiceV3Options : Service options
type AdminServiceV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAdminServiceV3UsingExternalConfig : constructs an instance of AdminServiceV3 with passed in options and external configuration.
func NewAdminServiceV3UsingExternalConfig(options *AdminServiceV3Options) (adminService *AdminServiceV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	adminService, err = NewAdminServiceV3(options)
	if err != nil {
		return
	}

	err = adminService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = adminService.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAdminServiceV3 : constructs an instance of AdminServiceV3 with passed in options.
func NewAdminServiceV3(options *AdminServiceV3Options) (service *AdminServiceV3, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &AdminServiceV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "adminService" suitable for processing requests.
func (adminService *AdminServiceV3) Clone() *AdminServiceV3 {
	if core.IsNil(adminService) {
		return nil
	}
	clone := *adminService
	clone.Service = adminService.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (adminService *AdminServiceV3) SetServiceURL(url string) error {
	return adminService.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (adminService *AdminServiceV3) GetServiceURL() string {
	return adminService.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (adminService *AdminServiceV3) SetDefaultHeaders(headers http.Header) {
	adminService.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (adminService *AdminServiceV3) SetEnableGzipCompression(enableGzip bool) {
	adminService.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (adminService *AdminServiceV3) GetEnableGzipCompression() bool {
	return adminService.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (adminService *AdminServiceV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	adminService.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (adminService *AdminServiceV3) DisableRetries() {
	adminService.Service.DisableRetries()
}

// GetSettings : Get settings
// Retrieve the settings of your service instance.
func (adminService *AdminServiceV3) GetSettings(getSettingsOptions *GetSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	return adminService.GetSettingsWithContext(context.Background(), getSettingsOptions)
}

// GetSettingsWithContext is an alternate form of the GetSettings method which supports a Context parameter
func (adminService *AdminServiceV3) GetSettingsWithContext(ctx context.Context, getSettingsOptions *GetSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSettingsOptions, "getSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/settings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "GetSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSettingsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getSettingsOptions.XCorrelationID))
	}
	if getSettingsOptions.XRequestID != nil {
		builder.AddHeader("X-Request-Id", fmt.Sprint(*getSettingsOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateSettings : Update settings
// Update the settings of your service instance.
func (adminService *AdminServiceV3) UpdateSettings(updateSettingsOptions *UpdateSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	return adminService.UpdateSettingsWithContext(context.Background(), updateSettingsOptions)
}

// UpdateSettingsWithContext is an alternate form of the UpdateSettings method which supports a Context parameter
func (adminService *AdminServiceV3) UpdateSettingsWithContext(ctx context.Context, updateSettingsOptions *UpdateSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSettingsOptions, "updateSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSettingsOptions, "updateSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/settings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "UpdateSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateSettingsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*updateSettingsOptions.XCorrelationID))
	}
	if updateSettingsOptions.XRequestID != nil {
		builder.AddHeader("X-Request-Id", fmt.Sprint(*updateSettingsOptions.XRequestID))
	}

	_, err = builder.SetBodyContentJSON(updateSettingsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostTestEvent : Create a test event
// Send a test event by using your configured Event Notifications instance.
func (adminService *AdminServiceV3) PostTestEvent(postTestEventOptions *PostTestEventOptions) (result *TestEvent, response *core.DetailedResponse, err error) {
	return adminService.PostTestEventWithContext(context.Background(), postTestEventOptions)
}

// PostTestEventWithContext is an alternate form of the PostTestEvent method which supports a Context parameter
func (adminService *AdminServiceV3) PostTestEventWithContext(ctx context.Context, postTestEventOptions *PostTestEventOptions) (result *TestEvent, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(postTestEventOptions, "postTestEventOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/test_event`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range postTestEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "PostTestEvent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if postTestEventOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*postTestEventOptions.XCorrelationID))
	}
	if postTestEventOptions.XRequestID != nil {
		builder.AddHeader("X-Request-Id", fmt.Sprint(*postTestEventOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTestEvent)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// EventNotifications : The Event Notifications settings.
type EventNotifications struct {
	// The Event Notifications Instance CRN.
	InstanceCrn *string `json:"instance_crn,omitempty"`

	// The updated_on date in ISO8601 format.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// Source.
	SourceID *string `json:"source_id,omitempty"`
}

// UnmarshalEventNotifications unmarshals an instance of EventNotifications from the specified map of raw messages.
func UnmarshalEventNotifications(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EventNotifications)
	err = core.UnmarshalPrimitive(m, "instance_crn", &obj.InstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_id", &obj.SourceID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSettingsOptions : The GetSettings options.
type GetSettingsOptions struct {
	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header  for the
	// corresponding response.  The same value is not used for downstream requests and retries of those requests.  If a
	// value of this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSettingsOptions : Instantiate GetSettingsOptions
func (*AdminServiceV3) NewGetSettingsOptions() *GetSettingsOptions {
	return &GetSettingsOptions{}
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetSettingsOptions) SetXCorrelationID(xCorrelationID string) *GetSettingsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *GetSettingsOptions) SetXRequestID(xRequestID string) *GetSettingsOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSettingsOptions) SetHeaders(param map[string]string) *GetSettingsOptions {
	options.Headers = param
	return options
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add     = "add"
	JSONPatchOperation_Op_Copy    = "copy"
	JSONPatchOperation_Op_Move    = "move"
	JSONPatchOperation_Op_Remove  = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test    = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*AdminServiceV3) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op:   core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectStorage : The Cloud Object Storage settings.
type ObjectStorage struct {
	// The EN Instance CRN.
	InstanceCrn *string `json:"instance_crn,omitempty"`

	// The bucket.
	Bucket *string `json:"bucket,omitempty"`

	// The bucket location.
	BucketLocation *string `json:"bucket_location,omitempty"`

	// The bucket endpoint.
	BucketEndpoint *string `json:"bucket_endpoint,omitempty"`

	// The updated_on date in ISO8601 format.
	UpdatedOn *string `json:"updated_on,omitempty"`
}

// UnmarshalObjectStorage unmarshals an instance of ObjectStorage from the specified map of raw messages.
func UnmarshalObjectStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectStorage)
	err = core.UnmarshalPrimitive(m, "instance_crn", &obj.InstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_location", &obj.BucketLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_endpoint", &obj.BucketEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostTestEventOptions : The PostTestEvent options.
type PostTestEventOptions struct {
	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header  for the
	// corresponding response.  The same value is not used for downstream requests and retries of those requests.  If a
	// value of this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostTestEventOptions : Instantiate PostTestEventOptions
func (*AdminServiceV3) NewPostTestEventOptions() *PostTestEventOptions {
	return &PostTestEventOptions{}
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *PostTestEventOptions) SetXCorrelationID(xCorrelationID string) *PostTestEventOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *PostTestEventOptions) SetXRequestID(xRequestID string) *PostTestEventOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostTestEventOptions) SetHeaders(param map[string]string) *PostTestEventOptions {
	options.Headers = param
	return options
}

// Settings : The settings.
type Settings struct {
	// The Event Notifications settings.
	EventNotifications *EventNotifications `json:"event_notifications,omitempty"`

	// The Cloud Object Storage settings.
	ObjectStorage *ObjectStorage `json:"object_storage,omitempty"`
}

// UnmarshalSettings unmarshals an instance of Settings from the specified map of raw messages.
func UnmarshalSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Settings)
	err = core.UnmarshalModel(m, "event_notifications", &obj.EventNotifications, UnmarshalEventNotifications)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "object_storage", &obj.ObjectStorage, UnmarshalObjectStorage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*AdminServiceV3) NewSettingsPatch(settings *Settings) (_patch []JSONPatchOperation) {
	if settings.EventNotifications != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/event_notifications"),
			Value: settings.EventNotifications,
		})
	}
	if settings.ObjectStorage != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/object_storage"),
			Value: settings.ObjectStorage,
		})
	}
	return
}

// TestEvent : The details of a test event response.
type TestEvent struct {
	// The indication of whether the event was received by Event Notifications.
	Success *bool `json:"success" validate:"required"`
}

// UnmarshalTestEvent unmarshals an instance of TestEvent from the specified map of raw messages.
func UnmarshalTestEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestEvent)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateSettingsOptions : The UpdateSettings options.
type UpdateSettingsOptions struct {
	// The request body to update your settings.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header  for the
	// corresponding response.  The same value is not used for downstream requests and retries of those requests.  If a
	// value of this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSettingsOptions : Instantiate UpdateSettingsOptions
func (*AdminServiceV3) NewUpdateSettingsOptions(body []JSONPatchOperation) *UpdateSettingsOptions {
	return &UpdateSettingsOptions{
		Body: body,
	}
}

// SetBody : Allow user to set Body
func (_options *UpdateSettingsOptions) SetBody(body []JSONPatchOperation) *UpdateSettingsOptions {
	_options.Body = body
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateSettingsOptions) SetXCorrelationID(xCorrelationID string) *UpdateSettingsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *UpdateSettingsOptions) SetXRequestID(xRequestID string) *UpdateSettingsOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSettingsOptions) SetHeaders(param map[string]string) *UpdateSettingsOptions {
	options.Headers = param
	return options
}
