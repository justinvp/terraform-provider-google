// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccLoggingMetric_loggingMetricBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLoggingMetricDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingMetric_loggingMetricBasicExample(context),
			},
			{
				ResourceName:      "google_logging_metric.logging_metric",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLoggingMetric_loggingMetricBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_metric" "logging_metric" {
  name = "my-(custom)/metric%{random_suffix}"
  filter = "resource.type=gae_app AND severity>=ERROR"
  metric_descriptor {
    metric_kind = "DELTA"
    value_type = "DISTRIBUTION"
    labels {
        key = "mass"
        value_type = "STRING"
        description = "amount of matter"
    }
  }
  value_extractor = "EXTRACT(jsonPayload.request)"
  label_extractors = { "mass": "EXTRACT(jsonPayload.request)" }
  bucket_options {
    linear_buckets {
      num_finite_buckets = 3
      width = 1
      offset = 1
    }
  }
}
`, context)
}

func testAccCheckLoggingMetricDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_logging_metric" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{LoggingBasePath}}projects/{{project}}/metrics/{{%name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("LoggingMetric still exists at %s", url)
		}
	}

	return nil
}
