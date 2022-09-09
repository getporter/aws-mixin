package aws

import (
	"strings"

	"get.porter.sh/porter/pkg"
)

const AWS_EXECUTION_ENV = "AWS_EXECUTION_ENV"

func (m *Mixin) SetUserAgent() {
	value := []string{pkg.UserAgent(), m.UserAgent()}

	if agentStr, ok := m.LookupEnv(AWS_EXECUTION_ENV); ok {
		value = append(value, agentStr)
	}

	m.Setenv(AWS_EXECUTION_ENV, strings.Join(value, " "))
}

func (m *Mixin) UserAgent() string {
	v := m.Version()
	return "getporter/" + v.Name + "/" + v.Version
}
