package files

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadYaml(yamlFile string, config interface{}) (interface{}, error) {
	if b, err := ioutil.ReadFile(yamlFile); err != nil {
		return nil, err
	} else {
		if err := yaml.Unmarshal(b, config); err != nil {
			return nil, err
		}
		return config, nil
	}
}
