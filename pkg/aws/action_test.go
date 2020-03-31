package aws

import (
	"io/ioutil"
	"sort"
	"testing"

	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMixin_UnmarshalInstallAction(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/install-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)

	require.Equal(t, 1, len(action.Steps))
	step := action.Steps[0]

	assert.Equal(t, "Provision VM", step.Description)
	require.NotEmpty(t, step.Outputs)
	assert.Equal(t, Output{"INSTANCE_ID", "$.Instances[0].InstanceId"}, step.Outputs[0])

	assert.Equal(t, "ec2", step.Service)
	assert.Equal(t, "run-instances", step.Operation)

	assert.Equal(t, []string{"myinst"}, step.Arguments)

	sort.Sort(step.Flags)
	assert.Equal(t, builder.Flags{
		builder.NewFlag("image-id", "ami-xxxxxxxx"),
		builder.NewFlag("instance-type", "t2.micro")}, step.Flags)

	assert.Equal(t, false, step.SuppressOutput)
	assert.Equal(t, false, step.SuppressesOutput())
}

func TestMixin_UnmarshalUpgradeAction(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/upgrade-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)

	require.Equal(t, 1, len(action.Steps))
	step := action.Steps[0]

	assert.Equal(t, "Tag VM", step.Description)
	require.Empty(t, step.Outputs)

	assert.Equal(t, "ec2", step.Service)
	assert.Equal(t, "create-tags", step.Operation)

	assert.Empty(t, step.Arguments)

	sort.Sort(step.Flags)
	assert.Equal(t, builder.Flags{
		builder.NewFlag("resources", "i-5203422c"),
		builder.NewFlag("tags", "Key=canary,Value=true")}, step.Flags)
}

func TestMixin_UnmarshalInvokeAction(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/invoke-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)

	require.Equal(t, 1, len(action.Steps))
	step := action.Steps[0]

	assert.Equal(t, "List Buckets", step.Description)
	assert.Equal(t, "s3api", step.Service)
	assert.Equal(t, "list-buckets", step.Operation)

	assert.Empty(t, step.Arguments)
	assert.Empty(t, step.Flags)

	require.Len(t, step.Outputs, 1)
	assert.Equal(t, Output{Name: "buckets", JsonPath: "$Buckets[*].Name"}, step.Outputs[0])

}

func TestMixin_UnmarshalUninstallAction(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/uninstall-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)

	require.Equal(t, 1, len(action.Steps))
	step := action.Steps[0]

	assert.Equal(t, "Deprovision VM", step.Description)
	require.Empty(t, step.Outputs)

	assert.Equal(t, "ec2", step.Service)
	assert.Equal(t, "terminate-instances", step.Operation)

	assert.Empty(t, step.Arguments)

	sort.Sort(step.Flags)
	assert.Equal(t, builder.Flags{
		builder.NewFlag("instance-ids", "i-5203422c i-5203422d")}, step.Flags)
}

func TestMixin_UnmarshalStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input.yaml")
	require.NoError(t, err)

	var step Step
	err = yaml.Unmarshal(b, &step)
	require.NoError(t, err)

	assert.Equal(t, "Provision VM", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, Output{"INSTANCE_ID", "$.Instances[0].InstanceId"}, step.Outputs[0])

	assert.Equal(t, "ec2", step.Service)
	assert.Equal(t, "run-instances", step.Operation)

	assert.Equal(t, []string{"myinst"}, step.Arguments)

	sort.Sort(step.Flags)
	assert.Equal(t, builder.Flags{
		builder.NewFlag("env", "FOO=BAR", "STUFF=THINGS"),
		builder.NewFlag("image-id", "ami-xxxxxxxx"),
		builder.NewFlag("instance-type", "t2.micro")}, step.Flags)
}

func TestMixin_UnmarshalInvalidStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input-invalid.yaml")
	require.NoError(t, err)

	var step Step
	err = yaml.Unmarshal(b, &step)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid yaml type for flag env")
}

func TestStep_SuppressesOutput(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input-suppress-output.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, true, step.SuppressOutput)
	assert.Equal(t, true, step.SuppressesOutput())
}
