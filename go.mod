module get.porter.sh/mixin/aws

go 1.13

require (
	get.porter.sh/porter v0.37.2-0.20210408130546-e8f54a713426
	github.com/ghodss/yaml v1.0.0
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.6.1
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/term v0.0.0-20210406210042-72f3dc4e9b72 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace get.porter.sh/porter => github.com/carolynvs/porter v0.37.1-0.20210408135241-db46041b5299

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
