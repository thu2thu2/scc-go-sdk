//go:build integration
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

package configmanagerv3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/configmanagerv3"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the configmanagerv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ConfigManagerV3 Integration Tests`, func() {
	const externalConfigFile = "../config_manager_v3.env"

	var (
		err                  error
		configManagerService *configmanagerv3.ConfigManagerV3
		serviceURL           string
		config               map[string]string
	)

	err = godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Error occured when loading external config file. Err: ", err)
	}

	ruleId := os.Getenv("RULE_ID")
	//homeAccountId := os.Getenv("Home_Account_ID")

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
			config, err = core.GetServiceProperties(configmanagerv3.DefaultServiceName)
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
			configManagerServiceOptions := &configmanagerv3.ConfigManagerV3Options{}

			configManagerService, err = configmanagerv3.NewConfigManagerV3UsingExternalConfig(configManagerServiceOptions)
			Expect(err).To(BeNil())
			Expect(configManagerService).ToNot(BeNil())
			Expect(configManagerService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			configManagerService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListRules - List all rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
			listRulesOptions := &configmanagerv3.ListRulesOptions{
				TypeQuery: core.StringPtr("system_defined"),
			}

			rules, response, err := configManagerService.ListRules(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})
	})

	Describe(`CreateRule - Create a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
			//additionalTargetAttributeModel := &configmanagerv3.AdditionalTargetAttribute{
			//	Name:     core.StringPtr("testString"),
			//	Operator: core.StringPtr("string_equals"),
			//	Value:    core.StringPtr("testString"),
			//}

			targetModel := &configmanagerv3.Target{
				ServiceName:        core.StringPtr("cloud-object-storage"),
				ServiceDisplayName: core.StringPtr("testString"),
				ResourceKind:       core.StringPtr("bucket"),
				//AdditionalTargetAttributes: []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			andModel := &configmanagerv3.And{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			orModel := &configmanagerv3.Or{
				Property: core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("testString"),
			}

			requiredConfigModel := &configmanagerv3.RequiredConfig{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []configmanagerv3.And{*andModel},
				Or:          []configmanagerv3.Or{*orModel},
			}

			parameterModel := &configmanagerv3.Parameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &configmanagerv3.Import{
				Parameters: []configmanagerv3.Parameter{*parameterModel},
			}

			createRuleOptions := &configmanagerv3.CreateRuleOptions{
				AccountID:      core.StringPtr("130003ea8bfa43c5aacea07a86da3000"),
				Description:    core.StringPtr("Example rule"),
				Target:         targetModel,
				RequiredConfig: requiredConfigModel,
				Labels:         []string{},
				Type:           core.StringPtr("user_defined"),
				Version:        core.StringPtr("1.0.0"),
				Import:         importModel,
				TypeQuery:      core.StringPtr("system_defined"),
			}

			rule, response, err := configManagerService.CreateRule(createRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())
		})
	})

	Describe(`GetRule - Get a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &configmanagerv3.GetRuleOptions{
				RuleID: core.StringPtr(ruleId),
			}

			rule, response, err := configManagerService.GetRule(getRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
	})

	Describe(`ReplaceRule - Update a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
			additionalTargetAttributeModel := &configmanagerv3.AdditionalTargetAttribute{
				Name:     core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("testString"),
			}

			targetModel := &configmanagerv3.Target{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ServiceDisplayName:         core.StringPtr("Cloud Object Storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			andModel := &configmanagerv3.And{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			orModel := &configmanagerv3.Or{
				Property: core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("testString"),
			}

			requiredConfigModel := &configmanagerv3.RequiredConfig{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []configmanagerv3.And{*andModel},
				Or:          []configmanagerv3.Or{*orModel},
			}

			parameterModel := &configmanagerv3.Parameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &configmanagerv3.Import{
				Parameters: []configmanagerv3.Parameter{*parameterModel},
			}

			replaceRuleOptions := &configmanagerv3.ReplaceRuleOptions{
				RuleID:         core.StringPtr(ruleId),
				IfMatch:        core.StringPtr("testString"),
				AccountID:      core.StringPtr("130003ea8bfa43c5aacea07a86da3000"),
				Description:    core.StringPtr("Example rule"),
				Target:         targetModel,
				RequiredConfig: requiredConfigModel,
				Labels:         []string{},
				Type:           core.StringPtr("user_defined"),
				Version:        core.StringPtr("1.0.1"),
				Import:         importModel,
			}

			rule, response, err := configManagerService.ReplaceRule(replaceRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
	})

	//Describe(`DeleteRule - Delete a custom rule`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
	//		deleteRuleOptions := &configmanagerv3.DeleteRuleOptions{
	//			RuleID: core.StringPtr(ruleId),
	//		}
	//
	//		response, err := configManagerService.DeleteRule(deleteRuleOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(204))
	//	})
	//})
})

//
// Utility functions are declared in the unit test file
//
