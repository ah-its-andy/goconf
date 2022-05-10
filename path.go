package goconf

import "strings"

var KeyDelimiter = "."

func CombinePath(pathSegments ...string) string {
	if len(pathSegments) == 0 {
		panic("pathSegments is empty")
	}
	paths := make([]string, 0)
	for _, path := range pathSegments {
		if len(path) == 0 {
			continue
		}
		paths = append(paths, path)
	}
	return strings.Join(paths, KeyDelimiter)
}

func GetSectionKey(path string) string {
	if len(path) == 0 {
		panic("path is empty")
	}
	lastDelimiterIndex := strings.LastIndex(path, KeyDelimiter)
	if lastDelimiterIndex == -1 {
		return path
	}
	return path[lastDelimiterIndex+1:]
}

func GetParentPath(path string) (string, bool) {
	if len(path) == 0 {
		panic("path is empty")
	}
	lastDelimiterIndex := strings.LastIndex(path, KeyDelimiter)
	if lastDelimiterIndex == -1 {
		return "", false
	} else {
		return path[:lastDelimiterIndex], true
	}
}
