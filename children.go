package goconf

func GetChildrenFromRoot(root Root, path string) []Section {
	seed := make([]string, 0)
	providers := root.GetProviders()
	for _, provider := range providers {
		keys := provider.GetChildKeys(path, seed...)
		seed = append(seed, keys...)
	}
	seed = distinctStringSlice(seed)
	sections := make([]Section, len(seed))
	for i, key := range seed {
		sections[i] = root.GetSection(key)
	}
	return sections
}

func distinctStringSlice(slice []string) []string {
	m := make(map[string]bool)
	for _, s := range slice {
		m[s] = true
	}
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}
