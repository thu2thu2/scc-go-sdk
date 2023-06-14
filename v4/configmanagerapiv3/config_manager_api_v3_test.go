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

package configmanagerapiv3_test

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
	"github.com/IBM/scc-go-sdk/v4/configmanagerapiv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ConfigManagerApiV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(configManagerApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(configManagerApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
				URL: "https://configmanagerapiv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(configManagerApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIG_MANAGER_API_URL":       "https://configmanagerapiv3/api",
				"CONFIG_MANAGER_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(&configmanagerapiv3.ConfigManagerApiV3Options{})
				Expect(configManagerApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := configManagerApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configManagerApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configManagerApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configManagerApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL: "https://testService/api",
				})
				Expect(configManagerApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configManagerApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configManagerApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configManagerApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configManagerApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(&configmanagerapiv3.ConfigManagerApiV3Options{})
				err := configManagerApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configManagerApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configManagerApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configManagerApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configManagerApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIG_MANAGER_API_URL":       "https://configmanagerapiv3/api",
				"CONFIG_MANAGER_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(&configmanagerapiv3.ConfigManagerApiV3Options{})

			It(`Instantiate service client with error`, func() {
				Expect(configManagerApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIG_MANAGER_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(&configmanagerapiv3.ConfigManagerApiV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(configManagerApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = configmanagerapiv3.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us.compliance.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = configmanagerapiv3.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us.compliance.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = configmanagerapiv3.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://eu.compliance.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = configmanagerapiv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions) - Operation response error`, func() {
		createRulePath := "/instances/testString/v3/rules"
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
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerapiv3.CreateRuleOptions)
				createRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerApiService.EnableRetries(0, 0)
				result, response, operationErr = configManagerApiService.CreateRule(createRuleOptionsModel)
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
		createRulePath := "/instances/testString/v3/rules"
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke CreateRule successfully with retries`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())
				configManagerApiService.EnableRetries(0, 0)

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerapiv3.CreateRuleOptions)
				createRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerApiService.CreateRuleWithContext(ctx, createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerApiService.DisableRetries()
				result, response, operationErr := configManagerApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerApiService.CreateRuleWithContext(ctx, createRuleOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerApiService.CreateRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerapiv3.CreateRuleOptions)
				createRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRule with error: Operation validation and request error`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerapiv3.CreateRuleOptions)
				createRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRuleOptions model with no property values
				createRuleOptionsModelNew := new(configmanagerapiv3.CreateRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerApiService.CreateRule(createRuleOptionsModelNew)
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
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(configmanagerapiv3.CreateRuleOptions)
				createRuleOptionsModel.InstanceID = core.StringPtr("testString")
				createRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				createRuleOptionsModel.TypeQuery = core.StringPtr("system_defined")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerApiService.CreateRule(createRuleOptionsModel)
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
	Describe(`ListRules(listRulesOptions *ListRulesOptions) - Operation response error`, func() {
		listRulesPath := "/instances/testString/v3/rules"
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
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerapiv3.ListRulesOptions)
				listRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerApiService.EnableRetries(0, 0)
				result, response, operationErr = configManagerApiService.ListRules(listRulesOptionsModel)
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
		listRulesPath := "/instances/testString/v3/rules"
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke ListRules successfully with retries`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())
				configManagerApiService.EnableRetries(0, 0)

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerapiv3.ListRulesOptions)
				listRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerApiService.ListRulesWithContext(ctx, listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerApiService.DisableRetries()
				result, response, operationErr := configManagerApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerApiService.ListRulesWithContext(ctx, listRulesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke ListRules successfully`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerApiService.ListRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerapiv3.ListRulesOptions)
				listRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRules with error: Operation validation and request error`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerapiv3.ListRulesOptions)
				listRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRulesOptions model with no property values
				listRulesOptionsModelNew := new(configmanagerapiv3.ListRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerApiService.ListRules(listRulesOptionsModelNew)
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
			It(`Invoke ListRules successfully`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configmanagerapiv3.ListRulesOptions)
				listRulesOptionsModel.InstanceID = core.StringPtr("testString")
				listRulesOptionsModel.TypeQuery = core.StringPtr("system_defined")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerApiService.ListRules(listRulesOptionsModel)
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
		getRulePath := "/instances/testString/v3/rules/testString"
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
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerApiService.EnableRetries(0, 0)
				result, response, operationErr = configManagerApiService.GetRule(getRuleOptionsModel)
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
		getRulePath := "/instances/testString/v3/rules/testString"
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke GetRule successfully with retries`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())
				configManagerApiService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerApiService.GetRuleWithContext(ctx, getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerApiService.DisableRetries()
				result, response, operationErr := configManagerApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerApiService.GetRuleWithContext(ctx, getRuleOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke GetRule successfully`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerApiService.GetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRule with error: Operation validation and request error`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleOptions model with no property values
				getRuleOptionsModelNew := new(configmanagerapiv3.GetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerApiService.GetRule(getRuleOptionsModelNew)
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
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configmanagerapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.InstanceID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerApiService.GetRule(getRuleOptionsModel)
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
	Describe(`AddRule(addRuleOptions *AddRuleOptions) - Operation response error`, func() {
		addRulePath := "/instances/testString/v3/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddRule with error: Operation response processing error`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the AddRuleOptions model
				addRuleOptionsModel := new(configmanagerapiv3.AddRuleOptions)
				addRuleOptionsModel.RuleID = core.StringPtr("testString")
				addRuleOptionsModel.InstanceID = core.StringPtr("testString")
				addRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				addRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configManagerApiService.AddRule(addRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configManagerApiService.EnableRetries(0, 0)
				result, response, operationErr = configManagerApiService.AddRule(addRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddRule(addRuleOptions *AddRuleOptions)`, func() {
		addRulePath := "/instances/testString/v3/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addRulePath))
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke AddRule successfully with retries`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())
				configManagerApiService.EnableRetries(0, 0)

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the AddRuleOptions model
				addRuleOptionsModel := new(configmanagerapiv3.AddRuleOptions)
				addRuleOptionsModel.RuleID = core.StringPtr("testString")
				addRuleOptionsModel.InstanceID = core.StringPtr("testString")
				addRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				addRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configManagerApiService.AddRuleWithContext(ctx, addRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configManagerApiService.DisableRetries()
				result, response, operationErr := configManagerApiService.AddRule(addRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configManagerApiService.AddRuleWithContext(ctx, addRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addRulePath))
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
					fmt.Fprintf(res, "%s", `{"rules": [{"creation_date": "CreationDate", "created_by": "CreatedBy", "modification_date": "ModificationDate", "modified_by": "ModifiedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"imports": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": ["AdditionalTargetAttributes"]}, "required_config": {"and": [{"property": "Property", "operator": "Operator"}], "or": [{"property": "Property", "operator": "Operator"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke AddRule successfully`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configManagerApiService.AddRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the AddRuleOptions model
				addRuleOptionsModel := new(configmanagerapiv3.AddRuleOptions)
				addRuleOptionsModel.RuleID = core.StringPtr("testString")
				addRuleOptionsModel.InstanceID = core.StringPtr("testString")
				addRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				addRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configManagerApiService.AddRule(addRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddRule with error: Operation validation and request error`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the AddRuleOptions model
				addRuleOptionsModel := new(configmanagerapiv3.AddRuleOptions)
				addRuleOptionsModel.RuleID = core.StringPtr("testString")
				addRuleOptionsModel.InstanceID = core.StringPtr("testString")
				addRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				addRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configManagerApiService.AddRule(addRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddRuleOptions model with no property values
				addRuleOptionsModelNew := new(configmanagerapiv3.AddRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configManagerApiService.AddRule(addRuleOptionsModelNew)
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
			It(`Invoke AddRule successfully`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}

				// Construct an instance of the AddRuleOptions model
				addRuleOptionsModel := new(configmanagerapiv3.AddRuleOptions)
				addRuleOptionsModel.RuleID = core.StringPtr("testString")
				addRuleOptionsModel.InstanceID = core.StringPtr("testString")
				addRuleOptionsModel.Rules = []configmanagerapiv3.Rule{*ruleModel}
				addRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configManagerApiService.AddRule(addRuleOptionsModel)
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
		deleteRulePath := "/instances/testString/v3/rules/testString"
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
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := configManagerApiService.DeleteRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configmanagerapiv3.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.InstanceID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = configManagerApiService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRule with error: Operation validation and request error`, func() {
				configManagerApiService, serviceErr := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configManagerApiService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configmanagerapiv3.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.InstanceID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configManagerApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := configManagerApiService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleOptions model with no property values
				deleteRuleOptionsModelNew := new(configmanagerapiv3.DeleteRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = configManagerApiService.DeleteRule(deleteRuleOptionsModelNew)
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
			configManagerApiService, _ := configmanagerapiv3.NewConfigManagerApiV3(&configmanagerapiv3.ConfigManagerApiV3Options{
				URL:           "http://configmanagerapiv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddRuleOptions successfully`, func() {
				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				Expect(importModel).ToNot(BeNil())
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")
				Expect(importModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(importModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(importModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(importModel.Type).To(Equal(core.StringPtr("string")))

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				Expect(parametersModel).ToNot(BeNil())
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}
				Expect(parametersModel.Imports).To(Equal([]configmanagerapiv3.Import{*importModel}))

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}
				Expect(targetModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.ServiceDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.ResourceKind).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.AdditionalTargetAttributes).To(Equal([]string{"testString"}))

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				Expect(andModel).ToNot(BeNil())
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")
				Expect(andModel.Property).To(Equal(core.StringPtr("testString")))
				Expect(andModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				Expect(orModel).ToNot(BeNil())
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")
				Expect(orModel.Property).To(Equal(core.StringPtr("testString")))
				Expect(orModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				Expect(requiredConfigModel).ToNot(BeNil())
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}
				Expect(requiredConfigModel.And).To(Equal([]configmanagerapiv3.And{*andModel}))
				Expect(requiredConfigModel.Or).To(Equal([]configmanagerapiv3.Or{*orModel}))

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}
				Expect(ruleModel.CreationDate).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.ModificationDate).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.ModifiedBy).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Type).To(Equal(core.StringPtr("user_defined")))
				Expect(ruleModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Import).To(Equal(parametersModel))
				Expect(ruleModel.Target).To(Equal(targetModel))
				Expect(ruleModel.RequiredConfig).To(Equal(requiredConfigModel))
				Expect(ruleModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the AddRuleOptions model
				ruleID := "testString"
				instanceID := "testString"
				addRuleOptionsModel := configManagerApiService.NewAddRuleOptions(ruleID, instanceID)
				addRuleOptionsModel.SetRuleID("testString")
				addRuleOptionsModel.SetInstanceID("testString")
				addRuleOptionsModel.SetRules([]configmanagerapiv3.Rule{*ruleModel})
				addRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addRuleOptionsModel).ToNot(BeNil())
				Expect(addRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(addRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(addRuleOptionsModel.Rules).To(Equal([]configmanagerapiv3.Rule{*ruleModel}))
				Expect(addRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRuleOptions successfully`, func() {
				// Construct an instance of the Import model
				importModel := new(configmanagerapiv3.Import)
				Expect(importModel).ToNot(BeNil())
				importModel.Name = core.StringPtr("testString")
				importModel.DisplayName = core.StringPtr("testString")
				importModel.Description = core.StringPtr("testString")
				importModel.Type = core.StringPtr("string")
				Expect(importModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(importModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(importModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(importModel.Type).To(Equal(core.StringPtr("string")))

				// Construct an instance of the Parameters model
				parametersModel := new(configmanagerapiv3.Parameters)
				Expect(parametersModel).ToNot(BeNil())
				parametersModel.Imports = []configmanagerapiv3.Import{*importModel}
				Expect(parametersModel.Imports).To(Equal([]configmanagerapiv3.Import{*importModel}))

				// Construct an instance of the Target model
				targetModel := new(configmanagerapiv3.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.ServiceName = core.StringPtr("testString")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("testString")
				targetModel.AdditionalTargetAttributes = []string{"testString"}
				Expect(targetModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.ServiceDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.ResourceKind).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.AdditionalTargetAttributes).To(Equal([]string{"testString"}))

				// Construct an instance of the And model
				andModel := new(configmanagerapiv3.And)
				Expect(andModel).ToNot(BeNil())
				andModel.Property = core.StringPtr("testString")
				andModel.Operator = core.StringPtr("testString")
				Expect(andModel.Property).To(Equal(core.StringPtr("testString")))
				Expect(andModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Or model
				orModel := new(configmanagerapiv3.Or)
				Expect(orModel).ToNot(BeNil())
				orModel.Property = core.StringPtr("testString")
				orModel.Operator = core.StringPtr("testString")
				Expect(orModel.Property).To(Equal(core.StringPtr("testString")))
				Expect(orModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RequiredConfig model
				requiredConfigModel := new(configmanagerapiv3.RequiredConfig)
				Expect(requiredConfigModel).ToNot(BeNil())
				requiredConfigModel.And = []configmanagerapiv3.And{*andModel}
				requiredConfigModel.Or = []configmanagerapiv3.Or{*orModel}
				Expect(requiredConfigModel.And).To(Equal([]configmanagerapiv3.And{*andModel}))
				Expect(requiredConfigModel.Or).To(Equal([]configmanagerapiv3.Or{*orModel}))

				// Construct an instance of the Rule model
				ruleModel := new(configmanagerapiv3.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.CreationDate = core.StringPtr("testString")
				ruleModel.CreatedBy = core.StringPtr("testString")
				ruleModel.ModificationDate = core.StringPtr("testString")
				ruleModel.ModifiedBy = core.StringPtr("testString")
				ruleModel.ID = core.StringPtr("testString")
				ruleModel.AccountID = core.StringPtr("testString")
				ruleModel.Description = core.StringPtr("testString")
				ruleModel.Type = core.StringPtr("user_defined")
				ruleModel.Version = core.StringPtr("testString")
				ruleModel.Import = parametersModel
				ruleModel.Target = targetModel
				ruleModel.RequiredConfig = requiredConfigModel
				ruleModel.Labels = []string{"testString"}
				Expect(ruleModel.CreationDate).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.ModificationDate).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.ModifiedBy).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Type).To(Equal(core.StringPtr("user_defined")))
				Expect(ruleModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Import).To(Equal(parametersModel))
				Expect(ruleModel.Target).To(Equal(targetModel))
				Expect(ruleModel.RequiredConfig).To(Equal(requiredConfigModel))
				Expect(ruleModel.Labels).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateRuleOptions model
				instanceID := "testString"
				createRuleOptionsModel := configManagerApiService.NewCreateRuleOptions(instanceID)
				createRuleOptionsModel.SetInstanceID("testString")
				createRuleOptionsModel.SetRules([]configmanagerapiv3.Rule{*ruleModel})
				createRuleOptionsModel.SetTypeQuery("system_defined")
				createRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRuleOptionsModel).ToNot(BeNil())
				Expect(createRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleOptionsModel.Rules).To(Equal([]configmanagerapiv3.Rule{*ruleModel}))
				Expect(createRuleOptionsModel.TypeQuery).To(Equal(core.StringPtr("system_defined")))
				Expect(createRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleOptions successfully`, func() {
				// Construct an instance of the DeleteRuleOptions model
				ruleID := "testString"
				instanceID := "testString"
				deleteRuleOptionsModel := configManagerApiService.NewDeleteRuleOptions(ruleID, instanceID)
				deleteRuleOptionsModel.SetRuleID("testString")
				deleteRuleOptionsModel.SetInstanceID("testString")
				deleteRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleOptionsModel).ToNot(BeNil())
				Expect(deleteRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRuleOptions successfully`, func() {
				// Construct an instance of the GetRuleOptions model
				ruleID := "testString"
				instanceID := "testString"
				getRuleOptionsModel := configManagerApiService.NewGetRuleOptions(ruleID, instanceID)
				getRuleOptionsModel.SetRuleID("testString")
				getRuleOptionsModel.SetInstanceID("testString")
				getRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleOptionsModel).ToNot(BeNil())
				Expect(getRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRulesOptions successfully`, func() {
				// Construct an instance of the ListRulesOptions model
				instanceID := "testString"
				listRulesOptionsModel := configManagerApiService.NewListRulesOptions(instanceID)
				listRulesOptionsModel.SetInstanceID("testString")
				listRulesOptionsModel.SetTypeQuery("system_defined")
				listRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRulesOptionsModel).ToNot(BeNil())
				Expect(listRulesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.TypeQuery).To(Equal(core.StringPtr("system_defined")))
				Expect(listRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRule successfully`, func() {
				description := "testString"
				_model, err := configManagerApiService.NewRule(description)
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
