package physicalfile

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ah-its-andy/goconf"
)

func Json(filePath string) *PhysicalFileSource {
	return PhysicalFile(filePath, loadYamlFile)
}

func loadJsonlFile(f *os.File) (map[string]*goconf.ExtractedValue, error) {
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var recv map[string]interface{}
	err = json.Unmarshal(fileBytes, &recv)
	if err != nil {
		return nil, err
	}

	return goconf.ExtractStructToMap(recv, ""), nil
}
