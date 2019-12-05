//go:generate packr2
package aws

import (
	"get.porter.sh/porter/pkg/context"
)

type Mixin struct {
	*context.Context
}

// New aws mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	return &Mixin{
		Context: context.New(),
	}, nil

}
