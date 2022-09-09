package aws

import (
	"strings"

	"get.porter.sh/porter/pkg"
)

// AWS_EXECUTION_ENV is the environment variable used by the aws CLI to set the
// user agent string.
const AWS_EXECUTION_ENV = "AWS_EXECUTION_ENV"

func (m *Mixin) SetUserAgent() {
	value := []string{pkg.UserAgent(), m.UserAgent()}

	if agentStr, ok := m.Config.LookupEnv(AWS_EXECUTION_ENV); ok {
		value = append(value, agentStr)
	}

	m.Config.Setenv(AWS_EXECUTION_ENV, strings.Join(value, " "))
}

func (m *Mixin) UserAgent() string {
	v := m.Version()
	return "getporter/" + v.Name + "/" + v.Version
}
