package aws

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMixin_Build(t *testing.T) {
	m := NewTestMixin(t)

	err  := m.Build()
	require.NoError(t, err, "Build failed")

	gotOutput := m.TestContext.GetOutput()

	wantOutput := `RUN apt-get update && apt-get install -y curl unzip python less groff
RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "/tmp/awscli-bundle.zip"
RUN unzip /tmp/awscli-bundle.zip -d /tmp
RUN /tmp/awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws
RUN rm -fr /tmp/awscli-bundle.zip /tmp/awscli-bundle
`

	assert.Equal(t, wantOutput, gotOutput)
}