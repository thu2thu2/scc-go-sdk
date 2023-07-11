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

package configmanagerv3_test

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
	"github.com/IBM/scc-go-sdk/v4/configmanagerv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ConfigManagerV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(configManagerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(configManagerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
				URL: "https://configmanagerv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(configManagerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIG_MANAGER_URL":       "https://configmanagerv3/api",
				"CONFIG_MANAGER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3UsingExternalConfig(&configmanagerv3.ConfigManagerV3Options{})
				Expect(configManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := configManagerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configManagerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configManagerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configManagerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3UsingExternalConfig(&configmanagerv3.ConfigManagerV3Options{
					URL: "https://testService/api",
				})
				Expect(configManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configManagerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configManagerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configManagerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configManagerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3UsingExternalConfig(&configmanagerv3.ConfigManagerV3Options{})
				err := configManagerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configManagerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configManagerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configManagerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configManagerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIG_MANAGER_URL":       "https://configmanagerv3/api",
				"CONFIG_MANAGER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3UsingExternalConfig(&configmanagerv3.ConfigManagerV3Options{})

			It(`Instantiate service client with error`, func() {
				Expect(configManagerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIG_MANAGER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3UsingExternalConfig(&configmanagerv3.ConfigManagerV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(configManagerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = configmanagerv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := configmanagerv3.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := configmanagerv3.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`ListRules(listRulesOptions *ListRulesOptions) - Operation response error`, func() {
		listRulesPath := "/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type_query"]).To(Equal([]string{"system_defined"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRules with error: Operation response processing error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerv3.ListRulesOptions)
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerService.EnableRetries(0, 0)
				result, response, operationErr = configManagerService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
		listRulesPath := "/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type_query"]).To(Equal([]string{"system_defined"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke ListRules successfully with retries`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())
				configManagerService.EnableRetries(0, 0)

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerv3.ListRulesOptions)
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerService.ListRulesWithContext(ctx, listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerService.DisableRetries()
				result, response, operationErr := configManagerService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerService.ListRulesWithContext(ctx, listRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type_query"]).To(Equal([]string{"system_defined"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke ListRules successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerService.ListRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerv3.ListRulesOptions)
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRules with error: Operation request error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerv3.ListRulesOptions)
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerService.ListRules(listRulesOptionsModel)
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
			It(`Invoke ListRules successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerv3.ListRulesOptions)
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerService.ListRules(listRulesOptionsModel)
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
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions) - Operation response error`, func() {
		createRulePath := "/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["type_query"]).To(Equal([]string{"system_defined"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRule with error: Operation response processing error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerv3.CreateRuleOptions)
				createRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerService.EnableRetries(0, 0)
				result, response, operationErr = configManagerService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
		createRulePath := "/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
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

					Expect(req.URL.Query()["type_query"]).To(Equal([]string{"system_defined"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke CreateRule successfully with retries`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())
				configManagerService.EnableRetries(0, 0)

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerv3.CreateRuleOptions)
				createRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerService.CreateRuleWithContext(ctx, createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerService.DisableRetries()
				result, response, operationErr := configManagerService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerService.CreateRuleWithContext(ctx, createRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
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

					Expect(req.URL.Query()["type_query"]).To(Equal([]string{"system_defined"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerService.CreateRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerv3.CreateRuleOptions)
				createRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRule with error: Operation validation and request error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerv3.CreateRuleOptions)
				createRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRuleOptions model with no property values
				createRuleOptionsModelNew := new(configmanagerv3.CreateRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerService.CreateRule(createRuleOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerv3.CreateRuleOptions)
				createRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerService.CreateRule(createRuleOptionsModel)
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
	Describe(`GetRule(getRuleOptions *GetRuleOptions) - Operation response error`, func() {
		getRulePath := "/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRule with error: Operation response processing error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerService.EnableRetries(0, 0)
				result, response, operationErr = configManagerService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
		getRulePath := "/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetRule successfully with retries`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())
				configManagerService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerService.GetRuleWithContext(ctx, getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerService.DisableRetries()
				result, response, operationErr := configManagerService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerService.GetRuleWithContext(ctx, getRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerService.GetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRule with error: Operation validation and request error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleOptions model with no property values
				getRuleOptionsModelNew := new(configmanagerv3.GetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerService.GetRule(getRuleOptionsModelNew)
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
			It(`Invoke GetRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerService.GetRule(getRuleOptionsModel)
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
	Describe(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) - Operation response error`, func() {
		replaceRulePath := "/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Ifmatch"]).ToNot(BeNil())
					Expect(req.Header["Ifmatch"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceRule with error: Operation response processing error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(configmanagerv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerService.EnableRetries(0, 0)
				result, response, operationErr = configManagerService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
		replaceRulePath := "/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
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

					Expect(req.Header["Ifmatch"]).ToNot(BeNil())
					Expect(req.Header["Ifmatch"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke ReplaceRule successfully with retries`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())
				configManagerService.EnableRetries(0, 0)

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(configmanagerv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerService.ReplaceRuleWithContext(ctx, replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerService.DisableRetries()
				result, response, operationErr := configManagerService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerService.ReplaceRuleWithContext(ctx, replaceRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
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

					Expect(req.Header["Ifmatch"]).ToNot(BeNil())
					Expect(req.Header["Ifmatch"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"property": "Property", "operator": "string_equals", "value": "Value"}], "or": [{"property": "Property", "operator": "string_equals", "value": "Value"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke ReplaceRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerService.ReplaceRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(configmanagerv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceRule with error: Operation validation and request error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(configmanagerv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceRuleOptions model with no property values
				replaceRuleOptionsModelNew := new(configmanagerv3.ReplaceRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerService.ReplaceRule(replaceRuleOptionsModelNew)
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
			It(`Invoke ReplaceRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(configmanagerv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.AccountID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerService.ReplaceRule(replaceRuleOptionsModel)
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
	Describe(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
		deleteRulePath := "/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRule successfully`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := configManagerService.DeleteRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configmanagerv3.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = configManagerService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRule with error: Operation validation and request error`, func() {
				configManagerService, serviceErr := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configmanagerv3.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := configManagerService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleOptions model with no property values
				deleteRuleOptionsModelNew := new(configmanagerv3.DeleteRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = configManagerService.DeleteRule(deleteRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			configManagerService, _ := configmanagerv3.NewConfigManagerV3(&configmanagerv3.ConfigManagerV3Options{
				URL:           "http://configmanagerv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateRuleOptions successfully`, func() {
				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				Expect(additionalTargetAttributeModel).ToNot(BeNil())
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")
				Expect(additionalTargetAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(additionalTargetAttributeModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(additionalTargetAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}
				Expect(targetModel.ServiceName).To(Equal(core.StringPtr("cloud-object-storage")))
				Expect(targetModel.ServiceDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.ResourceKind).To(Equal(core.StringPtr("bucket")))
				Expect(targetModel.AdditionalTargetAttributes).To(Equal([]configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}))

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				Expect(andModel).ToNot(BeNil())
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")
				Expect(andModel.Property).To(Equal(core.StringPtr("hard_quota")))
				Expect(andModel.Operator).To(Equal(core.StringPtr("num_equals")))
				Expect(andModel.Value).To(Equal(core.StringPtr("${hard_quota}")))

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				Expect(orModel).ToNot(BeNil())
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")
				Expect(orModel.Property).To(Equal(core.StringPtr("testString")))
				Expect(orModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(orModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				Expect(requiredConfigModel).ToNot(BeNil())
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}
				Expect(requiredConfigModel.Description).To(Equal(core.StringPtr("The Cloud Object Storage rule.")))
				Expect(requiredConfigModel.And).To(Equal([]configmanagerv3.And{*andModel}))
				Expect(requiredConfigModel.Or).To(Equal([]configmanagerv3.Or{*orModel}))

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				Expect(parameterModel).ToNot(BeNil())
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")
				Expect(parameterModel.Name).To(Equal(core.StringPtr("hard_quota")))
				Expect(parameterModel.DisplayName).To(Equal(core.StringPtr("The Cloud Object Storage bucket quota.")))
				Expect(parameterModel.Description).To(Equal(core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")))
				Expect(parameterModel.Type).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				Expect(importModel).ToNot(BeNil())
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}
				Expect(importModel.Parameters).To(Equal([]configmanagerv3.Parameter{*parameterModel}))

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsAccountID := "130003ea8bfa43c5aacea07a86da3000"
				createRuleOptionsDescription := "Example rule"
				var createRuleOptionsTarget *configmanagerv3.Target = nil
				var createRuleOptionsRequiredConfig *configmanagerv3.RequiredConfig = nil
				createRuleOptionsLabels := []string{}
				createRuleOptionsModel := configManagerService.NewCreateRuleOptions(createRuleOptionsAccountID, createRuleOptionsDescription, createRuleOptionsTarget, createRuleOptionsRequiredConfig, createRuleOptionsLabels)
				createRuleOptionsModel.SetAccountID("130003ea8bfa43c5aacea07a86da3000")
				createRuleOptionsModel.SetDescription("Example rule")
				createRuleOptionsModel.SetTarget(targetModel)
				createRuleOptionsModel.SetRequiredConfig(requiredConfigModel)
				createRuleOptionsModel.SetLabels([]string{})
				createRuleOptionsModel.SetType("user_defined")
				createRuleOptionsModel.SetVersion("1.0.0")
				createRuleOptionsModel.SetImport(importModel)
				createRuleOptionsModel.SetTypeQuery("system_defined")
				createRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRuleOptionsModel).ToNot(BeNil())
				Expect(createRuleOptionsModel.AccountID).To(Equal(core.StringPtr("130003ea8bfa43c5aacea07a86da3000")))
				Expect(createRuleOptionsModel.Description).To(Equal(core.StringPtr("Example rule")))
				Expect(createRuleOptionsModel.Target).To(Equal(targetModel))
				Expect(createRuleOptionsModel.RequiredConfig).To(Equal(requiredConfigModel))
				Expect(createRuleOptionsModel.Labels).To(Equal([]string{}))
				Expect(createRuleOptionsModel.Type).To(Equal(core.StringPtr("user_defined")))
				Expect(createRuleOptionsModel.Version).To(Equal(core.StringPtr("1.0.0")))
				Expect(createRuleOptionsModel.Import).To(Equal(importModel))
				Expect(createRuleOptionsModel.TypeQuery).To(Equal(core.StringPtr("system_defined")))
				Expect(createRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleOptions successfully`, func() {
				// Construct an instance of the DeleteRuleOptions model
				ruleID := "testString"
				deleteRuleOptionsModel := configManagerService.NewDeleteRuleOptions(ruleID)
				deleteRuleOptionsModel.SetRuleID("testString")
				deleteRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleOptionsModel).ToNot(BeNil())
				Expect(deleteRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRuleOptions successfully`, func() {
				// Construct an instance of the GetRuleOptions model
				ruleID := "testString"
				getRuleOptionsModel := configManagerService.NewGetRuleOptions(ruleID)
				getRuleOptionsModel.SetRuleID("testString")
				getRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleOptionsModel).ToNot(BeNil())
				Expect(getRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRulesOptions successfully`, func() {
				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := configManagerService.NewListRulesOptions()
				listRulesOptionsModel.SetTypeQuery("system_defined")
				listRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRulesOptionsModel).ToNot(BeNil())
				Expect(listRulesOptionsModel.TypeQuery).To(Equal(core.StringPtr("system_defined")))
				Expect(listRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceRuleOptions successfully`, func() {
				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(configmanagerv3.AdditionalTargetAttribute)
				Expect(additionalTargetAttributeModel).ToNot(BeNil())
				additionalTargetAttributeModel.Name = core.StringPtr("testString")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("testString")
				Expect(additionalTargetAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(additionalTargetAttributeModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(additionalTargetAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Target model
				targetModel := new(configmanagerv3.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}
				Expect(targetModel.ServiceName).To(Equal(core.StringPtr("cloud-object-storage")))
				Expect(targetModel.ServiceDisplayName).To(Equal(core.StringPtr("Cloud Object Storage")))
				Expect(targetModel.ResourceKind).To(Equal(core.StringPtr("bucket")))
				Expect(targetModel.AdditionalTargetAttributes).To(Equal([]configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}))

				// Construct an instance of the And model
				andModel := new(configmanagerv3.And)
				Expect(andModel).ToNot(BeNil())
				andModel.Property = core.StringPtr("hard_quota")
				andModel.Operator = core.StringPtr("num_equals")
				andModel.Value = core.StringPtr("${hard_quota}")
				Expect(andModel.Property).To(Equal(core.StringPtr("hard_quota")))
				Expect(andModel.Operator).To(Equal(core.StringPtr("num_equals")))
				Expect(andModel.Value).To(Equal(core.StringPtr("${hard_quota}")))

				// Construct an instance of the Or model
				orModel := new(configmanagerv3.Or)
				Expect(orModel).ToNot(BeNil())
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("string_equals")
				orModel.Value = core.StringPtr("testString")
				Expect(orModel.Property).To(Equal(core.StringPtr("testString")))
				Expect(orModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(orModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerv3.RequiredConfig)
				Expect(requiredConfigModel).ToNot(BeNil())
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []configmanagerv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerv3.Or{*orModel}
				Expect(requiredConfigModel.Description).To(Equal(core.StringPtr("The Cloud Object Storage rule.")))
				Expect(requiredConfigModel.And).To(Equal([]configmanagerv3.And{*andModel}))
				Expect(requiredConfigModel.Or).To(Equal([]configmanagerv3.Or{*orModel}))

				// Construct an instance of the Parameter model
				parameterModel := new(configmanagerv3.Parameter)
				Expect(parameterModel).ToNot(BeNil())
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")
				Expect(parameterModel.Name).To(Equal(core.StringPtr("hard_quota")))
				Expect(parameterModel.DisplayName).To(Equal(core.StringPtr("The Cloud Object Storage bucket quota.")))
				Expect(parameterModel.Description).To(Equal(core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")))
				Expect(parameterModel.Type).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the Import model
				importModel := new(configmanagerv3.Import)
				Expect(importModel).ToNot(BeNil())
				importModel.Parameters = []configmanagerv3.Parameter{*parameterModel}
				Expect(importModel.Parameters).To(Equal([]configmanagerv3.Parameter{*parameterModel}))

				// Construct an instance of the ReplaceRuleOptions model
				ruleID := "testString"
				ifMatch := "testString"
				replaceRuleOptionsAccountID := "130003ea8bfa43c5aacea07a86da3000"
				replaceRuleOptionsDescription := "Example rule"
				var replaceRuleOptionsTarget *configmanagerv3.Target = nil
				var replaceRuleOptionsRequiredConfig *configmanagerv3.RequiredConfig = nil
				replaceRuleOptionsLabels := []string{}
				replaceRuleOptionsModel := configManagerService.NewReplaceRuleOptions(ruleID, ifMatch, replaceRuleOptionsAccountID, replaceRuleOptionsDescription, replaceRuleOptionsTarget, replaceRuleOptionsRequiredConfig, replaceRuleOptionsLabels)
				replaceRuleOptionsModel.SetRuleID("testString")
				replaceRuleOptionsModel.SetIfMatch("testString")
				replaceRuleOptionsModel.SetAccountID("130003ea8bfa43c5aacea07a86da3000")
				replaceRuleOptionsModel.SetDescription("Example rule")
				replaceRuleOptionsModel.SetTarget(targetModel)
				replaceRuleOptionsModel.SetRequiredConfig(requiredConfigModel)
				replaceRuleOptionsModel.SetLabels([]string{})
				replaceRuleOptionsModel.SetType("user_defined")
				replaceRuleOptionsModel.SetVersion("1.0.1")
				replaceRuleOptionsModel.SetImport(importModel)
				replaceRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceRuleOptionsModel).ToNot(BeNil())
				Expect(replaceRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.AccountID).To(Equal(core.StringPtr("130003ea8bfa43c5aacea07a86da3000")))
				Expect(replaceRuleOptionsModel.Description).To(Equal(core.StringPtr("Example rule")))
				Expect(replaceRuleOptionsModel.Target).To(Equal(targetModel))
				Expect(replaceRuleOptionsModel.RequiredConfig).To(Equal(requiredConfigModel))
				Expect(replaceRuleOptionsModel.Labels).To(Equal([]string{}))
				Expect(replaceRuleOptionsModel.Type).To(Equal(core.StringPtr("user_defined")))
				Expect(replaceRuleOptionsModel.Version).To(Equal(core.StringPtr("1.0.1")))
				Expect(replaceRuleOptionsModel.Import).To(Equal(importModel))
				Expect(replaceRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTarget successfully`, func() {
				serviceName := "testString"
				resourceKind := "testString"
				_model, err := configManagerService.NewTarget(serviceName, resourceKind)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
