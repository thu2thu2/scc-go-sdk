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

package adminservicev3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/adminservicev3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the adminservicev3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AdminServiceV3 Integration Tests`, func() {
	const externalConfigFile = "../admin_service_v3.env"

	var (
		err                 error
		adminServiceService *adminservicev3.AdminServiceV3
		serviceURL          string
		config              map[string]string
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
			config, err = core.GetServiceProperties(adminservicev3.DefaultServiceName)
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
			adminServiceServiceOptions := &adminservicev3.AdminServiceV3Options{}

			adminServiceService, err = adminservicev3.NewAdminServiceV3UsingExternalConfig(adminServiceServiceOptions)
			Expect(err).To(BeNil())
			Expect(adminServiceService).ToNot(BeNil())
			Expect(adminServiceService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			adminServiceService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetInstanceAccess - Get access level for an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceAccess(getInstanceAccessOptions *GetInstanceAccessOptions)`, func() {
			getInstanceAccessOptions := &adminservicev3.GetInstanceAccessOptions{
				InstanceID: core.StringPtr("testString"),
			}

			instanceAccess, response, err := adminServiceService.GetInstanceAccess(getInstanceAccessOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceAccess).ToNot(BeNil())
		})
	})

	Describe(`GetInstanceSettings - Retrieves instance settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceSettings(getInstanceSettingsOptions *GetInstanceSettingsOptions)`, func() {
			getInstanceSettingsOptions := &adminservicev3.GetInstanceSettingsOptions{
				InstanceID: core.StringPtr("testString"),
			}

			instanceSettings, response, err := adminServiceService.GetInstanceSettings(getInstanceSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceSettings).ToNot(BeNil())
		})
	})

	Describe(`UpdateInstanceSettings - Patch instance settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceSettings(updateInstanceSettingsOptions *UpdateInstanceSettingsOptions)`, func() {
			jsonPatchOperationModel := &adminservicev3.JSONPatchOperation{
				Op:    core.StringPtr("add"),
				Path:  core.StringPtr("testString"),
				From:  core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateInstanceSettingsOptions := &adminservicev3.UpdateInstanceSettingsOptions{
				InstanceID:         core.StringPtr("testString"),
				JSONPatchOperation: []adminservicev3.JSONPatchOperation{*jsonPatchOperationModel},
			}

			instanceSettings, response, err := adminServiceService.UpdateInstanceSettings(updateInstanceSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceSettings).ToNot(BeNil())
		})
	})

	Describe(`PostInstanceTestEvent - Instnace send test event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostInstanceTestEvent(postInstanceTestEventOptions *PostInstanceTestEventOptions)`, func() {
			postInstanceTestEventOptions := &adminservicev3.PostInstanceTestEventOptions{
				InstanceID: core.StringPtr("testString"),
			}

			testEvent, response, err := adminServiceService.PostInstanceTestEvent(postInstanceTestEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testEvent).ToNot(BeNil())
		})
	})

	Describe(`GetInstancePlans - Retrieves instance plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstancePlans(getInstancePlansOptions *GetInstancePlansOptions)`, func() {
			getInstancePlansOptions := &adminservicev3.GetInstancePlansOptions{
				InstanceID: core.StringPtr("testString"),
			}

			instancePlan, response, err := adminServiceService.GetInstancePlans(getInstancePlansOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instancePlan).ToNot(BeNil())
		})
	})

	Describe(`PostInstancePlans - Create instance plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostInstancePlans(postInstancePlansOptions *PostInstancePlansOptions)`, func() {
			postInstancePlansOptions := &adminservicev3.PostInstancePlansOptions{
				InstanceID: core.StringPtr("testString"),
				Name:       core.StringPtr("Standard"),
			}

			instancePlan, response, err := adminServiceService.PostInstancePlans(postInstancePlansOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instancePlan).ToNot(BeNil())
		})
	})

	Describe(`SetInstancePlans - Update instance plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetInstancePlans(setInstancePlansOptions *SetInstancePlansOptions)`, func() {
			setInstancePlansOptions := &adminservicev3.SetInstancePlansOptions{
				InstanceID: core.StringPtr("testString"),
				Name:       core.StringPtr("Standard"),
			}

			instancePlan, response, err := adminServiceService.SetInstancePlans(setInstancePlansOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instancePlan).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
