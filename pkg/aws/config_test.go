package aws

import (
	"testing"

	"get.porter.sh/mixin/aws/pkg"
	"github.com/stretchr/testify/require"
)

func TestSetUserAgent(t *testing.T) {
	pkg.Commit = "abc123"
	pkg.Version = "v1.2.3"

	m := NewTestMixin(t)
	m.SetUserAgent()

	expected := "getporter/aws/" + pkg.Version
	require.Contains(t, m.Getenv(AWS_EXECUTION_ENV), expected)
}
