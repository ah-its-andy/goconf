package goconf

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var _ Section = (*ConfigurationSection)(nil)

type ConfigurationSection struct {
	root Root
	path string
}

func NewSection(root Root, path string) *ConfigurationSection {
	return &ConfigurationSection{
		root: root,
		path: path,
	}
}

func (section *ConfigurationSection) GetName() string {
	return GetSectionKey(section.path)
}
func (section *ConfigurationSection) GetPath() string {
	return section.path
}
func (section *ConfigurationSection) GetValue() (string, bool) {
	return section.root.GetString(section.path)
}

func (section *ConfigurationSection) GetString(name string) (string, bool) {
	return section.root.GetString(CombinePath(section.path, name))
}

func (section *ConfigurationSection) GetExtracted(name string) (*ExtractedValue, bool) {
	return section.root.GetExtracted(CombinePath(section.path, name))
}
func (section *ConfigurationSection) GetSection(name string) Section {
	return section.root.GetSection(CombinePath(section.path, name))
}
func (section *ConfigurationSection) GetChildren() []Section {
	return GetChildrenFromRoot(section.root, section.path)
}

func (section *ConfigurationSection) GetRaw() (interface{}, bool) {
	if extracted, ok := section.root.GetExtracted(section.path); ok {
		return extracted.Raw, true
	}
	return nil, false
}
func (section *ConfigurationSection) Bind(recv interface{}) error {
	if extracted, ok := section.root.GetExtracted(section.path); !ok {
		return fmt.Errorf("section %s not found", section.path)
	} else {
		rawType := reflect.TypeOf(extracted.Raw)

		if rawType.Kind() == reflect.Map {
			return bindMap(recv, extracted.Raw.(map[string]interface{}))
		} else {
			return fmt.Errorf("Not supported type %s", rawType.Kind())
		}
	}
}

func bindMap(recv interface{}, mapValue map[string]interface{}) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &recv,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(mapValue)
	if err != nil {
		return err
	}
	return nil
}
