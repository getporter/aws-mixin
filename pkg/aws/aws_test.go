package aws

import (
	"bytes"
	"io/ioutil"
	"os"
	"sort"
	"testing"

	"github.com/deislabs/porter/pkg/test"
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
	assert.Equal(t, Flags{
		NewFlag("image-id", "ami-xxxxxxxx"),
		NewFlag("instance-type", "t2.micro")}, step.Flags)
}

func TestMixin_UnmarshalUpgradelAction(t *testing.T) {
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
	assert.Equal(t, Flags{
		NewFlag("resources", "i-5203422c"),
		NewFlag("tags", "Key=canary,Value=true")}, step.Flags)
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
	assert.Equal(t, Flags{
		NewFlag("instance-ids", "i-5203422c i-5203422d")}, step.Flags)
}

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

func TestMixin_Execute(t *testing.T) {
	testcases := []struct {
		name    string
		wantCmd string
		step    Step
	}{
		{"args, no flags", "aws s3 mb s3://mybucket --output json",
			Step{Service: "s3", Operation: "mb", Arguments: []string{"s3://mybucket"}},
		},
		{"no args, with flags", "aws ec2 run-instances --image-id ami-xxxxxxxx --instance-type t2.micro --output json",
			Step{Service: "ec2", Operation: "run-instances", Flags: Flags{NewFlag("instance-type", "t2.micro"), NewFlag("image-id", "ami-xxxxxxxx")}},
		},
		{"args and flag", "aws ec2 run-instances myinst --image-id ami-xxxxxxxx --instance-type t2.micro --output json",
			Step{Service: "ec2", Operation: "run-instances", Arguments: []string{"myinst"}, Flags: []Flag{NewFlag("image-id", "ami-xxxxxxxx"), NewFlag("instance-type", "t2.micro")}},
		},
		{"repeated flag", "aws ec2 run-instances --env FOO=BAR --env STUFF=THINGS --output json",
			Step{Service: "ec2", Operation: "run-instances", Flags: Flags{NewFlag("env", "FOO=BAR", "STUFF=THINGS")}},
		},
	}

	defer os.Unsetenv(test.ExpectedCommandEnv)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv(test.ExpectedCommandEnv, tc.wantCmd)

			action := Action{Steps: []Steps{{tc.step}}}
			b, err := yaml.Marshal(action)
			require.NoError(t, err)

			y := string(b)
			t.Log(y)

			h := NewTestMixin(t)
			h.In = bytes.NewReader(b)

			err = h.Execute()

			require.NoError(t, err)
		})
	}
}
