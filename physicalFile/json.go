package physicalfile

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/ah-its-andy/goconf"
)

func Json(filePath string) *PhysicalFileSource {
	return PhysicalFile(filePath, ResolvePhysicalFile(loadJsonlFile))
}

func JsonReader(reader io.Reader) *PhysicalFileSource {
	return FromReader(reader, ResolvePhysicalFile(loadJsonlFile))
}

func loadJsonlFile(reader io.Reader) (map[string]*goconf.ExtractedValue, error) {
	fileBytes, err := ioutil.ReadAll(reader)
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
