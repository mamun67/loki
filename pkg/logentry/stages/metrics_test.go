package stages

import (
	"strings"
	"testing"
	"time"

	"github.com/cortexproject/cortex/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	testutil "github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/prometheus/common/model"
)

var labelFoo = model.LabelSet(map[model.LabelName]model.LabelValue{"foo": "bar", "bar": "foo"})
var labelFu = model.LabelSet(map[model.LabelName]model.LabelValue{"fu": "baz", "baz": "fu"})

func Test_withMetric(t *testing.T) {
	cfg := map[string]interface{}{
		"metrics": map[string]interface{}{
			"total_keys": map[string]interface{}{
				"type":        "Counter",
				"source":      "length(keys(@))",
				"description": "the total keys per doc",
			},
			"keys_per_line": map[string]interface{}{
				"type":        "Histogram",
				"source":      "length(keys(@))",
				"description": "keys per doc",
				"buckets":     []float64{1, 3, 5, 10},
			},
			"numeric_float": map[string]interface{}{
				"type":        "Gauge",
				"source":      "numeric.float",
				"description": "numeric_float",
			},
			"numeric_integer": map[string]interface{}{
				"type":        "Gauge",
				"source":      "numeric.integer",
				"description": "numeric.integer",
			},
			"numeric_string": map[string]interface{}{
				"type":        "Gauge",
				"source":      "numeric.string",
				"description": "numeric.string",
			},
			"contains_warn": map[string]interface{}{
				"type":        "Counter",
				"source":      "contains(values(@),'WARN')",
				"description": "contains_warn",
			},
			"contains_false": map[string]interface{}{
				"type":        "Counter",
				"source":      "contains(keys(@),'nope')",
				"description": "contains_false",
			},
			"unconvertible": map[string]interface{}{
				"type":        "Counter",
				"source":      "values(@)",
				"description": "unconvertible",
			},
		},
	}
	registry := prometheus.NewRegistry()
	metricStage, err := New(util.Logger, StageTypeJSON, cfg, registry)
	if err != nil {
		t.Fatalf("failed to create stage with metrics: %v", err)
	}
	var ts = time.Now()
	var entry = logFixture

	metricStage.Process(labelFoo, &ts, &entry)
	metricStage.Process(labelFu, &ts, &entry)

	if err := testutil.GatherAndCompare(registry,
		strings.NewReader(goldenMetrics), metricNames(cfg)...); err != nil {
		t.Fatalf("missmatch metrics: %v", err)
	}
}

func metricNames(cfg map[string]interface{}) []string {
	metrics := cfg["metrics"].(map[string]interface{})
	result := make([]string, 0, len(cfg))
	for name := range metrics {
		result = append(result, name)
	}
	return result
}

const goldenMetrics = `# HELP contains_false contains_false
# TYPE contains_false counter
contains_false{bar="foo",foo="bar"} 0.0
contains_false{baz="fu",fu="baz"} 0.0
# HELP contains_warn contains_warn
# TYPE contains_warn counter
contains_warn{bar="foo",foo="bar"} 1.0
contains_warn{baz="fu",fu="baz"} 1.0
# HELP keys_per_line keys per doc
# TYPE keys_per_line histogram
keys_per_line_bucket{bar="foo",foo="bar",le="1.0"} 0.0
keys_per_line_bucket{bar="foo",foo="bar",le="3.0"} 0.0
keys_per_line_bucket{bar="foo",foo="bar",le="5.0"} 0.0
keys_per_line_bucket{bar="foo",foo="bar",le="10.0"} 1.0
keys_per_line_bucket{bar="foo",foo="bar",le="+Inf"} 1.0
keys_per_line_sum{bar="foo",foo="bar"} 8.0
keys_per_line_count{bar="foo",foo="bar"} 1.0
keys_per_line_bucket{baz="fu",fu="baz",le="1.0"} 0.0
keys_per_line_bucket{baz="fu",fu="baz",le="3.0"} 0.0
keys_per_line_bucket{baz="fu",fu="baz",le="5.0"} 0.0
keys_per_line_bucket{baz="fu",fu="baz",le="10.0"} 1.0
keys_per_line_bucket{baz="fu",fu="baz",le="+Inf"} 1.0
keys_per_line_sum{baz="fu",fu="baz"} 8.0
keys_per_line_count{baz="fu",fu="baz"} 1.0
# HELP numeric_float numeric_float
# TYPE numeric_float gauge
numeric_float{bar="foo",foo="bar"} 12.34
numeric_float{baz="fu",fu="baz"} 12.34
# HELP numeric_integer numeric.integer
# TYPE numeric_integer gauge
numeric_integer{bar="foo",foo="bar"} 123.0
numeric_integer{baz="fu",fu="baz"} 123.0
# HELP numeric_string numeric.string
# TYPE numeric_string gauge
numeric_string{bar="foo",foo="bar"} 123.0
numeric_string{baz="fu",fu="baz"} 123.0
# HELP total_keys the total keys per doc
# TYPE total_keys counter
total_keys{bar="foo",foo="bar"} 8.0
total_keys{baz="fu",fu="baz"} 8.0
# HELP unconvertible unconvertible
# TYPE unconvertible counter
unconvertible{bar="foo",foo="bar"} 0.0
unconvertible{baz="fu",fu="baz"} 0.0
`
