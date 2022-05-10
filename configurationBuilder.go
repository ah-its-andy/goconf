package goconf

var _ Builder = (*ConfigurationBuilder)(nil)

type ConfigurationBuilder struct {
	sources    []Source
	properties map[string]interface{}
}

func NewBuilder() Builder {
	return &ConfigurationBuilder{
		properties: make(map[string]interface{}),
		sources:    make([]Source, 0),
	}
}

func (builder *ConfigurationBuilder) Properties() map[string]interface{} {
	return builder.properties
}
func (builder *ConfigurationBuilder) Sources() []Source {
	return builder.sources
}
func (builder *ConfigurationBuilder) AddSource(source Source) Builder {
	builder.sources = append(builder.sources, source)
	return builder
}
func (builder *ConfigurationBuilder) BuildRoot() Root {
	providers := make([]Provider, 0)
	for _, source := range builder.sources {
		providers = append(providers, source.BuildProvider(builder))
	}
	return NewRoot(providers)
}
