// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/cloud-go-sdk/configmanagerapiv3"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the configmanagerapiv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ConfigManagerApiV3 Integration Tests`, func() {
	const externalConfigFile = "../config_manager_api_v3.env"

	var (
		err          error
		configManagerApiService *configmanagerapiv3.ConfigManagerApiV3
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(configmanagerapiv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			configManagerApiServiceOptions := &configmanagerapiv3.ConfigManagerApiV3Options{}

			configManagerApiService, err = configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(configManagerApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(configManagerApiService).ToNot(BeNil())
			Expect(configManagerApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			configManagerApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateRule - Create a user defined rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
			importModel := &configmanagerapiv3.Import{
				Name: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Type: core.StringPtr("string"),
			}

			parametersModel := &configmanagerapiv3.Parameters{
				Imports: []configmanagerapiv3.Import{*importModel},
			}

			targetModel := &configmanagerapiv3.Target{
				ServiceName: core.StringPtr("testString"),
				ServiceDisplayName: core.StringPtr("testString"),
				ResourceKind: core.StringPtr("testString"),
				AdditionalTargetAttributes: []string{"testString"},
			}

			andModel := &configmanagerapiv3.And{
				Property: core.StringPtr("testString"),
				Operator: core.StringPtr("testString"),
			}

			orModel := &configmanagerapiv3.Or{
				Property: core.StringPtr("testString"),
				Operator: core.StringPtr("testString"),
			}

			requiredConfigModel := &configmanagerapiv3.RequiredConfig{
				And: []configmanagerapiv3.And{*andModel},
				Or: []configmanagerapiv3.Or{*orModel},
			}

			ruleModel := &configmanagerapiv3.Rule{
				CreationDate: core.StringPtr("testString"),
				CreatedBy: core.StringPtr("testString"),
				ModificationDate: core.StringPtr("testString"),
				ModifiedBy: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AccountID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Type: core.StringPtr("user_defined"),
				Version: core.StringPtr("testString"),
				Import: parametersModel,
				Target: targetModel,
				RequiredConfig: requiredConfigModel,
				Labels: []string{"testString"},
			}

			createRuleOptions := &configmanagerapiv3.CreateRuleOptions{
				InstanceID: core.StringPtr("testString"),
				Rules: []configmanagerapiv3.Rule{*ruleModel},
				TypeQuery: core.StringPtr("system_defined"),
			}

			rules, response, err := configManagerApiService.CreateRule(createRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rules).ToNot(BeNil())
		})
	})

	Describe(`ListRules - Retrieve all rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
			listRulesOptions := &configmanagerapiv3.ListRulesOptions{
				InstanceID: core.StringPtr("testString"),
				TypeQuery: core.StringPtr("system_defined"),
			}

			rules, response, err := configManagerApiService.ListRules(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})
	})

	Describe(`GetRule - Retrieve a specific user defined rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &configmanagerapiv3.GetRuleOptions{
				RuleID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
			}

			rules, response, err := configManagerApiService.GetRule(getRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})
	})

	Describe(`AddRule - Update a specific user defined rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddRule(addRuleOptions *AddRuleOptions)`, func() {
			importModel := &configmanagerapiv3.Import{
				Name: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Type: core.StringPtr("string"),
			}

			parametersModel := &configmanagerapiv3.Parameters{
				Imports: []configmanagerapiv3.Import{*importModel},
			}

			targetModel := &configmanagerapiv3.Target{
				ServiceName: core.StringPtr("testString"),
				ServiceDisplayName: core.StringPtr("testString"),
				ResourceKind: core.StringPtr("testString"),
				AdditionalTargetAttributes: []string{"testString"},
			}

			andModel := &configmanagerapiv3.And{
				Property: core.StringPtr("testString"),
				Operator: core.StringPtr("testString"),
			}

			orModel := &configmanagerapiv3.Or{
				Property: core.StringPtr("testString"),
				Operator: core.StringPtr("testString"),
			}

			requiredConfigModel := &configmanagerapiv3.RequiredConfig{
				And: []configmanagerapiv3.And{*andModel},
				Or: []configmanagerapiv3.Or{*orModel},
			}

			ruleModel := &configmanagerapiv3.Rule{
				CreationDate: core.StringPtr("testString"),
				CreatedBy: core.StringPtr("testString"),
				ModificationDate: core.StringPtr("testString"),
				ModifiedBy: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AccountID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Type: core.StringPtr("user_defined"),
				Version: core.StringPtr("testString"),
				Import: parametersModel,
				Target: targetModel,
				RequiredConfig: requiredConfigModel,
				Labels: []string{"testString"},
			}

			addRuleOptions := &configmanagerapiv3.AddRuleOptions{
				RuleID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				Rules: []configmanagerapiv3.Rule{*ruleModel},
			}

			rules, response, err := configManagerApiService.AddRule(addRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})
	})

	Describe(`DeleteRule - Delete a specific user defined rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
			deleteRuleOptions := &configmanagerapiv3.DeleteRuleOptions{
				RuleID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
			}

			response, err := configManagerApiService.DeleteRule(deleteRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
