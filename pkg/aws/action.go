package aws

import (
	"github.com/deislabs/porter/pkg/exec/builder"
)

var _ builder.ExecutableAction = Action{}

type Action struct {
	Steps []Steps // using UnmarshalYAML so that we don't need a custom type per action
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - aws: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var steps []Steps
	results, err := builder.UnmarshalAction(unmarshal, &steps)
	if err != nil {
		return err
	}

	for _, result := range results {
		step := result.(*[]Steps)
		a.Steps = append(a.Steps, *step...)
	}
	return nil
}

func (a Action) GetSteps() []builder.ExecutableStep {
	steps := make([]builder.ExecutableStep, len(a.Steps))
	for i := range a.Steps {
		steps[i] = a.Steps[i]
	}

	return steps
}

var _ builder.ExecutableStep = Step{}
var _ builder.StepWithOutputs = Step{}

type Steps struct {
	Step `yaml:"aws"`
}

type Step struct {
	Description string        `yaml:"description"`
	Service     string        `yaml:"service"`
	Operation   string        `yaml:"operation"`
	Arguments   []string      `yaml:"arguments,omitempty"`
	Flags       builder.Flags `yaml:"flags,omitempty"`
	Outputs     []Output      `yaml:"outputs,omitempty"`
}

func (s Step) GetCommand() string {
	return "aws"
}

func (s Step) GetArguments() []string {
	args := make([]string, 0, len(s.Arguments)+2)

	// Specify the Service and Operation
	args = append(args, s.Service)
	args = append(args, s.Operation)

	// Append the positional arguments
	args = append(args, s.Arguments...)

	return args
}

func (s Step) GetFlags() builder.Flags {
	// Always request json formatted output
	return append(s.Flags, builder.NewFlag("output", "json"))
}

func (s Step) GetOutputs() []builder.Output {
	outputs := make([]builder.Output, len(s.Outputs))
	for i := range s.Outputs {
		outputs[i] = s.Outputs[i]
	}
	return outputs
}

var _ builder.OutputJsonPath = Output{}

type Output struct {
	Name     string `yaml:"name"`
	JsonPath string `yaml:"jsonPath"`
}

func (o Output) GetName() string {
	return o.Name
}

func (o Output) GetJsonPath() string {
	return o.JsonPath
}
