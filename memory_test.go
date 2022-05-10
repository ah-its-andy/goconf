package goconf_test

import (
	"os"
	"testing"

	"github.com/ah-its-andy/goconf"
	"github.com/stretchr/testify/assert"
)

func TestMemorySourceConstructor(t *testing.T) {
	source := goconf.Memory(nil)
	assert.NotNil(t, source)
	assert.Equal(t, 0, len(source.InitialData))
}

func TestEnvironmentVariableConstructor(t *testing.T) {
	os.Setenv("GOCONF_TEST_ENV", "test")
	defer os.Unsetenv("GOCONF_TEST_ENV")
	source := goconf.EnvironmentVariable("GOCONF_")
	assert.NotNil(t, source)
	assert.Equal(t, 1, len(source.InitialData))
	assert.Equal(t, "test", source.InitialData["GOCONF_TEST_ENV"])
}

func TestBuildMemoryProvider(t *testing.T) {
	source := goconf.Memory(map[string]string{
		"key": "value",
	})
	provider := source.BuildProvider(nil)
	assert.NotNil(t, provider)
	v, ok := provider.GetString("key")
	assert.True(t, ok)
	assert.Equal(t, "value", v)
}

func TestBuildMemoryProviderShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expected")
		}
	}()

	goconf.NewMemoryProvider(nil)
}

func TestBuildMemoryProviderAddShouldPassed(t *testing.T) {
	source := goconf.Memory(nil)
	provider := source.BuildProvider(nil).(*goconf.MemoryProvider)
	provider.Add("key", "value")
	v, ok := provider.GetString("key")
	assert.True(t, ok)
	assert.Equal(t, "value", v)
}
