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

package adminserviceapiv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/cloud-go-sdk/adminserviceapiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Admin Service API service.
//
// The following configuration properties are assumed to be defined:
// ADMIN_SERVICE_API_URL=<service base url>
// ADMIN_SERVICE_API_AUTH_TYPE=iam
// ADMIN_SERVICE_API_APIKEY=<IAM apikey>
// ADMIN_SERVICE_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`AdminServiceApiV1 Examples Tests`, func() {

	const externalConfigFile = "../admin_service_api_v1.env"

	var (
		adminServiceApiService *adminserviceapiv1.AdminServiceApiV1
		config       map[string]string
	)

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
			config, err = core.GetServiceProperties(adminserviceapiv1.DefaultServiceName)
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

			adminServiceApiServiceOptions := &adminserviceapiv1.AdminServiceApiV1Options{}

			adminServiceApiService, err = adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(adminServiceApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(adminServiceApiService).ToNot(BeNil())
		})
	})

	Describe(`AdminServiceApiV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := adminServiceApiService.NewGetSettingsOptions()

			settings, response, err := adminServiceApiService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})
		It(`UpdateSettings request example`, func() {
			fmt.Println("\nUpdateSettings() result:")
			// begin-update_settings

			jsonPatchOperationModel := &adminserviceapiv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateSettingsOptions := adminServiceApiService.NewUpdateSettingsOptions(
				[]adminserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			settings, response, err := adminServiceApiService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})
		It(`PostTestEvent request example`, func() {
			fmt.Println("\nPostTestEvent() result:")
			// begin-post_test_event

			postTestEventOptions := adminServiceApiService.NewPostTestEventOptions()

			testEvent, response, err := adminServiceApiService.PostTestEvent(postTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testEvent).ToNot(BeNil())
		})
	})
})
