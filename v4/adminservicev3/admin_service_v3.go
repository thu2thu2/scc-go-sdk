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

// AdminServiceV3 : APIs for the SCC Services
//
// API Version: 3.0.0
type AdminServiceV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "http://localhost:8080"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "admin_service"

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

// GetInstanceAccess : Get access level for an instance
// Retrieve access information about the caller for the specified instance.
func (adminService *AdminServiceV3) GetInstanceAccess(getInstanceAccessOptions *GetInstanceAccessOptions) (result *InstanceAccess, response *core.DetailedResponse, err error) {
	return adminService.GetInstanceAccessWithContext(context.Background(), getInstanceAccessOptions)
}

// GetInstanceAccessWithContext is an alternate form of the GetInstanceAccess method which supports a Context parameter
func (adminService *AdminServiceV3) GetInstanceAccessWithContext(ctx context.Context, getInstanceAccessOptions *GetInstanceAccessOptions) (result *InstanceAccess, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceAccessOptions, "getInstanceAccessOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceAccessOptions, "getInstanceAccessOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceAccessOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/access`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceAccessOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "GetInstanceAccess")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceAccess)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetInstanceSettings : Retrieves instance settings
// Retrieves the instance of an account.
func (adminService *AdminServiceV3) GetInstanceSettings(getInstanceSettingsOptions *GetInstanceSettingsOptions) (result *InstanceSettings, response *core.DetailedResponse, err error) {
	return adminService.GetInstanceSettingsWithContext(context.Background(), getInstanceSettingsOptions)
}

// GetInstanceSettingsWithContext is an alternate form of the GetInstanceSettings method which supports a Context parameter
func (adminService *AdminServiceV3) GetInstanceSettingsWithContext(ctx context.Context, getInstanceSettingsOptions *GetInstanceSettingsOptions) (result *InstanceSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstanceSettingsOptions, "getInstanceSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstanceSettingsOptions, "getInstanceSettingsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstanceSettingsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/settings`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstanceSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "GetInstanceSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateInstanceSettings : Patch instance settings
// Update instance settings.
func (adminService *AdminServiceV3) UpdateInstanceSettings(updateInstanceSettingsOptions *UpdateInstanceSettingsOptions) (result *InstanceSettings, response *core.DetailedResponse, err error) {
	return adminService.UpdateInstanceSettingsWithContext(context.Background(), updateInstanceSettingsOptions)
}

// UpdateInstanceSettingsWithContext is an alternate form of the UpdateInstanceSettings method which supports a Context parameter
func (adminService *AdminServiceV3) UpdateInstanceSettingsWithContext(ctx context.Context, updateInstanceSettingsOptions *UpdateInstanceSettingsOptions) (result *InstanceSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateInstanceSettingsOptions, "updateInstanceSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateInstanceSettingsOptions, "updateInstanceSettingsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateInstanceSettingsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/settings`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateInstanceSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "UpdateInstanceSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateInstanceSettingsOptions.JSONPatchOperation)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostInstanceTestEvent : Instnace send test event
// Send a test event using your configured Event Notifications instance.Send a test event using your configured Event
// Notifications instance.
func (adminService *AdminServiceV3) PostInstanceTestEvent(postInstanceTestEventOptions *PostInstanceTestEventOptions) (result *TestEvent, response *core.DetailedResponse, err error) {
	return adminService.PostInstanceTestEventWithContext(context.Background(), postInstanceTestEventOptions)
}

// PostInstanceTestEventWithContext is an alternate form of the PostInstanceTestEvent method which supports a Context parameter
func (adminService *AdminServiceV3) PostInstanceTestEventWithContext(ctx context.Context, postInstanceTestEventOptions *PostInstanceTestEventOptions) (result *TestEvent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postInstanceTestEventOptions, "postInstanceTestEventOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postInstanceTestEventOptions, "postInstanceTestEventOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *postInstanceTestEventOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/test_event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postInstanceTestEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "PostInstanceTestEvent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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

// GetInstancePlans : Retrieves instance plan
// Retrieves the plan of an instance.
func (adminService *AdminServiceV3) GetInstancePlans(getInstancePlansOptions *GetInstancePlansOptions) (result *InstancePlan, response *core.DetailedResponse, err error) {
	return adminService.GetInstancePlansWithContext(context.Background(), getInstancePlansOptions)
}

// GetInstancePlansWithContext is an alternate form of the GetInstancePlans method which supports a Context parameter
func (adminService *AdminServiceV3) GetInstancePlansWithContext(ctx context.Context, getInstancePlansOptions *GetInstancePlansOptions) (result *InstancePlan, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInstancePlansOptions, "getInstancePlansOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInstancePlansOptions, "getInstancePlansOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getInstancePlansOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/plans`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInstancePlansOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "GetInstancePlans")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstancePlan)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostInstancePlans : Create instance plan
// Create instance plan.
func (adminService *AdminServiceV3) PostInstancePlans(postInstancePlansOptions *PostInstancePlansOptions) (result *InstancePlan, response *core.DetailedResponse, err error) {
	return adminService.PostInstancePlansWithContext(context.Background(), postInstancePlansOptions)
}

// PostInstancePlansWithContext is an alternate form of the PostInstancePlans method which supports a Context parameter
func (adminService *AdminServiceV3) PostInstancePlansWithContext(ctx context.Context, postInstancePlansOptions *PostInstancePlansOptions) (result *InstancePlan, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postInstancePlansOptions, "postInstancePlansOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postInstancePlansOptions, "postInstancePlansOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *postInstancePlansOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/plans`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postInstancePlansOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "PostInstancePlans")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postInstancePlansOptions.Name != nil {
		body["name"] = postInstancePlansOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstancePlan)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetInstancePlans : Update instance plan
// Change instance plan.
func (adminService *AdminServiceV3) SetInstancePlans(setInstancePlansOptions *SetInstancePlansOptions) (result *InstancePlan, response *core.DetailedResponse, err error) {
	return adminService.SetInstancePlansWithContext(context.Background(), setInstancePlansOptions)
}

// SetInstancePlansWithContext is an alternate form of the SetInstancePlans method which supports a Context parameter
func (adminService *AdminServiceV3) SetInstancePlansWithContext(ctx context.Context, setInstancePlansOptions *SetInstancePlansOptions) (result *InstancePlan, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setInstancePlansOptions, "setInstancePlansOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setInstancePlansOptions, "setInstancePlansOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *setInstancePlansOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminService.Service.Options.URL, `/instances/{instance_id}/v3/plans`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setInstancePlansOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("admin_service", "V3", "SetInstancePlans")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setInstancePlansOptions.Name != nil {
		body["name"] = setInstancePlansOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstancePlan)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// EnCrn : Event Notifications settings.
type EnCrn struct {
	// EN Instance CRN.
	InstanceCrn *string `json:"instance_crn,omitempty"`

	// Name of the source.
	Name *string `json:"name,omitempty"`

	// Description of the source.
	Description *string `json:"description,omitempty"`
}

// UnmarshalEnCrn unmarshals an instance of EnCrn from the specified map of raw messages.
func UnmarshalEnCrn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnCrn)
	err = core.UnmarshalPrimitive(m, "instance_crn", &obj.InstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetInstanceAccessOptions : The GetInstanceAccess options.
type GetInstanceAccessOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceAccessOptions : Instantiate GetInstanceAccessOptions
func (*AdminServiceV3) NewGetInstanceAccessOptions(instanceID string) *GetInstanceAccessOptions {
	return &GetInstanceAccessOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceAccessOptions) SetInstanceID(instanceID string) *GetInstanceAccessOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceAccessOptions) SetHeaders(param map[string]string) *GetInstanceAccessOptions {
	options.Headers = param
	return options
}

// GetInstancePlansOptions : The GetInstancePlans options.
type GetInstancePlansOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstancePlansOptions : Instantiate GetInstancePlansOptions
func (*AdminServiceV3) NewGetInstancePlansOptions(instanceID string) *GetInstancePlansOptions {
	return &GetInstancePlansOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstancePlansOptions) SetInstanceID(instanceID string) *GetInstancePlansOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstancePlansOptions) SetHeaders(param map[string]string) *GetInstancePlansOptions {
	options.Headers = param
	return options
}

// GetInstanceSettingsOptions : The GetInstanceSettings options.
type GetInstanceSettingsOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInstanceSettingsOptions : Instantiate GetInstanceSettingsOptions
func (*AdminServiceV3) NewGetInstanceSettingsOptions(instanceID string) *GetInstanceSettingsOptions {
	return &GetInstanceSettingsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetInstanceSettingsOptions) SetInstanceID(instanceID string) *GetInstanceSettingsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInstanceSettingsOptions) SetHeaders(param map[string]string) *GetInstanceSettingsOptions {
	options.Headers = param
	return options
}

// InstanceAccess : The caller's level of access for the specified home instance.
type InstanceAccess struct {
	// The caller's level of access for the specified home instance.
	Admin *InstanceAccessLevels `json:"admin" validate:"required"`
}

// UnmarshalInstanceAccess unmarshals an instance of InstanceAccess from the specified map of raw messages.
func UnmarshalInstanceAccess(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceAccess)
	err = core.UnmarshalModel(m, "admin", &obj.Admin, UnmarshalInstanceAccessLevels)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceAccessLevels : The caller's level of access for the specified home instance.
type InstanceAccessLevels struct {
	// User access to read properties.
	Read *bool `json:"read" validate:"required"`

	// User access to update properties.
	Update *bool `json:"update" validate:"required"`
}

// UnmarshalInstanceAccessLevels unmarshals an instance of InstanceAccessLevels from the specified map of raw messages.
func UnmarshalInstanceAccessLevels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceAccessLevels)
	err = core.UnmarshalPrimitive(m, "read", &obj.Read)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update", &obj.Update)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstancePlan : Instance plan.
type InstancePlan struct {
	// Plan name, such as Trial, Standard.
	Name *string `json:"name" validate:"required"`
}

// NewInstancePlan : Instantiate InstancePlan (Generic Model Constructor)
func (*AdminServiceV3) NewInstancePlan(name string) (_model *InstancePlan, err error) {
	_model = &InstancePlan{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalInstancePlan unmarshals an instance of InstancePlan from the specified map of raw messages.
func UnmarshalInstancePlan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstancePlan)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceSettings : Instance settings.
type InstanceSettings struct {
	// Location settings.
	Location *LocationID `json:"location,omitempty"`

	// Event Notifications settings.
	EventNotifications *EnCrn `json:"event_notifications,omitempty"`
}

// UnmarshalInstanceSettings unmarshals an instance of InstanceSettings from the specified map of raw messages.
func UnmarshalInstanceSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceSettings)
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocationID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "event_notifications", &obj.EventNotifications, UnmarshalEnCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

// LocationID : Location settings.
type LocationID struct {
	// Location ID.
	ID *string `json:"id" validate:"required"`
}

// Constants associated with the LocationID.ID property.
// Location ID.
const (
	LocationID_ID_CaTor = "ca-tor"
	LocationID_ID_Eu    = "eu"
	LocationID_ID_EuFr2 = "eu-fr2"
	LocationID_ID_Us    = "us"
)

// UnmarshalLocationID unmarshals an instance of LocationID from the specified map of raw messages.
func UnmarshalLocationID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LocationID)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostInstancePlansOptions : The PostInstancePlans options.
type PostInstancePlansOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Plan name, such as Trial, Standard.
	Name *string `json:"name" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostInstancePlansOptions : Instantiate PostInstancePlansOptions
func (*AdminServiceV3) NewPostInstancePlansOptions(instanceID string, name string) *PostInstancePlansOptions {
	return &PostInstancePlansOptions{
		InstanceID: core.StringPtr(instanceID),
		Name:       core.StringPtr(name),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *PostInstancePlansOptions) SetInstanceID(instanceID string) *PostInstancePlansOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *PostInstancePlansOptions) SetName(name string) *PostInstancePlansOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostInstancePlansOptions) SetHeaders(param map[string]string) *PostInstancePlansOptions {
	options.Headers = param
	return options
}

// PostInstanceTestEventOptions : The PostInstanceTestEvent options.
type PostInstanceTestEventOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostInstanceTestEventOptions : Instantiate PostInstanceTestEventOptions
func (*AdminServiceV3) NewPostInstanceTestEventOptions(instanceID string) *PostInstanceTestEventOptions {
	return &PostInstanceTestEventOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *PostInstanceTestEventOptions) SetInstanceID(instanceID string) *PostInstanceTestEventOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostInstanceTestEventOptions) SetHeaders(param map[string]string) *PostInstanceTestEventOptions {
	options.Headers = param
	return options
}

// SetInstancePlansOptions : The SetInstancePlans options.
type SetInstancePlansOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Plan name, such as Trial, Standard.
	Name *string `json:"name" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetInstancePlansOptions : Instantiate SetInstancePlansOptions
func (*AdminServiceV3) NewSetInstancePlansOptions(instanceID string, name string) *SetInstancePlansOptions {
	return &SetInstancePlansOptions{
		InstanceID: core.StringPtr(instanceID),
		Name:       core.StringPtr(name),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *SetInstancePlansOptions) SetInstanceID(instanceID string) *SetInstancePlansOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *SetInstancePlansOptions) SetName(name string) *SetInstancePlansOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetInstancePlansOptions) SetHeaders(param map[string]string) *SetInstancePlansOptions {
	options.Headers = param
	return options
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

// UpdateInstanceSettingsOptions : The UpdateInstanceSettings options.
type UpdateInstanceSettingsOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// New instance settings.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateInstanceSettingsOptions : Instantiate UpdateInstanceSettingsOptions
func (*AdminServiceV3) NewUpdateInstanceSettingsOptions(instanceID string, jsonPatchOperation []JSONPatchOperation) *UpdateInstanceSettingsOptions {
	return &UpdateInstanceSettingsOptions{
		InstanceID:         core.StringPtr(instanceID),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateInstanceSettingsOptions) SetInstanceID(instanceID string) *UpdateInstanceSettingsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *UpdateInstanceSettingsOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *UpdateInstanceSettingsOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateInstanceSettingsOptions) SetHeaders(param map[string]string) *UpdateInstanceSettingsOptions {
	options.Headers = param
	return options
}
