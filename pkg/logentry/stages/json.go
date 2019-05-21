package stages

import (
	"encoding/json"
	//"encoding/json"
	//"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jmespath/go-jmespath"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
)

// Config Errors
const (
	ErrExpressionsRequired  = "JMES expression is required"
	ErrCouldNotCompileJMES  = "could not compile JMES expression"
	ErrEmptyJSONStageConfig = "empty json stage configuration"
)

type JSONConfig struct {
	Expressions map[string]string `mapstructure:"expressions"`
}

// validateJSONConfig validates a json config and returns a map of necessary jmespath expressions.
func validateJSONConfig(c *JSONConfig) (map[string]*jmespath.JMESPath, error) {
	if c == nil {
		return nil, errors.New(ErrEmptyJSONStageConfig)
	}

	if len(c.Expressions) == 0 {
		return nil, errors.New(ErrExpressionsRequired)
	}

	expressions := map[string]*jmespath.JMESPath{}

	for n, e := range c.Expressions {
		var err error
		jmes := e
		// If there is no expression, use the name as the expression.
		if e == "" {
			jmes = n
		}
		expressions[n], err = jmespath.Compile(jmes)
		if err != nil {
			return nil, errors.Wrap(err, ErrCouldNotCompileJMES)
		}
	}
	return expressions, nil
}

// jsonStage extracts log data via json parsing.
type jsonStage struct {
	cfg         *JSONConfig
	expressions map[string]*jmespath.JMESPath
	logger      log.Logger
}

// newJSONStage creates a new json mutator from a config.
func newJSONStage(logger log.Logger, config interface{}) (*jsonStage, error) {
	cfg, err := parseJSONConfig(config)
	if err != nil {
		return nil, err
	}
	expressions, err := validateJSONConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &jsonStage{
		cfg:         cfg,
		expressions: expressions,
		logger:      log.With(logger, "component", "mutator", "type", "json"),
	}, nil
}

func parseJSONConfig(config interface{}) (*JSONConfig, error) {
	cfg := &JSONConfig{}
	err := mapstructure.Decode(config, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Process implements Mutator
func (j *jsonStage) Process(labels model.LabelSet, extracted map[string]interface{}, t *time.Time, entry *string) {
	if entry == nil {
		level.Debug(j.logger).Log("msg", "cannot parse a nil entry")
		return
	}

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(*entry), &data); err != nil {
		level.Debug(j.logger).Log("msg", "failed to unmarshal log line", "err", err)
		return
	}

	for n, e := range j.expressions {
		r, err := e.Search(data)
		if err != nil {
			level.Debug(j.logger).Log("msg", "failed to search JMES expression", "err", err)
			continue
		}

		//TODO do we want to pass arrays out in a non-marshaled format? Downstream stages would need to be updated to accept this
		switch r.(type) {
		case float64:
			// All numbers in JSON are unmarshaled to float64.
			extracted[n] = r
		case string:
			extracted[n] = r
		case bool:
			extracted[n] = r
		default:
			// If the value wasn't a string or a number, marshal it back to json
			jm, err := json.Marshal(r)
			if err != nil {
				level.Debug(j.logger).Log("msg", "failed to marshal complex type back to string", "err", err)
				continue
			}
			extracted[n] = string(jm)
		}
	}

}
