package aws

type Step struct {
	Description string            `yaml:"description"`
	Service     string            `yaml:"service"`
	Operation   string            `yaml:"operation"`
	Arguments   []string          `yaml:"arguments"`
	Flags       map[string]string `yaml:"flags"`
	Outputs     []Output          `yaml:"outputs"`
}

type Output struct {
	Name     string `yaml:"name"`
	JsonPath string `yaml:"jsonPath"`
}
