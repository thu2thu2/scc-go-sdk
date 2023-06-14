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

package adminservicev3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/adminservicev3"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the admin-service service.
//
// The following configuration properties are assumed to be defined:
// ADMIN_SERVICE_URL=<service base url>
// ADMIN_SERVICE_AUTH_TYPE=iam
// ADMIN_SERVICE_APIKEY=<IAM apikey>
// ADMIN_SERVICE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`AdminServiceV3 Examples Tests`, func() {

	const externalConfigFile = "../admin_service_v3.env"

	var (
		adminServiceService *adminservicev3.AdminServiceV3
		config              map[string]string
	)

	err := godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Error occured when loading external config file. Err: ", err)
	}

	accountInstanceId := os.Getenv("ACCOUNT_INSTANCE_ID")

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
			config, err = core.GetServiceProperties(adminservicev3.DefaultServiceName)
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

			adminServiceServiceOptions := &adminservicev3.AdminServiceV3Options{}

			adminServiceService, err = adminservicev3.NewAdminServiceV3UsingExternalConfig(adminServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(adminServiceService).ToNot(BeNil())
		})
	})

	Describe(`AdminServiceV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceAccess request example`, func() {
			fmt.Println("\nGetInstanceAccess() result:")
			// begin-get_instance_access

			getInstanceAccessOptions := adminServiceService.NewGetInstanceAccessOptions(
				accountInstanceId,
			)

			instanceAccess, response, err := adminServiceService.GetInstanceAccess(getInstanceAccessOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceAccess, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_access

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceAccess).ToNot(BeNil())
		})
		It(`GetInstanceSettings request example`, func() {
			fmt.Println("\nGetInstanceSettings() result:")
			// begin-get_instance_settings

			getInstanceSettingsOptions := adminServiceService.NewGetInstanceSettingsOptions(
				accountInstanceId,
			)

			instanceSettings, response, err := adminServiceService.GetInstanceSettings(getInstanceSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceSettings).ToNot(BeNil())
		})
		It(`UpdateInstanceSettings request example`, func() {
			fmt.Println("\nUpdateInstanceSettings() result:")
			// begin-update_instance_settings

			jsonPatchOperationModel := &adminservicev3.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateInstanceSettingsOptions := adminServiceService.NewUpdateInstanceSettingsOptions(
				accountInstanceId,
				[]adminservicev3.JSONPatchOperation{*jsonPatchOperationModel},
			)

			instanceSettings, response, err := adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceSettings).ToNot(BeNil())
		})
		It(`PostInstanceTestEvent request example`, func() {
			fmt.Println("\nPostInstanceTestEvent() result:")
			// begin-post_instance_test_event

			postInstanceTestEventOptions := adminServiceService.NewPostInstanceTestEventOptions(
				accountInstanceId,
			)

			testEvent, response, err := adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-post_instance_test_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testEvent).ToNot(BeNil())
		})
		//It(`GetInstancePlans request example`, func() {
		//	fmt.Println("\nGetInstancePlans() result:")
		//	// begin-get_instance_plans
		//
		//	getInstancePlansOptions := adminServiceService.NewGetInstancePlansOptions(
		//		accountInstanceId,
		//	)
		//
		//	instancePlan, response, err := adminServiceService.GetInstancePlans(getInstancePlansOptions)
		//	if err != nil {
		//		panic(err)
		//	}
		//	b, _ := json.MarshalIndent(instancePlan, "", "  ")
		//	fmt.Println(string(b))
		//
		//	// end-get_instance_plans
		//
		//	Expect(err).To(BeNil())
		//	Expect(response.StatusCode).To(Equal(200))
		//	Expect(instancePlan).ToNot(BeNil())
		//})
		//It(`PostInstancePlans request example`, func() {
		//	fmt.Println("\nPostInstancePlans() result:")
		//	// begin-post_instance_plans
		//
		//	postInstancePlansOptions := adminServiceService.NewPostInstancePlansOptions(
		//		accountInstanceId,
		//		"Standard",
		//	)
		//
		//	instancePlan, response, err := adminServiceService.PostInstancePlans(postInstancePlansOptions)
		//	if err != nil {
		//		panic(err)
		//	}
		//	b, _ := json.MarshalIndent(instancePlan, "", "  ")
		//	fmt.Println(string(b))
		//
		//	// end-post_instance_plans
		//
		//	Expect(err).To(BeNil())
		//	Expect(response.StatusCode).To(Equal(200))
		//	Expect(instancePlan).ToNot(BeNil())
		//})
		//It(`SetInstancePlans request example`, func() {
		//	fmt.Println("\nSetInstancePlans() result:")
		//	// begin-set_instance_plans
		//
		//	setInstancePlansOptions := adminServiceService.NewSetInstancePlansOptions(
		//		accountInstanceId,
		//		"Standard",
		//	)
		//
		//	instancePlan, response, err := adminServiceService.SetInstancePlans(setInstancePlansOptions)
		//	if err != nil {
		//		panic(err)
		//	}
		//	b, _ := json.MarshalIndent(instancePlan, "", "  ")
		//	fmt.Println(string(b))
		//
		//	// end-set_instance_plans
		//
		//	Expect(err).To(BeNil())
		//	Expect(response.StatusCode).To(Equal(200))
		//	Expect(instancePlan).ToNot(BeNil())
		//})
	})
})
