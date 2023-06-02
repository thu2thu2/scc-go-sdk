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

// Package resultsreportsapiv3 : Operations and models for the ResultsReportsApiV3 service
package resultsreportsapiv3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/scc-go-sdk/v4/common"
	"github.com/go-openapi/strfmt"
)

// ResultsReportsApiV3 : Security and Compliance Center Results/Reports API
//
// API Version: 3.0.0
type ResultsReportsApiV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.compliance.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "results_reports_api"

const ParameterizedServiceURL = "https://{environment}.cloud.ibm.com"

var defaultUrlVariables = map[string]string{
	"environment": "us-south.compliance",
}

// ResultsReportsApiV3Options : Service options
type ResultsReportsApiV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewResultsReportsApiV3UsingExternalConfig : constructs an instance of ResultsReportsApiV3 with passed in options and external configuration.
func NewResultsReportsApiV3UsingExternalConfig(options *ResultsReportsApiV3Options) (resultsReportsApi *ResultsReportsApiV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	resultsReportsApi, err = NewResultsReportsApiV3(options)
	if err != nil {
		return
	}

	err = resultsReportsApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = resultsReportsApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewResultsReportsApiV3 : constructs an instance of ResultsReportsApiV3 with passed in options.
func NewResultsReportsApiV3(options *ResultsReportsApiV3Options) (service *ResultsReportsApiV3, err error) {
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

	service = &ResultsReportsApiV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "resultsReportsApi" suitable for processing requests.
func (resultsReportsApi *ResultsReportsApiV3) Clone() *ResultsReportsApiV3 {
	if core.IsNil(resultsReportsApi) {
		return nil
	}
	clone := *resultsReportsApi
	clone.Service = resultsReportsApi.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (resultsReportsApi *ResultsReportsApiV3) SetServiceURL(url string) error {
	return resultsReportsApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (resultsReportsApi *ResultsReportsApiV3) GetServiceURL() string {
	return resultsReportsApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (resultsReportsApi *ResultsReportsApiV3) SetDefaultHeaders(headers http.Header) {
	resultsReportsApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (resultsReportsApi *ResultsReportsApiV3) SetEnableGzipCompression(enableGzip bool) {
	resultsReportsApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (resultsReportsApi *ResultsReportsApiV3) GetEnableGzipCompression() bool {
	return resultsReportsApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (resultsReportsApi *ResultsReportsApiV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	resultsReportsApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (resultsReportsApi *ResultsReportsApiV3) DisableRetries() {
	resultsReportsApi.Service.DisableRetries()
}

// GetLatestReports : Return the latest distinct reports grouped by profile ID, scope ID and attachment ID
// Retrieves the latest distinct reports grouped by profile ID, scope ID and attachment ID.
func (resultsReportsApi *ResultsReportsApiV3) GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions) (result *GetLatestReportsResponse, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetLatestReportsWithContext(context.Background(), getLatestReportsOptions)
}

// GetLatestReportsWithContext is an alternate form of the GetLatestReports method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetLatestReportsWithContext(ctx context.Context, getLatestReportsOptions *GetLatestReportsOptions) (result *GetLatestReportsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLatestReportsOptions, "getLatestReportsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLatestReportsOptions, "getLatestReportsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getLatestReportsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/latest`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLatestReportsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetLatestReports")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLatestReportsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getLatestReportsOptions.XCorrelationID))
	}

	if getLatestReportsOptions.HomeAccountID != nil {
		builder.AddQuery("home_account_id", fmt.Sprint(*getLatestReportsOptions.HomeAccountID))
	}
	if getLatestReportsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*getLatestReportsOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetLatestReportsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListReports : List reports
// Retrieves a page of reports filtered by the specified parameters.
func (resultsReportsApi *ResultsReportsApiV3) ListReports(listReportsOptions *ListReportsOptions) (result *ReportPage, response *core.DetailedResponse, err error) {
	return resultsReportsApi.ListReportsWithContext(context.Background(), listReportsOptions)
}

// ListReportsWithContext is an alternate form of the ListReports method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) ListReportsWithContext(ctx context.Context, listReportsOptions *ListReportsOptions) (result *ReportPage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listReportsOptions, "listReportsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listReportsOptions, "listReportsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listReportsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listReportsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "ListReports")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listReportsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*listReportsOptions.XCorrelationID))
	}

	if listReportsOptions.HomeAccountID != nil {
		builder.AddQuery("home_account_id", fmt.Sprint(*listReportsOptions.HomeAccountID))
	}
	if listReportsOptions.AttachmentID != nil {
		builder.AddQuery("attachment_id", fmt.Sprint(*listReportsOptions.AttachmentID))
	}
	if listReportsOptions.GroupID != nil {
		builder.AddQuery("group_id", fmt.Sprint(*listReportsOptions.GroupID))
	}
	if listReportsOptions.ProfileID != nil {
		builder.AddQuery("profile_id", fmt.Sprint(*listReportsOptions.ProfileID))
	}
	if listReportsOptions.ScopeID != nil {
		builder.AddQuery("scope_id", fmt.Sprint(*listReportsOptions.ScopeID))
	}
	if listReportsOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listReportsOptions.Type))
	}
	if listReportsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listReportsOptions.Start))
	}
	if listReportsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listReportsOptions.Limit))
	}
	if listReportsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listReportsOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReportPage)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportsProfiles : Retrieve a unique list of profiles
// Gets a list of profile groups from the set of reports, without duplicates, to be used as a filtering option.
func (resultsReportsApi *ResultsReportsApiV3) GetReportsProfiles(getReportsProfilesOptions *GetReportsProfilesOptions) (result *GetProfilesResponse, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportsProfilesWithContext(context.Background(), getReportsProfilesOptions)
}

// GetReportsProfilesWithContext is an alternate form of the GetReportsProfiles method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportsProfilesWithContext(ctx context.Context, getReportsProfilesOptions *GetReportsProfilesOptions) (result *GetProfilesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportsProfilesOptions, "getReportsProfilesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportsProfilesOptions, "getReportsProfilesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportsProfilesOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/profiles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportsProfilesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportsProfiles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportsProfilesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportsProfilesOptions.XCorrelationID))
	}

	if getReportsProfilesOptions.HomeAccountID != nil {
		builder.AddQuery("home_account_id", fmt.Sprint(*getReportsProfilesOptions.HomeAccountID))
	}
	if getReportsProfilesOptions.ReportID != nil {
		builder.AddQuery("report_id", fmt.Sprint(*getReportsProfilesOptions.ReportID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetProfilesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportsScopes : Retrieve a unique list of scopes
// Gets a list of scopes from the set of reports, without duplicates, to be used as a filtering option.
func (resultsReportsApi *ResultsReportsApiV3) GetReportsScopes(getReportsScopesOptions *GetReportsScopesOptions) (result *GetScopesResponse, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportsScopesWithContext(context.Background(), getReportsScopesOptions)
}

// GetReportsScopesWithContext is an alternate form of the GetReportsScopes method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportsScopesWithContext(ctx context.Context, getReportsScopesOptions *GetReportsScopesOptions) (result *GetScopesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportsScopesOptions, "getReportsScopesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportsScopesOptions, "getReportsScopesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportsScopesOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/scopes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportsScopesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportsScopes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportsScopesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportsScopesOptions.XCorrelationID))
	}

	if getReportsScopesOptions.HomeAccountID != nil {
		builder.AddQuery("home_account_id", fmt.Sprint(*getReportsScopesOptions.HomeAccountID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetScopesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReport : Retrieve a single report
// Retrieves the report by the specified `report_id`.
func (resultsReportsApi *ResultsReportsApiV3) GetReport(getReportOptions *GetReportOptions) (result *Report, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportWithContext(context.Background(), getReportOptions)
}

// GetReportWithContext is an alternate form of the GetReport method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportWithContext(ctx context.Context, getReportOptions *GetReportOptions) (result *Report, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportOptions, "getReportOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportOptions, "getReportOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportOptions.InstanceID,
		"report_id":   *getReportOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReport)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportSummary : Retrieve a report summary
// Gets the complete summarized information for a single report.
func (resultsReportsApi *ResultsReportsApiV3) GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions) (result *ReportSummary, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportSummaryWithContext(context.Background(), getReportSummaryOptions)
}

// GetReportSummaryWithContext is an alternate form of the GetReportSummary method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportSummaryWithContext(ctx context.Context, getReportSummaryOptions *GetReportSummaryOptions) (result *ReportSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportSummaryOptions, "getReportSummaryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportSummaryOptions, "getReportSummaryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportSummaryOptions.InstanceID,
		"report_id":   *getReportSummaryOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/summary`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportSummaryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportSummary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportSummaryOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportSummaryOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReportSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportEvaluation : Get evaluation details for a given report
// Retrieves evaluation details of a report by the specified `report_id`.
func (resultsReportsApi *ResultsReportsApiV3) GetReportEvaluation(getReportEvaluationOptions *GetReportEvaluationOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportEvaluationWithContext(context.Background(), getReportEvaluationOptions)
}

// GetReportEvaluationWithContext is an alternate form of the GetReportEvaluation method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportEvaluationWithContext(ctx context.Context, getReportEvaluationOptions *GetReportEvaluationOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportEvaluationOptions, "getReportEvaluationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportEvaluationOptions, "getReportEvaluationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportEvaluationOptions.InstanceID,
		"report_id":   *getReportEvaluationOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/download`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportEvaluationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportEvaluation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/csv")
	if getReportEvaluationOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportEvaluationOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resultsReportsApi.Service.Request(request, &result)

	return
}

// GetReportControls : Get controls for a given report
// Gets a sorted and filtered list of controls for the specified report.
func (resultsReportsApi *ResultsReportsApiV3) GetReportControls(getReportControlsOptions *GetReportControlsOptions) (result *GetReportControlsResponse, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportControlsWithContext(context.Background(), getReportControlsOptions)
}

// GetReportControlsWithContext is an alternate form of the GetReportControls method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportControlsWithContext(ctx context.Context, getReportControlsOptions *GetReportControlsOptions) (result *GetReportControlsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportControlsOptions, "getReportControlsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportControlsOptions, "getReportControlsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportControlsOptions.InstanceID,
		"report_id":   *getReportControlsOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/controls`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportControlsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportControls")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportControlsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportControlsOptions.XCorrelationID))
	}

	if getReportControlsOptions.ControlID != nil {
		builder.AddQuery("control_id", fmt.Sprint(*getReportControlsOptions.ControlID))
	}
	if getReportControlsOptions.ControlName != nil {
		builder.AddQuery("control_name", fmt.Sprint(*getReportControlsOptions.ControlName))
	}
	if getReportControlsOptions.ControlDescription != nil {
		builder.AddQuery("control_description", fmt.Sprint(*getReportControlsOptions.ControlDescription))
	}
	if getReportControlsOptions.ControlCategory != nil {
		builder.AddQuery("control_category", fmt.Sprint(*getReportControlsOptions.ControlCategory))
	}
	if getReportControlsOptions.Status != nil {
		builder.AddQuery("status", fmt.Sprint(*getReportControlsOptions.Status))
	}
	if getReportControlsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*getReportControlsOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetReportControlsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportRule : Retrieve a single rule in a report
// Retrieves the rule by the specified `report_id` and `rule_id`.
func (resultsReportsApi *ResultsReportsApiV3) GetReportRule(getReportRuleOptions *GetReportRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportRuleWithContext(context.Background(), getReportRuleOptions)
}

// GetReportRuleWithContext is an alternate form of the GetReportRule method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportRuleWithContext(ctx context.Context, getReportRuleOptions *GetReportRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportRuleOptions, "getReportRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportRuleOptions, "getReportRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportRuleOptions.InstanceID,
		"report_id":   *getReportRuleOptions.ReportID,
		"rule_id":     *getReportRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportRuleOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportRuleOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
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

// ListReportEvaluations : List report evaluations
// Gets a paginated list of evaluations for the specified report.
func (resultsReportsApi *ResultsReportsApiV3) ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) (result *EvaluationPage, response *core.DetailedResponse, err error) {
	return resultsReportsApi.ListReportEvaluationsWithContext(context.Background(), listReportEvaluationsOptions)
}

// ListReportEvaluationsWithContext is an alternate form of the ListReportEvaluations method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) ListReportEvaluationsWithContext(ctx context.Context, listReportEvaluationsOptions *ListReportEvaluationsOptions) (result *EvaluationPage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listReportEvaluationsOptions, "listReportEvaluationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listReportEvaluationsOptions, "listReportEvaluationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listReportEvaluationsOptions.InstanceID,
		"report_id":   *listReportEvaluationsOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/evaluations`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listReportEvaluationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "ListReportEvaluations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listReportEvaluationsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*listReportEvaluationsOptions.XCorrelationID))
	}

	if listReportEvaluationsOptions.AssessmentID != nil {
		builder.AddQuery("assessment_id", fmt.Sprint(*listReportEvaluationsOptions.AssessmentID))
	}
	if listReportEvaluationsOptions.ComponentID != nil {
		builder.AddQuery("component_id", fmt.Sprint(*listReportEvaluationsOptions.ComponentID))
	}
	if listReportEvaluationsOptions.TargetID != nil {
		builder.AddQuery("target_id", fmt.Sprint(*listReportEvaluationsOptions.TargetID))
	}
	if listReportEvaluationsOptions.TargetName != nil {
		builder.AddQuery("target_name", fmt.Sprint(*listReportEvaluationsOptions.TargetName))
	}
	if listReportEvaluationsOptions.Status != nil {
		builder.AddQuery("status", fmt.Sprint(*listReportEvaluationsOptions.Status))
	}
	if listReportEvaluationsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listReportEvaluationsOptions.Start))
	}
	if listReportEvaluationsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listReportEvaluationsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEvaluationPage)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListReportResources : List report resources
// Gets a paginated list of resources for the specified report.
func (resultsReportsApi *ResultsReportsApiV3) ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) (result *ResourcePage, response *core.DetailedResponse, err error) {
	return resultsReportsApi.ListReportResourcesWithContext(context.Background(), listReportResourcesOptions)
}

// ListReportResourcesWithContext is an alternate form of the ListReportResources method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) ListReportResourcesWithContext(ctx context.Context, listReportResourcesOptions *ListReportResourcesOptions) (result *ResourcePage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listReportResourcesOptions, "listReportResourcesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listReportResourcesOptions, "listReportResourcesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listReportResourcesOptions.InstanceID,
		"report_id":   *listReportResourcesOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/resources`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listReportResourcesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "ListReportResources")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listReportResourcesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*listReportResourcesOptions.XCorrelationID))
	}

	if listReportResourcesOptions.ID != nil {
		builder.AddQuery("id", fmt.Sprint(*listReportResourcesOptions.ID))
	}
	if listReportResourcesOptions.ResourceName != nil {
		builder.AddQuery("resource_name", fmt.Sprint(*listReportResourcesOptions.ResourceName))
	}
	if listReportResourcesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listReportResourcesOptions.AccountID))
	}
	if listReportResourcesOptions.ComponentID != nil {
		builder.AddQuery("component_id", fmt.Sprint(*listReportResourcesOptions.ComponentID))
	}
	if listReportResourcesOptions.Status != nil {
		builder.AddQuery("status", fmt.Sprint(*listReportResourcesOptions.Status))
	}
	if listReportResourcesOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listReportResourcesOptions.Start))
	}
	if listReportResourcesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listReportResourcesOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourcePage)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportTags : Get tags associated with the report
// Gets a list of tags for the specified report.
func (resultsReportsApi *ResultsReportsApiV3) GetReportTags(getReportTagsOptions *GetReportTagsOptions) (result *GetTagsResponse, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportTagsWithContext(context.Background(), getReportTagsOptions)
}

// GetReportTagsWithContext is an alternate form of the GetReportTags method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportTagsWithContext(ctx context.Context, getReportTagsOptions *GetReportTagsOptions) (result *GetTagsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportTagsOptions, "getReportTagsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportTagsOptions, "getReportTagsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"report_id": *getReportTagsOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/v3/reports/{report_id}/tags`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportTagsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportTags")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportTagsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportTagsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTagsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetReportViolationsDrift : Get report violations drift
// Gets a list of report violation data points for the specified report and timeframe.
func (resultsReportsApi *ResultsReportsApiV3) GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions) (result *GetReportViolationsDriftResult, response *core.DetailedResponse, err error) {
	return resultsReportsApi.GetReportViolationsDriftWithContext(context.Background(), getReportViolationsDriftOptions)
}

// GetReportViolationsDriftWithContext is an alternate form of the GetReportViolationsDrift method which supports a Context parameter
func (resultsReportsApi *ResultsReportsApiV3) GetReportViolationsDriftWithContext(ctx context.Context, getReportViolationsDriftOptions *GetReportViolationsDriftOptions) (result *GetReportViolationsDriftResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReportViolationsDriftOptions, "getReportViolationsDriftOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReportViolationsDriftOptions, "getReportViolationsDriftOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getReportViolationsDriftOptions.InstanceID,
		"report_id":   *getReportViolationsDriftOptions.ReportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resultsReportsApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resultsReportsApi.Service.Options.URL, `/instances/{instance_id}/v3/reports/{report_id}/violations_drift`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReportViolationsDriftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("results_reports_api", "V3", "GetReportViolationsDrift")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getReportViolationsDriftOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-Id", fmt.Sprint(*getReportViolationsDriftOptions.XCorrelationID))
	}

	if getReportViolationsDriftOptions.ScanTimeDuration != nil {
		builder.AddQuery("scan_time_duration", fmt.Sprint(*getReportViolationsDriftOptions.ScanTimeDuration))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resultsReportsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetReportViolationsDriftResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Account : The account associated to a report.
type Account struct {
	// The account ID.
	ID *string `json:"id,omitempty"`

	// The account name.
	Name *string `json:"name,omitempty"`

	// The account type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalAccount unmarshals an instance of Account from the specified map of raw messages.
func UnmarshalAccount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Account)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// Assessment : A control specification assessment.
type Assessment struct {
	// The assessment ID.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// The assessment type.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// The assessment method.
	AssessmentMethod *string `json:"assessment_method,omitempty"`

	// The assessment description.
	AssessmentDescription *string `json:"assessment_description,omitempty"`

	// The number of parameters of this assessment.
	ParameterCount *int64 `json:"parameter_count,omitempty"`

	// The list of parameters of this assessment.
	Parameters []Parameter `json:"parameters,omitempty"`
}

// UnmarshalAssessment unmarshals an instance of Assessment from the specified map of raw messages.
func UnmarshalAssessment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Assessment)
	err = core.UnmarshalPrimitive(m, "assessment_id", &obj.AssessmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_type", &obj.AssessmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_method", &obj.AssessmentMethod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_description", &obj.AssessmentDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_count", &obj.ParameterCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalParameter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Attachment : The attachment associated to a report.
type Attachment struct {
	// The attachment ID.
	ID *string `json:"id,omitempty"`
}

// UnmarshalAttachment unmarshals an instance of Attachment from the specified map of raw messages.
func UnmarshalAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attachment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ComplianceScore : A compliance score.
type ComplianceScore struct {
	// The number of passed evaluations.
	Passed *int64 `json:"passed,omitempty"`

	// The total number of evaluations.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The percent of passed evaluations out of the total.
	Percent *int64 `json:"percent,omitempty"`
}

// UnmarshalComplianceScore unmarshals an instance of ComplianceScore from the specified map of raw messages.
func UnmarshalComplianceScore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComplianceScore)
	err = core.UnmarshalPrimitive(m, "passed", &obj.Passed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "percent", &obj.Percent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ComplianceStats : Compliance stats.
type ComplianceStats struct {
	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of checks.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of compliant checks.
	CompliantCount *int64 `json:"compliant_count,omitempty"`

	// The number of not compliant checks.
	NotCompliantCount *int64 `json:"not_compliant_count,omitempty"`

	// The number of checks unable to perform.
	UnableToPerformCount *int64 `json:"unable_to_perform_count,omitempty"`

	// The number of checks requiring a user evaluation.
	UserEvaluationRequiredCount *int64 `json:"user_evaluation_required_count,omitempty"`
}

// Constants associated with the ComplianceStats.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	ComplianceStats_Status_Compliant              = "compliant"
	ComplianceStats_Status_NotCompliant           = "not_compliant"
	ComplianceStats_Status_UnableToPerform        = "unable_to_perform"
	ComplianceStats_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalComplianceStats unmarshals an instance of ComplianceStats from the specified map of raw messages.
func UnmarshalComplianceStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComplianceStats)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "compliant_count", &obj.CompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_compliant_count", &obj.NotCompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unable_to_perform_count", &obj.UnableToPerformCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_evaluation_required_count", &obj.UserEvaluationRequiredCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlSpecificationWithStats : A control specification with compliance stats.
type ControlSpecificationWithStats struct {
	// The control specification ID.
	ID *string `json:"id,omitempty"`

	// The component ID.
	ComponentID *string `json:"component_id,omitempty"`

	// The component description.
	Description *string `json:"description,omitempty"`

	// The environment.
	Environment *string `json:"environment,omitempty"`

	// The control specification responsibility.
	Responsibility *string `json:"responsibility,omitempty"`

	// The list of assessments.
	Assessments []Assessment `json:"assessments,omitempty"`

	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of checks.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of compliant checks.
	CompliantCount *int64 `json:"compliant_count,omitempty"`

	// The number of not compliant checks.
	NotCompliantCount *int64 `json:"not_compliant_count,omitempty"`

	// The number of checks unable to perform.
	UnableToPerformCount *int64 `json:"unable_to_perform_count,omitempty"`

	// The number of checks requiring a user evaluation.
	UserEvaluationRequiredCount *int64 `json:"user_evaluation_required_count,omitempty"`
}

// Constants associated with the ControlSpecificationWithStats.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	ControlSpecificationWithStats_Status_Compliant              = "compliant"
	ControlSpecificationWithStats_Status_NotCompliant           = "not_compliant"
	ControlSpecificationWithStats_Status_UnableToPerform        = "unable_to_perform"
	ControlSpecificationWithStats_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalControlSpecificationWithStats unmarshals an instance of ControlSpecificationWithStats from the specified map of raw messages.
func UnmarshalControlSpecificationWithStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlSpecificationWithStats)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "component_id", &obj.ComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "responsibility", &obj.Responsibility)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "assessments", &obj.Assessments, UnmarshalAssessment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "compliant_count", &obj.CompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_compliant_count", &obj.NotCompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unable_to_perform_count", &obj.UnableToPerformCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_evaluation_required_count", &obj.UserEvaluationRequiredCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlWithStats : A control with compliance stats.
type ControlWithStats struct {
	// The control ID.
	ID *string `json:"id,omitempty"`

	// The control library ID.
	ControlLibraryID *string `json:"control_library_id,omitempty"`

	// The control library version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The control name.
	ControlName *string `json:"control_name,omitempty"`

	// The control description.
	ControlDescription *string `json:"control_description,omitempty"`

	// The control category.
	ControlCategory *string `json:"control_category,omitempty"`

	// The control path.
	ControlPath *string `json:"control_path,omitempty"`

	// The list of specifications in this page.
	ControlSpecifications []ControlSpecificationWithStats `json:"control_specifications,omitempty"`

	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of checks.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of compliant checks.
	CompliantCount *int64 `json:"compliant_count,omitempty"`

	// The number of not compliant checks.
	NotCompliantCount *int64 `json:"not_compliant_count,omitempty"`

	// The number of checks unable to perform.
	UnableToPerformCount *int64 `json:"unable_to_perform_count,omitempty"`

	// The number of checks requiring a user evaluation.
	UserEvaluationRequiredCount *int64 `json:"user_evaluation_required_count,omitempty"`
}

// Constants associated with the ControlWithStats.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	ControlWithStats_Status_Compliant              = "compliant"
	ControlWithStats_Status_NotCompliant           = "not_compliant"
	ControlWithStats_Status_UnableToPerform        = "unable_to_perform"
	ControlWithStats_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalControlWithStats unmarshals an instance of ControlWithStats from the specified map of raw messages.
func UnmarshalControlWithStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlWithStats)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_id", &obj.ControlLibraryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_version", &obj.ControlLibraryVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_name", &obj.ControlName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_description", &obj.ControlDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_category", &obj.ControlCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_path", &obj.ControlPath)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_specifications", &obj.ControlSpecifications, UnmarshalControlSpecificationWithStats)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "compliant_count", &obj.CompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_compliant_count", &obj.NotCompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unable_to_perform_count", &obj.UnableToPerformCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_evaluation_required_count", &obj.UserEvaluationRequiredCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvalDetails : Evaluation details.
type EvalDetails struct {
	// The evaluation properties.
	Properties []Property `json:"properties,omitempty"`
}

// UnmarshalEvalDetails unmarshals an instance of EvalDetails from the specified map of raw messages.
func UnmarshalEvalDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvalDetails)
	err = core.UnmarshalModel(m, "properties", &obj.Properties, UnmarshalProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvalStats : Evaluation stats.
type EvalStats struct {
	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of evaluations.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of passed evaluations.
	PassCount *int64 `json:"pass_count,omitempty"`

	// The number of failed evaluations.
	FailureCount *int64 `json:"failure_count,omitempty"`

	// The number of evaluations that ended with errors (started but not finished).
	ErrorCount *int64 `json:"error_count,omitempty"`

	// The number of completed evaluations (passed and failed).
	CompletedCount *int64 `json:"completed_count,omitempty"`
}

// Constants associated with the EvalStats.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	EvalStats_Status_Compliant              = "compliant"
	EvalStats_Status_NotCompliant           = "not_compliant"
	EvalStats_Status_UnableToPerform        = "unable_to_perform"
	EvalStats_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalEvalStats unmarshals an instance of EvalStats from the specified map of raw messages.
func UnmarshalEvalStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvalStats)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pass_count", &obj.PassCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failure_count", &obj.FailureCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_count", &obj.ErrorCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_count", &obj.CompletedCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Evaluation : A control specification assessment evaluation.
type Evaluation struct {
	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The ID of the report associated to this evaluation.
	ReportID *string `json:"report_id,omitempty"`

	// The control ID.
	ControlID *string `json:"control_id,omitempty"`

	// The component ID.
	ComponentID *string `json:"component_id,omitempty"`

	// A control specification assessment.
	Assessment *Assessment `json:"assessment,omitempty"`

	// The time the evaluation was made.
	EvaluateTime *string `json:"evaluate_time,omitempty"`

	// An evaluation target.
	Target *Target `json:"target,omitempty"`

	// The allowed values of an evaluation status.
	Status *string `json:"status,omitempty"`

	// The reason for the evaluation failure.
	Reason *string `json:"reason,omitempty"`

	// Evaluation details.
	Details *EvalDetails `json:"details,omitempty"`
}

// Constants associated with the Evaluation.Status property.
// The allowed values of an evaluation status.
const (
	Evaluation_Status_Error   = "error"
	Evaluation_Status_Failure = "failure"
	Evaluation_Status_Pass    = "pass"
	Evaluation_Status_Skipped = "skipped"
)

// UnmarshalEvaluation unmarshals an instance of Evaluation from the specified map of raw messages.
func UnmarshalEvaluation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Evaluation)
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_id", &obj.ControlID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "component_id", &obj.ComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "assessment", &obj.Assessment, UnmarshalAssessment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "evaluate_time", &obj.EvaluateTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "details", &obj.Details, UnmarshalEvalDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvaluationPage : A page of assessment evaluations.
type EvaluationPage struct {
	// The total number of resources in the collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The requested page limit.
	Limit *int64 `json:"limit" validate:"required"`

	// The token of the next page when present.
	Start *string `json:"start,omitempty"`

	// A page reference.
	First *PageHRef `json:"first" validate:"required"`

	// A page reference.
	Next *PageHRef `json:"next,omitempty"`

	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The list of evaluations in this page.
	Evaluations []Evaluation `json:"evaluations,omitempty"`
}

// UnmarshalEvaluationPage unmarshals an instance of EvaluationPage from the specified map of raw messages.
func UnmarshalEvaluationPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvaluationPage)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageHRef)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageHRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evaluations", &obj.Evaluations, UnmarshalEvaluation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *EvaluationPage) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// GetLatestReportsOptions : The GetLatestReports options.
type GetLatestReportsOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The ID of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// Sorts results by using a valid sort field. To learn more, see
	// [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLatestReportsOptions : Instantiate GetLatestReportsOptions
func (*ResultsReportsApiV3) NewGetLatestReportsOptions(instanceID string) *GetLatestReportsOptions {
	return &GetLatestReportsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetLatestReportsOptions) SetInstanceID(instanceID string) *GetLatestReportsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetLatestReportsOptions) SetXCorrelationID(xCorrelationID string) *GetLatestReportsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHomeAccountID : Allow user to set HomeAccountID
func (_options *GetLatestReportsOptions) SetHomeAccountID(homeAccountID string) *GetLatestReportsOptions {
	_options.HomeAccountID = core.StringPtr(homeAccountID)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *GetLatestReportsOptions) SetSort(sort string) *GetLatestReportsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLatestReportsOptions) SetHeaders(param map[string]string) *GetLatestReportsOptions {
	options.Headers = param
	return options
}

// GetLatestReportsResponse : The response body of the `get_latest_reports` operation.
type GetLatestReportsResponse struct {
	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// Compliance stats.
	ControlsSummary *ComplianceStats `json:"controls_summary,omitempty"`

	// Evaluation stats.
	EvaluationsSummary *EvalStats `json:"evaluations_summary,omitempty"`

	// A compliance score.
	Score *ComplianceScore `json:"score,omitempty"`

	// A list of reports.
	Reports []Report `json:"reports,omitempty"`
}

// UnmarshalGetLatestReportsResponse unmarshals an instance of GetLatestReportsResponse from the specified map of raw messages.
func UnmarshalGetLatestReportsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetLatestReportsResponse)
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls_summary", &obj.ControlsSummary, UnmarshalComplianceStats)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evaluations_summary", &obj.EvaluationsSummary, UnmarshalEvalStats)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "score", &obj.Score, UnmarshalComplianceScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reports", &obj.Reports, UnmarshalReport)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetProfilesResponse : The response body of the `get_profiles` operation.
type GetProfilesResponse struct {
	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// A list of profile groups.
	Profiles []Profile `json:"profiles,omitempty"`
}

// UnmarshalGetProfilesResponse unmarshals an instance of GetProfilesResponse from the specified map of raw messages.
func UnmarshalGetProfilesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetProfilesResponse)
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "profiles", &obj.Profiles, UnmarshalProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetReportControlsOptions : The GetReportControls options.
type GetReportControlsOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The ID of the control.
	ControlID *string `json:"control_id,omitempty"`

	// The name of the control.
	ControlName *string `json:"control_name,omitempty"`

	// The description of the control.
	ControlDescription *string `json:"control_description,omitempty"`

	// A control category value.
	ControlCategory *string `json:"control_category,omitempty"`

	// A compliance status value.
	Status *string `json:"status,omitempty"`

	// Sorts results by using a valid sort field. To learn more, see
	// [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
	Sort *string `json:"sort,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetReportControlsOptions.Status property.
// A compliance status value.
const (
	GetReportControlsOptions_Status_Compliant              = "compliant"
	GetReportControlsOptions_Status_NotCompliant           = "not_compliant"
	GetReportControlsOptions_Status_UnableToPerform        = "unable_to_perform"
	GetReportControlsOptions_Status_UserEvaluationRequired = "user_evaluation_required"
)

// NewGetReportControlsOptions : Instantiate GetReportControlsOptions
func (*ResultsReportsApiV3) NewGetReportControlsOptions(instanceID string, reportID string) *GetReportControlsOptions {
	return &GetReportControlsOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportControlsOptions) SetInstanceID(instanceID string) *GetReportControlsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportControlsOptions) SetReportID(reportID string) *GetReportControlsOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetControlID : Allow user to set ControlID
func (_options *GetReportControlsOptions) SetControlID(controlID string) *GetReportControlsOptions {
	_options.ControlID = core.StringPtr(controlID)
	return _options
}

// SetControlName : Allow user to set ControlName
func (_options *GetReportControlsOptions) SetControlName(controlName string) *GetReportControlsOptions {
	_options.ControlName = core.StringPtr(controlName)
	return _options
}

// SetControlDescription : Allow user to set ControlDescription
func (_options *GetReportControlsOptions) SetControlDescription(controlDescription string) *GetReportControlsOptions {
	_options.ControlDescription = core.StringPtr(controlDescription)
	return _options
}

// SetControlCategory : Allow user to set ControlCategory
func (_options *GetReportControlsOptions) SetControlCategory(controlCategory string) *GetReportControlsOptions {
	_options.ControlCategory = core.StringPtr(controlCategory)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *GetReportControlsOptions) SetStatus(status string) *GetReportControlsOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *GetReportControlsOptions) SetSort(sort string) *GetReportControlsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportControlsOptions) SetXCorrelationID(xCorrelationID string) *GetReportControlsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportControlsOptions) SetHeaders(param map[string]string) *GetReportControlsOptions {
	options.Headers = param
	return options
}

// GetReportControlsResponse : A list of controls.
type GetReportControlsResponse struct {
	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of checks.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of compliant checks.
	CompliantCount *int64 `json:"compliant_count,omitempty"`

	// The number of not compliant checks.
	NotCompliantCount *int64 `json:"not_compliant_count,omitempty"`

	// The number of checks unable to perform.
	UnableToPerformCount *int64 `json:"unable_to_perform_count,omitempty"`

	// The number of checks requiring a user evaluation.
	UserEvaluationRequiredCount *int64 `json:"user_evaluation_required_count,omitempty"`

	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The list of controls in the report.
	Controls []ControlWithStats `json:"controls,omitempty"`
}

// Constants associated with the GetReportControlsResponse.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	GetReportControlsResponse_Status_Compliant              = "compliant"
	GetReportControlsResponse_Status_NotCompliant           = "not_compliant"
	GetReportControlsResponse_Status_UnableToPerform        = "unable_to_perform"
	GetReportControlsResponse_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalGetReportControlsResponse unmarshals an instance of GetReportControlsResponse from the specified map of raw messages.
func UnmarshalGetReportControlsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetReportControlsResponse)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "compliant_count", &obj.CompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_compliant_count", &obj.NotCompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unable_to_perform_count", &obj.UnableToPerformCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_evaluation_required_count", &obj.UserEvaluationRequiredCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalControlWithStats)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetReportEvaluationOptions : The GetReportEvaluation options.
type GetReportEvaluationOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportEvaluationOptions : Instantiate GetReportEvaluationOptions
func (*ResultsReportsApiV3) NewGetReportEvaluationOptions(instanceID string, reportID string) *GetReportEvaluationOptions {
	return &GetReportEvaluationOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportEvaluationOptions) SetInstanceID(instanceID string) *GetReportEvaluationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportEvaluationOptions) SetReportID(reportID string) *GetReportEvaluationOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportEvaluationOptions) SetXCorrelationID(xCorrelationID string) *GetReportEvaluationOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportEvaluationOptions) SetHeaders(param map[string]string) *GetReportEvaluationOptions {
	options.Headers = param
	return options
}

// GetReportOptions : The GetReport options.
type GetReportOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportOptions : Instantiate GetReportOptions
func (*ResultsReportsApiV3) NewGetReportOptions(instanceID string, reportID string) *GetReportOptions {
	return &GetReportOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportOptions) SetInstanceID(instanceID string) *GetReportOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportOptions) SetReportID(reportID string) *GetReportOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportOptions) SetXCorrelationID(xCorrelationID string) *GetReportOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportOptions) SetHeaders(param map[string]string) *GetReportOptions {
	options.Headers = param
	return options
}

// GetReportRuleOptions : The GetReportRule options.
type GetReportRuleOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The ID of a rule in a report.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportRuleOptions : Instantiate GetReportRuleOptions
func (*ResultsReportsApiV3) NewGetReportRuleOptions(instanceID string, reportID string, ruleID string) *GetReportRuleOptions {
	return &GetReportRuleOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
		RuleID:     core.StringPtr(ruleID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportRuleOptions) SetInstanceID(instanceID string) *GetReportRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportRuleOptions) SetReportID(reportID string) *GetReportRuleOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *GetReportRuleOptions) SetRuleID(ruleID string) *GetReportRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportRuleOptions) SetXCorrelationID(xCorrelationID string) *GetReportRuleOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportRuleOptions) SetHeaders(param map[string]string) *GetReportRuleOptions {
	options.Headers = param
	return options
}

// GetReportSummaryOptions : The GetReportSummary options.
type GetReportSummaryOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportSummaryOptions : Instantiate GetReportSummaryOptions
func (*ResultsReportsApiV3) NewGetReportSummaryOptions(instanceID string, reportID string) *GetReportSummaryOptions {
	return &GetReportSummaryOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportSummaryOptions) SetInstanceID(instanceID string) *GetReportSummaryOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportSummaryOptions) SetReportID(reportID string) *GetReportSummaryOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportSummaryOptions) SetXCorrelationID(xCorrelationID string) *GetReportSummaryOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportSummaryOptions) SetHeaders(param map[string]string) *GetReportSummaryOptions {
	options.Headers = param
	return options
}

// GetReportTagsOptions : The GetReportTags options.
type GetReportTagsOptions struct {
	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportTagsOptions : Instantiate GetReportTagsOptions
func (*ResultsReportsApiV3) NewGetReportTagsOptions(reportID string) *GetReportTagsOptions {
	return &GetReportTagsOptions{
		ReportID: core.StringPtr(reportID),
	}
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportTagsOptions) SetReportID(reportID string) *GetReportTagsOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportTagsOptions) SetXCorrelationID(xCorrelationID string) *GetReportTagsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportTagsOptions) SetHeaders(param map[string]string) *GetReportTagsOptions {
	options.Headers = param
	return options
}

// GetReportViolationsDriftOptions : The GetReportViolationsDrift options.
type GetReportViolationsDriftOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The duration of the `scan_time` timestamp in number of days.
	ScanTimeDuration *int64 `json:"scan_time_duration,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportViolationsDriftOptions : Instantiate GetReportViolationsDriftOptions
func (*ResultsReportsApiV3) NewGetReportViolationsDriftOptions(instanceID string, reportID string) *GetReportViolationsDriftOptions {
	return &GetReportViolationsDriftOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportViolationsDriftOptions) SetInstanceID(instanceID string) *GetReportViolationsDriftOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportViolationsDriftOptions) SetReportID(reportID string) *GetReportViolationsDriftOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetScanTimeDuration : Allow user to set ScanTimeDuration
func (_options *GetReportViolationsDriftOptions) SetScanTimeDuration(scanTimeDuration int64) *GetReportViolationsDriftOptions {
	_options.ScanTimeDuration = core.Int64Ptr(scanTimeDuration)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportViolationsDriftOptions) SetXCorrelationID(xCorrelationID string) *GetReportViolationsDriftOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportViolationsDriftOptions) SetHeaders(param map[string]string) *GetReportViolationsDriftOptions {
	options.Headers = param
	return options
}

// GetReportViolationsDriftResult : The response body of the `get_report_violations_drift` operation.
type GetReportViolationsDriftResult struct {
	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// A list of report violations data points.
	DataPoints []ReportViolationDataPoint `json:"data_points,omitempty"`
}

// UnmarshalGetReportViolationsDriftResult unmarshals an instance of GetReportViolationsDriftResult from the specified map of raw messages.
func UnmarshalGetReportViolationsDriftResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetReportViolationsDriftResult)
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_points", &obj.DataPoints, UnmarshalReportViolationDataPoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetReportsProfilesOptions : The GetReportsProfiles options.
type GetReportsProfilesOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The ID of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The ID of the report.
	ReportID *string `json:"report_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportsProfilesOptions : Instantiate GetReportsProfilesOptions
func (*ResultsReportsApiV3) NewGetReportsProfilesOptions(instanceID string) *GetReportsProfilesOptions {
	return &GetReportsProfilesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportsProfilesOptions) SetInstanceID(instanceID string) *GetReportsProfilesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportsProfilesOptions) SetXCorrelationID(xCorrelationID string) *GetReportsProfilesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHomeAccountID : Allow user to set HomeAccountID
func (_options *GetReportsProfilesOptions) SetHomeAccountID(homeAccountID string) *GetReportsProfilesOptions {
	_options.HomeAccountID = core.StringPtr(homeAccountID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *GetReportsProfilesOptions) SetReportID(reportID string) *GetReportsProfilesOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportsProfilesOptions) SetHeaders(param map[string]string) *GetReportsProfilesOptions {
	options.Headers = param
	return options
}

// GetReportsScopesOptions : The GetReportsScopes options.
type GetReportsScopesOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The ID of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReportsScopesOptions : Instantiate GetReportsScopesOptions
func (*ResultsReportsApiV3) NewGetReportsScopesOptions(instanceID string) *GetReportsScopesOptions {
	return &GetReportsScopesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetReportsScopesOptions) SetInstanceID(instanceID string) *GetReportsScopesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetReportsScopesOptions) SetXCorrelationID(xCorrelationID string) *GetReportsScopesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHomeAccountID : Allow user to set HomeAccountID
func (_options *GetReportsScopesOptions) SetHomeAccountID(homeAccountID string) *GetReportsScopesOptions {
	_options.HomeAccountID = core.StringPtr(homeAccountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReportsScopesOptions) SetHeaders(param map[string]string) *GetReportsScopesOptions {
	options.Headers = param
	return options
}

// GetScopesResponse : The response body of the `get_scopes` operation.
type GetScopesResponse struct {
	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// A list of scopes.
	Scopes []Scope `json:"scopes,omitempty"`
}

// UnmarshalGetScopesResponse unmarshals an instance of GetScopesResponse from the specified map of raw messages.
func UnmarshalGetScopesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetScopesResponse)
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scopes", &obj.Scopes, UnmarshalScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTagsResponse : The response body of the `get_tags` operation.
type GetTagsResponse struct {
	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The resource tags.
	Tags *Tags `json:"tags,omitempty"`
}

// UnmarshalGetTagsResponse unmarshals an instance of GetTagsResponse from the specified map of raw messages.
func UnmarshalGetTagsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTagsResponse)
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalTags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListReportEvaluationsOptions : The ListReportEvaluations options.
type ListReportEvaluationsOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The ID of the assessment.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// The ID of component.
	ComponentID *string `json:"component_id,omitempty"`

	// The ID of the evaluation target.
	TargetID *string `json:"target_id,omitempty"`

	// The name of the evaluation target.
	TargetName *string `json:"target_name,omitempty"`

	// An evaluation status value.
	Status *string `json:"status,omitempty"`

	// Determine what resource to start the page on or after.
	Start *string `json:"start,omitempty"`

	// How many resources to return, unless the offset and limit are such that the response is  the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListReportEvaluationsOptions.Status property.
// An evaluation status value.
const (
	ListReportEvaluationsOptions_Status_Error   = "error"
	ListReportEvaluationsOptions_Status_Failure = "failure"
	ListReportEvaluationsOptions_Status_Pass    = "pass"
	ListReportEvaluationsOptions_Status_Skipped = "skipped"
)

// NewListReportEvaluationsOptions : Instantiate ListReportEvaluationsOptions
func (*ResultsReportsApiV3) NewListReportEvaluationsOptions(instanceID string, reportID string) *ListReportEvaluationsOptions {
	return &ListReportEvaluationsOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListReportEvaluationsOptions) SetInstanceID(instanceID string) *ListReportEvaluationsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *ListReportEvaluationsOptions) SetReportID(reportID string) *ListReportEvaluationsOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetAssessmentID : Allow user to set AssessmentID
func (_options *ListReportEvaluationsOptions) SetAssessmentID(assessmentID string) *ListReportEvaluationsOptions {
	_options.AssessmentID = core.StringPtr(assessmentID)
	return _options
}

// SetComponentID : Allow user to set ComponentID
func (_options *ListReportEvaluationsOptions) SetComponentID(componentID string) *ListReportEvaluationsOptions {
	_options.ComponentID = core.StringPtr(componentID)
	return _options
}

// SetTargetID : Allow user to set TargetID
func (_options *ListReportEvaluationsOptions) SetTargetID(targetID string) *ListReportEvaluationsOptions {
	_options.TargetID = core.StringPtr(targetID)
	return _options
}

// SetTargetName : Allow user to set TargetName
func (_options *ListReportEvaluationsOptions) SetTargetName(targetName string) *ListReportEvaluationsOptions {
	_options.TargetName = core.StringPtr(targetName)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ListReportEvaluationsOptions) SetStatus(status string) *ListReportEvaluationsOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListReportEvaluationsOptions) SetStart(start string) *ListReportEvaluationsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListReportEvaluationsOptions) SetLimit(limit int64) *ListReportEvaluationsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListReportEvaluationsOptions) SetXCorrelationID(xCorrelationID string) *ListReportEvaluationsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListReportEvaluationsOptions) SetHeaders(param map[string]string) *ListReportEvaluationsOptions {
	options.Headers = param
	return options
}

// ListReportResourcesOptions : The ListReportResources options.
type ListReportResourcesOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the scan associated to a report.
	ReportID *string `json:"report_id" validate:"required,ne="`

	// The ID of the resource in the `list_report_resources` operation.
	ID *string `json:"id,omitempty"`

	// The name of the resource.
	ResourceName *string `json:"resource_name,omitempty"`

	// The ID of the account owning a resource.
	AccountID *string `json:"account_id,omitempty"`

	// The ID of component.
	ComponentID *string `json:"component_id,omitempty"`

	// A compliance status value.
	Status *string `json:"status,omitempty"`

	// Determine what resource to start the page on or after.
	Start *string `json:"start,omitempty"`

	// How many resources to return, unless the offset and limit are such that the response is  the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListReportResourcesOptions.Status property.
// A compliance status value.
const (
	ListReportResourcesOptions_Status_Compliant              = "compliant"
	ListReportResourcesOptions_Status_NotCompliant           = "not_compliant"
	ListReportResourcesOptions_Status_UnableToPerform        = "unable_to_perform"
	ListReportResourcesOptions_Status_UserEvaluationRequired = "user_evaluation_required"
)

// NewListReportResourcesOptions : Instantiate ListReportResourcesOptions
func (*ResultsReportsApiV3) NewListReportResourcesOptions(instanceID string, reportID string) *ListReportResourcesOptions {
	return &ListReportResourcesOptions{
		InstanceID: core.StringPtr(instanceID),
		ReportID:   core.StringPtr(reportID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListReportResourcesOptions) SetInstanceID(instanceID string) *ListReportResourcesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetReportID : Allow user to set ReportID
func (_options *ListReportResourcesOptions) SetReportID(reportID string) *ListReportResourcesOptions {
	_options.ReportID = core.StringPtr(reportID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ListReportResourcesOptions) SetID(id string) *ListReportResourcesOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetResourceName : Allow user to set ResourceName
func (_options *ListReportResourcesOptions) SetResourceName(resourceName string) *ListReportResourcesOptions {
	_options.ResourceName = core.StringPtr(resourceName)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ListReportResourcesOptions) SetAccountID(accountID string) *ListReportResourcesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetComponentID : Allow user to set ComponentID
func (_options *ListReportResourcesOptions) SetComponentID(componentID string) *ListReportResourcesOptions {
	_options.ComponentID = core.StringPtr(componentID)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ListReportResourcesOptions) SetStatus(status string) *ListReportResourcesOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListReportResourcesOptions) SetStart(start string) *ListReportResourcesOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListReportResourcesOptions) SetLimit(limit int64) *ListReportResourcesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListReportResourcesOptions) SetXCorrelationID(xCorrelationID string) *ListReportResourcesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListReportResourcesOptions) SetHeaders(param map[string]string) *ListReportResourcesOptions {
	options.Headers = param
	return options
}

// ListReportsOptions : The ListReports options.
type ListReportsOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this headers is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-Id,omitempty"`

	// The ID of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The ID of the attachment.
	AttachmentID *string `json:"attachment_id,omitempty"`

	// The report group id.
	GroupID *string `json:"group_id,omitempty"`

	// The ID of the profile.
	ProfileID *string `json:"profile_id,omitempty"`

	// The ID of the scope.
	ScopeID *string `json:"scope_id,omitempty"`

	// The type of the scan.
	Type *string `json:"type,omitempty"`

	// Determine what resource to start the page on or after.
	Start *string `json:"start,omitempty"`

	// How many resources to return, unless the offset and limit are such that the response is  the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// Sorts results by using a valid sort field. To learn more, see
	// [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListReportsOptions.Type property.
// The type of the scan.
const (
	ListReportsOptions_Type_Ondemand  = "ondemand"
	ListReportsOptions_Type_Scheduled = "scheduled"
)

// NewListReportsOptions : Instantiate ListReportsOptions
func (*ResultsReportsApiV3) NewListReportsOptions(instanceID string) *ListReportsOptions {
	return &ListReportsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListReportsOptions) SetInstanceID(instanceID string) *ListReportsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListReportsOptions) SetXCorrelationID(xCorrelationID string) *ListReportsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHomeAccountID : Allow user to set HomeAccountID
func (_options *ListReportsOptions) SetHomeAccountID(homeAccountID string) *ListReportsOptions {
	_options.HomeAccountID = core.StringPtr(homeAccountID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ListReportsOptions) SetAttachmentID(attachmentID string) *ListReportsOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetGroupID : Allow user to set GroupID
func (_options *ListReportsOptions) SetGroupID(groupID string) *ListReportsOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetProfileID : Allow user to set ProfileID
func (_options *ListReportsOptions) SetProfileID(profileID string) *ListReportsOptions {
	_options.ProfileID = core.StringPtr(profileID)
	return _options
}

// SetScopeID : Allow user to set ScopeID
func (_options *ListReportsOptions) SetScopeID(scopeID string) *ListReportsOptions {
	_options.ScopeID = core.StringPtr(scopeID)
	return _options
}

// SetType : Allow user to set Type
func (_options *ListReportsOptions) SetType(typeVar string) *ListReportsOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListReportsOptions) SetStart(start string) *ListReportsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListReportsOptions) SetLimit(limit int64) *ListReportsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListReportsOptions) SetSort(sort string) *ListReportsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListReportsOptions) SetHeaders(param map[string]string) *ListReportsOptions {
	options.Headers = param
	return options
}

// PageHRef : A page reference.
type PageHRef struct {
	// A URL for the first and next page.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPageHRef unmarshals an instance of PageHRef from the specified map of raw messages.
func UnmarshalPageHRef(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PageHRef)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Parameter : A parameter.
type Parameter struct {
	// The parameter name.
	ParameterName *string `json:"parameter_name,omitempty"`

	// The parameter display name.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// The parameter type.
	ParameterType *string `json:"parameter_type,omitempty"`

	// A property value.
	ParameterValue interface{} `json:"parameter_value,omitempty"`
}

// UnmarshalParameter unmarshals an instance of Parameter from the specified map of raw messages.
func UnmarshalParameter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Parameter)
	err = core.UnmarshalPrimitive(m, "parameter_name", &obj.ParameterName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_display_name", &obj.ParameterDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_type", &obj.ParameterType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_value", &obj.ParameterValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Profile : A profile.
type Profile struct {
	// The profile ID.
	ID *string `json:"id,omitempty"`

	// The profile name.
	Name *string `json:"name,omitempty"`

	// The profile version.
	Version *string `json:"version,omitempty"`
}

// UnmarshalProfile unmarshals an instance of Profile from the specified map of raw messages.
func UnmarshalProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Profile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Property : A property.
type Property struct {
	// The property name.
	Property *string `json:"property,omitempty"`

	// The property description.
	PropertyDescription *string `json:"property_description,omitempty"`

	// The property operator.
	Operator *string `json:"operator,omitempty"`

	// A property value.
	ExpectedValue interface{} `json:"expected_value,omitempty"`

	// A property value.
	FoundValue interface{} `json:"found_value,omitempty"`
}

// UnmarshalProperty unmarshals an instance of Property from the specified map of raw messages.
func UnmarshalProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Property)
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property_description", &obj.PropertyDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_value", &obj.ExpectedValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "found_value", &obj.FoundValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Report : A report.
type Report struct {
	// The ID of this report.
	ID *string `json:"id,omitempty"`

	// The group ID (combined profile, scope, and attachment IDs) associated to this report.
	GroupID *string `json:"group_id,omitempty"`

	// When the report was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// When the scan was run.
	ScanTime *strfmt.DateTime `json:"scan_time,omitempty"`

	// The type of the scan.
	Type *string `json:"type,omitempty"`

	// The COS object associated to this report.
	CosObject *string `json:"cos_object,omitempty"`

	// The account associated to a report.
	Account *Account `json:"account,omitempty"`

	// A profile.
	Profile *Profile `json:"profile,omitempty"`

	// A scope.
	Scope *Scope `json:"scope,omitempty"`

	// The attachment associated to a report.
	Attachment *Attachment `json:"attachment,omitempty"`
}

// UnmarshalReport unmarshals an instance of Report from the specified map of raw messages.
func UnmarshalReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Report)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scan_time", &obj.ScanTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cos_object", &obj.CosObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account", &obj.Account, UnmarshalAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "profile", &obj.Profile, UnmarshalProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scope", &obj.Scope, UnmarshalScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment", &obj.Attachment, UnmarshalAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReportPage : A page of reports.
type ReportPage struct {
	// The total number of resources in the collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The requested page limit.
	Limit *int64 `json:"limit" validate:"required"`

	// The token of the next page when present.
	Start *string `json:"start,omitempty"`

	// A page reference.
	First *PageHRef `json:"first" validate:"required"`

	// A page reference.
	Next *PageHRef `json:"next,omitempty"`

	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The list of reports in this page.
	Reports []Report `json:"reports,omitempty"`
}

// UnmarshalReportPage unmarshals an instance of ReportPage from the specified map of raw messages.
func UnmarshalReportPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReportPage)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageHRef)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageHRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reports", &obj.Reports, UnmarshalReport)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ReportPage) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// ReportSummary : A report summary.
type ReportSummary struct {
	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The account associated to a report.
	Account *Account `json:"account,omitempty"`

	// A compliance score.
	Score *ComplianceScore `json:"score,omitempty"`

	// Compliance stats.
	Controls *ComplianceStats `json:"controls,omitempty"`

	// Evaluation stats.
	Evaluations *EvalStats `json:"evaluations,omitempty"`

	// A resource summary.
	Resources *ResourceSummary `json:"resources,omitempty"`
}

// UnmarshalReportSummary unmarshals an instance of ReportSummary from the specified map of raw messages.
func UnmarshalReportSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReportSummary)
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account", &obj.Account, UnmarshalAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "score", &obj.Score, UnmarshalComplianceScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalComplianceStats)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evaluations", &obj.Evaluations, UnmarshalEvalStats)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReportViolationDataPoint : A report violation data point.
type ReportViolationDataPoint struct {
	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The group ID (combined profile, scope, and attachment IDs) associated to this report.
	ReportGroupID *string `json:"report_group_id,omitempty"`

	// When the scan was run.
	ScanTime *strfmt.DateTime `json:"scan_time,omitempty"`

	// Compliance stats.
	Controls *ComplianceStats `json:"controls,omitempty"`
}

// UnmarshalReportViolationDataPoint unmarshals an instance of ReportViolationDataPoint from the specified map of raw messages.
func UnmarshalReportViolationDataPoint(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReportViolationDataPoint)
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_group_id", &obj.ReportGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scan_time", &obj.ScanTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalComplianceStats)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resource : A resource.
type Resource struct {
	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The resource CRN.
	ID *string `json:"id,omitempty"`

	// The resource name.
	ResourceName *string `json:"resource_name,omitempty"`

	// The id of the component.
	ComponentID *string `json:"component_id,omitempty"`

	// The environment.
	Environment *string `json:"environment,omitempty"`

	// The account associated to a report.
	Account *Account `json:"account,omitempty"`

	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of evaluations.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of passed evaluations.
	PassCount *int64 `json:"pass_count,omitempty"`

	// The number of failed evaluations.
	FailureCount *int64 `json:"failure_count,omitempty"`

	// The number of evaluations that ended with errors (started but not finished).
	ErrorCount *int64 `json:"error_count,omitempty"`

	// The number of completed evaluations (passed and failed).
	CompletedCount *int64 `json:"completed_count,omitempty"`
}

// Constants associated with the Resource.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	Resource_Status_Compliant              = "compliant"
	Resource_Status_NotCompliant           = "not_compliant"
	Resource_Status_UnableToPerform        = "unable_to_perform"
	Resource_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "component_id", &obj.ComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account", &obj.Account, UnmarshalAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pass_count", &obj.PassCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failure_count", &obj.FailureCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_count", &obj.ErrorCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_count", &obj.CompletedCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourcePage : A page of resource evaluation summaries.
type ResourcePage struct {
	// The total number of resources in the collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The requested page limit.
	Limit *int64 `json:"limit" validate:"required"`

	// The token of the next page when present.
	Start *string `json:"start,omitempty"`

	// A page reference.
	First *PageHRef `json:"first" validate:"required"`

	// A page reference.
	Next *PageHRef `json:"next,omitempty"`

	// The id of the home account.
	HomeAccountID *string `json:"home_account_id,omitempty"`

	// The id of the report.
	ReportID *string `json:"report_id,omitempty"`

	// The list of resource evaluation summaries in this page.
	Resources []Resource `json:"resources,omitempty"`
}

// UnmarshalResourcePage unmarshals an instance of ResourcePage from the specified map of raw messages.
func UnmarshalResourcePage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourcePage)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageHRef)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageHRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "home_account_id", &obj.HomeAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ResourcePage) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// ResourceSummary : A resource summary.
type ResourceSummary struct {
	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of checks.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of compliant checks.
	CompliantCount *int64 `json:"compliant_count,omitempty"`

	// The number of not compliant checks.
	NotCompliantCount *int64 `json:"not_compliant_count,omitempty"`

	// The number of checks unable to perform.
	UnableToPerformCount *int64 `json:"unable_to_perform_count,omitempty"`

	// The number of checks requiring a user evaluation.
	UserEvaluationRequiredCount *int64 `json:"user_evaluation_required_count,omitempty"`

	// The top 10 resources with the most failures.
	TopFailed []ResourceSummaryItem `json:"top_failed,omitempty"`
}

// Constants associated with the ResourceSummary.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	ResourceSummary_Status_Compliant              = "compliant"
	ResourceSummary_Status_NotCompliant           = "not_compliant"
	ResourceSummary_Status_UnableToPerform        = "unable_to_perform"
	ResourceSummary_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalResourceSummary unmarshals an instance of ResourceSummary from the specified map of raw messages.
func UnmarshalResourceSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceSummary)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "compliant_count", &obj.CompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_compliant_count", &obj.NotCompliantCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unable_to_perform_count", &obj.UnableToPerformCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_evaluation_required_count", &obj.UserEvaluationRequiredCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "top_failed", &obj.TopFailed, UnmarshalResourceSummaryItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceSummaryItem : A resource summary item.
type ResourceSummaryItem struct {
	// The resource name.
	Name *string `json:"name,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty"`

	// The service managing the resource.
	Service *string `json:"service,omitempty"`

	// The resource tags.
	Tags *Tags `json:"tags,omitempty"`

	// The account owning the resource.
	Account *string `json:"account,omitempty"`

	// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	Status *string `json:"status,omitempty"`

	// The total number of evaluations.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of passed evaluations.
	PassCount *int64 `json:"pass_count,omitempty"`

	// The number of failed evaluations.
	FailureCount *int64 `json:"failure_count,omitempty"`

	// The number of evaluations that ended with errors (started but not finished).
	ErrorCount *int64 `json:"error_count,omitempty"`

	// The number of completed evaluations (passed and failed).
	CompletedCount *int64 `json:"completed_count,omitempty"`
}

// Constants associated with the ResourceSummaryItem.Status property.
// The allowed values of an aggregated status for controls, specifications, assessments, and resources.
const (
	ResourceSummaryItem_Status_Compliant              = "compliant"
	ResourceSummaryItem_Status_NotCompliant           = "not_compliant"
	ResourceSummaryItem_Status_UnableToPerform        = "unable_to_perform"
	ResourceSummaryItem_Status_UserEvaluationRequired = "user_evaluation_required"
)

// UnmarshalResourceSummaryItem unmarshals an instance of ResourceSummaryItem from the specified map of raw messages.
func UnmarshalResourceSummaryItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceSummaryItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service", &obj.Service)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalTags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account", &obj.Account)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pass_count", &obj.PassCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failure_count", &obj.FailureCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_count", &obj.ErrorCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_count", &obj.CompletedCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : A rule.
type Rule struct {
	// The rule ID.
	ID *string `json:"id,omitempty"`

	// The rule type.
	Type *string `json:"type,omitempty"`

	// The rule description.
	Description *string `json:"description,omitempty"`

	// The rule version.
	Version *string `json:"version,omitempty"`

	// The rule account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The data the rule was created.
	CreationDate *strfmt.DateTime `json:"creation_date,omitempty"`

	// The ID of the user that created this rule.
	CreatedBy *string `json:"created_by,omitempty"`

	// The data the rule was modified.
	ModificationDate *strfmt.DateTime `json:"modification_date,omitempty"`

	// The ID of the user that modified this rule.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// The rule labels.
	Labels []string `json:"labels,omitempty"`
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
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
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Scope : A scope.
type Scope struct {
	// The scope ID.
	ID *string `json:"id,omitempty"`

	// The scope type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalScope unmarshals an instance of Scope from the specified map of raw messages.
func UnmarshalScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scope)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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

// Tags : The resource tags.
type Tags struct {
	// The user tags.
	User []string `json:"user,omitempty"`

	// The access tags.
	Access []string `json:"access,omitempty"`

	// The service tags.
	Service []string `json:"service,omitempty"`
}

// UnmarshalTags unmarshals an instance of Tags from the specified map of raw messages.
func UnmarshalTags(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tags)
	err = core.UnmarshalPrimitive(m, "user", &obj.User)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "access", &obj.Access)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service", &obj.Service)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Target : An evaluation target.
type Target struct {
	// The target ID.
	ID *string `json:"id,omitempty"`

	// The target account id.
	AccountID *string `json:"account_id,omitempty"`

	// The target resource CRN.
	ResourceCrn *string `json:"resource_crn,omitempty"`

	// The target resource name.
	ResourceName *string `json:"resource_name,omitempty"`

	// The target service name.
	ServiceName *string `json:"service_name,omitempty"`
}

// UnmarshalTarget unmarshals an instance of Target from the specified map of raw messages.
func UnmarshalTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Target)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReportsPager can be used to simplify the use of the "ListReports" method.
type ReportsPager struct {
	hasNext     bool
	options     *ListReportsOptions
	client      *ResultsReportsApiV3
	pageContext struct {
		next *string
	}
}

// NewReportsPager returns a new ReportsPager instance.
func (resultsReportsApi *ResultsReportsApiV3) NewReportsPager(options *ListReportsOptions) (pager *ReportsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListReportsOptions = *options
	pager = &ReportsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  resultsReportsApi,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ReportsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ReportsPager) GetNextWithContext(ctx context.Context) (page []Report, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListReportsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Reports

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ReportsPager) GetAllWithContext(ctx context.Context) (allItems []Report, err error) {
	for pager.HasNext() {
		var nextPage []Report
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ReportsPager) GetNext() (page []Report, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ReportsPager) GetAll() (allItems []Report, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ReportEvaluationsPager can be used to simplify the use of the "ListReportEvaluations" method.
type ReportEvaluationsPager struct {
	hasNext     bool
	options     *ListReportEvaluationsOptions
	client      *ResultsReportsApiV3
	pageContext struct {
		next *string
	}
}

// NewReportEvaluationsPager returns a new ReportEvaluationsPager instance.
func (resultsReportsApi *ResultsReportsApiV3) NewReportEvaluationsPager(options *ListReportEvaluationsOptions) (pager *ReportEvaluationsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListReportEvaluationsOptions = *options
	pager = &ReportEvaluationsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  resultsReportsApi,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ReportEvaluationsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ReportEvaluationsPager) GetNextWithContext(ctx context.Context) (page []Evaluation, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListReportEvaluationsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Evaluations

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ReportEvaluationsPager) GetAllWithContext(ctx context.Context) (allItems []Evaluation, err error) {
	for pager.HasNext() {
		var nextPage []Evaluation
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ReportEvaluationsPager) GetNext() (page []Evaluation, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ReportEvaluationsPager) GetAll() (allItems []Evaluation, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ReportResourcesPager can be used to simplify the use of the "ListReportResources" method.
type ReportResourcesPager struct {
	hasNext     bool
	options     *ListReportResourcesOptions
	client      *ResultsReportsApiV3
	pageContext struct {
		next *string
	}
}

// NewReportResourcesPager returns a new ReportResourcesPager instance.
func (resultsReportsApi *ResultsReportsApiV3) NewReportResourcesPager(options *ListReportResourcesOptions) (pager *ReportResourcesPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListReportResourcesOptions = *options
	pager = &ReportResourcesPager{
		hasNext: true,
		options: &optionsCopy,
		client:  resultsReportsApi,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ReportResourcesPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ReportResourcesPager) GetNextWithContext(ctx context.Context) (page []Resource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListReportResourcesWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ReportResourcesPager) GetAllWithContext(ctx context.Context) (allItems []Resource, err error) {
	for pager.HasNext() {
		var nextPage []Resource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ReportResourcesPager) GetNext() (page []Resource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ReportResourcesPager) GetAll() (allItems []Resource, err error) {
	return pager.GetAllWithContext(context.Background())
}
