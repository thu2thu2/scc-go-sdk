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

// Package configmanagerapiv3 : Operations and models for the ConfigManagerApiV3 service
package configmanagerapiv3

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

// ConfigManagerApiV3 : This is an API for the Config Manager
//
// API Version: 3.0.0
type ConfigManagerApiV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us.compliance.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "config_manager_api"

// ConfigManagerApiV3Options : Service options
type ConfigManagerApiV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewConfigManagerApiV3UsingExternalConfig : constructs an instance of ConfigManagerApiV3 with passed in options and external configuration.
func NewConfigManagerApiV3UsingExternalConfig(options *ConfigManagerApiV3Options) (configManagerApi *ConfigManagerApiV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	configManagerApi, err = NewConfigManagerApiV3(options)
	if err != nil {
		return
	}

	err = configManagerApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = configManagerApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewConfigManagerApiV3 : constructs an instance of ConfigManagerApiV3 with passed in options.
func NewConfigManagerApiV3(options *ConfigManagerApiV3Options) (service *ConfigManagerApiV3, err error) {
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

	service = &ConfigManagerApiV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-south": "https://us.compliance.cloud.ibm.com",
		"us-east":  "https://us.compliance.cloud.ibm.com",
		"eu-de":    "https://eu.compliance.cloud.ibm.com",
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "configManagerApi" suitable for processing requests.
func (configManagerApi *ConfigManagerApiV3) Clone() *ConfigManagerApiV3 {
	if core.IsNil(configManagerApi) {
		return nil
	}
	clone := *configManagerApi
	clone.Service = configManagerApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (configManagerApi *ConfigManagerApiV3) SetServiceURL(url string) error {
	return configManagerApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (configManagerApi *ConfigManagerApiV3) GetServiceURL() string {
	return configManagerApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (configManagerApi *ConfigManagerApiV3) SetDefaultHeaders(headers http.Header) {
	configManagerApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (configManagerApi *ConfigManagerApiV3) SetEnableGzipCompression(enableGzip bool) {
	configManagerApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (configManagerApi *ConfigManagerApiV3) GetEnableGzipCompression() bool {
	return configManagerApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (configManagerApi *ConfigManagerApiV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	configManagerApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (configManagerApi *ConfigManagerApiV3) DisableRetries() {
	configManagerApi.Service.DisableRetries()
}

// CreateRule : Create a user defined rule
// Create a user defined rule.
func (configManagerApi *ConfigManagerApiV3) CreateRule(createRuleOptions *CreateRuleOptions) (result *Rules, response *core.DetailedResponse, err error) {
	return configManagerApi.CreateRuleWithContext(context.Background(), createRuleOptions)
}

// CreateRuleWithContext is an alternate form of the CreateRule method which supports a Context parameter
func (configManagerApi *ConfigManagerApiV3) CreateRuleWithContext(ctx context.Context, createRuleOptions *CreateRuleOptions) (result *Rules, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRuleOptions, "createRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRuleOptions, "createRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createRuleOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManagerApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManagerApi.Service.Options.URL, `/instances/{instance_id}/v3/rules`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager_api", "V3", "CreateRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createRuleOptions.TypeQuery != nil {
		builder.AddQuery("type_query", fmt.Sprint(*createRuleOptions.TypeQuery))
	}

	body := make(map[string]interface{})
	if createRuleOptions.Rules != nil {
		body["rules"] = createRuleOptions.Rules
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
	response, err = configManagerApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRules)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRules : Retrieve all rules
// Retrieve all rules.
func (configManagerApi *ConfigManagerApiV3) ListRules(listRulesOptions *ListRulesOptions) (result *Rules, response *core.DetailedResponse, err error) {
	return configManagerApi.ListRulesWithContext(context.Background(), listRulesOptions)
}

// ListRulesWithContext is an alternate form of the ListRules method which supports a Context parameter
func (configManagerApi *ConfigManagerApiV3) ListRulesWithContext(ctx context.Context, listRulesOptions *ListRulesOptions) (result *Rules, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listRulesOptions, "listRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listRulesOptions, "listRulesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listRulesOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManagerApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManagerApi.Service.Options.URL, `/instances/{instance_id}/v3/rules`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager_api", "V3", "ListRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listRulesOptions.TypeQuery != nil {
		builder.AddQuery("type_query", fmt.Sprint(*listRulesOptions.TypeQuery))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configManagerApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRules)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRule : Retrieve a specific user defined rule
// Retrieve a specific user defined rule.
func (configManagerApi *ConfigManagerApiV3) GetRule(getRuleOptions *GetRuleOptions) (result *Rules, response *core.DetailedResponse, err error) {
	return configManagerApi.GetRuleWithContext(context.Background(), getRuleOptions)
}

// GetRuleWithContext is an alternate form of the GetRule method which supports a Context parameter
func (configManagerApi *ConfigManagerApiV3) GetRuleWithContext(ctx context.Context, getRuleOptions *GetRuleOptions) (result *Rules, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRuleOptions, "getRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRuleOptions, "getRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":     *getRuleOptions.RuleID,
		"instance_id": *getRuleOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManagerApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManagerApi.Service.Options.URL, `/instances/{instance_id}/v3/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager_api", "V3", "GetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configManagerApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRules)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddRule : Update a specific user defined rule
// Update a specific user defined rule.
func (configManagerApi *ConfigManagerApiV3) AddRule(addRuleOptions *AddRuleOptions) (result *Rules, response *core.DetailedResponse, err error) {
	return configManagerApi.AddRuleWithContext(context.Background(), addRuleOptions)
}

// AddRuleWithContext is an alternate form of the AddRule method which supports a Context parameter
func (configManagerApi *ConfigManagerApiV3) AddRuleWithContext(ctx context.Context, addRuleOptions *AddRuleOptions) (result *Rules, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addRuleOptions, "addRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addRuleOptions, "addRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":     *addRuleOptions.RuleID,
		"instance_id": *addRuleOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManagerApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManagerApi.Service.Options.URL, `/instances/{instance_id}/v3/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager_api", "V3", "AddRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	//builder.AddHeader("If-Match", "2-39de1d64d11e67f7b92d5b0b647708c5")

	body := make(map[string]interface{})
	if addRuleOptions.Rules != nil {
		body["rules"] = addRuleOptions.Rules
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
	response, err = configManagerApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRules)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRule : Delete a specific user defined rule
// Delete a specific user defined rule.
func (configManagerApi *ConfigManagerApiV3) DeleteRule(deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	return configManagerApi.DeleteRuleWithContext(context.Background(), deleteRuleOptions)
}

// DeleteRuleWithContext is an alternate form of the DeleteRule method which supports a Context parameter
func (configManagerApi *ConfigManagerApiV3) DeleteRuleWithContext(ctx context.Context, deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRuleOptions, "deleteRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRuleOptions, "deleteRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":     *deleteRuleOptions.RuleID,
		"instance_id": *deleteRuleOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManagerApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManagerApi.Service.Options.URL, `/instances/{instance_id}/v3/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager_api", "V3", "DeleteRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = configManagerApi.Service.Request(request, nil)

	return
}

// AddRuleOptions : The AddRule options.
type AddRuleOptions struct {
	// The ID of the corresponding rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// List of rules.
	Rules []Rule `json:"rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddRuleOptions : Instantiate AddRuleOptions
func (*ConfigManagerApiV3) NewAddRuleOptions(ruleID string, instanceID string) *AddRuleOptions {
	return &AddRuleOptions{
		RuleID:     core.StringPtr(ruleID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetRuleID : Allow user to set RuleID
func (_options *AddRuleOptions) SetRuleID(ruleID string) *AddRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *AddRuleOptions) SetInstanceID(instanceID string) *AddRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *AddRuleOptions) SetRules(rules []Rule) *AddRuleOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddRuleOptions) SetHeaders(param map[string]string) *AddRuleOptions {
	options.Headers = param
	return options
}

// And : AND required configurations.
type And struct {
	// Property.
	Property *string `json:"property,omitempty"`

	// Operator.
	Operator *string `json:"operator,omitempty"`
}

// UnmarshalAnd unmarshals an instance of And from the specified map of raw messages.
func UnmarshalAnd(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(And)
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRuleOptions : The CreateRule options.
type CreateRuleOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// List of rules.
	Rules []Rule `json:"rules,omitempty"`

	// List user_defined or system_defined rules only.
	TypeQuery *string `json:"type_query,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRuleOptions : Instantiate CreateRuleOptions
func (*ConfigManagerApiV3) NewCreateRuleOptions(instanceID string) *CreateRuleOptions {
	return &CreateRuleOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateRuleOptions) SetInstanceID(instanceID string) *CreateRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *CreateRuleOptions) SetRules(rules []Rule) *CreateRuleOptions {
	_options.Rules = rules
	return _options
}

// SetTypeQuery : Allow user to set TypeQuery
func (_options *CreateRuleOptions) SetTypeQuery(typeQuery string) *CreateRuleOptions {
	_options.TypeQuery = core.StringPtr(typeQuery)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRuleOptions) SetHeaders(param map[string]string) *CreateRuleOptions {
	options.Headers = param
	return options
}

// DeleteRuleOptions : The DeleteRule options.
type DeleteRuleOptions struct {
	// The ID of the corresponding rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRuleOptions : Instantiate DeleteRuleOptions
func (*ConfigManagerApiV3) NewDeleteRuleOptions(ruleID string, instanceID string) *DeleteRuleOptions {
	return &DeleteRuleOptions{
		RuleID:     core.StringPtr(ruleID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetRuleID : Allow user to set RuleID
func (_options *DeleteRuleOptions) SetRuleID(ruleID string) *DeleteRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteRuleOptions) SetInstanceID(instanceID string) *DeleteRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRuleOptions) SetHeaders(param map[string]string) *DeleteRuleOptions {
	options.Headers = param
	return options
}

// GetRuleOptions : The GetRule options.
type GetRuleOptions struct {
	// The ID of the corresponding rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRuleOptions : Instantiate GetRuleOptions
func (*ConfigManagerApiV3) NewGetRuleOptions(ruleID string, instanceID string) *GetRuleOptions {
	return &GetRuleOptions{
		RuleID:     core.StringPtr(ruleID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetRuleID : Allow user to set RuleID
func (_options *GetRuleOptions) SetRuleID(ruleID string) *GetRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetRuleOptions) SetInstanceID(instanceID string) *GetRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRuleOptions) SetHeaders(param map[string]string) *GetRuleOptions {
	options.Headers = param
	return options
}

// Import : Imports.
type Import struct {
	// Name.
	Name *string `json:"name,omitempty"`

	// Display name.
	DisplayName *string `json:"display_name,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the Import.Type property.
// Type.
const (
	Import_Type_Boolean    = "boolean"
	Import_Type_General    = "general"
	Import_Type_IpList     = "ip_list"
	Import_Type_Numeric    = "numeric"
	Import_Type_String     = "string"
	Import_Type_StringList = "string_list"
	Import_Type_Timestamp  = "timestamp"
)

// UnmarshalImport unmarshals an instance of Import from the specified map of raw messages.
func UnmarshalImport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Import)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListRulesOptions : The ListRules options.
type ListRulesOptions struct {
	// The ID of the managing instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// List user_defined or system_defined rules only.
	TypeQuery *string `json:"type_query,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRulesOptions : Instantiate ListRulesOptions
func (*ConfigManagerApiV3) NewListRulesOptions(instanceID string) *ListRulesOptions {
	return &ListRulesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListRulesOptions) SetInstanceID(instanceID string) *ListRulesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTypeQuery : Allow user to set TypeQuery
func (_options *ListRulesOptions) SetTypeQuery(typeQuery string) *ListRulesOptions {
	_options.TypeQuery = core.StringPtr(typeQuery)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListRulesOptions) SetHeaders(param map[string]string) *ListRulesOptions {
	options.Headers = param
	return options
}

// Or : OR required configurations.
type Or struct {
	// Property.
	Property *string `json:"property,omitempty"`

	// Operator.
	Operator *string `json:"operator,omitempty"`
}

// UnmarshalOr unmarshals an instance of Or from the specified map of raw messages.
func UnmarshalOr(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Or)
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Parameters : List of import parameters.
type Parameters struct {
	// Import parameters.
	Imports []Import `json:"imports,omitempty"`
}

// UnmarshalParameters unmarshals an instance of Parameters from the specified map of raw messages.
func UnmarshalParameters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Parameters)
	err = core.UnmarshalModel(m, "imports", &obj.Imports, UnmarshalImport)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RequiredConfig : Required Configurations.
type RequiredConfig struct {
	// AND required configurations.
	And []And `json:"and,omitempty"`

	// OR required configurations.
	Or []Or `json:"or,omitempty"`
}

// UnmarshalRequiredConfig unmarshals an instance of RequiredConfig from the specified map of raw messages.
func UnmarshalRequiredConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RequiredConfig)
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalAnd)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalOr)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : Rule.
type Rule struct {
	// Creation date.
	CreationDate *string `json:"creation_date,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Modification date.
	ModificationDate *string `json:"modification_date,omitempty"`

	// Modified by.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// Rule ID.
	ID *string `json:"id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Details of a rules response.
	Description *string `json:"description" validate:"required"`

	// Rule type (user_defined or system_defined).
	Type *string `json:"type,omitempty"`

	// Verison number of rule.
	Version *string `json:"version,omitempty"`

	// List of import parameters.
	Import *Parameters `json:"import,omitempty"`

	// Service target.
	Target *Target `json:"target,omitempty"`

	// Required Configurations.
	RequiredConfig *RequiredConfig `json:"required_config,omitempty"`

	// Labels.
	Labels []string `json:"labels,omitempty"`
}

// Constants associated with the Rule.Type property.
// Rule type (user_defined or system_defined).
const (
	Rule_Type_SystemDefined = "system_defined"
	Rule_Type_UserDefined   = "user_defined"
)

// NewRule : Instantiate Rule (Generic Model Constructor)
func (*ConfigManagerApiV3) NewRule(description string) (_model *Rule, err error) {
	_model = &Rule{
		Description: core.StringPtr(description),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "creation_date", &obj.CreationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modification_date", &obj.ModificationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "import", &obj.Import, UnmarshalParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_config", &obj.RequiredConfig, UnmarshalRequiredConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rules : Rule.
type Rules struct {
	// List of rules.
	Rules []Rule `json:"rules,omitempty"`
}

// UnmarshalRules unmarshals an instance of Rules from the specified map of raw messages.
func UnmarshalRules(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rules)
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Target : Service target.
type Target struct {
	// Service name.
	ServiceName *string `json:"service_name,omitempty"`

	// Service display name.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// Resource kind.
	ResourceKind *string `json:"resource_kind,omitempty"`

	// Supported properties.
	AdditionalTargetAttributes []string `json:"additional_target_attributes,omitempty"`
}

// UnmarshalTarget unmarshals an instance of Target from the specified map of raw messages.
func UnmarshalTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Target)
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_display_name", &obj.ServiceDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_kind", &obj.ResourceKind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "additional_target_attributes", &obj.AdditionalTargetAttributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
