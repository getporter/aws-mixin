package aws

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"get.porter.sh/porter/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

func TestMixin_Execute(t *testing.T) {

	testcases := []struct {
		name        string
		file        string
		wantOutput  string
		wantCommand string
	}{
		{"install", "testdata/install-input.yaml", "INSTANCE_ID",
			"aws ec2 run-instances myinst --image-id ami-xxxxxxxx --instance-type t2.micro --output json"},
		{"upgrade", "testdata/upgrade-input.yaml", "",
			"aws ec2 create-tags --output json --resources i-5203422c --tags Key=canary,Value=true"},
		{"invoke", "testdata/invoke-input.yaml", "buckets",
			"aws s3api list-buckets --output json"},
		{"uninstall", "testdata/uninstall-input.yaml", "",
			`aws ec2 terminate-instances --instance-ids i-5203422c i-5203422d --output json`},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := NewTestMixin(t)

			m.Setenv(test.ExpectedCommandEnv, tc.wantCommand)
			mixinInputB, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			err = m.Execute()
			require.NoError(t, err, "execute failed")

			if tc.wantOutput == "" {
				outputs, _ := m.FileSystem.ReadDir("/cnab/app/porter/outputs")
				assert.Empty(t, outputs, "expected no outputs to be created")
			} else {
				wantPath := path.Join("/cnab/app/porter/outputs", tc.wantOutput)
				exists, _ := m.FileSystem.Exists(wantPath)
				assert.True(t, exists, "output file was not created %s", wantPath)
			}
		})
	}
}
