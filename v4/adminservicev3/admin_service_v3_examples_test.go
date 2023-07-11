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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Admin Service service.
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
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := adminServiceService.NewGetSettingsOptions()

			settings, response, err := adminServiceService.GetSettings(getSettingsOptions)
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

			jsonPatchOperationModel := &adminservicev3.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateSettingsOptions := adminServiceService.NewUpdateSettingsOptions(
				[]adminservicev3.JSONPatchOperation{*jsonPatchOperationModel},
			)

			settings, response, err := adminServiceService.UpdateSettings(updateSettingsOptions)
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

			postTestEventOptions := adminServiceService.NewPostTestEventOptions()

			testEvent, response, err := adminServiceService.PostTestEvent(postTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})
	})
})
