package stages

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"

	"github.com/grafana/loki/pkg/logentry/metric"
)

const customPrefix = "promtail_custom_"

const (
	ErrEmptyMetricsStageConfig = "empty metric stage configuration"
)

// MetricConfig is a single metrics configuration.
type MetricConfig struct {
	MetricType  string    `mapstructure:"type"`
	Description string    `mapstructure:"description"`
	Source      *string   `mapstructure:"source"`
	Buckets     []float64 `mapstructure:"buckets"`
}

// MetricsConfig is a set of configured metrics.
type MetricsConfig map[string]MetricConfig

func validateMetricsConfig(cfg MetricsConfig) error {
	if cfg == nil {
		return errors.New(ErrEmptyMetricsStageConfig)
	}
	for name, config := range cfg {
		//If the source is not defined, default to the metric name
		if config.Source == nil {
			cp := config
			nm := name
			cp.Source = &nm
			cfg[name] = cp
		}
	}
	return nil
}

// newMetric creates a new set of metrics to process for each log entry
func newMetric(config interface{}, registry prometheus.Registerer) (*metricStage, error) {
	cfgs := &MetricsConfig{}
	err := mapstructure.Decode(config, cfgs)
	if err != nil {
		return nil, err
	}
	err = validateMetricsConfig(*cfgs)
	if err != nil {
		return nil, err
	}
	metrics := map[string]prometheus.Collector{}
	for name, cfg := range *cfgs {
		var collector prometheus.Collector

		switch strings.ToLower(cfg.MetricType) {
		case MetricTypeCounter:
			collector = metric.NewCounters(customPrefix+name, cfg.Description)
		case MetricTypeGauge:
			collector = metric.NewGauges(customPrefix+name, cfg.Description)
		case MetricTypeHistogram:
			collector = metric.NewHistograms(customPrefix+name, cfg.Description, cfg.Buckets)
		}
		if collector != nil {
			registry.MustRegister(collector)
			metrics[name] = collector
		}
	}
	return &metricStage{
		cfg:     *cfgs,
		metrics: metrics,
	}, nil
}

type metricStage struct {
	cfg     MetricsConfig
	metrics map[string]prometheus.Collector
}

func (m *metricStage) Process(labels model.LabelSet, extracted map[string]interface{}, t *time.Time, entry *string) {
	for name, collector := range m.metrics {
		if v, ok := extracted[*m.cfg[name].Source]; ok {
			switch vec := collector.(type) {
			case *metric.Counters:
				recordCounter(vec.With(labels), v)
			case *metric.Gauges:
				recordGauge(vec.With(labels), v)
			case *metric.Histograms:
				recordHistogram(vec.With(labels), v)
			}
		}
	}
}

func recordCounter(counter prometheus.Counter, v interface{}) {
	f, err := getFloat(v)
	if err != nil || f < 0 {
		return
	}
	counter.Add(f)
}

func recordGauge(gauge prometheus.Gauge, v interface{}) {
	f, err := getFloat(v)
	if err != nil {
		return
	}
	//todo Gauge we be able to add,inc,dec,set
	gauge.Add(f)
}

func recordHistogram(histogram prometheus.Histogram, v interface{}) {
	f, err := getFloat(v)
	if err != nil {
		return
	}
	histogram.Observe(f)
}

func getFloat(unk interface{}) (float64, error) {

	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	case bool:
		if i {
			return float64(1), nil
		}
		return float64(0), nil
	default:
		return math.NaN(), fmt.Errorf("Can't convert %v to float64", unk)
	}
}
