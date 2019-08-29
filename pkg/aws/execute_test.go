package aws

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/deislabs/porter/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

func TestMixin_Execute(t *testing.T) {
	m := NewTestMixin(t)

	testcases := []struct {
		name        string
		file        string
		wantCommand string
	}{
		{"install", "testdata/install-input.yaml", "aws ec2 run-instances myinst --image-id ami-xxxxxxxx --instance-type t2.micro --output json"},
		{"upgrade", "testdata/upgrade-input.yaml", "aws ec2 create-tags --output json --resources i-5203422c --tags Key=canary,Value=true"},
		{"invoke", "testdata/invoke-input.yaml", "aws s3api list-buckets --output json"},
		{"uninstall", "testdata/uninstall-input.yaml", `aws ec2 terminate-instances --instance-ids "i-5203422c i-5203422d" --output json`},
	}

	defer os.Unsetenv(test.ExpectedCommandEnv)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv(test.ExpectedCommandEnv, tc.wantCommand)
			mixinInputB, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			err = m.Execute()
			require.NoError(t, err, "execute failed")
		})
	}
}
