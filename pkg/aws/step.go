package aws

import "github.com/pkg/errors"

type Step struct {
	Description string   `yaml:"description"`
	Service     string   `yaml:"service"`
	Operation   string   `yaml:"operation"`
	Arguments   []string `yaml:"arguments"`
	Flags       Flags    `yaml:"flags"`
	Outputs     []Output `yaml:"outputs"`
}

type Flags []Flag
type Flag struct {
	Name   string
	Values []string
}

func NewFlag(name string, values ...string) Flag {
	f := Flag{
		Name:   name,
		Values: make([]string, len(values)),
	}
	copy(f.Values, values)
	return f
}

// UnmarshalYAML takes any yaml in this form
func (flags *Flags) UnmarshalYAML(unmarshal func(interface{}) error) error {
	flagMap := map[interface{}]interface{}{}
	err := unmarshal(&flagMap)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal yaml into Step.Flags")
	}

	*flags = make(Flags, 0, len(flagMap))
	for k, v := range flagMap {
		f := Flag{}
		f.Name = k.(string)

		switch t := v.(type) {
		case string:
			f.Values = make([]string, 1)
			f.Values[0] = v.(string)
		case []interface{}:
			f.Values = make([]string, len(t))
			for i := range t {
				iv, ok := t[i].(string)
				if !ok {
					return errors.Errorf("invalid yaml type for flag %s: %T", f.Name, t[i])
				}
				f.Values[i] = iv
			}
		default:
			return errors.Errorf("invalid yaml type for flag %s: %T", f.Name, t)
		}

		*flags = append(*flags, f)
	}
	return nil
}

type Output struct {
	Name     string `yaml:"name"`
	JsonPath string `yaml:"jsonPath"`
}
