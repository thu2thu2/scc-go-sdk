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

package adminservicev3_test

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
	"github.com/IBM/scc-go-sdk/v4/adminservicev3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe(`AdminServiceV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(adminServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(adminServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
				URL: "https://adminservicev3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(adminServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMIN_SERVICE_URL":       "https://adminservicev3/api",
				"ADMIN_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3UsingExternalConfig(&adminservicev3.AdminServiceV3Options{})
				Expect(adminServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := adminServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3UsingExternalConfig(&adminservicev3.AdminServiceV3Options{
					URL: "https://testService/api",
				})
				Expect(adminServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := adminServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3UsingExternalConfig(&adminservicev3.AdminServiceV3Options{})
				err := adminServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := adminServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMIN_SERVICE_URL":       "https://adminservicev3/api",
				"ADMIN_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3UsingExternalConfig(&adminservicev3.AdminServiceV3Options{})

			It(`Instantiate service client with error`, func() {
				Expect(adminServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMIN_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3UsingExternalConfig(&adminservicev3.AdminServiceV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(adminServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = adminservicev3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := adminservicev3.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := adminservicev3.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminservicev3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "InstanceCrn", "updated_on": "UpdatedOn", "source_id": "SourceID"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "UpdatedOn"}}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminservicev3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "InstanceCrn", "updated_on": "UpdatedOn", "source_id": "SourceID"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "UpdatedOn"}}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminservicev3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminservicev3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.GetSettings(getSettingsOptionsModel)
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
			It(`Invoke GetSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminservicev3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.GetSettings(getSettingsOptionsModel)
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
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions) - Operation response error`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSettings with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(adminservicev3.UpdateSettingsOptions)
				updateSettingsOptionsModel.Body = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "InstanceCrn", "updated_on": "UpdatedOn", "source_id": "SourceID"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "UpdatedOn"}}`)
				}))
			})
			It(`Invoke UpdateSettings successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(adminservicev3.UpdateSettingsOptions)
				updateSettingsOptionsModel.Body = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "InstanceCrn", "updated_on": "UpdatedOn", "source_id": "SourceID"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "UpdatedOn"}}`)
				}))
			})
			It(`Invoke UpdateSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.UpdateSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(adminservicev3.UpdateSettingsOptions)
				updateSettingsOptionsModel.Body = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSettings with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(adminservicev3.UpdateSettingsOptions)
				updateSettingsOptionsModel.Body = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSettingsOptions model with no property values
				updateSettingsOptionsModelNew := new(adminservicev3.UpdateSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.UpdateSettings(updateSettingsOptionsModelNew)
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
			It(`Invoke UpdateSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(adminservicev3.UpdateSettingsOptions)
				updateSettingsOptionsModel.Body = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.UpdateSettings(updateSettingsOptionsModel)
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
	Describe(`PostTestEvent(postTestEventOptions *PostTestEventOptions) - Operation response error`, func() {
		postTestEventPath := "/test_event"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostTestEvent with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(adminservicev3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("testString")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostTestEvent(postTestEventOptions *PostTestEventOptions)`, func() {
		postTestEventPath := "/test_event"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"success": false}`)
				}))
			})
			It(`Invoke PostTestEvent successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(adminservicev3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("testString")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.PostTestEventWithContext(ctx, postTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.PostTestEventWithContext(ctx, postTestEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"success": false}`)
				}))
			})
			It(`Invoke PostTestEvent successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.PostTestEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(adminservicev3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("testString")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostTestEvent with error: Operation request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(adminservicev3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("testString")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.PostTestEvent(postTestEventOptionsModel)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke PostTestEvent successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(adminservicev3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("testString")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.PostTestEvent(postTestEventOptionsModel)
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
			adminServiceService, _ := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
				URL:           "http://adminservicev3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := adminServiceService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetXCorrelationID("testString")
				getSettingsOptionsModel.SetXRequestID("testString")
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := adminServiceService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostTestEventOptions successfully`, func() {
				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := adminServiceService.NewPostTestEventOptions()
				postTestEventOptionsModel.SetXCorrelationID("testString")
				postTestEventOptionsModel.SetXRequestID("testString")
				postTestEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postTestEventOptionsModel).ToNot(BeNil())
				Expect(postTestEventOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(postTestEventOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(postTestEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSettingsPatch successfully`, func() {
				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(adminservicev3.EventNotifications)
				eventNotificationsModel.InstanceCrn = core.StringPtr("testString")
				eventNotificationsModel.UpdatedOn = core.StringPtr("testString")
				eventNotificationsModel.SourceID = core.StringPtr("testString")

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(adminservicev3.ObjectStorage)
				objectStorageModel.InstanceCrn = core.StringPtr("testString")
				objectStorageModel.Bucket = core.StringPtr("testString")
				objectStorageModel.BucketLocation = core.StringPtr("testString")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settings := new(adminservicev3.Settings)
				settings.EventNotifications = eventNotificationsModel
				settings.ObjectStorage = objectStorageModel

				settingsPatch := adminServiceService.NewSettingsPatch(settings)
				Expect(settingsPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(adminservicev3.JSONPatchOperation).Path
				}
				Expect(settingsPatch).To(MatchAllElements(_path, Elements{
					"/event_notifications": MatchAllFields(Fields{
						"Op":    PointTo(Equal(adminservicev3.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/event_notifications")),
						"From":  BeNil(),
						"Value": Equal(settings.EventNotifications),
					}),
					"/object_storage": MatchAllFields(Fields{
						"Op":    PointTo(Equal(adminservicev3.JSONPatchOperation_Op_Add)),
						"Path":  PointTo(Equal("/object_storage")),
						"From":  BeNil(),
						"Value": Equal(settings.ObjectStorage),
					}),
				}))
			})
			It(`Invoke NewUpdateSettingsOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateSettingsOptions model
				body := []adminservicev3.JSONPatchOperation{}
				updateSettingsOptionsModel := adminServiceService.NewUpdateSettingsOptions(body)
				updateSettingsOptionsModel.SetBody([]adminservicev3.JSONPatchOperation{*jsonPatchOperationModel})
				updateSettingsOptionsModel.SetXCorrelationID("testString")
				updateSettingsOptionsModel.SetXRequestID("testString")
				updateSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSettingsOptionsModel).ToNot(BeNil())
				Expect(updateSettingsOptionsModel.Body).To(Equal([]adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateSettingsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(updateSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
