package goconf

var _ Root = (*ConfigurationRoot)(nil)

type ConfigurationRoot struct {
	Providers []Provider
}

func NewRoot(providers []Provider) *ConfigurationRoot {
	return &ConfigurationRoot{
		Providers: providers,
	}
}

func (root *ConfigurationRoot) GetProviders() []Provider {
	return root.Providers
}
func (root *ConfigurationRoot) Reload() error {
	for _, provider := range root.Providers {
		err := provider.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func (root *ConfigurationRoot) GetString(name string) (string, bool) {
	return GetConfiguration(root.Providers, name)
}
func (root *ConfigurationRoot) GetExtracted(name string) (*ExtractedValue, bool) {
	return GetConfigExtractedValue(root.Providers, name)
}

func (root *ConfigurationRoot) GetSection(name string) Section {
	return NewSection(root, name)
}
func (root *ConfigurationRoot) GetChildren() []Section {
	return GetChildrenFromRoot(root, "")
}

func GetConfiguration(providers []Provider, name string) (string, bool) {
	for i := len(providers) - 1; i >= 0; i-- {
		provider := providers[i]
		if val, ok := provider.GetString(name); ok {
			return val, ok
		}
	}
	return "", false
}

func GetConfigExtractedValue(providers []Provider, name string) (*ExtractedValue, bool) {
	for i := len(providers) - 1; i >= 0; i-- {
		provider := providers[i]
		if val, ok := provider.GetExtracted(name); ok {
			return val, ok
		}
	}
	return nil, false
}
