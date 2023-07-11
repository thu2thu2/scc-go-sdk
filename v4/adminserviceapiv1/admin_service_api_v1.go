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

// Package adminserviceapiv1 : Operations and models for the AdminServiceApiV1 service
package adminserviceapiv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/cloud-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// AdminServiceApiV1 : This is an API for the Admin Service
//
// API Version: 1.0.0
type AdminServiceApiV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "admin_service_api"

const ParameterizedServiceURL = "https://{environment}.cloud.ibm.com/instances/{instance_id}/v3"

var defaultUrlVariables = map[string]string{
	"environment": "us-south.compliance",
	"instance_id": "instance_id",
}

// AdminServiceApiV1Options : Service options
type AdminServiceApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAdminServiceApiV1UsingExternalConfig : constructs an instance of AdminServiceApiV1 with passed in options and external configuration.
func NewAdminServiceApiV1UsingExternalConfig(options *AdminServiceApiV1Options) (adminServiceApi *AdminServiceApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	adminServiceApi, err = NewAdminServiceApiV1(options)
	if err != nil {
		return
	}

	err = adminServiceApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = adminServiceApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAdminServiceApiV1 : constructs an instance of AdminServiceApiV1 with passed in options.
func NewAdminServiceApiV1(options *AdminServiceApiV1Options) (service *AdminServiceApiV1, err error) {
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

	service = &AdminServiceApiV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "adminServiceApi" suitable for processing requests.
func (adminServiceApi *AdminServiceApiV1) Clone() *AdminServiceApiV1 {
	if core.IsNil(adminServiceApi) {
		return nil
	}
	clone := *adminServiceApi
	clone.Service = adminServiceApi.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (adminServiceApi *AdminServiceApiV1) SetServiceURL(url string) error {
	return adminServiceApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (adminServiceApi *AdminServiceApiV1) GetServiceURL() string {
	return adminServiceApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (adminServiceApi *AdminServiceApiV1) SetDefaultHeaders(headers http.Header) {
	adminServiceApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (adminServiceApi *AdminServiceApiV1) SetEnableGzipCompression(enableGzip bool) {
	adminServiceApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (adminServiceApi *AdminServiceApiV1) GetEnableGzipCompression() bool {
	return adminServiceApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (adminServiceApi *AdminServiceApiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	adminServiceApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (adminServiceApi *AdminServiceApiV1) DisableRetries() {
	adminServiceApi.Service.DisableRetries()
}

// GetSettings : Retrieves settings
// Retrieves the settings of an instance.
func (adminServiceApi *AdminServiceApiV1) GetSettings(getSettingsOptions *GetSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	return adminServiceApi.GetSettingsWithContext(context.Background(), getSettingsOptions)
}

// GetSettingsWithContext is an alternate form of the GetSettings method which supports a Context parameter
func (adminServiceApi *AdminServiceApiV1) GetSettingsWithContext(ctx context.Context, getSettingsOptions *GetSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSettingsOptions, "getSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminServiceApi.Service.Options.URL, `/settings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service_api", "V1", "GetSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminServiceApi.Service.Request(request, &rawResponse)
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

// UpdateSettings : Patch settings
// Update settings.
func (adminServiceApi *AdminServiceApiV1) UpdateSettings(updateSettingsOptions *UpdateSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	return adminServiceApi.UpdateSettingsWithContext(context.Background(), updateSettingsOptions)
}

// UpdateSettingsWithContext is an alternate form of the UpdateSettings method which supports a Context parameter
func (adminServiceApi *AdminServiceApiV1) UpdateSettingsWithContext(ctx context.Context, updateSettingsOptions *UpdateSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = adminServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminServiceApi.Service.Options.URL, `/settings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service_api", "V1", "UpdateSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateSettingsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminServiceApi.Service.Request(request, &rawResponse)
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

// PostTestEvent : Send test event
// Send a test event using your configured Event Notifications instance.
func (adminServiceApi *AdminServiceApiV1) PostTestEvent(postTestEventOptions *PostTestEventOptions) (result *TestEvent, response *core.DetailedResponse, err error) {
	return adminServiceApi.PostTestEventWithContext(context.Background(), postTestEventOptions)
}

// PostTestEventWithContext is an alternate form of the PostTestEvent method which supports a Context parameter
func (adminServiceApi *AdminServiceApiV1) PostTestEventWithContext(ctx context.Context, postTestEventOptions *PostTestEventOptions) (result *TestEvent, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(postTestEventOptions, "postTestEventOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminServiceApi.Service.Options.URL, `/test_event`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range postTestEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service_api", "V1", "PostTestEvent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminServiceApi.Service.Request(request, &rawResponse)
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

// EventNotifications : Event Notifications settings.
type EventNotifications struct {
	// EN Instance CRN.
	InstanceCrn *string `json:"instance_crn,omitempty"`

	// Modified date in ISO8601 format.
	Modified *string `json:"modified,omitempty"`

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
	err = core.UnmarshalPrimitive(m, "modified", &obj.Modified)
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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSettingsOptions : Instantiate GetSettingsOptions
func (*AdminServiceApiV1) NewGetSettingsOptions() *GetSettingsOptions {
	return &GetSettingsOptions{}
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
	JSONPatchOperation_Op_Add = "add"
	JSONPatchOperation_Op_Copy = "copy"
	JSONPatchOperation_Op_Move = "move"
	JSONPatchOperation_Op_Remove = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*AdminServiceApiV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op: core.StringPtr(op),
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

// ObjectStorage : Object Storage settings.
type ObjectStorage struct {
	// EN Instance CRN.
	InstanceCrn *string `json:"instance_crn,omitempty"`

	// Bucket.
	Bucket *string `json:"bucket,omitempty"`

	// Bucket location.
	BucketLocation *string `json:"bucket_location,omitempty"`

	// Bucket endpoint.
	BucketEndpoint *string `json:"bucket_endpoint,omitempty"`

	// Modified date in ISO8601 format.
	Modified *string `json:"modified,omitempty"`
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
	err = core.UnmarshalPrimitive(m, "modified", &obj.Modified)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostTestEventOptions : The PostTestEvent options.
type PostTestEventOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostTestEventOptions : Instantiate PostTestEventOptions
func (*AdminServiceApiV1) NewPostTestEventOptions() *PostTestEventOptions {
	return &PostTestEventOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *PostTestEventOptions) SetHeaders(param map[string]string) *PostTestEventOptions {
	options.Headers = param
	return options
}

// Settings : Settings.
type Settings struct {
	// Event Notifications settings.
	EventNotifications *EventNotifications `json:"event_notifications,omitempty"`

	// Object Storage settings.
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

func (*AdminServiceApiV1) NewSettingsPatch(settings *Settings) (_patch []JSONPatchOperation) {
	if (settings.EventNotifications != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/event_notifications"),
			Value: settings.EventNotifications,
		})
	}
	if (settings.ObjectStorage != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/object_storage"),
			Value: settings.ObjectStorage,
		})
	}
	return
}

// TestEvent : The details of a test event response.
type TestEvent struct {
	// Indicates whether the event was received by Event Notifications.
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
	// New settings.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSettingsOptions : Instantiate UpdateSettingsOptions
func (*AdminServiceApiV1) NewUpdateSettingsOptions(body []JSONPatchOperation) *UpdateSettingsOptions {
	return &UpdateSettingsOptions{
		Body: body,
	}
}

// SetBody : Allow user to set Body
func (_options *UpdateSettingsOptions) SetBody(body []JSONPatchOperation) *UpdateSettingsOptions {
	_options.Body = body
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSettingsOptions) SetHeaders(param map[string]string) *UpdateSettingsOptions {
	options.Headers = param
	return options
}
