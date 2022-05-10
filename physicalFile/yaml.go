package physicalfile

import (
	"io/ioutil"
	"os"

	"github.com/ah-its-andy/goconf"
	yaml "gopkg.in/yaml.v2"
)

func Yaml(filePath string) *PhysicalFileSource {
	return PhysicalFile(filePath, loadYamlFile)
}

func loadYamlFile(f *os.File) (map[string]*goconf.ExtractedValue, error) {
	fileBytes, err := ioutil.ReadAll(f)
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
