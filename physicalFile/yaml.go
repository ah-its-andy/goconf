package physicalfile

import (
	"io"
	"io/ioutil"

	"github.com/ah-its-andy/goconf"
	yaml "gopkg.in/yaml.v2"
)

func Yaml(filePath string) *PhysicalFileSource {
	return PhysicalFile(filePath, ResolvePhysicalFile(loadYamlFile))
}

func YamlReader(reader io.Reader) *PhysicalFileSource {
	return FromReader(reader, ResolvePhysicalFile(loadYamlFile))
}

func loadYamlFile(reader io.Reader) (map[string]*goconf.ExtractedValue, error) {
	fileBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var recv map[string]interface{}
	err = yaml.Unmarshal(fileBytes, &recv)
	if err != nil {
		return nil, err
	}

	return goconf.ExtractStructToMap(recv, ""), nil
}
