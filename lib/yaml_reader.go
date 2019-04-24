package lib

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type YamlReader struct {
}

func NewYamlReader(file string) string {
	jsonFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	j, err := yaml.YAMLToJSON(jsonFile)
	if err != nil {
		panic(err)
	}

	return string(j)
}
