package aws

import (
	"get.porter.sh/porter/pkg/portercontext"
)

type Mixin struct {
	*portercontext.Context
}

// New aws mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	m := &Mixin{
		Context: portercontext.New(),
	}
	m.SetUserAgent()

	return m, nil
}
