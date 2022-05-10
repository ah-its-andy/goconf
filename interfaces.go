package goconf

import (
	"fmt"
	"reflect"
)

type Configuration interface {
	GetString(name string) (string, bool)
	GetExtracted(name string) (*ExtractedValue, bool)
	GetSection(name string) Section
	GetChildren() []Section
}

type Section interface {
	Configuration

	GetName() string
	GetPath() string
	GetValue() (string, bool)
	GetRaw() (interface{}, bool)
	Bind(recv interface{}) error
}

type Provider interface {
	Load() error
	// 尝试获取一个配置key的值
	GetString(name string) (string, bool)
	GetExtracted(name string) (*ExtractedValue, bool)
	// 基于此返回给定父路径的直接后代配置键
	GetChildKeys(path string, earlierKeys ...string) []string
}

type Builder interface {
	Properties() map[string]interface{}
	Sources() []Source
	AddSource(source Source) Builder
	BuildRoot() Root
}

type Source interface {
	BuildProvider(builder Builder) Provider
}

type Root interface {
	Configuration

	Reload() error
	GetProviders() []Provider
}

type ExtractedValue struct {
	Name  string
	Value string
	Raw   interface{}
}

func NewExtractedValue(name string, value string, raw interface{}) *ExtractedValue {
	return &ExtractedValue{
		Name:  name,
		Value: value,
		Raw:   raw,
	}
}

func ExtractStructToMap(data interface{}, path string) map[string]*ExtractedValue {
	ret := make(map[string]*ExtractedValue)
	valueType := reflect.TypeOf(data)
	if valueType.Kind() == reflect.Map {
		ret[path] = NewExtractedValue(path, "", data)
		subMap := ExtractMap(data, path)
		for subKey, subValue := range subMap {
			ret[subKey] = subValue
		}
	} else if valueType.Kind() == reflect.Slice {
		ret[path] = NewExtractedValue(path, "", data)
		sliceValue := reflect.ValueOf(data)
		for i := 0; i < sliceValue.Len(); i++ {
			structValue := sliceValue.Index(i)
			sliceKey := CombinePath(path, fmt.Sprintf("$%d", i))
			subMap := ExtractStructToMap(structValue.Interface(), sliceKey)
			for subKey, subValue := range subMap {
				ret[subKey] = subValue
			}
		}
	} else if valueType.Kind() == reflect.Ptr {
		ptrValue := reflect.ValueOf(data)
		if ptrValue.IsNil() {
			return nil
		}
		subMap := ExtractStructToMap(ptrValue.Elem().Interface(), path)
		for subKey, subValue := range subMap {
			ret[subKey] = subValue
		}
	} else {
		ret[path] = NewExtractedValue(path, fmt.Sprintf("%v", data), data)

	}
	return ret
}

func ExtractMap(data interface{}, path string) map[string]*ExtractedValue {
	ret := make(map[string]*ExtractedValue)
	mapRange := reflect.ValueOf(data).MapRange()
	for mapRange.Next() {
		var subPath string
		key := mapRange.Key()
		v := mapRange.Value()
		if k, ok := key.Interface().(string); ok {
			subPath = CombinePath(path, k)
		} else {
			subPath = CombinePath(path, fmt.Sprintf("%v", key.Interface()))
		}
		subMap := ExtractStructToMap(v.Interface(), subPath)
		if subMap == nil || len(subMap) == 0 {
			continue
		}
		for subKey, subValue := range subMap {
			ret[subKey] = subValue
		}
		continue
	}

	return ret
}
