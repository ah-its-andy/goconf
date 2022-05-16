# goconf
Configuration providers for go inspired by .net core configuration libs. Extensible and easy to use.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ah-its-andy/goconf?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/ah-its-andy/goconf)](https://github.com/gookit/goutil)
[![GoDoc](https://godoc.org/github.com/ah-its-andy/goconf?status.svg)](https://pkg.go.dev/github.com/ah-its-andy/goconf)
[![Go Report Card](https://goreportcard.com/badge/github.com/ah-its-andy/goconf)](https://goreportcard.com/report/github.com/ah-its-andy/goconf)
[![Unit-Tests](https://github.com/ah-its-andy/goconf/workflows/Unit-Tests/badge.svg)](https://github.com/ah-its-andy/goconf/actions)

## Supportives:
- Memory source
- Environment variables
- JSON file
- YAML file
- INI file

## Features:
- Multi providers support (merge sources)
- Extensible providers support (new providers can be added easily, even only one function to implement a new provider using file from local filesystem)
- Get value by full path with key delimiter (e.g. `application.bind_addr.port`)
- Cast value to specialized type with `TypeConversionFunc`
- Bind configuration section to struct

## Usage
- package: `github.com/ah-its-andy/goconf`
```go
// initialize on application startup
goconf.Init(func(b goconf.Builder) {
	b.AddSource(physicalfile.Yaml(/*yaml file path, absolute or relative  both supported*/)))
     .AddSource(physicalfile.Json(/*json file path, absolute or relative  both supported*/))
     .AddSource(goconf.EnvironmentVariable(/*prefix for filter environment variables*/))
     .AddSource(goconf.Memory(/*config map*/))
})

// use it anywhere
bindAddr, ok := goconf.GetString("application.bind_addr.addr") //Get string value

bindAddrWithDefault := goconf.GetStringOrDefault("application.bind_addr.addr", "default value") //returns default value when key is not found

castValue, ok := goconf.Cast("application.bind_addr.port", goconf.IntConversion) //cast value to int

castValueWithDefault := goconf.CastOrDefault("application.bind_addr.port", 0 /*default value*/, goconf.IntConversion) //cast value to int, returns default value when key is not found

section:= gocinf.GetSection("application") //get section

var application fakeStruct.Application
err := section.Bind(&application) //bind section to struct

```

## Refer
- [github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) for binding struct
- [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) for yaml file support
- [github.com/stretchr/testify](https://github.com/stretchr/testify) for testing
