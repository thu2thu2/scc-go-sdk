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

package resultsv3_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/resultsv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ResultsV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(resultsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(resultsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
				URL: "https://resultsv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(resultsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESULTS_URL":       "https://resultsv3/api",
				"RESULTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resultsService, serviceErr := resultsv3.NewResultsV3UsingExternalConfig(&resultsv3.ResultsV3Options{})
				Expect(resultsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := resultsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resultsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resultsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resultsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resultsService, serviceErr := resultsv3.NewResultsV3UsingExternalConfig(&resultsv3.ResultsV3Options{
					URL: "https://testService/api",
				})
				Expect(resultsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resultsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := resultsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resultsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resultsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resultsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resultsService, serviceErr := resultsv3.NewResultsV3UsingExternalConfig(&resultsv3.ResultsV3Options{})
				err := resultsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resultsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := resultsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resultsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resultsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resultsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESULTS_URL":       "https://resultsv3/api",
				"RESULTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			resultsService, serviceErr := resultsv3.NewResultsV3UsingExternalConfig(&resultsv3.ResultsV3Options{})

			It(`Instantiate service client with error`, func() {
				Expect(resultsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESULTS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			resultsService, serviceErr := resultsv3.NewResultsV3UsingExternalConfig(&resultsv3.ResultsV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(resultsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = resultsv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := resultsv3.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := resultsv3.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions) - Operation response error`, func() {
		getLatestReportsPath := "/reports/latest"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestReportsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLatestReports with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(resultsv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("testString")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions)`, func() {
		getLatestReportsPath := "/reports/latest"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "controls_summary": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations_summary": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_at": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "scope": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c"}}]}`)
				}))
			})
			It(`Invoke GetLatestReports successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(resultsv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("testString")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetLatestReportsWithContext(ctx, getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetLatestReportsWithContext(ctx, getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "controls_summary": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations_summary": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_at": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "scope": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c"}}]}`)
				}))
			})
			It(`Invoke GetLatestReports successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetLatestReports(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(resultsv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("testString")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLatestReports with error: Operation request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(resultsv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("testString")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetLatestReports successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(resultsv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("testString")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReports(listReportsOptions *ListReportsOptions) - Operation response error`, func() {
		listReportsPath := "/reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["attachment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["scope_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"scheduled"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReports with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(resultsv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.ScopeID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("testString")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.ListReports(listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.ListReports(listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReports(listReportsOptions *ListReportsOptions)`, func() {
		listReportsPath := "/reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["attachment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["scope_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"scheduled"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_at": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "scope": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c"}}]}`)
				}))
			})
			It(`Invoke ListReports successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(resultsv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.ScopeID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("testString")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.ListReportsWithContext(ctx, listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.ListReports(listReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.ListReportsWithContext(ctx, listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["attachment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["scope_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"scheduled"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_at": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "scope": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c"}}]}`)
				}))
			})
			It(`Invoke ListReports successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.ListReports(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(resultsv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.ScopeID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("testString")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.ListReports(listReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReports with error: Operation request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(resultsv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.ScopeID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("testString")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.ListReports(listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListReports successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(resultsv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.ScopeID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("testString")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.ListReports(listReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resultsv3.ReportPage)
				nextObject := new(resultsv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(resultsv3.ReportPage)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(resultsv3.ReportPage)
				nextObject := new(resultsv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"reports":[{"id":"44a5-a292-32114fa73558","group_id":"55b6-b3A4-432250b84669","created_at":"2022-08-15T12:30:01Z","scan_time":"2022-08-15T12:30:01Z","type":"scheduled","cos_object":"crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7","instance_id":"84644a08-31b6-4988-b504-49a46ca69ccd","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"profile":{"id":"44a5-a292-32114fa73558","name":"IBM FS Cloud","version":"0.1"},"scope":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","type":"account"},"attachment":{"id":"531fc3e28bfc43c5a2cea07786d93f5c"}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"reports":[{"id":"44a5-a292-32114fa73558","group_id":"55b6-b3A4-432250b84669","created_at":"2022-08-15T12:30:01Z","scan_time":"2022-08-15T12:30:01Z","type":"scheduled","cos_object":"crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7","instance_id":"84644a08-31b6-4988-b504-49a46ca69ccd","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"profile":{"id":"44a5-a292-32114fa73558","name":"IBM FS Cloud","version":"0.1"},"scope":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","type":"account"},"attachment":{"id":"531fc3e28bfc43c5a2cea07786d93f5c"}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReportsPager.GetNext successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				listReportsOptionsModel := &resultsv3.ListReportsOptions{
					XCorrelationID: core.StringPtr("testString"),
					AttachmentID:   core.StringPtr("testString"),
					GroupID:        core.StringPtr("testString"),
					ProfileID:      core.StringPtr("testString"),
					ScopeID:        core.StringPtr("testString"),
					Type:           core.StringPtr("scheduled"),
					Limit:          core.Int64Ptr(int64(10)),
					Sort:           core.StringPtr("testString"),
				}

				pager, err := resultsService.NewReportsPager(listReportsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resultsv3.Report
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReportsPager.GetAll successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				listReportsOptionsModel := &resultsv3.ListReportsOptions{
					XCorrelationID: core.StringPtr("testString"),
					AttachmentID:   core.StringPtr("testString"),
					GroupID:        core.StringPtr("testString"),
					ProfileID:      core.StringPtr("testString"),
					ScopeID:        core.StringPtr("testString"),
					Type:           core.StringPtr("scheduled"),
					Limit:          core.Int64Ptr(int64(10)),
					Sort:           core.StringPtr("testString"),
				}

				pager, err := resultsService.NewReportsPager(listReportsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetReportsProfiles(getReportsProfilesOptions *GetReportsProfilesOptions) - Operation response error`, func() {
		getReportsProfilesPath := "/reports/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportsProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["report_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportsProfiles with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportsProfilesOptions model
				getReportsProfilesOptionsModel := new(resultsv3.GetReportsProfilesOptions)
				getReportsProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.ReportID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportsProfiles(getReportsProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportsProfiles(getReportsProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportsProfiles(getReportsProfilesOptions *GetReportsProfilesOptions)`, func() {
		getReportsProfilesPath := "/reports/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportsProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["report_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "profiles": [{"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}]}`)
				}))
			})
			It(`Invoke GetReportsProfiles successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportsProfilesOptions model
				getReportsProfilesOptionsModel := new(resultsv3.GetReportsProfilesOptions)
				getReportsProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.ReportID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportsProfilesWithContext(ctx, getReportsProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportsProfiles(getReportsProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportsProfilesWithContext(ctx, getReportsProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportsProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["report_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "profiles": [{"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}]}`)
				}))
			})
			It(`Invoke GetReportsProfiles successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportsProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportsProfilesOptions model
				getReportsProfilesOptionsModel := new(resultsv3.GetReportsProfilesOptions)
				getReportsProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.ReportID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportsProfiles(getReportsProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportsProfiles with error: Operation request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportsProfilesOptions model
				getReportsProfilesOptionsModel := new(resultsv3.GetReportsProfilesOptions)
				getReportsProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.ReportID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportsProfiles(getReportsProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportsProfiles successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportsProfilesOptions model
				getReportsProfilesOptionsModel := new(resultsv3.GetReportsProfilesOptions)
				getReportsProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.ReportID = core.StringPtr("testString")
				getReportsProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportsProfiles(getReportsProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportsScopes(getReportsScopesOptions *GetReportsScopesOptions) - Operation response error`, func() {
		getReportsScopesPath := "/reports/scopes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportsScopesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportsScopes with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportsScopesOptions model
				getReportsScopesOptionsModel := new(resultsv3.GetReportsScopesOptions)
				getReportsScopesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportsScopes(getReportsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportsScopes(getReportsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportsScopes(getReportsScopesOptions *GetReportsScopesOptions)`, func() {
		getReportsScopesPath := "/reports/scopes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportsScopesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "scopes": [{"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}]}`)
				}))
			})
			It(`Invoke GetReportsScopes successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportsScopesOptions model
				getReportsScopesOptionsModel := new(resultsv3.GetReportsScopesOptions)
				getReportsScopesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportsScopesWithContext(ctx, getReportsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportsScopes(getReportsScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportsScopesWithContext(ctx, getReportsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportsScopesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "scopes": [{"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}]}`)
				}))
			})
			It(`Invoke GetReportsScopes successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportsScopes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportsScopesOptions model
				getReportsScopesOptionsModel := new(resultsv3.GetReportsScopesOptions)
				getReportsScopesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportsScopes(getReportsScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportsScopes with error: Operation request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportsScopesOptions model
				getReportsScopesOptionsModel := new(resultsv3.GetReportsScopesOptions)
				getReportsScopesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportsScopes(getReportsScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportsScopes successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportsScopesOptions model
				getReportsScopesOptionsModel := new(resultsv3.GetReportsScopesOptions)
				getReportsScopesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportsScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportsScopes(getReportsScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReport(getReportOptions *GetReportOptions) - Operation response error`, func() {
		getReportPath := "/reports/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReport with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(resultsv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReport(getReportOptions *GetReportOptions)`, func() {
		getReportPath := "/reports/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_at": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "scope": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c"}}`)
				}))
			})
			It(`Invoke GetReport successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(resultsv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportWithContext(ctx, getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportWithContext(ctx, getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_at": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "scope": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "type": "account"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c"}}`)
				}))
			})
			It(`Invoke GetReport successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(resultsv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReport with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(resultsv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportOptions model with no property values
				getReportOptionsModelNew := new(resultsv3.GetReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReport(getReportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReport successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(resultsv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions) - Operation response error`, func() {
		getReportSummaryPath := "/reports/testString/summary"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportSummaryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportSummary with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(resultsv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions)`, func() {
		getReportSummaryPath := "/reports/testString/summary"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "isntance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "resources": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "top_failed": [{"name": "my-bucket", "id": "531fc3e28bfc43c5a2cea07786d93f5c", "service": "cloud-object-storage", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}, "account": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}}`)
				}))
			})
			It(`Invoke GetReportSummary successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(resultsv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportSummaryWithContext(ctx, getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportSummaryWithContext(ctx, getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "isntance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "resources": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "top_failed": [{"name": "my-bucket", "id": "531fc3e28bfc43c5a2cea07786d93f5c", "service": "cloud-object-storage", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}, "account": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}}`)
				}))
			})
			It(`Invoke GetReportSummary successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(resultsv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportSummary with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(resultsv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportSummaryOptions model with no property values
				getReportSummaryOptionsModelNew := new(resultsv3.GetReportSummaryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReportSummary(getReportSummaryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportSummary successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(resultsv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportEvaluation(getReportEvaluationOptions *GetReportEvaluationOptions)`, func() {
		getReportEvaluationPath := "/reports/testString/download"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportEvaluationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for exclude_summary query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/csv")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetReportEvaluation successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(resultsv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportEvaluationWithContext(ctx, getReportEvaluationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportEvaluationWithContext(ctx, getReportEvaluationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportEvaluationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for exclude_summary query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/csv")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetReportEvaluation successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportEvaluation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(resultsv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportEvaluation with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(resultsv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportEvaluationOptions model with no property values
				getReportEvaluationOptionsModelNew := new(resultsv3.GetReportEvaluationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReportEvaluation(getReportEvaluationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportEvaluation successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(resultsv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportControls(getReportControlsOptions *GetReportControlsOptions) - Operation response error`, func() {
		getReportControlsPath := "/reports/testString/controls"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportControlsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["control_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_description"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_category"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportControls with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(resultsv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportControls(getReportControlsOptions *GetReportControlsOptions)`, func() {
		getReportControlsPath := "/reports/testString/controls"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["control_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_description"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_category"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "home_account_id": "HomeAccountID", "report_id": "ReportID", "controls": [{"id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_version": "v1.2.3", "control_name": "Password Management", "control_description": "Password Management", "control_category": "Access Control", "control_path": "AC-2(a)", "control_specifications": [{"id": "18d32a4430e54c81a6668952609763b2", "component_id": "cloud-object_storage", "description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "environment": "ibm cloud", "responsibility": "user", "assessments": [{"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}]}`)
				}))
			})
			It(`Invoke GetReportControls successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(resultsv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportControlsWithContext(ctx, getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportControlsWithContext(ctx, getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["control_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_description"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_category"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "home_account_id": "HomeAccountID", "report_id": "ReportID", "controls": [{"id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_version": "v1.2.3", "control_name": "Password Management", "control_description": "Password Management", "control_category": "Access Control", "control_path": "AC-2(a)", "control_specifications": [{"id": "18d32a4430e54c81a6668952609763b2", "component_id": "cloud-object_storage", "description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "environment": "ibm cloud", "responsibility": "user", "assessments": [{"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}]}`)
				}))
			})
			It(`Invoke GetReportControls successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportControls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(resultsv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportControls with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(resultsv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportControlsOptions model with no property values
				getReportControlsOptionsModelNew := new(resultsv3.GetReportControlsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReportControls(getReportControlsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportControls successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(resultsv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportRule(getReportRuleOptions *GetReportRuleOptions) - Operation response error`, func() {
		getReportRulePath := "/reports/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportRule with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(resultsv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("testString")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportRule(getReportRuleOptions *GetReportRuleOptions)`, func() {
		getReportRulePath := "/reports/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "rule-7b0560a4-df94-4629-bb76-680f3155ddda", "type": "user_defined/system_defined", "description": "rule", "version": "1.2.3", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "created_at": "2022-08-15T12:30:01Z", "created_by": "IBMid-12345", "updated_at": "2022-08-15T12:30:01Z", "updated_by": "IBMid-12345", "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetReportRule successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(resultsv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("testString")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportRuleWithContext(ctx, getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportRuleWithContext(ctx, getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "rule-7b0560a4-df94-4629-bb76-680f3155ddda", "type": "user_defined/system_defined", "description": "rule", "version": "1.2.3", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "created_at": "2022-08-15T12:30:01Z", "created_by": "IBMid-12345", "updated_at": "2022-08-15T12:30:01Z", "updated_by": "IBMid-12345", "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetReportRule successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(resultsv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("testString")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportRule with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(resultsv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("testString")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportRuleOptions model with no property values
				getReportRuleOptionsModelNew := new(resultsv3.GetReportRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReportRule(getReportRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportRule successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(resultsv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("testString")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) - Operation response error`, func() {
		listReportEvaluationsPath := "/reports/testString/evaluations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["assessment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"failure"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReportEvaluations with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(resultsv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions)`, func() {
		listReportEvaluationsPath := "/reports/testString/evaluations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["assessment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"failure"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "evaluations": [{"home_account_id": "be200c80cabc456e91139e4152327456", "report_id": "44a5-a292-32114fa73558", "control_id": "28016c95-b389-447f-8a05-eabe1ad7fd24", "component_id": "cloud-object_storage", "assessment": {"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}, "evaluate_time": "2022-06-30T11:03:44.630150782Z", "target": {"id": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "resource_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "resource_name": "mybucket", "service_name": "cloud-object-storage"}, "status": "failure", "reason": "One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met", "details": {"properties": [{"property": "allowed_network", "property_description": "A description for this property", "operator": "string_equals", "expected_value": "anyValue", "found_value": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke ListReportEvaluations successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(resultsv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.ListReportEvaluationsWithContext(ctx, listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.ListReportEvaluationsWithContext(ctx, listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["assessment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"failure"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "evaluations": [{"home_account_id": "be200c80cabc456e91139e4152327456", "report_id": "44a5-a292-32114fa73558", "control_id": "28016c95-b389-447f-8a05-eabe1ad7fd24", "component_id": "cloud-object_storage", "assessment": {"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}, "evaluate_time": "2022-06-30T11:03:44.630150782Z", "target": {"id": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "resource_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "resource_name": "mybucket", "service_name": "cloud-object-storage"}, "status": "failure", "reason": "One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met", "details": {"properties": [{"property": "allowed_network", "property_description": "A description for this property", "operator": "string_equals", "expected_value": "anyValue", "found_value": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke ListReportEvaluations successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.ListReportEvaluations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(resultsv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReportEvaluations with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(resultsv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListReportEvaluationsOptions model with no property values
				listReportEvaluationsOptionsModelNew := new(resultsv3.ListReportEvaluationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.ListReportEvaluations(listReportEvaluationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListReportEvaluations successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(resultsv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resultsv3.EvaluationPage)
				nextObject := new(resultsv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(resultsv3.EvaluationPage)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(resultsv3.EvaluationPage)
				nextObject := new(resultsv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"evaluations":[{"home_account_id":"be200c80cabc456e91139e4152327456","report_id":"44a5-a292-32114fa73558","control_id":"28016c95-b389-447f-8a05-eabe1ad7fd24","component_id":"cloud-object_storage","assessment":{"assessment_id":"382c2b06-e6b2-43ee-b189-c1c7743b67ee","assessment_type":"ibm-cloud-rule","assessment_method":"ibm-cloud-rule","assessment_description":"Check whether Cloud Object Storage is accessible only by using private endpoints","parameter_count":1,"parameters":[{"parameter_name":"location","parameter_display_name":"Location","parameter_type":"string","parameter_value":"anyValue"}]},"evaluate_time":"2022-06-30T11:03:44.630150782Z","target":{"id":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","account_id":"59bcbfa6ea2f006b4ed7094c1a08dcdd","resource_crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","resource_name":"mybucket","service_name":"cloud-object-storage"},"status":"failure","reason":"One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met","details":{"properties":[{"property":"allowed_network","property_description":"A description for this property","operator":"string_equals","expected_value":"anyValue","found_value":"anyValue"}]}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"evaluations":[{"home_account_id":"be200c80cabc456e91139e4152327456","report_id":"44a5-a292-32114fa73558","control_id":"28016c95-b389-447f-8a05-eabe1ad7fd24","component_id":"cloud-object_storage","assessment":{"assessment_id":"382c2b06-e6b2-43ee-b189-c1c7743b67ee","assessment_type":"ibm-cloud-rule","assessment_method":"ibm-cloud-rule","assessment_description":"Check whether Cloud Object Storage is accessible only by using private endpoints","parameter_count":1,"parameters":[{"parameter_name":"location","parameter_display_name":"Location","parameter_type":"string","parameter_value":"anyValue"}]},"evaluate_time":"2022-06-30T11:03:44.630150782Z","target":{"id":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","account_id":"59bcbfa6ea2f006b4ed7094c1a08dcdd","resource_crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","resource_name":"mybucket","service_name":"cloud-object-storage"},"status":"failure","reason":"One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met","details":{"properties":[{"property":"allowed_network","property_description":"A description for this property","operator":"string_equals","expected_value":"anyValue","found_value":"anyValue"}]}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReportEvaluationsPager.GetNext successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				listReportEvaluationsOptionsModel := &resultsv3.ListReportEvaluationsOptions{
					ReportID:       core.StringPtr("testString"),
					AssessmentID:   core.StringPtr("testString"),
					ComponentID:    core.StringPtr("testString"),
					TargetID:       core.StringPtr("testString"),
					TargetName:     core.StringPtr("testString"),
					Status:         core.StringPtr("failure"),
					Limit:          core.Int64Ptr(int64(10)),
					XCorrelationID: core.StringPtr("testString"),
				}

				pager, err := resultsService.NewReportEvaluationsPager(listReportEvaluationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resultsv3.Evaluation
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReportEvaluationsPager.GetAll successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				listReportEvaluationsOptionsModel := &resultsv3.ListReportEvaluationsOptions{
					ReportID:       core.StringPtr("testString"),
					AssessmentID:   core.StringPtr("testString"),
					ComponentID:    core.StringPtr("testString"),
					TargetID:       core.StringPtr("testString"),
					TargetName:     core.StringPtr("testString"),
					Status:         core.StringPtr("failure"),
					Limit:          core.Int64Ptr(int64(10)),
					XCorrelationID: core.StringPtr("testString"),
				}

				pager, err := resultsService.NewReportEvaluationsPager(listReportEvaluationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) - Operation response error`, func() {
		listReportResourcesPath := "/reports/testString/resources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReportResources with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(resultsv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions)`, func() {
		listReportResourcesPath := "/reports/testString/resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "resources": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "id": "crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::", "resource_name": "jeff's key", "component_id": "cloud-object_storage", "environment": "ibm cloud", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}`)
				}))
			})
			It(`Invoke ListReportResources successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(resultsv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.ListReportResourcesWithContext(ctx, listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.ListReportResourcesWithContext(ctx, listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "resources": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "id": "crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::", "resource_name": "jeff's key", "component_id": "cloud-object_storage", "environment": "ibm cloud", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}`)
				}))
			})
			It(`Invoke ListReportResources successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.ListReportResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(resultsv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReportResources with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(resultsv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListReportResourcesOptions model with no property values
				listReportResourcesOptionsModelNew := new(resultsv3.ListReportResourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.ListReportResources(listReportResourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListReportResources successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(resultsv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resultsv3.ResourcePage)
				nextObject := new(resultsv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(resultsv3.ResourcePage)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(resultsv3.ResourcePage)
				nextObject := new(resultsv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1,"resources":[{"report_id":"30b434b3-cb08-4845-af10-7a8fc682b6a8","id":"crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::","resource_name":"jeff's key","component_id":"cloud-object_storage","environment":"ibm cloud","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"status":"compliant","total_count":140,"pass_count":123,"failure_count":12,"error_count":5,"completed_count":135}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"report_id":"30b434b3-cb08-4845-af10-7a8fc682b6a8","id":"crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::","resource_name":"jeff's key","component_id":"cloud-object_storage","environment":"ibm cloud","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"status":"compliant","total_count":140,"pass_count":123,"failure_count":12,"error_count":5,"completed_count":135}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReportResourcesPager.GetNext successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				listReportResourcesOptionsModel := &resultsv3.ListReportResourcesOptions{
					ReportID:       core.StringPtr("testString"),
					ID:             core.StringPtr("testString"),
					ResourceName:   core.StringPtr("testString"),
					AccountID:      core.StringPtr("testString"),
					ComponentID:    core.StringPtr("testString"),
					Status:         core.StringPtr("compliant"),
					Limit:          core.Int64Ptr(int64(10)),
					XCorrelationID: core.StringPtr("testString"),
				}

				pager, err := resultsService.NewReportResourcesPager(listReportResourcesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resultsv3.Resource
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReportResourcesPager.GetAll successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				listReportResourcesOptionsModel := &resultsv3.ListReportResourcesOptions{
					ReportID:       core.StringPtr("testString"),
					ID:             core.StringPtr("testString"),
					ResourceName:   core.StringPtr("testString"),
					AccountID:      core.StringPtr("testString"),
					ComponentID:    core.StringPtr("testString"),
					Status:         core.StringPtr("compliant"),
					Limit:          core.Int64Ptr(int64(10)),
					XCorrelationID: core.StringPtr("testString"),
				}

				pager, err := resultsService.NewReportResourcesPager(listReportResourcesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetReportTags(getReportTagsOptions *GetReportTagsOptions) - Operation response error`, func() {
		getReportTagsPath := "/reports/testString/tags"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportTagsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportTags with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(resultsv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportTags(getReportTagsOptions *GetReportTagsOptions)`, func() {
		getReportTagsPath := "/reports/testString/tags"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportTagsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "ReportID", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}}`)
				}))
			})
			It(`Invoke GetReportTags successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(resultsv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportTagsWithContext(ctx, getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportTagsWithContext(ctx, getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportTagsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "ReportID", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}}`)
				}))
			})
			It(`Invoke GetReportTags successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportTags(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(resultsv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportTags with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(resultsv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportTagsOptions model with no property values
				getReportTagsOptionsModelNew := new(resultsv3.GetReportTagsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReportTags(getReportTagsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportTags successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(resultsv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions) - Operation response error`, func() {
		getReportViolationsDriftPath := "/reports/testString/violations_drift"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportViolationsDriftPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["scan_time_duration"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportViolationsDrift with error: Operation response processing error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(resultsv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resultsService.EnableRetries(0, 0)
				result, response, operationErr = resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions)`, func() {
		getReportViolationsDriftPath := "/reports/testString/violations_drift"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportViolationsDriftPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["scan_time_duration"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "report_id": "ReportID", "data_points": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "report_group_id": "55b6-b3A4-432250b84669", "scan_time": "2022-08-15T12:30:01Z", "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}}]}`)
				}))
			})
			It(`Invoke GetReportViolationsDrift successfully with retries`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())
				resultsService.EnableRetries(0, 0)

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(resultsv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resultsService.GetReportViolationsDriftWithContext(ctx, getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resultsService.DisableRetries()
				result, response, operationErr := resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resultsService.GetReportViolationsDriftWithContext(ctx, getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportViolationsDriftPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["scan_time_duration"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "report_id": "ReportID", "data_points": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "report_group_id": "55b6-b3A4-432250b84669", "scan_time": "2022-08-15T12:30:01Z", "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}}]}`)
				}))
			})
			It(`Invoke GetReportViolationsDrift successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resultsService.GetReportViolationsDrift(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(resultsv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportViolationsDrift with error: Operation validation and request error`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(resultsv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resultsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportViolationsDriftOptions model with no property values
				getReportViolationsDriftOptionsModelNew := new(resultsv3.GetReportViolationsDriftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReportViolationsDrift successfully`, func() {
				resultsService, serviceErr := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resultsService).ToNot(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(resultsv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resultsService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			resultsService, _ := resultsv3.NewResultsV3(&resultsv3.ResultsV3Options{
				URL:           "http://resultsv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetLatestReportsOptions successfully`, func() {
				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := resultsService.NewGetLatestReportsOptions()
				getLatestReportsOptionsModel.SetXCorrelationID("testString")
				getLatestReportsOptionsModel.SetSort("testString")
				getLatestReportsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLatestReportsOptionsModel).ToNot(BeNil())
				Expect(getLatestReportsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestReportsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(getLatestReportsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportControlsOptions successfully`, func() {
				// Construct an instance of the GetReportControlsOptions model
				reportID := "testString"
				getReportControlsOptionsModel := resultsService.NewGetReportControlsOptions(reportID)
				getReportControlsOptionsModel.SetReportID("testString")
				getReportControlsOptionsModel.SetControlID("testString")
				getReportControlsOptionsModel.SetControlName("testString")
				getReportControlsOptionsModel.SetControlDescription("testString")
				getReportControlsOptionsModel.SetControlCategory("testString")
				getReportControlsOptionsModel.SetStatus("compliant")
				getReportControlsOptionsModel.SetSort("testString")
				getReportControlsOptionsModel.SetXCorrelationID("testString")
				getReportControlsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportControlsOptionsModel).ToNot(BeNil())
				Expect(getReportControlsOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlName).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlDescription).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlCategory).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.Status).To(Equal(core.StringPtr("compliant")))
				Expect(getReportControlsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportEvaluationOptions successfully`, func() {
				// Construct an instance of the GetReportEvaluationOptions model
				reportID := "testString"
				getReportEvaluationOptionsModel := resultsService.NewGetReportEvaluationOptions(reportID)
				getReportEvaluationOptionsModel.SetReportID("testString")
				getReportEvaluationOptionsModel.SetXCorrelationID("testString")
				getReportEvaluationOptionsModel.SetExcludeSummary(true)
				getReportEvaluationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportEvaluationOptionsModel).ToNot(BeNil())
				Expect(getReportEvaluationOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportEvaluationOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportEvaluationOptionsModel.ExcludeSummary).To(Equal(core.BoolPtr(true)))
				Expect(getReportEvaluationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportOptions successfully`, func() {
				// Construct an instance of the GetReportOptions model
				reportID := "testString"
				getReportOptionsModel := resultsService.NewGetReportOptions(reportID)
				getReportOptionsModel.SetReportID("testString")
				getReportOptionsModel.SetXCorrelationID("testString")
				getReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportOptionsModel).ToNot(BeNil())
				Expect(getReportOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportRuleOptions successfully`, func() {
				// Construct an instance of the GetReportRuleOptions model
				reportID := "testString"
				ruleID := "testString"
				getReportRuleOptionsModel := resultsService.NewGetReportRuleOptions(reportID, ruleID)
				getReportRuleOptionsModel.SetReportID("testString")
				getReportRuleOptionsModel.SetRuleID("testString")
				getReportRuleOptionsModel.SetXCorrelationID("testString")
				getReportRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportRuleOptionsModel).ToNot(BeNil())
				Expect(getReportRuleOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getReportRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportSummaryOptions successfully`, func() {
				// Construct an instance of the GetReportSummaryOptions model
				reportID := "testString"
				getReportSummaryOptionsModel := resultsService.NewGetReportSummaryOptions(reportID)
				getReportSummaryOptionsModel.SetReportID("testString")
				getReportSummaryOptionsModel.SetXCorrelationID("testString")
				getReportSummaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportSummaryOptionsModel).ToNot(BeNil())
				Expect(getReportSummaryOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportSummaryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportSummaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportTagsOptions successfully`, func() {
				// Construct an instance of the GetReportTagsOptions model
				reportID := "testString"
				getReportTagsOptionsModel := resultsService.NewGetReportTagsOptions(reportID)
				getReportTagsOptionsModel.SetReportID("testString")
				getReportTagsOptionsModel.SetXCorrelationID("testString")
				getReportTagsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportTagsOptionsModel).ToNot(BeNil())
				Expect(getReportTagsOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportTagsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportTagsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportViolationsDriftOptions successfully`, func() {
				// Construct an instance of the GetReportViolationsDriftOptions model
				reportID := "testString"
				getReportViolationsDriftOptionsModel := resultsService.NewGetReportViolationsDriftOptions(reportID)
				getReportViolationsDriftOptionsModel.SetReportID("testString")
				getReportViolationsDriftOptionsModel.SetScanTimeDuration(int64(0))
				getReportViolationsDriftOptionsModel.SetXCorrelationID("testString")
				getReportViolationsDriftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportViolationsDriftOptionsModel).ToNot(BeNil())
				Expect(getReportViolationsDriftOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportViolationsDriftOptionsModel.ScanTimeDuration).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getReportViolationsDriftOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportViolationsDriftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportsProfilesOptions successfully`, func() {
				// Construct an instance of the GetReportsProfilesOptions model
				getReportsProfilesOptionsModel := resultsService.NewGetReportsProfilesOptions()
				getReportsProfilesOptionsModel.SetXCorrelationID("testString")
				getReportsProfilesOptionsModel.SetReportID("testString")
				getReportsProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportsProfilesOptionsModel).ToNot(BeNil())
				Expect(getReportsProfilesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportsProfilesOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportsProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportsScopesOptions successfully`, func() {
				// Construct an instance of the GetReportsScopesOptions model
				getReportsScopesOptionsModel := resultsService.NewGetReportsScopesOptions()
				getReportsScopesOptionsModel.SetXCorrelationID("testString")
				getReportsScopesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportsScopesOptionsModel).ToNot(BeNil())
				Expect(getReportsScopesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportsScopesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReportEvaluationsOptions successfully`, func() {
				// Construct an instance of the ListReportEvaluationsOptions model
				reportID := "testString"
				listReportEvaluationsOptionsModel := resultsService.NewListReportEvaluationsOptions(reportID)
				listReportEvaluationsOptionsModel.SetReportID("testString")
				listReportEvaluationsOptionsModel.SetAssessmentID("testString")
				listReportEvaluationsOptionsModel.SetComponentID("testString")
				listReportEvaluationsOptionsModel.SetTargetID("testString")
				listReportEvaluationsOptionsModel.SetTargetName("testString")
				listReportEvaluationsOptionsModel.SetStatus("failure")
				listReportEvaluationsOptionsModel.SetStart("testString")
				listReportEvaluationsOptionsModel.SetLimit(int64(10))
				listReportEvaluationsOptionsModel.SetXCorrelationID("testString")
				listReportEvaluationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReportEvaluationsOptionsModel).ToNot(BeNil())
				Expect(listReportEvaluationsOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.ComponentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.TargetID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.TargetName).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.Status).To(Equal(core.StringPtr("failure")))
				Expect(listReportEvaluationsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listReportEvaluationsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReportResourcesOptions successfully`, func() {
				// Construct an instance of the ListReportResourcesOptions model
				reportID := "testString"
				listReportResourcesOptionsModel := resultsService.NewListReportResourcesOptions(reportID)
				listReportResourcesOptionsModel.SetReportID("testString")
				listReportResourcesOptionsModel.SetID("testString")
				listReportResourcesOptionsModel.SetResourceName("testString")
				listReportResourcesOptionsModel.SetAccountID("testString")
				listReportResourcesOptionsModel.SetComponentID("testString")
				listReportResourcesOptionsModel.SetStatus("compliant")
				listReportResourcesOptionsModel.SetStart("testString")
				listReportResourcesOptionsModel.SetLimit(int64(10))
				listReportResourcesOptionsModel.SetXCorrelationID("testString")
				listReportResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReportResourcesOptionsModel).ToNot(BeNil())
				Expect(listReportResourcesOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.ResourceName).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.ComponentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.Status).To(Equal(core.StringPtr("compliant")))
				Expect(listReportResourcesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listReportResourcesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReportsOptions successfully`, func() {
				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := resultsService.NewListReportsOptions()
				listReportsOptionsModel.SetXCorrelationID("testString")
				listReportsOptionsModel.SetAttachmentID("testString")
				listReportsOptionsModel.SetGroupID("testString")
				listReportsOptionsModel.SetProfileID("testString")
				listReportsOptionsModel.SetScopeID("testString")
				listReportsOptionsModel.SetType("scheduled")
				listReportsOptionsModel.SetStart("testString")
				listReportsOptionsModel.SetLimit(int64(10))
				listReportsOptionsModel.SetSort("testString")
				listReportsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReportsOptionsModel).ToNot(BeNil())
				Expect(listReportsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.Type).To(Equal(core.StringPtr("scheduled")))
				Expect(listReportsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listReportsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
