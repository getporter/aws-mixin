package aws

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
)

func TestFlags_Sort(t *testing.T) {
	flags := Flags{
		NewFlag("b", "1"),
		NewFlag("a", "2"),
		NewFlag("c", "3"),
	}

	sort.Sort(flags)

	assert.Equal(t, "a", flags[0].Name)
	assert.Equal(t, "b", flags[1].Name)
	assert.Equal(t, "c", flags[2].Name)
}

func TestMixin_UnmarshalStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input.yaml")
	require.NoError(t, err)

	var step Steps
	err = yaml.Unmarshal(b, &step)
	require.NoError(t, err)

	assert.Equal(t, "Provision VM", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, Output{"INSTANCE_ID", "$.Instances[0].InstanceId"}, step.Outputs[0])

	assert.Equal(t, "ec2", step.Service)
	assert.Equal(t, "run-instances", step.Operation)

	assert.Equal(t, []string{"myinst"}, step.Arguments)

	sort.Sort(step.Flags)
	assert.Equal(t, Flags{
		NewFlag("env", "FOO=BAR", "STUFF=THINGS"),
		NewFlag("image-id", "ami-xxxxxxxx"),
		NewFlag("instance-type", "t2.micro")}, step.Flags)
}

func TestMixin_UnmarshalInvalidStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input-invalid.yaml")
	require.NoError(t, err)

	var step Steps
	err = yaml.Unmarshal(b, &step)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid yaml type for flag env")
}
