package aws

import (
	"fmt"

	"github.com/deislabs/porter-aws/pkg"
)

func (m *Mixin) PrintVersion() {
	fmt.Fprintf(m.Out, "aws mixin %s (%s)\n", pkg.Version, pkg.Commit)
}
