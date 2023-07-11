//go:build examples
// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/configmanagerv3"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Config Manager service.
//
// The following configuration properties are assumed to be defined:
// CONFIG_MANAGER_URL=<service base url>
// CONFIG_MANAGER_AUTH_TYPE=iam
// CONFIG_MANAGER_APIKEY=<IAM apikey>
// CONFIG_MANAGER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`ConfigManagerV3 Examples Tests`, func() {

	const externalConfigFile = "../config_manager_v3.env"

	var (
		configManagerService *configmanagerv3.ConfigManagerV3
		config               map[string]string
	)

	err := godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Error occured when loading external config file. Err: ", err)
	}

	ruleId := os.Getenv("RULE_ID")
	//homeAccountId := os.Getenv("Home_Account_ID")

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(configmanagerv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			configManagerServiceOptions := &configmanagerv3.ConfigManagerV3Options{}

			configManagerService, err = configmanagerv3.NewConfigManagerV3UsingExternalConfig(configManagerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(configManagerService).ToNot(BeNil())
		})
	})

	Describe(`ConfigManagerV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules

			listRulesOptions := configManagerService.NewListRulesOptions()
			listRulesOptions.SetTypeQuery("system_defined")

			rules, response, err := configManagerService.ListRules(listRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rules, "", "  ")
			fmt.Println(string(b))

			// end-list_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

			//additionalTargetAttributeModel := &configmanagerv3.AdditionalTargetAttribute{}

			targetModel := &configmanagerv3.Target{
				ServiceName:  core.StringPtr("cloud-object-storage"),
				ResourceKind: core.StringPtr("bucket"),
				//AdditionalTargetAttributes: []configmanagerv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			andModel := &configmanagerv3.And{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &configmanagerv3.RequiredConfig{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []configmanagerv3.And{*andModel},
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

			createRuleOptions := configManagerService.NewCreateRuleOptions(
				"130003ea8bfa43c5aacea07a86da3000",
				"Example rule",
				targetModel,
				requiredConfigModel,
				[]string{},
			)
			createRuleOptions.SetVersion("1.0.0")
			createRuleOptions.SetImport(importModel)
			createRuleOptions.SetTypeQuery("system_defined")

			rule, response, err := configManagerService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())
		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := configManagerService.NewGetRuleOptions(
				ruleId,
			)

			rule, response, err := configManagerService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
		It(`ReplaceRule request example`, func() {
			fmt.Println("\nReplaceRule() result:")
			// begin-replace_rule

			additionalTargetAttributeModel := &configmanagerv3.AdditionalTargetAttribute{}

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

			requiredConfigModel := &configmanagerv3.RequiredConfig{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []configmanagerv3.And{*andModel},
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

			replaceRuleOptions := configManagerService.NewReplaceRuleOptions(
				ruleId,
				"testString",
				"130003ea8bfa43c5aacea07a86da3000",
				"Example rule",
				targetModel,
				requiredConfigModel,
				[]string{},
			)
			replaceRuleOptions.SetType("user_defined")
			replaceRuleOptions.SetVersion("1.0.1")
			replaceRuleOptions.SetImport(importModel)

			rule, response, err := configManagerService.ReplaceRule(replaceRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-replace_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
		//It(`DeleteRule request example`, func() {
		//	// begin-delete_rule
		//
		//	deleteRuleOptions := configManagerService.NewDeleteRuleOptions(
		//		"testString",
		//	)
		//
		//	response, err := configManagerService.DeleteRule(deleteRuleOptions)
		//	if err != nil {
		//		panic(err)
		//	}
		//	if response.StatusCode != 204 {
		//		fmt.Printf("\nUnexpected response status code received from DeleteRule(): %d\n", response.StatusCode)
		//	}
		//
		//	// end-delete_rule
		//
		//	Expect(err).To(BeNil())
		//	Expect(response.StatusCode).To(Equal(204))
		//})
	})
})
