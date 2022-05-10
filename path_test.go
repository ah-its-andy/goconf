package goconf_test

import (
	"testing"

	"github.com/ah-its-andy/goconf"
	"github.com/stretchr/testify/assert"
)

func TestCombinePathShouldPassed(t *testing.T) {
	paths := []string{"a", "b", "c"}
	expected := "a.b.c"
	actual := goconf.CombinePath(paths...)
	assert.Equal(t, expected, actual)
}

func TestCombinePathShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	goconf.CombinePath()
}

func TestGetSectionKeyShouldPassed(t *testing.T) {
	path := "a.b.c"
	expected := "c"
	actual := goconf.GetSectionKey(path)
	assert.Equal(t, expected, actual)
}

func TestGetSectionNoKeyDelimiterShouldPassed(t *testing.T) {
	path := "a:b"
	expected := "a:b"
	actual := goconf.GetSectionKey(path)
	assert.Equal(t, expected, actual)
	assert.NotSame(t, &expected, &actual)
}

func TestGetSectionShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	goconf.GetSectionKey("")
}

func TestGetParentPathShouldPassed(t *testing.T) {
	path := "a.b.c"
	expected := "a.b"
	actual, ok := goconf.GetParentPath(path)
	assert.True(t, ok)
	assert.Equal(t, expected, actual)
}

func TestGetParentFalseShouldPassed(t *testing.T) {
	path := "a"
	expected := ""
	actual, ok := goconf.GetParentPath(path)
	assert.False(t, ok)
	assert.Equal(t, expected, actual)
}

func TestGetParentShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	goconf.GetParentPath("")
}
