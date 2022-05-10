package goconf

import (
	"os"
	"strings"
)

type MemoryConfSource struct {
	InitialData map[string]string
}

func Memory(initialData map[string]string) *MemoryConfSource {
	if initialData == nil {
		initialData = make(map[string]string)
	}
	return &MemoryConfSource{
		InitialData: initialData,
	}
}

func EnvironmentVariable(prefix string) *MemoryConfSource {
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

func (source *MemoryConfSource) BuildProvider(builder Builder) Provider {
	return NewMemoryConfProvider(source)
}

type MemoryConfProvider struct {
	ConfigurationProvider

	source *MemoryConfSource
}

func NewMemoryConfProvider(source *MemoryConfSource) *MemoryConfProvider {
	provider := &MemoryConfProvider{
		ConfigurationProvider: *NewConfigurationProvider(),
		source:                source,
	}
	if source == nil || source.InitialData == nil {
		return provider
	}
	for key, value := range source.InitialData {
		provider.Data[key] = NewExtractedValue(key, value, value)
	}
	return provider
}

func (provider *MemoryConfProvider) Add(key, value string) {
	provider.Data[key] = NewExtractedValue(key, value, value)
}
