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

package adminserviceapiv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/cloud-go-sdk/adminserviceapiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the adminserviceapiv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AdminServiceApiV1 Integration Tests`, func() {
	const externalConfigFile = "../admin_service_api_v1.env"

	var (
		err          error
		adminServiceApiService *adminserviceapiv1.AdminServiceApiV1
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
			config, err = core.GetServiceProperties(adminserviceapiv1.DefaultServiceName)
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
			adminServiceApiServiceOptions := &adminserviceapiv1.AdminServiceApiV1Options{}

			adminServiceApiService, err = adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(adminServiceApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(adminServiceApiService).ToNot(BeNil())
			Expect(adminServiceApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			adminServiceApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSettings - Retrieves settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{
			}

			settings, response, err := adminServiceApiService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})
	})

	Describe(`UpdateSettings - Patch settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			jsonPatchOperationModel := &adminserviceapiv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateSettingsOptions := &adminserviceapiv1.UpdateSettingsOptions{
				Body: []adminserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			settings, response, err := adminServiceApiService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})
	})

	Describe(`PostTestEvent - Send test event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostTestEvent(postTestEventOptions *PostTestEventOptions)`, func() {
			postTestEventOptions := &adminserviceapiv1.PostTestEventOptions{
			}

			testEvent, response, err := adminServiceApiService.PostTestEvent(postTestEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testEvent).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
