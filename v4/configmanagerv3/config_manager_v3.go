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

// Package configmanagerv3 : Operations and models for the ConfigManagerV3 service
package configmanagerv3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/scc-go-sdk/v4/common"
	"github.com/go-openapi/strfmt"
)

// ConfigManagerV3 : The Security and Compliance Center API reference.
//
// API Version: 3.0.0
type ConfigManagerV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "config_manager"

const ParameterizedServiceURL = "https://{environment}.cloud.ibm.com/instances/{instance_id}/v3"

var defaultUrlVariables = map[string]string{
	"environment": "us-south.compliance",
	"instance_id": "instance_id",
}

// ConfigManagerV3Options : Service options
type ConfigManagerV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewConfigManagerV3UsingExternalConfig : constructs an instance of ConfigManagerV3 with passed in options and external configuration.
func NewConfigManagerV3UsingExternalConfig(options *ConfigManagerV3Options) (configManager *ConfigManagerV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	configManager, err = NewConfigManagerV3(options)
	if err != nil {
		return
	}

	err = configManager.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = configManager.Service.SetServiceURL(options.URL)
	}
	return
}

// NewConfigManagerV3 : constructs an instance of ConfigManagerV3 with passed in options.
func NewConfigManagerV3(options *ConfigManagerV3Options) (service *ConfigManagerV3, err error) {
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

	service = &ConfigManagerV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "configManager" suitable for processing requests.
func (configManager *ConfigManagerV3) Clone() *ConfigManagerV3 {
	if core.IsNil(configManager) {
		return nil
	}
	clone := *configManager
	clone.Service = configManager.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (configManager *ConfigManagerV3) SetServiceURL(url string) error {
	return configManager.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (configManager *ConfigManagerV3) GetServiceURL() string {
	return configManager.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (configManager *ConfigManagerV3) SetDefaultHeaders(headers http.Header) {
	configManager.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (configManager *ConfigManagerV3) SetEnableGzipCompression(enableGzip bool) {
	configManager.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (configManager *ConfigManagerV3) GetEnableGzipCompression() bool {
	return configManager.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (configManager *ConfigManagerV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	configManager.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (configManager *ConfigManagerV3) DisableRetries() {
	configManager.Service.DisableRetries()
}

// ListRules : List all rules
// Retrieve all the rules that you use to evaluate your resources.
func (configManager *ConfigManagerV3) ListRules(listRulesOptions *ListRulesOptions) (result *Rules, response *core.DetailedResponse, err error) {
	return configManager.ListRulesWithContext(context.Background(), listRulesOptions)
}

// ListRulesWithContext is an alternate form of the ListRules method which supports a Context parameter
func (configManager *ConfigManagerV3) ListRulesWithContext(ctx context.Context, listRulesOptions *ListRulesOptions) (result *Rules, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRulesOptions, "listRulesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManager.Service.Options.URL, `/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager", "V3", "ListRules")
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
	response, err = configManager.Service.Request(request, &rawResponse)
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

// CreateRule : Create a custom rule
// Create a custom rule to target the exact configuration properties that you need to evaluate your resources for
// compliance.
func (configManager *ConfigManagerV3) CreateRule(createRuleOptions *CreateRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configManager.CreateRuleWithContext(context.Background(), createRuleOptions)
}

// CreateRuleWithContext is an alternate form of the CreateRule method which supports a Context parameter
func (configManager *ConfigManagerV3) CreateRuleWithContext(ctx context.Context, createRuleOptions *CreateRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRuleOptions, "createRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRuleOptions, "createRuleOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManager.Service.Options.URL, `/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager", "V3", "CreateRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createRuleOptions.TypeQuery != nil {
		builder.AddQuery("type_query", fmt.Sprint(*createRuleOptions.TypeQuery))
	}

	body := make(map[string]interface{})
	if createRuleOptions.AccountID != nil {
		body["account_id"] = createRuleOptions.AccountID
	}
	if createRuleOptions.Description != nil {
		body["description"] = createRuleOptions.Description
	}
	if createRuleOptions.Target != nil {
		body["target"] = createRuleOptions.Target
	}
	if createRuleOptions.RequiredConfig != nil {
		body["required_config"] = createRuleOptions.RequiredConfig
	}
	if createRuleOptions.Labels != nil {
		body["labels"] = createRuleOptions.Labels
	}
	if createRuleOptions.Type != nil {
		body["type"] = createRuleOptions.Type
	}
	if createRuleOptions.Version != nil {
		body["version"] = createRuleOptions.Version
	}
	if createRuleOptions.Import != nil {
		body["import"] = createRuleOptions.Import
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
	response, err = configManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRule : Get a custom rule
// Retrieve a user-defined rule that you use to evaluate your resources.
func (configManager *ConfigManagerV3) GetRule(getRuleOptions *GetRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configManager.GetRuleWithContext(context.Background(), getRuleOptions)
}

// GetRuleWithContext is an alternate form of the GetRule method which supports a Context parameter
func (configManager *ConfigManagerV3) GetRuleWithContext(ctx context.Context, getRuleOptions *GetRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRuleOptions, "getRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRuleOptions, "getRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *getRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManager.Service.Options.URL, `/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager", "V3", "GetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceRule : Update a custom rule
// Update a user-defined rule that you use to evaluate your resources.
func (configManager *ConfigManagerV3) ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configManager.ReplaceRuleWithContext(context.Background(), replaceRuleOptions)
}

// ReplaceRuleWithContext is an alternate form of the ReplaceRule method which supports a Context parameter
func (configManager *ConfigManagerV3) ReplaceRuleWithContext(ctx context.Context, replaceRuleOptions *ReplaceRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceRuleOptions, "replaceRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceRuleOptions, "replaceRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *replaceRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManager.Service.Options.URL, `/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager", "V3", "ReplaceRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddHeader("If-Match", "W/\"1-41e98decc9a311cea98c1702f39da727\"")
	if replaceRuleOptions.IfMatch != nil {
		builder.AddHeader("IfMatch", fmt.Sprint(*replaceRuleOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if replaceRuleOptions.AccountID != nil {
		body["account_id"] = replaceRuleOptions.AccountID
	}
	if replaceRuleOptions.Description != nil {
		body["description"] = replaceRuleOptions.Description
	}
	if replaceRuleOptions.Target != nil {
		body["target"] = replaceRuleOptions.Target
	}
	if replaceRuleOptions.RequiredConfig != nil {
		body["required_config"] = replaceRuleOptions.RequiredConfig
	}
	if replaceRuleOptions.Labels != nil {
		body["labels"] = replaceRuleOptions.Labels
	}
	if replaceRuleOptions.Type != nil {
		body["type"] = replaceRuleOptions.Type
	}
	if replaceRuleOptions.Version != nil {
		body["version"] = replaceRuleOptions.Version
	}
	if replaceRuleOptions.Import != nil {
		body["import"] = replaceRuleOptions.Import
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
	response, err = configManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRule : Delete a custom rule
// Delete a user-defined rule that you no longer require.
func (configManager *ConfigManagerV3) DeleteRule(deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	return configManager.DeleteRuleWithContext(context.Background(), deleteRuleOptions)
}

// DeleteRuleWithContext is an alternate form of the DeleteRule method which supports a Context parameter
func (configManager *ConfigManagerV3) DeleteRuleWithContext(ctx context.Context, deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRuleOptions, "deleteRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRuleOptions, "deleteRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *deleteRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configManager.Service.Options.URL, `/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("config_manager", "V3", "DeleteRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = configManager.Service.Request(request, nil)

	return
}

// AdditionalTargetAttribute : The additional target attribute of the service.
type AdditionalTargetAttribute struct {
	// The additional target attribute name.
	Name *string `json:"name,omitempty"`

	// The operator.
	Operator *string `json:"operator,omitempty"`

	// The value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the AdditionalTargetAttribute.Operator property.
// The operator.
const (
	AdditionalTargetAttribute_Operator_DaysLessThan         = "days_less_than"
	AdditionalTargetAttribute_Operator_IpsEquals            = "ips_equals"
	AdditionalTargetAttribute_Operator_IpsInRange           = "ips_in_range"
	AdditionalTargetAttribute_Operator_IpsNotEquals         = "ips_not_equals"
	AdditionalTargetAttribute_Operator_IsEmpty              = "is_empty"
	AdditionalTargetAttribute_Operator_IsFalse              = "is_false"
	AdditionalTargetAttribute_Operator_IsNotEmpty           = "is_not_empty"
	AdditionalTargetAttribute_Operator_IsTrue               = "is_true"
	AdditionalTargetAttribute_Operator_NumEquals            = "num_equals"
	AdditionalTargetAttribute_Operator_NumGreaterThan       = "num_greater_than"
	AdditionalTargetAttribute_Operator_NumGreaterThanEquals = "num_greater_than_equals"
	AdditionalTargetAttribute_Operator_NumLessThan          = "num_less_than"
	AdditionalTargetAttribute_Operator_NumLessThanEquals    = "num_less_than_equals"
	AdditionalTargetAttribute_Operator_NumNotEquals         = "num_not_equals"
	AdditionalTargetAttribute_Operator_StringContains       = "string_contains"
	AdditionalTargetAttribute_Operator_StringEquals         = "string_equals"
	AdditionalTargetAttribute_Operator_StringMatch          = "string_match"
	AdditionalTargetAttribute_Operator_StringNotContains    = "string_not_contains"
	AdditionalTargetAttribute_Operator_StringNotEquals      = "string_not_equals"
	AdditionalTargetAttribute_Operator_StringNotMatch       = "string_not_match"
	AdditionalTargetAttribute_Operator_StringsAllowed       = "strings_allowed"
	AdditionalTargetAttribute_Operator_StringsInList        = "strings_in_list"
	AdditionalTargetAttribute_Operator_StringsRequired      = "strings_required"
)

// UnmarshalAdditionalTargetAttribute unmarshals an instance of AdditionalTargetAttribute from the specified map of raw messages.
func UnmarshalAdditionalTargetAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AdditionalTargetAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
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

// And : The `AND` required configurations.
type And struct {
	// The property.
	Property *string `json:"property,omitempty"`

	// The operator.
	Operator *string `json:"operator,omitempty"`

	// The value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the And.Operator property.
// The operator.
const (
	And_Operator_DaysLessThan         = "days_less_than"
	And_Operator_IpsEquals            = "ips_equals"
	And_Operator_IpsInRange           = "ips_in_range"
	And_Operator_IpsNotEquals         = "ips_not_equals"
	And_Operator_IsEmpty              = "is_empty"
	And_Operator_IsFalse              = "is_false"
	And_Operator_IsNotEmpty           = "is_not_empty"
	And_Operator_IsTrue               = "is_true"
	And_Operator_NumEquals            = "num_equals"
	And_Operator_NumGreaterThan       = "num_greater_than"
	And_Operator_NumGreaterThanEquals = "num_greater_than_equals"
	And_Operator_NumLessThan          = "num_less_than"
	And_Operator_NumLessThanEquals    = "num_less_than_equals"
	And_Operator_NumNotEquals         = "num_not_equals"
	And_Operator_StringContains       = "string_contains"
	And_Operator_StringEquals         = "string_equals"
	And_Operator_StringMatch          = "string_match"
	And_Operator_StringNotContains    = "string_not_contains"
	And_Operator_StringNotEquals      = "string_not_equals"
	And_Operator_StringNotMatch       = "string_not_match"
	And_Operator_StringsAllowed       = "strings_allowed"
	And_Operator_StringsInList        = "strings_in_list"
	And_Operator_StringsRequired      = "strings_required"
)

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
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRuleOptions : The CreateRule options.
type CreateRuleOptions struct {
	// Account identification value.
	AccountID *string `json:"account_id" validate:"required"`

	// Rule description.
	Description *string `json:"description" validate:"required"`

	// The rule target.
	Target *Target `json:"target" validate:"required"`

	// The required configurations.
	RequiredConfig *RequiredConfig `json:"required_config" validate:"required"`

	// List of labels corresponding to rule.
	Labels []string `json:"labels" validate:"required"`

	// Rule type (user_defined or system_defined).
	Type *string `json:"type,omitempty"`

	// Rule verison number.
	Version *string `json:"version,omitempty"`

	// The collection of import parameters.
	Import *Import `json:"import,omitempty"`

	// The list of only user defined, or system defined rules.
	TypeQuery *string `json:"type_query,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateRuleOptions.Type property.
// Rule type (user_defined or system_defined).
const (
	CreateRuleOptions_Type_SystemDefined = "system_defined"
	CreateRuleOptions_Type_UserDefined   = "user_defined"
)

// NewCreateRuleOptions : Instantiate CreateRuleOptions
func (*ConfigManagerV3) NewCreateRuleOptions(accountID string, description string, target *Target, requiredConfig *RequiredConfig, labels []string) *CreateRuleOptions {
	return &CreateRuleOptions{
		AccountID:      core.StringPtr(accountID),
		Description:    core.StringPtr(description),
		Target:         target,
		RequiredConfig: requiredConfig,
		Labels:         labels,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *CreateRuleOptions) SetAccountID(accountID string) *CreateRuleOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateRuleOptions) SetDescription(description string) *CreateRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *CreateRuleOptions) SetTarget(target *Target) *CreateRuleOptions {
	_options.Target = target
	return _options
}

// SetRequiredConfig : Allow user to set RequiredConfig
func (_options *CreateRuleOptions) SetRequiredConfig(requiredConfig *RequiredConfig) *CreateRuleOptions {
	_options.RequiredConfig = requiredConfig
	return _options
}

// SetLabels : Allow user to set Labels
func (_options *CreateRuleOptions) SetLabels(labels []string) *CreateRuleOptions {
	_options.Labels = labels
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateRuleOptions) SetType(typeVar string) *CreateRuleOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreateRuleOptions) SetVersion(version string) *CreateRuleOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetImport : Allow user to set Import
func (_options *CreateRuleOptions) SetImport(importVar *Import) *CreateRuleOptions {
	_options.Import = importVar
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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRuleOptions : Instantiate DeleteRuleOptions
func (*ConfigManagerV3) NewDeleteRuleOptions(ruleID string) *DeleteRuleOptions {
	return &DeleteRuleOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (_options *DeleteRuleOptions) SetRuleID(ruleID string) *DeleteRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRuleOptions : Instantiate GetRuleOptions
func (*ConfigManagerV3) NewGetRuleOptions(ruleID string) *GetRuleOptions {
	return &GetRuleOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (_options *GetRuleOptions) SetRuleID(ruleID string) *GetRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRuleOptions) SetHeaders(param map[string]string) *GetRuleOptions {
	options.Headers = param
	return options
}

// Import : The collection of import parameters.
type Import struct {
	// The list of import parameters.
	Parameters []Parameter `json:"parameters,omitempty"`
}

// UnmarshalImport unmarshals an instance of Import from the specified map of raw messages.
func UnmarshalImport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Import)
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalParameter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListRulesOptions : The ListRules options.
type ListRulesOptions struct {
	// The list of only user defined, or system defined rules.
	TypeQuery *string `json:"type_query,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRulesOptions : Instantiate ListRulesOptions
func (*ConfigManagerV3) NewListRulesOptions() *ListRulesOptions {
	return &ListRulesOptions{}
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

// Or : The `OR` required configurations.
type Or struct {
	// The property.
	Property *string `json:"property,omitempty"`

	// The operator.
	Operator *string `json:"operator,omitempty"`

	// The value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the Or.Operator property.
// The operator.
const (
	Or_Operator_DaysLessThan         = "days_less_than"
	Or_Operator_IpsEquals            = "ips_equals"
	Or_Operator_IpsInRange           = "ips_in_range"
	Or_Operator_IpsNotEquals         = "ips_not_equals"
	Or_Operator_IsEmpty              = "is_empty"
	Or_Operator_IsFalse              = "is_false"
	Or_Operator_IsNotEmpty           = "is_not_empty"
	Or_Operator_IsTrue               = "is_true"
	Or_Operator_NumEquals            = "num_equals"
	Or_Operator_NumGreaterThan       = "num_greater_than"
	Or_Operator_NumGreaterThanEquals = "num_greater_than_equals"
	Or_Operator_NumLessThan          = "num_less_than"
	Or_Operator_NumLessThanEquals    = "num_less_than_equals"
	Or_Operator_NumNotEquals         = "num_not_equals"
	Or_Operator_StringContains       = "string_contains"
	Or_Operator_StringEquals         = "string_equals"
	Or_Operator_StringMatch          = "string_match"
	Or_Operator_StringNotContains    = "string_not_contains"
	Or_Operator_StringNotEquals      = "string_not_equals"
	Or_Operator_StringNotMatch       = "string_not_match"
	Or_Operator_StringsAllowed       = "strings_allowed"
	Or_Operator_StringsInList        = "strings_in_list"
	Or_Operator_StringsRequired      = "strings_required"
)

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
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Parameter : The rule import parameter.
type Parameter struct {
	// The import parameter name.
	Name *string `json:"name,omitempty"`

	// The display name of the property.
	DisplayName *string `json:"display_name,omitempty"`

	// The propery description.
	Description *string `json:"description,omitempty"`

	// The property type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the Parameter.Type property.
// The property type.
const (
	Parameter_Type_Boolean    = "boolean"
	Parameter_Type_General    = "general"
	Parameter_Type_IpList     = "ip_list"
	Parameter_Type_Numeric    = "numeric"
	Parameter_Type_String     = "string"
	Parameter_Type_StringList = "string_list"
	Parameter_Type_Timestamp  = "timestamp"
)

// UnmarshalParameter unmarshals an instance of Parameter from the specified map of raw messages.
func UnmarshalParameter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Parameter)
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

// ReplaceRuleOptions : The ReplaceRule options.
type ReplaceRuleOptions struct {
	// The ID of the corresponding rule.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// Compares a supplied `Etag` value with the version that is stored for the requested resource. If the values match,
	// the server allows the request method to continue.
	//
	// To find the `Etag` value, run a GET request on the resource that you want to modify, and check the response headers.
	IfMatch *string `json:"IfMatch" validate:"required"`

	// Account identification value.
	AccountID *string `json:"account_id" validate:"required"`

	// Rule description.
	Description *string `json:"description" validate:"required"`

	// The rule target.
	Target *Target `json:"target" validate:"required"`

	// The required configurations.
	RequiredConfig *RequiredConfig `json:"required_config" validate:"required"`

	// List of labels corresponding to rule.
	Labels []string `json:"labels" validate:"required"`

	// Rule type (user_defined or system_defined).
	Type *string `json:"type,omitempty"`

	// Rule verison number.
	Version *string `json:"version,omitempty"`

	// The collection of import parameters.
	Import *Import `json:"import,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceRuleOptions.Type property.
// Rule type (user_defined or system_defined).
const (
	ReplaceRuleOptions_Type_SystemDefined = "system_defined"
	ReplaceRuleOptions_Type_UserDefined   = "user_defined"
)

// NewReplaceRuleOptions : Instantiate ReplaceRuleOptions
func (*ConfigManagerV3) NewReplaceRuleOptions(ruleID string, ifMatch string, accountID string, description string, target *Target, requiredConfig *RequiredConfig, labels []string) *ReplaceRuleOptions {
	return &ReplaceRuleOptions{
		RuleID:         core.StringPtr(ruleID),
		IfMatch:        core.StringPtr(ifMatch),
		AccountID:      core.StringPtr(accountID),
		Description:    core.StringPtr(description),
		Target:         target,
		RequiredConfig: requiredConfig,
		Labels:         labels,
	}
}

// SetRuleID : Allow user to set RuleID
func (_options *ReplaceRuleOptions) SetRuleID(ruleID string) *ReplaceRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *ReplaceRuleOptions) SetIfMatch(ifMatch string) *ReplaceRuleOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ReplaceRuleOptions) SetAccountID(accountID string) *ReplaceRuleOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ReplaceRuleOptions) SetDescription(description string) *ReplaceRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *ReplaceRuleOptions) SetTarget(target *Target) *ReplaceRuleOptions {
	_options.Target = target
	return _options
}

// SetRequiredConfig : Allow user to set RequiredConfig
func (_options *ReplaceRuleOptions) SetRequiredConfig(requiredConfig *RequiredConfig) *ReplaceRuleOptions {
	_options.RequiredConfig = requiredConfig
	return _options
}

// SetLabels : Allow user to set Labels
func (_options *ReplaceRuleOptions) SetLabels(labels []string) *ReplaceRuleOptions {
	_options.Labels = labels
	return _options
}

// SetType : Allow user to set Type
func (_options *ReplaceRuleOptions) SetType(typeVar string) *ReplaceRuleOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *ReplaceRuleOptions) SetVersion(version string) *ReplaceRuleOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetImport : Allow user to set Import
func (_options *ReplaceRuleOptions) SetImport(importVar *Import) *ReplaceRuleOptions {
	_options.Import = importVar
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceRuleOptions) SetHeaders(param map[string]string) *ReplaceRuleOptions {
	options.Headers = param
	return options
}

// RequiredConfig : The required configurations.
type RequiredConfig struct {
	// The required config description.
	Description *string `json:"description,omitempty"`

	// The `AND` required configurations.
	And []And `json:"and,omitempty"`

	// The `OR` required configurations.
	Or []Or `json:"or,omitempty"`
}

// UnmarshalRequiredConfig unmarshals an instance of RequiredConfig from the specified map of raw messages.
func UnmarshalRequiredConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RequiredConfig)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
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

// Rule : Rule reponse corresponding to account instance.
type Rule struct {
	// The date when the rule was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The user who created the rule.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The date when the rule was modified.
	UpdatedOn *strfmt.DateTime `json:"updated_on" validate:"required"`

	// The user who modified the rule.
	UpdatedBy *string `json:"updated_by" validate:"required"`

	// The rule ID.
	ID *string `json:"id" validate:"required"`

	// The account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The details of a rule's response.
	Description *string `json:"description" validate:"required"`

	// The rule type (allowable values are `user_defined` or `system_defined`).
	Type *string `json:"type" validate:"required"`

	// The version number of a rule.
	Version *string `json:"version" validate:"required"`

	// The collection of import parameters.
	Import *Import `json:"import,omitempty"`

	// The rule target.
	Target *Target `json:"target" validate:"required"`

	// The required configurations.
	RequiredConfig *RequiredConfig `json:"required_config" validate:"required"`

	// The list of labels.
	Labels []string `json:"labels" validate:"required"`
}

// Constants associated with the Rule.Type property.
// The rule type (allowable values are `user_defined` or `system_defined`).
const (
	Rule_Type_SystemDefined = "system_defined"
	Rule_Type_UserDefined   = "user_defined"
)

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
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
	err = core.UnmarshalModel(m, "import", &obj.Import, UnmarshalImport)
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

// Rules : Collection of rule corresponding to account instance.
type Rules struct {
	// The list of rules.
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

// Target : The rule target.
type Target struct {
	// The target service name.
	ServiceName *string `json:"service_name" validate:"required"`

	// The display name of the target service.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// The target resource kind.
	ResourceKind *string `json:"resource_kind" validate:"required"`

	// The list of targets supported properties.
	AdditionalTargetAttributes []AdditionalTargetAttribute `json:"additional_target_attributes,omitempty"`
}

// NewTarget : Instantiate Target (Generic Model Constructor)
func (*ConfigManagerV3) NewTarget(serviceName string, resourceKind string) (_model *Target, err error) {
	_model = &Target{
		ServiceName:  core.StringPtr(serviceName),
		ResourceKind: core.StringPtr(resourceKind),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
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
	err = core.UnmarshalModel(m, "additional_target_attributes", &obj.AdditionalTargetAttributes, UnmarshalAdditionalTargetAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
