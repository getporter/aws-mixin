package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	m := NewTestMixin(t)

	err := m.Build()
	require.NoError(t, err, "Build failed")

	gotOutput := m.TestContext.GetOutput()

	wantOutput := `RUN apt-get update && apt-get install -y --no-install-recommends curl unzip libc6 less groff
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-$(uname -m).zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN ./aws/install
RUN rm -fr awscliv2.zip ./aws
`

	assert.Equal(t, wantOutput, gotOutput)
}
