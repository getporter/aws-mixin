package aws

import (
	"get.porter.sh/mixin/aws/pkg"
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	return version.PrintVersion(m.Config.Context, opts, m.Version())
}

func (m *Mixin) Version() mixin.Metadata {
	return mixin.Metadata{
		Name: "aws",
		VersionInfo: pkgmgmt.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Porter Authors",
		},
	}
}
