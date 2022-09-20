package aws

import (
	"get.porter.sh/porter/pkg/runtime"
)

type Mixin struct {
	// Config is a specialized context with additional runtime settings.
	runtime.RuntimeConfig
}

// New aws mixin client, initialized with useful defaults.
func New() *Mixin {
	m := &Mixin{
		RuntimeConfig: runtime.NewConfig(),
	}
	m.SetUserAgent()

	return m
}
