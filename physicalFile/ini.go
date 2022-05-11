package physicalfile

import (
	"bufio"
	"io"
	"strings"

	"github.com/ah-its-andy/goconf"
)

func Ini(filePath string) *PhysicalFileSource {
	return PhysicalFile(filePath, ResolvePhysicalFile(loadIniFile))
}

func IniReader(reader io.Reader) *PhysicalFileSource {
	return FromReader(reader, ResolvePhysicalFile(loadIniFile))
}

func loadIniFile(reader io.Reader) (map[string]*goconf.ExtractedValue, error) {
	data, err := ResolveIni(reader)
	if err != nil {
		return nil, err
	}

	return goconf.ExtractStructToMap(data, ""), nil
}

func ResolveIni(reader io.Reader) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	scanner := bufio.NewScanner(reader)
	var sectionName string
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t")
		if len(line) == 0 {
			continue
		}
		if string(line[0]) == "[" && string(line[len(line)-1]) == "]" {
			// section
			sectionName = line[1 : len(line)-1]
		} else if string(line[0]) == ";" {
			// comment
			continue
		} else {
			// key-value
			index := strings.Index(line, "=")
			if index == -1 {
				continue
			}

			key := strings.Trim(line[:index], " \t")
			value := strings.Trim(line[index+1:], " \t")

			if sectionName == "" {
				ret[key] = value
			} else {
				if _, ok := ret[sectionName]; !ok {
					ret[sectionName] = make(map[string]interface{})
				}
				ret[sectionName].(map[string]interface{})[key] = value
			}
		}
	}
	return ret, scanner.Err()
}
