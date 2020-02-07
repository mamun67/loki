package logql

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestResult_RecordMetrics(t *testing.T) {
	for _, f := range prometheus.ExponentialBuckets(0.125, 2, 10) {
		// fmt.Print(humanize.Bytes(uint64(f)))
		fmt.Print(f)

		fmt.Print(" ")

	}
}

func Test_queryType(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  string
	}{
		{"bad", "ddd", ""},
		{"limited", `{app="foo"}`, typeLimited},
		{"limited multi label", `{app="foo" ,fuzz=~"foo"}`, typeLimited},
		{"filter", `{app="foo"} |= "foo"`, typeFilter},
		{"metrics", `rate({app="foo"} |= "foo"[5m])`, typeMetric},
		{"filters", `{app="foo"} |= "foo" |= "f" != "b"`, typeFilter},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryType(tt.query); got != tt.want {
				t.Errorf("queryType() = %v, want %v", got, tt.want)
			}
		})
	}
}
