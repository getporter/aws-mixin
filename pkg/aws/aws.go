package aws

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Action struct {
	Steps []Steps // using UnmarshalYAML so that we don't need a custom type per action
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - aws: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	actionMap := map[interface{}][]interface{}{}
	err := unmarshal(&actionMap)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal yaml into an action map of aws steps")
	}

	for _, stepMaps := range actionMap {
		b, err := yaml.Marshal(stepMaps)
		if err != nil {
			return err
		}

		var steps []Steps
		err = yaml.Unmarshal(b, &steps)
		if err != nil {
			return err
		}

		a.Steps = append(a.Steps, steps...)
	}

	return nil
}

type Steps struct {
	Step `yaml:"aws"`
}

func (m *Mixin) Execute() error {
	payload, err := m.getPayloadData()
	if err != nil {
		return err
	}

	var action Action
	err = yaml.Unmarshal(payload, &action)
	if err != nil {
		return err
	}
	if len(action.Steps) != 1 {
		return errors.Errorf("expected a single step, but got %d", len(action.Steps))
	}
	step := action.Steps[0]

	// Always output json so that we can query it for outputs afterwards
	step.Flags = append(step.Flags, NewFlag("output", "json"))

	fmt.Fprintf(m.Out, "Starting operation: %s\n", step.Description)

	args := make([]string, 0, 2+len(step.Arguments)+len(step.Flags)*2)

	// Specify the aws service and command to run
	args = append(args, step.Service)
	args = append(args, step.Operation)

	// Append the positional arguments
	for _, arg := range step.Arguments {
		args = append(args, arg)
	}

	// Append the flags to the argument list
	sort.Sort(step.Flags)
	for _, flag := range step.Flags {
		for _, value := range flag.Values {
			args = append(args, fmt.Sprintf("--%s", flag.Name))
			args = append(args, value)
		}
	}

	cmd := m.NewCommand("aws", args...)
	cmd.Stdout = m.Out
	cmd.Stderr = m.Err

	prettyCmd := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))
	if m.Debug {
		fmt.Fprintln(m.Out, prettyCmd)
	}

	err = cmd.Start()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("couldn't run command %s", prettyCmd))
	}

	err = cmd.Wait()

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error running command %s", prettyCmd))
	}
	fmt.Fprintf(m.Out, "Finished operation: %s\n", step.Description)

	for _, output := range step.Outputs {
		//TODO populate the output
		v := "SOME VALUE"

		err := m.WriteMixinOutputToFile(output.Name, []byte(v))
		if err != nil {
			return errors.Wrapf(err, "unable to write output '%s'", output.Name)
		}
	}
	return nil
}
