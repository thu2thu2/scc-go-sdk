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

package resultsreportsapiv3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/resultsreportsapiv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the resultsreportsapiv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ResultsReportsApiV3 Integration Tests`, func() {
	const externalConfigFile = "../results_reports_api_v3.env"

	var (
		err                      error
		resultsReportsApiService *resultsreportsapiv3.ResultsReportsApiV3
		serviceURL               string
		config                   map[string]string
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
			config, err = core.GetServiceProperties(resultsreportsapiv3.DefaultServiceName)
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
			resultsReportsApiServiceOptions := &resultsreportsapiv3.ResultsReportsApiV3Options{}

			resultsReportsApiService, err = resultsreportsapiv3.NewResultsReportsApiV3UsingExternalConfig(resultsReportsApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(resultsReportsApiService).ToNot(BeNil())
			Expect(resultsReportsApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			resultsReportsApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetLatestReports - Return the latest distinct reports grouped by profile ID, scope ID and attachment ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions)`, func() {
			getLatestReportsOptions := &resultsreportsapiv3.GetLatestReportsOptions{
				InstanceID:     core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
				HomeAccountID:  core.StringPtr("testString"),
				Sort:           core.StringPtr("testString"),
			}

			getLatestReportsResponse, response, err := resultsReportsApiService.GetLatestReports(getLatestReportsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getLatestReportsResponse).ToNot(BeNil())
		})
	})

	Describe(`ListReports - List reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) with pagination`, func() {
			listReportsOptions := &resultsreportsapiv3.ListReportsOptions{
				InstanceID:     core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
				HomeAccountID:  core.StringPtr("testString"),
				AttachmentID:   core.StringPtr("testString"),
				GroupID:        core.StringPtr("testString"),
				ProfileID:      core.StringPtr("testString"),
				ScopeID:        core.StringPtr("testString"),
				Type:           core.StringPtr("scheduled"),
				Start:          core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				Sort:           core.StringPtr("testString"),
			}

			listReportsOptions.Start = nil
			listReportsOptions.Limit = core.Int64Ptr(1)

			var allResults []resultsreportsapiv3.Report
			for {
				reportPage, response, err := resultsReportsApiService.ListReports(listReportsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reportPage).ToNot(BeNil())
				allResults = append(allResults, reportPage.Reports...)

				listReportsOptions.Start, err = reportPage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) using ReportsPager`, func() {
			listReportsOptions := &resultsreportsapiv3.ListReportsOptions{
				InstanceID:     core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
				HomeAccountID:  core.StringPtr("testString"),
				AttachmentID:   core.StringPtr("testString"),
				GroupID:        core.StringPtr("testString"),
				ProfileID:      core.StringPtr("testString"),
				ScopeID:        core.StringPtr("testString"),
				Type:           core.StringPtr("scheduled"),
				Limit:          core.Int64Ptr(int64(10)),
				Sort:           core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := resultsReportsApiService.NewReportsPager(listReportsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []resultsreportsapiv3.Report
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = resultsReportsApiService.NewReportsPager(listReportsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReports() returned a total of %d item(s) using ReportsPager.\n", len(allResults))
		})
	})

	Describe(`GetReportsProfiles - Retrieve a unique list of profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportsProfiles(getReportsProfilesOptions *GetReportsProfilesOptions)`, func() {
			getReportsProfilesOptions := &resultsreportsapiv3.GetReportsProfilesOptions{
				InstanceID:     core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
				HomeAccountID:  core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
			}

			getProfilesResponse, response, err := resultsReportsApiService.GetReportsProfiles(getReportsProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getProfilesResponse).ToNot(BeNil())
		})
	})

	Describe(`GetReportsScopes - Retrieve a unique list of scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportsScopes(getReportsScopesOptions *GetReportsScopesOptions)`, func() {
			getReportsScopesOptions := &resultsreportsapiv3.GetReportsScopesOptions{
				InstanceID:     core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
				HomeAccountID:  core.StringPtr("testString"),
			}

			getScopesResponse, response, err := resultsReportsApiService.GetReportsScopes(getReportsScopesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getScopesResponse).ToNot(BeNil())
		})
	})

	Describe(`GetReport - Retrieve a single report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReport(getReportOptions *GetReportOptions)`, func() {
			getReportOptions := &resultsreportsapiv3.GetReportOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
			}

			report, response, err := resultsReportsApiService.GetReport(getReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
	})

	Describe(`GetReportSummary - Retrieve a report summary`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions)`, func() {
			getReportSummaryOptions := &resultsreportsapiv3.GetReportSummaryOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
			}

			reportSummary, response, err := resultsReportsApiService.GetReportSummary(getReportSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
	})

	Describe(`GetReportEvaluation - Get evaluation details for a given report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportEvaluation(getReportEvaluationOptions *GetReportEvaluationOptions)`, func() {
			getReportEvaluationOptions := &resultsreportsapiv3.GetReportEvaluationOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
			}

			result, response, err := resultsReportsApiService.GetReportEvaluation(getReportEvaluationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetReportControls - Get controls for a given report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportControls(getReportControlsOptions *GetReportControlsOptions)`, func() {
			getReportControlsOptions := &resultsreportsapiv3.GetReportControlsOptions{
				InstanceID:         core.StringPtr("testString"),
				ReportID:           core.StringPtr("testString"),
				ControlID:          core.StringPtr("testString"),
				ControlName:        core.StringPtr("testString"),
				ControlDescription: core.StringPtr("testString"),
				ControlCategory:    core.StringPtr("testString"),
				Status:             core.StringPtr("compliant"),
				Sort:               core.StringPtr("testString"),
				XCorrelationID:     core.StringPtr("testString"),
			}

			getReportControlsResponse, response, err := resultsReportsApiService.GetReportControls(getReportControlsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getReportControlsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetReportRule - Retrieve a single rule in a report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportRule(getReportRuleOptions *GetReportRuleOptions)`, func() {
			getReportRuleOptions := &resultsreportsapiv3.GetReportRuleOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				RuleID:         core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
			}

			rule, response, err := resultsReportsApiService.GetReportRule(getReportRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
	})

	Describe(`ListReportEvaluations - List report evaluations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) with pagination`, func() {
			listReportEvaluationsOptions := &resultsreportsapiv3.ListReportEvaluationsOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				AssessmentID:   core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				TargetID:       core.StringPtr("testString"),
				TargetName:     core.StringPtr("testString"),
				Status:         core.StringPtr("failure"),
				Start:          core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				XCorrelationID: core.StringPtr("testString"),
			}

			listReportEvaluationsOptions.Start = nil
			listReportEvaluationsOptions.Limit = core.Int64Ptr(1)

			var allResults []resultsreportsapiv3.Evaluation
			for {
				evaluationPage, response, err := resultsReportsApiService.ListReportEvaluations(listReportEvaluationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(evaluationPage).ToNot(BeNil())
				allResults = append(allResults, evaluationPage.Evaluations...)

				listReportEvaluationsOptions.Start, err = evaluationPage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportEvaluationsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) using ReportEvaluationsPager`, func() {
			listReportEvaluationsOptions := &resultsreportsapiv3.ListReportEvaluationsOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				AssessmentID:   core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				TargetID:       core.StringPtr("testString"),
				TargetName:     core.StringPtr("testString"),
				Status:         core.StringPtr("failure"),
				Limit:          core.Int64Ptr(int64(10)),
				XCorrelationID: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := resultsReportsApiService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []resultsreportsapiv3.Evaluation
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = resultsReportsApiService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportEvaluations() returned a total of %d item(s) using ReportEvaluationsPager.\n", len(allResults))
		})
	})

	Describe(`ListReportResources - List report resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) with pagination`, func() {
			listReportResourcesOptions := &resultsreportsapiv3.ListReportResourcesOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
				ResourceName:   core.StringPtr("testString"),
				AccountID:      core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				Status:         core.StringPtr("compliant"),
				Start:          core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				XCorrelationID: core.StringPtr("testString"),
			}

			listReportResourcesOptions.Start = nil
			listReportResourcesOptions.Limit = core.Int64Ptr(1)

			var allResults []resultsreportsapiv3.Resource
			for {
				resourcePage, response, err := resultsReportsApiService.ListReportResources(listReportResourcesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(resourcePage).ToNot(BeNil())
				allResults = append(allResults, resourcePage.Resources...)

				listReportResourcesOptions.Start, err = resourcePage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportResourcesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) using ReportResourcesPager`, func() {
			listReportResourcesOptions := &resultsreportsapiv3.ListReportResourcesOptions{
				InstanceID:     core.StringPtr("testString"),
				ReportID:       core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
				ResourceName:   core.StringPtr("testString"),
				AccountID:      core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				Status:         core.StringPtr("compliant"),
				Limit:          core.Int64Ptr(int64(10)),
				XCorrelationID: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := resultsReportsApiService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []resultsreportsapiv3.Resource
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = resultsReportsApiService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportResources() returned a total of %d item(s) using ReportResourcesPager.\n", len(allResults))
		})
	})

	Describe(`GetReportTags - Get tags associated with the report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportTags(getReportTagsOptions *GetReportTagsOptions)`, func() {
			getReportTagsOptions := &resultsreportsapiv3.GetReportTagsOptions{
				ReportID:       core.StringPtr("testString"),
				XCorrelationID: core.StringPtr("testString"),
			}

			getTagsResponse, response, err := resultsReportsApiService.GetReportTags(getReportTagsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTagsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetReportViolationsDrift - Get report violations drift`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions)`, func() {
			getReportViolationsDriftOptions := &resultsreportsapiv3.GetReportViolationsDriftOptions{
				InstanceID:       core.StringPtr("testString"),
				ReportID:         core.StringPtr("testString"),
				ScanTimeDuration: core.Int64Ptr(int64(0)),
				XCorrelationID:   core.StringPtr("testString"),
			}

			getReportViolationsDriftResult, response, err := resultsReportsApiService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getReportViolationsDriftResult).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
