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

package configmanagerapiv3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/configmanagerapiv3"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Config Manager API service.
//
// The following configuration properties are assumed to be defined:
// CONFIG_MANAGER_API_URL=<service base url>
// CONFIG_MANAGER_API_AUTH_TYPE=iam
// CONFIG_MANAGER_API_APIKEY=<IAM apikey>
// CONFIG_MANAGER_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`ConfigManagerApiV3 Examples Tests`, func() {

	const externalConfigFile = "../config_manager_api_v3.env"

	var (
		configManagerApiService *configmanagerapiv3.ConfigManagerApiV3
		config                  map[string]string
	)

	err := godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Error occured when loading external config file. Err: ", err)
	}

	accountInstanceId := os.Getenv("ACCOUNT_INSTANCE_ID")
	ruleId := os.Getenv("RULE_ID")

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
			config, err = core.GetServiceProperties(configmanagerapiv3.DefaultServiceName)
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

			configManagerApiServiceOptions := &configmanagerapiv3.ConfigManagerApiV3Options{}

			configManagerApiService, err = configmanagerapiv3.NewConfigManagerApiV3UsingExternalConfig(configManagerApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(configManagerApiService).ToNot(BeNil())
		})
	})

	Describe(`ConfigManagerApiV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

			createRuleOptions := configManagerApiService.NewCreateRuleOptions(
				accountInstanceId,
			)
			createRuleOptions.SetTypeQuery("system_defined")

			ruleArray := make([]configmanagerapiv3.Rule, 1)
			desc := "test-rule"
			ruleArray[0] = configmanagerapiv3.Rule{Description: &desc}
			createRuleOptions.SetRules(ruleArray)

			rules, response, err := configManagerApiService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rules, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rules).ToNot(BeNil())
		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules

			listRulesOptions := configManagerApiService.NewListRulesOptions(
				accountInstanceId,
			)
			listRulesOptions.SetTypeQuery("system_defined")

			rules, response, err := configManagerApiService.ListRules(listRulesOptions)
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
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := configManagerApiService.NewGetRuleOptions(
				ruleId,
				accountInstanceId,
			)

			rules, response, err := configManagerApiService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rules, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})
		It(`AddRule request example`, func() {
			fmt.Println("\nAddRule() result:")
			// begin-add_rule

			addRuleOptions := configManagerApiService.NewAddRuleOptions(
				ruleId,
				accountInstanceId,
			)

			rules, response, err := configManagerApiService.AddRule(addRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rules, "", "  ")
			fmt.Println(string(b))

			// end-add_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())
		})

		// Only run the following when the test rule is no longer needed

		//It(`DeleteRule request example`, func() {
		//	// begin-delete_rule
		//
		//	deleteRuleOptions := configManagerApiService.NewDeleteRuleOptions(
		//		ruleId,
		//		accountInstanceId,
		//	)
		//
		//	response, err := configManagerApiService.DeleteRule(deleteRuleOptions)
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
