package converter

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1alpha3"
	"gopkg.in/inf.v0"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

type SLOConverter struct {
}

func NewSLOConverter() *SLOConverter {
	return &SLOConverter{}
}

type SLO struct {
	Objectives []*Objective `yaml:"objectives" json:"objectives"`
	TotalScore Score        `yaml:"total_score" json:"total_score"`
}

type Score struct {
	Pass    string `yaml:"pass" json:"pass"`
	Warning string `yaml:"warning" json:"warning"`
}

type Objective struct {
	Name    string     `yaml:"sli" json:"sli"`
	KeySLI  bool       `yaml:"key_sli,omitempty" json:"key_sli,omitempty"`
	Weight  int        `yaml:"weight,omitempty" json:"weight,omitempty"`
	Warning []Criteria `yaml:"warning,omitempty" json:"warning,omitempty"`
	Pass    []Criteria `yaml:"pass,omitempty" json:"pass,omitempty"`
}

type Criteria struct {
	Operators []string `yaml:"criteria,omitempty" json:"criteria,omitempty"`
}

func (o *Objective) hasSupportedCriteria() bool {
	return len(o.Pass) > 1 || len(o.Warning) > 1
}

func (c *SLOConverter) Convert(fileContent []byte, analysisDef string, namespace string) (string, error) {
	//check that provider and namespace is set
	if analysisDef == "" || namespace == "" {
		return "", fmt.Errorf("missing arguments: 'definition' and 'namespace' needs to be set for conversion")
	}

	// unmarshall content
	content := &SLO{}
	err := yaml.Unmarshal(fileContent, content)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling file content: %s", err.Error())
	}

	// convert
	analysisDefinition, err := c.convertSLO(content, analysisDef, namespace)
	if err != nil {
		return "", err
	}

	// marshal AnalysisDefinition to Yaml
	yamlData, err := yaml.Marshal(analysisDefinition)
	if err != nil {
		return "", fmt.Errorf("error marshalling data: %s", err.Error())
	}

	return string(yamlData), nil
}

func (c *SLOConverter) convertSLO(sloContent *SLO, name string, namespace string) (*metricsapi.AnalysisDefinition, error) {
	// define resulting AnalysisDefinition with easy conversions
	passPercentage, err := removePercentage(sloContent.TotalScore.Pass)
	if err != nil {
		return nil, err
	}
	warnPercentage, err := removePercentage(sloContent.TotalScore.Warning)
	if err != nil {
		return nil, err
	}
	definition := &metricsapi.AnalysisDefinition{
		TypeMeta: v1.TypeMeta{
			Kind:       "AnalysisDefinition",
			APIVersion: "metrics.keptn.sh/v1alpha3",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		Spec: metricsapi.AnalysisDefinitionSpec{
			TotalScore: metricsapi.TotalScore{
				PassPercentage:    passPercentage,
				WarningPercentage: warnPercentage,
			},
			// create a slice of size of len(sloContent.Objectives), but reserve capacity for
			// double the size, as some objectives may be twice there (conversion of criteria with logical AND)
			Objectives: make([]metricsapi.Objective, len(sloContent.Objectives), len(sloContent.Objectives)*2),
		},
	}

	// convert objectives one after another
	indexObjectives := 0
	for _, o := range sloContent.Objectives {
		target, err := setupTarget(o)
		if err != nil {
			return nil, err
		}
		objective := metricsapi.Objective{
			AnalysisValueTemplateRef: metricsapi.ObjectReference{
				Name:      o.Name,
				Namespace: namespace,
			},
			KeyObjective: o.KeySLI,
			Weight:       o.Weight,
			Target:       *target,
		}
		definition.Spec.Objectives[indexObjectives] = objective
		indexObjectives++
	}
	return definition, nil
}

// removes % symbol from the scoring values and converts to numeric value
func removePercentage(str string) (int, error) {
	t := strings.ReplaceAll(str, "%", "")
	f, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return 0, err
	}
	return int(math.Round(f)), nil
}

// creates and sets up the target struct from objective
// TODO refactor this function in a follow-up + weight distribution
// nolint:gocognit,gocyclo
func setupTarget(o *Objective) (*metricsapi.Target, error) {
	target := &metricsapi.Target{}
	// remove criteria, which contain % in their operators
	o = cleanupObjective(o)
	// skip objective target conversion if it has criteria combined with logical OR -> not supported
	// this way the SLO will become "informative"
	if o.hasSupportedCriteria() {
		return target, nil
	}

	// if warning criteria are not defined, negate the pass criteria to create fail criteria
	if len(o.Warning) == 0 {
		if len(o.Pass) > 0 {
			if len(o.Pass[0].Operators) > 0 {
				// TODO cover use cases with multiple operators (create new objectives)
				op, err := newOperator(o.Pass[0].Operators[0])
				if err != nil {
					return target, err
				}
				target.Failure = op
				return target, nil
			}
		}
	}

	// if warning criteria are defined, create new criteria with the following logic:
	// !(warn criteria) -> fail criteria
	// !(pass criteria) -> warn criteria
	var err error
	if len(o.Pass) > 0 {
		if len(o.Pass[0].Operators) > 0 {
			// TODO cover use cases with multiple operators (create new objectives)
			op, err := newOperator(o.Pass[0].Operators[0])
			if err != nil {
				return target, err
			}
			target.Warning = op
		}
		if len(o.Warning[0].Operators) > 0 {
			// TODO cover use cases with multiple operators (create new objectives)
			op, err := newOperator(o.Warning[0].Operators[0])
			if err != nil {
				return target, err
			}
			target.Failure = op
		}
	}

	return target, err
}

func cleanupObjective(o *Objective) *Objective {
	o.Pass = cleanupCriteria(o.Pass)
	o.Warning = cleanupCriteria(o.Warning)
	return o
}

// remove % operators from criterium structure
// if criteria did have only % operators, remove it from strucutre
func cleanupCriteria(criteria []Criteria) []Criteria {
	newCriteria := make([]Criteria, 0, len(criteria))
	for _, c := range criteria {
		operators := make([]string, 0, len(c.Operators))
		for _, op := range c.Operators {
			// keep only criteria with real values, not percentage
			if !strings.Contains(op, "%") {
				operators = append(operators, op)
			}
		}
		// if criterium does have operator, store it
		if len(operators) > 0 {
			newCriteria = append(newCriteria, Criteria{Operators: operators})
		}
	}

	return newCriteria
}

// create operator for Target
func newOperator(op string) (*metricsapi.Operator, error) {
	// remove whitespaces
	op = strings.Replace(op, " ", "", -1)

	operators := []string{"<=", "<", ">=", ">"}
	for _, operator := range operators {
		if strings.HasPrefix(op, operator) {
			return createOperator(operator, strings.TrimPrefix(op, operator))
		}
	}

	return &metricsapi.Operator{}, fmt.Errorf("invalid operator: '%s'", op)
}

// checks and negates the existing operator
func createOperator(op string, value string) (*metricsapi.Operator, error) {
	dec := inf.NewDec(1, 0)
	_, ok := dec.SetString(value)
	if !ok {
		return nil, fmt.Errorf("unable to convert value '%s' to decimal", value)
	}
	if op == "<=" {
		return &metricsapi.Operator{
			GreaterThan: &metricsapi.OperatorValue{
				FixedValue: *resource.NewDecimalQuantity(*dec, resource.DecimalSI),
			},
		}, nil
	} else if op == "<" {
		return &metricsapi.Operator{
			GreaterThanOrEqual: &metricsapi.OperatorValue{
				FixedValue: *resource.NewDecimalQuantity(*dec, resource.DecimalSI),
			},
		}, nil
	} else if op == ">=" {
		return &metricsapi.Operator{
			LessThan: &metricsapi.OperatorValue{
				FixedValue: *resource.NewDecimalQuantity(*dec, resource.DecimalSI),
			},
		}, nil
	} else if op == ">" {
		return &metricsapi.Operator{
			LessThanOrEqual: &metricsapi.OperatorValue{
				FixedValue: *resource.NewDecimalQuantity(*dec, resource.DecimalSI),
			},
		}, nil
	}

	return &metricsapi.Operator{}, fmt.Errorf("invalid operator: '%s'", op)
}
