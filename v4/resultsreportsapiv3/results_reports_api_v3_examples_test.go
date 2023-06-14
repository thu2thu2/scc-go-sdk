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

package resultsreportsapiv3_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/resultsreportsapiv3"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Results/Reports API service.
//
// The following configuration properties are assumed to be defined:
// RESULTS_REPORTS_API_URL=<service base url>
// RESULTS_REPORTS_API_AUTH_TYPE=iam
// RESULTS_REPORTS_API_APIKEY=<IAM apikey>
// RESULTS_REPORTS_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`ResultsReportsApiV3 Examples Tests`, func() {

	const externalConfigFile = "../results_reports_api_v3.env"

	var (
		resultsReportsApiService *resultsreportsapiv3.ResultsReportsApiV3
		config                   map[string]string
	)

	err := godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Error occured when loading external config file. Err: ", err)
	}

	accountInstanceId := os.Getenv("ACCOUNT_INSTANCE_ID")
	reportId := os.Getenv("REPORT_ID")
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
			config, err = core.GetServiceProperties(resultsreportsapiv3.DefaultServiceName)
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

			resultsReportsApiServiceOptions := &resultsreportsapiv3.ResultsReportsApiV3Options{}

			resultsReportsApiService, err = resultsreportsapiv3.NewResultsReportsApiV3UsingExternalConfig(resultsReportsApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(resultsReportsApiService).ToNot(BeNil())
		})
	})

	Describe(`ResultsReportsApiV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestReports request example`, func() {
			fmt.Println("\nGetLatestReports() result:")
			// begin-get_latest_reports

			getLatestReportsOptions := resultsReportsApiService.NewGetLatestReportsOptions(
				accountInstanceId,
			)

			getLatestReportsResponse, response, err := resultsReportsApiService.GetLatestReports(getLatestReportsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getLatestReportsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_reports

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getLatestReportsResponse).ToNot(BeNil())
		})
		It(`ListReports request example`, func() {
			fmt.Println("\nListReports() result:")
			// begin-list_reports
			listReportsOptions := &resultsreportsapiv3.ListReportsOptions{
				InstanceID:     core.StringPtr(accountInstanceId),
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

			pager, err := resultsReportsApiService.NewReportsPager(listReportsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []resultsreportsapiv3.Report
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_reports
		})
		It(`GetReportsProfiles request example`, func() {
			fmt.Println("\nGetReportsProfiles() result:")
			// begin-get_reports_profiles

			getReportsProfilesOptions := resultsReportsApiService.NewGetReportsProfilesOptions(
				accountInstanceId,
			)

			getProfilesResponse, response, err := resultsReportsApiService.GetReportsProfiles(getReportsProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getProfilesResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_reports_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getProfilesResponse).ToNot(BeNil())
		})
		It(`GetReportsScopes request example`, func() {
			fmt.Println("\nGetReportsScopes() result:")
			// begin-get_reports_scopes

			getReportsScopesOptions := resultsReportsApiService.NewGetReportsScopesOptions(
				accountInstanceId,
			)

			getScopesResponse, response, err := resultsReportsApiService.GetReportsScopes(getReportsScopesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getScopesResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_reports_scopes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getScopesResponse).ToNot(BeNil())
		})
		It(`GetReport request example`, func() {
			fmt.Println("\nGetReport() result:")
			// begin-get_report

			getReportOptions := resultsReportsApiService.NewGetReportOptions(
				accountInstanceId,
				reportId,
			)

			report, response, err := resultsReportsApiService.GetReport(getReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-get_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
		It(`GetReportSummary request example`, func() {
			fmt.Println("\nGetReportSummary() result:")
			// begin-get_report_summary

			getReportSummaryOptions := resultsReportsApiService.NewGetReportSummaryOptions(
				accountInstanceId,
				reportId,
			)

			reportSummary, response, err := resultsReportsApiService.GetReportSummary(getReportSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_report_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
		It(`GetReportEvaluation request example`, func() {
			fmt.Println("\nGetReportEvaluation() result:")
			// begin-get_report_evaluation

			getReportEvaluationOptions := resultsReportsApiService.NewGetReportEvaluationOptions(
				accountInstanceId,
				reportId,
			)

			result, response, err := resultsReportsApiService.GetReportEvaluation(getReportEvaluationOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-get_report_evaluation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`GetReportControls request example`, func() {
			fmt.Println("\nGetReportControls() result:")
			// begin-get_report_controls

			getReportControlsOptions := resultsReportsApiService.NewGetReportControlsOptions(
				accountInstanceId,
				reportId,
			)
			getReportControlsOptions.SetStatus("compliant")

			getReportControlsResponse, response, err := resultsReportsApiService.GetReportControls(getReportControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getReportControlsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_report_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getReportControlsResponse).ToNot(BeNil())
		})
		It(`GetReportRule request example`, func() {
			fmt.Println("\nGetReportRule() result:")
			// begin-get_report_rule

			getReportRuleOptions := resultsReportsApiService.NewGetReportRuleOptions(
				accountInstanceId,
				reportId,
				ruleId,
			)

			rule, response, err := resultsReportsApiService.GetReportRule(getReportRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_report_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
		It(`ListReportEvaluations request example`, func() {
			fmt.Println("\nListReportEvaluations() result:")
			// begin-list_report_evaluations
			listReportEvaluationsOptions := &resultsreportsapiv3.ListReportEvaluationsOptions{
				InstanceID:     core.StringPtr(accountInstanceId),
				ReportID:       core.StringPtr(reportId),
				AssessmentID:   core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				TargetID:       core.StringPtr("testString"),
				TargetName:     core.StringPtr("testString"),
				Status:         core.StringPtr("failure"),
				Limit:          core.Int64Ptr(int64(10)),
				XCorrelationID: core.StringPtr("testString"),
			}

			pager, err := resultsReportsApiService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []resultsreportsapiv3.Evaluation
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_report_evaluations
		})
		It(`ListReportResources request example`, func() {
			fmt.Println("\nListReportResources() result:")
			// begin-list_report_resources
			listReportResourcesOptions := &resultsreportsapiv3.ListReportResourcesOptions{
				InstanceID:     core.StringPtr(accountInstanceId),
				ReportID:       core.StringPtr(reportId),
				ID:             core.StringPtr("testString"),
				ResourceName:   core.StringPtr("testString"),
				AccountID:      core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				Status:         core.StringPtr("compliant"),
				Limit:          core.Int64Ptr(int64(10)),
				XCorrelationID: core.StringPtr("testString"),
			}

			pager, err := resultsReportsApiService.NewReportResourcesPager(listReportResourcesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []resultsreportsapiv3.Resource
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_report_resources
		})
		It(`GetReportTags request example`, func() {
			fmt.Println("\nGetReportTags() result:")
			// begin-get_report_tags

			getReportTagsOptions := resultsReportsApiService.NewGetReportTagsOptions(
				reportId,
			)

			getTagsResponse, response, err := resultsReportsApiService.GetReportTags(getReportTagsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getTagsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_report_tags

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTagsResponse).ToNot(BeNil())
		})
		It(`GetReportViolationsDrift request example`, func() {
			fmt.Println("\nGetReportViolationsDrift() result:")
			// begin-get_report_violations_drift

			getReportViolationsDriftOptions := resultsReportsApiService.NewGetReportViolationsDriftOptions(
				accountInstanceId,
				reportId,
			)

			getReportViolationsDriftResult, response, err := resultsReportsApiService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getReportViolationsDriftResult, "", "  ")
			fmt.Println(string(b))

			// end-get_report_violations_drift

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getReportViolationsDriftResult).ToNot(BeNil())
		})
	})
})
