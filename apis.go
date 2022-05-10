package goconf

import (
	"strconv"
	"strings"
)

// Do not use this directly.
var initialized Root

type TypeConversionFunc func(string) (interface{}, error)

var (
	IntConversion TypeConversionFunc = func(s string) (interface{}, error) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		return int(v), nil
	}

	FloatConversion TypeConversionFunc = func(s string) (interface{}, error) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		return v, nil
	}

	BooleanConversion TypeConversionFunc = func(s string) (interface{}, error) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return strings.ToLower(s) == "true", nil
		}
		return v > 0, nil
	}
)

func Init(fn func(Builder)) {
	builder := NewBuilder()
	fn(builder)
	initialized = builder.BuildRoot()
	err := initialized.Reload()
	if err != nil {
		panic(err)
	}
}

func panicIfNotInitialized() {
	if initialized == nil {
		panic("configuration is not initialized")
	}
}

func CastOrDefault(name string, defaultValue interface{}, fn TypeConversionFunc) interface{} {
	if v, ok := Cast(name, fn); ok {
		return v
	} else {
		return defaultValue
	}
}

func Cast(name string, fn TypeConversionFunc) (interface{}, bool) {
	if v, ok := GetString(name); ok {
		if value, err := fn(v); err == nil {
			return value, true
		}
	}

	return nil, false
}

func GetString(name string) (string, bool) {
	panicIfNotInitialized()

	return initialized.GetString(name)
}

func GetStringOrDefault(name string, defaultValue string) string {
	panicIfNotInitialized()

	if value, ok := initialized.GetString(name); ok {
		return value
	} else {
		return defaultValue
	}
}

func GetSection(name string) Section {
	panicIfNotInitialized()
	return initialized.GetSection(name)
}

func ConfigRoot() Root {
	panicIfNotInitialized()
	return initialized
}
