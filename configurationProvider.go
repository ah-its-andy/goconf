package goconf

import (
	"sort"
	"strings"
)

var _ Provider = (*ConfigurationProvider)(nil)

type ConfigurationProvider struct {
	Data map[string]*ExtractedValue
}

func NewConfigurationProvider() *ConfigurationProvider {
	return &ConfigurationProvider{
		Data: make(map[string]*ExtractedValue),
	}
}

func (provider *ConfigurationProvider) Load() error {
	return nil
}

// 尝试获取一个配置key的值
func (provider *ConfigurationProvider) GetString(name string) (string, bool) {
	if v, ok := provider.GetExtracted(name); ok {
		return v.Value, true
	}
	return "", false
}

func (provider *ConfigurationProvider) GetExtracted(name string) (*ExtractedValue, bool) {
	value, ok := provider.Data[name]
	if value == nil {
		return nil, false
	}
	return value, ok
}

// 基于此返回给定父路径的直接后代配置键
func (provider *ConfigurationProvider) GetChildKeys(path string, earlierKeys ...string) []string {
	results := make([]string, 0)
	if len(path) == 0 {
		for key, _ := range provider.Data {
			results = append(results, Segment(key, 0))
		}
	} else {
		for key, _ := range provider.Data {
			if len(key) > len(path) &&
				strings.HasPrefix(key, path) &&
				string(key[len(path)]) == KeyDelimiter {
				results = append(results, Segment(key, len(path)+1))
			}
		}
	}
	results = append(results, earlierKeys...)
	sort.Strings(results)
	return results
}

func Segment(key string, prefixLength int) string {
	indexOf := strings.Index(key[prefixLength:], KeyDelimiter)
	if indexOf < 0 {
		return key[prefixLength:]
	} else {
		return key[prefixLength : indexOf-prefixLength]
	}
}
