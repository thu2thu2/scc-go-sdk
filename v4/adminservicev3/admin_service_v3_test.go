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
	Describe(`GetInstanceAccess(getInstanceAccessOptions *GetInstanceAccessOptions) - Operation response error`, func() {
		getInstanceAccessPath := "/instances/testString/v3/access"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceAccessPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceAccess with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstanceAccessOptions model
				getInstanceAccessOptionsModel := new(adminservicev3.GetInstanceAccessOptions)
				getInstanceAccessOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceAccessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceAccess(getInstanceAccessOptions *GetInstanceAccessOptions)`, func() {
		getInstanceAccessPath := "/instances/testString/v3/access"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceAccessPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"admin": {"read": true, "update": true}}`)
				}))
			})
			It(`Invoke GetInstanceAccess successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceAccessOptions model
				getInstanceAccessOptionsModel := new(adminservicev3.GetInstanceAccessOptions)
				getInstanceAccessOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceAccessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.GetInstanceAccessWithContext(ctx, getInstanceAccessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.GetInstanceAccessWithContext(ctx, getInstanceAccessOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceAccessPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"admin": {"read": true, "update": true}}`)
				}))
			})
			It(`Invoke GetInstanceAccess successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.GetInstanceAccess(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceAccessOptions model
				getInstanceAccessOptionsModel := new(adminservicev3.GetInstanceAccessOptions)
				getInstanceAccessOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceAccessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceAccess with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstanceAccessOptions model
				getInstanceAccessOptionsModel := new(adminservicev3.GetInstanceAccessOptions)
				getInstanceAccessOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceAccessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceAccessOptions model with no property values
				getInstanceAccessOptionsModelNew := new(adminservicev3.GetInstanceAccessOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModelNew)
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
			It(`Invoke GetInstanceAccess successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstanceAccessOptions model
				getInstanceAccessOptionsModel := new(adminservicev3.GetInstanceAccessOptions)
				getInstanceAccessOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceAccessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.GetInstanceAccess(getInstanceAccessOptionsModel)
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
	Describe(`GetInstanceSettings(getInstanceSettingsOptions *GetInstanceSettingsOptions) - Operation response error`, func() {
		getInstanceSettingsPath := "/instances/testString/v3/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceSettings with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstanceSettingsOptions model
				getInstanceSettingsOptionsModel := new(adminservicev3.GetInstanceSettingsOptions)
				getInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceSettings(getInstanceSettingsOptions *GetInstanceSettingsOptions)`, func() {
		getInstanceSettingsPath := "/instances/testString/v3/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "ca-tor"}, "event_notifications": {"instance_crn": "InstanceCrn", "name": "Name", "description": "Description"}}`)
				}))
			})
			It(`Invoke GetInstanceSettings successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetInstanceSettingsOptions model
				getInstanceSettingsOptionsModel := new(adminservicev3.GetInstanceSettingsOptions)
				getInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.GetInstanceSettingsWithContext(ctx, getInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.GetInstanceSettingsWithContext(ctx, getInstanceSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstanceSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "ca-tor"}, "event_notifications": {"instance_crn": "InstanceCrn", "name": "Name", "description": "Description"}}`)
				}))
			})
			It(`Invoke GetInstanceSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.GetInstanceSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceSettingsOptions model
				getInstanceSettingsOptionsModel := new(adminservicev3.GetInstanceSettingsOptions)
				getInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstanceSettings with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstanceSettingsOptions model
				getInstanceSettingsOptionsModel := new(adminservicev3.GetInstanceSettingsOptions)
				getInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceSettingsOptions model with no property values
				getInstanceSettingsOptionsModelNew := new(adminservicev3.GetInstanceSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModelNew)
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
			It(`Invoke GetInstanceSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstanceSettingsOptions model
				getInstanceSettingsOptionsModel := new(adminservicev3.GetInstanceSettingsOptions)
				getInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				getInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.GetInstanceSettings(getInstanceSettingsOptionsModel)
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
	Describe(`UpdateInstanceSettings(updateInstanceSettingsOptions *UpdateInstanceSettingsOptions) - Operation response error`, func() {
		updateInstanceSettingsPath := "/instances/testString/v3/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceSettings with error: Operation response processing error`, func() {
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

				// Construct an instance of the UpdateInstanceSettingsOptions model
				updateInstanceSettingsOptionsModel := new(adminservicev3.UpdateInstanceSettingsOptions)
				updateInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				updateInstanceSettingsOptionsModel.JSONPatchOperation = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceSettings(updateInstanceSettingsOptions *UpdateInstanceSettingsOptions)`, func() {
		updateInstanceSettingsPath := "/instances/testString/v3/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceSettingsPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "ca-tor"}, "event_notifications": {"instance_crn": "InstanceCrn", "name": "Name", "description": "Description"}}`)
				}))
			})
			It(`Invoke UpdateInstanceSettings successfully with retries`, func() {
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

				// Construct an instance of the UpdateInstanceSettingsOptions model
				updateInstanceSettingsOptionsModel := new(adminservicev3.UpdateInstanceSettingsOptions)
				updateInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				updateInstanceSettingsOptionsModel.JSONPatchOperation = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.UpdateInstanceSettingsWithContext(ctx, updateInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.UpdateInstanceSettingsWithContext(ctx, updateInstanceSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateInstanceSettingsPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "ca-tor"}, "event_notifications": {"instance_crn": "InstanceCrn", "name": "Name", "description": "Description"}}`)
				}))
			})
			It(`Invoke UpdateInstanceSettings successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.UpdateInstanceSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(adminservicev3.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateInstanceSettingsOptions model
				updateInstanceSettingsOptionsModel := new(adminservicev3.UpdateInstanceSettingsOptions)
				updateInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				updateInstanceSettingsOptionsModel.JSONPatchOperation = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateInstanceSettings with error: Operation validation and request error`, func() {
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

				// Construct an instance of the UpdateInstanceSettingsOptions model
				updateInstanceSettingsOptionsModel := new(adminservicev3.UpdateInstanceSettingsOptions)
				updateInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				updateInstanceSettingsOptionsModel.JSONPatchOperation = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceSettingsOptions model with no property values
				updateInstanceSettingsOptionsModelNew := new(adminservicev3.UpdateInstanceSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModelNew)
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
			It(`Invoke UpdateInstanceSettings successfully`, func() {
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

				// Construct an instance of the UpdateInstanceSettingsOptions model
				updateInstanceSettingsOptionsModel := new(adminservicev3.UpdateInstanceSettingsOptions)
				updateInstanceSettingsOptionsModel.InstanceID = core.StringPtr("testString")
				updateInstanceSettingsOptionsModel.JSONPatchOperation = []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}
				updateInstanceSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptionsModel)
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
	Describe(`PostInstanceTestEvent(postInstanceTestEventOptions *PostInstanceTestEventOptions) - Operation response error`, func() {
		postInstanceTestEventPath := "/instances/testString/v3/test_event"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postInstanceTestEventPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostInstanceTestEvent with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostInstanceTestEventOptions model
				postInstanceTestEventOptionsModel := new(adminservicev3.PostInstanceTestEventOptions)
				postInstanceTestEventOptionsModel.InstanceID = core.StringPtr("testString")
				postInstanceTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostInstanceTestEvent(postInstanceTestEventOptions *PostInstanceTestEventOptions)`, func() {
		postInstanceTestEventPath := "/instances/testString/v3/test_event"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postInstanceTestEventPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false}`)
				}))
			})
			It(`Invoke PostInstanceTestEvent successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the PostInstanceTestEventOptions model
				postInstanceTestEventOptionsModel := new(adminservicev3.PostInstanceTestEventOptions)
				postInstanceTestEventOptionsModel.InstanceID = core.StringPtr("testString")
				postInstanceTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.PostInstanceTestEventWithContext(ctx, postInstanceTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.PostInstanceTestEventWithContext(ctx, postInstanceTestEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postInstanceTestEventPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": false}`)
				}))
			})
			It(`Invoke PostInstanceTestEvent successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.PostInstanceTestEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostInstanceTestEventOptions model
				postInstanceTestEventOptionsModel := new(adminservicev3.PostInstanceTestEventOptions)
				postInstanceTestEventOptionsModel.InstanceID = core.StringPtr("testString")
				postInstanceTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostInstanceTestEvent with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostInstanceTestEventOptions model
				postInstanceTestEventOptionsModel := new(adminservicev3.PostInstanceTestEventOptions)
				postInstanceTestEventOptionsModel.InstanceID = core.StringPtr("testString")
				postInstanceTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostInstanceTestEventOptions model with no property values
				postInstanceTestEventOptionsModelNew := new(adminservicev3.PostInstanceTestEventOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModelNew)
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
			It(`Invoke PostInstanceTestEvent successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostInstanceTestEventOptions model
				postInstanceTestEventOptionsModel := new(adminservicev3.PostInstanceTestEventOptions)
				postInstanceTestEventOptionsModel.InstanceID = core.StringPtr("testString")
				postInstanceTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptionsModel)
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
	Describe(`GetInstancePlans(getInstancePlansOptions *GetInstancePlansOptions) - Operation response error`, func() {
		getInstancePlansPath := "/instances/testString/v3/plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePlansPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstancePlans with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstancePlansOptions model
				getInstancePlansOptionsModel := new(adminservicev3.GetInstancePlansOptions)
				getInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				getInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.GetInstancePlans(getInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.GetInstancePlans(getInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstancePlans(getInstancePlansOptions *GetInstancePlansOptions)`, func() {
		getInstancePlansPath := "/instances/testString/v3/plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePlansPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name"}`)
				}))
			})
			It(`Invoke GetInstancePlans successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetInstancePlansOptions model
				getInstancePlansOptionsModel := new(adminservicev3.GetInstancePlansOptions)
				getInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				getInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.GetInstancePlansWithContext(ctx, getInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.GetInstancePlans(getInstancePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.GetInstancePlansWithContext(ctx, getInstancePlansOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getInstancePlansPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name"}`)
				}))
			})
			It(`Invoke GetInstancePlans successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.GetInstancePlans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstancePlansOptions model
				getInstancePlansOptionsModel := new(adminservicev3.GetInstancePlansOptions)
				getInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				getInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.GetInstancePlans(getInstancePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInstancePlans with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstancePlansOptions model
				getInstancePlansOptionsModel := new(adminservicev3.GetInstancePlansOptions)
				getInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				getInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.GetInstancePlans(getInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstancePlansOptions model with no property values
				getInstancePlansOptionsModelNew := new(adminservicev3.GetInstancePlansOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.GetInstancePlans(getInstancePlansOptionsModelNew)
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
			It(`Invoke GetInstancePlans successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the GetInstancePlansOptions model
				getInstancePlansOptionsModel := new(adminservicev3.GetInstancePlansOptions)
				getInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				getInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.GetInstancePlans(getInstancePlansOptionsModel)
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
	Describe(`PostInstancePlans(postInstancePlansOptions *PostInstancePlansOptions) - Operation response error`, func() {
		postInstancePlansPath := "/instances/testString/v3/plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postInstancePlansPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostInstancePlans with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostInstancePlansOptions model
				postInstancePlansOptionsModel := new(adminservicev3.PostInstancePlansOptions)
				postInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				postInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				postInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.PostInstancePlans(postInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.PostInstancePlans(postInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostInstancePlans(postInstancePlansOptions *PostInstancePlansOptions)`, func() {
		postInstancePlansPath := "/instances/testString/v3/plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postInstancePlansPath))
					Expect(req.Method).To(Equal("POST"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name"}`)
				}))
			})
			It(`Invoke PostInstancePlans successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the PostInstancePlansOptions model
				postInstancePlansOptionsModel := new(adminservicev3.PostInstancePlansOptions)
				postInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				postInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				postInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.PostInstancePlansWithContext(ctx, postInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.PostInstancePlans(postInstancePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.PostInstancePlansWithContext(ctx, postInstancePlansOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postInstancePlansPath))
					Expect(req.Method).To(Equal("POST"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name"}`)
				}))
			})
			It(`Invoke PostInstancePlans successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.PostInstancePlans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostInstancePlansOptions model
				postInstancePlansOptionsModel := new(adminservicev3.PostInstancePlansOptions)
				postInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				postInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				postInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.PostInstancePlans(postInstancePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostInstancePlans with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostInstancePlansOptions model
				postInstancePlansOptionsModel := new(adminservicev3.PostInstancePlansOptions)
				postInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				postInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				postInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.PostInstancePlans(postInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostInstancePlansOptions model with no property values
				postInstancePlansOptionsModelNew := new(adminservicev3.PostInstancePlansOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.PostInstancePlans(postInstancePlansOptionsModelNew)
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
			It(`Invoke PostInstancePlans successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the PostInstancePlansOptions model
				postInstancePlansOptionsModel := new(adminservicev3.PostInstancePlansOptions)
				postInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				postInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				postInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.PostInstancePlans(postInstancePlansOptionsModel)
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
	Describe(`SetInstancePlans(setInstancePlansOptions *SetInstancePlansOptions) - Operation response error`, func() {
		setInstancePlansPath := "/instances/testString/v3/plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setInstancePlansPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetInstancePlans with error: Operation response processing error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the SetInstancePlansOptions model
				setInstancePlansOptionsModel := new(adminservicev3.SetInstancePlansOptions)
				setInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				setInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				setInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceService.SetInstancePlans(setInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceService.SetInstancePlans(setInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetInstancePlans(setInstancePlansOptions *SetInstancePlansOptions)`, func() {
		setInstancePlansPath := "/instances/testString/v3/plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setInstancePlansPath))
					Expect(req.Method).To(Equal("PUT"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name"}`)
				}))
			})
			It(`Invoke SetInstancePlans successfully with retries`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())
				adminServiceService.EnableRetries(0, 0)

				// Construct an instance of the SetInstancePlansOptions model
				setInstancePlansOptionsModel := new(adminservicev3.SetInstancePlansOptions)
				setInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				setInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				setInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceService.SetInstancePlansWithContext(ctx, setInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceService.DisableRetries()
				result, response, operationErr := adminServiceService.SetInstancePlans(setInstancePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceService.SetInstancePlansWithContext(ctx, setInstancePlansOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setInstancePlansPath))
					Expect(req.Method).To(Equal("PUT"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name"}`)
				}))
			})
			It(`Invoke SetInstancePlans successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceService.SetInstancePlans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetInstancePlansOptions model
				setInstancePlansOptionsModel := new(adminservicev3.SetInstancePlansOptions)
				setInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				setInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				setInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceService.SetInstancePlans(setInstancePlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetInstancePlans with error: Operation validation and request error`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the SetInstancePlansOptions model
				setInstancePlansOptionsModel := new(adminservicev3.SetInstancePlansOptions)
				setInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				setInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				setInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceService.SetInstancePlans(setInstancePlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetInstancePlansOptions model with no property values
				setInstancePlansOptionsModelNew := new(adminservicev3.SetInstancePlansOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceService.SetInstancePlans(setInstancePlansOptionsModelNew)
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
			It(`Invoke SetInstancePlans successfully`, func() {
				adminServiceService, serviceErr := adminservicev3.NewAdminServiceV3(&adminservicev3.AdminServiceV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceService).ToNot(BeNil())

				// Construct an instance of the SetInstancePlansOptions model
				setInstancePlansOptionsModel := new(adminservicev3.SetInstancePlansOptions)
				setInstancePlansOptionsModel.InstanceID = core.StringPtr("testString")
				setInstancePlansOptionsModel.Name = core.StringPtr("Standard")
				setInstancePlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceService.SetInstancePlans(setInstancePlansOptionsModel)
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
			It(`Invoke NewGetInstanceAccessOptions successfully`, func() {
				// Construct an instance of the GetInstanceAccessOptions model
				instanceID := "testString"
				getInstanceAccessOptionsModel := adminServiceService.NewGetInstanceAccessOptions(instanceID)
				getInstanceAccessOptionsModel.SetInstanceID("testString")
				getInstanceAccessOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceAccessOptionsModel).ToNot(BeNil())
				Expect(getInstanceAccessOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getInstanceAccessOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstancePlansOptions successfully`, func() {
				// Construct an instance of the GetInstancePlansOptions model
				instanceID := "testString"
				getInstancePlansOptionsModel := adminServiceService.NewGetInstancePlansOptions(instanceID)
				getInstancePlansOptionsModel.SetInstanceID("testString")
				getInstancePlansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstancePlansOptionsModel).ToNot(BeNil())
				Expect(getInstancePlansOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getInstancePlansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInstanceSettingsOptions successfully`, func() {
				// Construct an instance of the GetInstanceSettingsOptions model
				instanceID := "testString"
				getInstanceSettingsOptionsModel := adminServiceService.NewGetInstanceSettingsOptions(instanceID)
				getInstanceSettingsOptionsModel.SetInstanceID("testString")
				getInstanceSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceSettingsOptionsModel).ToNot(BeNil())
				Expect(getInstanceSettingsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getInstanceSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInstancePlan successfully`, func() {
				name := "testString"
				_model, err := adminServiceService.NewInstancePlan(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := adminServiceService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostInstancePlansOptions successfully`, func() {
				// Construct an instance of the PostInstancePlansOptions model
				instanceID := "testString"
				postInstancePlansOptionsName := "Standard"
				postInstancePlansOptionsModel := adminServiceService.NewPostInstancePlansOptions(instanceID, postInstancePlansOptionsName)
				postInstancePlansOptionsModel.SetInstanceID("testString")
				postInstancePlansOptionsModel.SetName("Standard")
				postInstancePlansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postInstancePlansOptionsModel).ToNot(BeNil())
				Expect(postInstancePlansOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(postInstancePlansOptionsModel.Name).To(Equal(core.StringPtr("Standard")))
				Expect(postInstancePlansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostInstanceTestEventOptions successfully`, func() {
				// Construct an instance of the PostInstanceTestEventOptions model
				instanceID := "testString"
				postInstanceTestEventOptionsModel := adminServiceService.NewPostInstanceTestEventOptions(instanceID)
				postInstanceTestEventOptionsModel.SetInstanceID("testString")
				postInstanceTestEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postInstanceTestEventOptionsModel).ToNot(BeNil())
				Expect(postInstanceTestEventOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(postInstanceTestEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetInstancePlansOptions successfully`, func() {
				// Construct an instance of the SetInstancePlansOptions model
				instanceID := "testString"
				setInstancePlansOptionsName := "Standard"
				setInstancePlansOptionsModel := adminServiceService.NewSetInstancePlansOptions(instanceID, setInstancePlansOptionsName)
				setInstancePlansOptionsModel.SetInstanceID("testString")
				setInstancePlansOptionsModel.SetName("Standard")
				setInstancePlansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setInstancePlansOptionsModel).ToNot(BeNil())
				Expect(setInstancePlansOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(setInstancePlansOptionsModel.Name).To(Equal(core.StringPtr("Standard")))
				Expect(setInstancePlansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceSettingsOptions successfully`, func() {
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

				// Construct an instance of the UpdateInstanceSettingsOptions model
				instanceID := "testString"
				jsonPatchOperation := []adminservicev3.JSONPatchOperation{}
				updateInstanceSettingsOptionsModel := adminServiceService.NewUpdateInstanceSettingsOptions(instanceID, jsonPatchOperation)
				updateInstanceSettingsOptionsModel.SetInstanceID("testString")
				updateInstanceSettingsOptionsModel.SetJSONPatchOperation([]adminservicev3.JSONPatchOperation{*jsonPatchOperationModel})
				updateInstanceSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceSettingsOptionsModel).ToNot(BeNil())
				Expect(updateInstanceSettingsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateInstanceSettingsOptionsModel.JSONPatchOperation).To(Equal([]adminservicev3.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateInstanceSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
