package goconf

import (
	"os"
	"strings"
)

type MemorySource struct {
	InitialData map[string]string
}

func Memory(initialData map[string]string) *MemorySource {
	if initialData == nil {
		initialData = make(map[string]string)
	}
	return &MemorySource{
		InitialData: initialData,
	}
}

func EnvironmentVariable(prefix string) *MemorySource {
	data := make(map[string]string)
	for _, env := range os.Environ() {
		if len(prefix) > 0 && !strings.HasPrefix(env, prefix) {
			continue
		}
		kv := strings.SplitN(env, "=", 2)
		if len(kv) != 2 {
			continue
		}
		data[kv[0]] = kv[1]
	}
	return Memory(data)
}

func (source *MemorySource) BuildProvider(builder Builder) Provider {
	return NewMemoryProvider(source)
}

type MemoryProvider struct {
	ConfigurationProvider

	source *MemorySource
}

func NewMemoryProvider(source *MemorySource) *MemoryProvider {
	if source == nil {
		panic("goconf.NewMemoryProvider: source is nil")
	}
	provider := &MemoryProvider{
		ConfigurationProvider: *NewConfigurationProvider(),
		source:                source,
	}
	if source.InitialData == nil {
		return provider
	}
	for key, value := range source.InitialData {
		provider.Data[key] = NewExtractedValue(key, value, value)
	}
	return provider
}

func (provider *MemoryProvider) Add(key, value string) {
	provider.Data[key] = NewExtractedValue(key, value, value)
}
