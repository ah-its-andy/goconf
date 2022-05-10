package goconf_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/ah-its-andy/goconf"
	physicalfile "github.com/ah-its-andy/goconf/physicalFile"
	"github.com/stretchr/testify/assert"
)

func newTestReader() io.Reader {
	return bytes.NewReader([]byte(
		`name: "test"
age: 18
sec:
   key: "value"
# comment`))
}

func initTestConfig(reader io.Reader) {
	goconf.Init(func(b goconf.Builder) {
		b.AddSource(physicalfile.YamlReader(reader))
	})
}

func TestGetSectionShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())
	sec := goconf.GetSection("sec")
	assert.NotNil(t, sec)

	v, ok := sec.GetString("key")
	assert.True(t, ok)
	assert.Equal(t, "value", v)
}

func TestSectionBindMapShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())
	sec := goconf.GetSection("sec")
	assert.NotNil(t, sec)

	var recv map[string]interface{}
	err := sec.Bind(&recv)
	assert.Nil(t, err)
	assert.Equal(t, "value", recv["key"])
}

type bindTestStrcut struct {
	Key string `json:"key"`
}

func TestSectionBindStructShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())
	sec := goconf.GetSection("sec")
	assert.NotNil(t, sec)

	var recv bindTestStrcut
	err := sec.Bind(&recv)
	assert.Nil(t, err)
	assert.Equal(t, "value", recv.Key)
}

func TestGetStringShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())

	v, ok := goconf.GetString("name")
	assert.True(t, ok)
	assert.Equal(t, "test", v)

	v, ok = goconf.GetString("age")
	assert.True(t, ok)
	assert.Equal(t, "18", v)
}

func TestGetStringNotExistsShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())

	v, ok := goconf.GetString("notexists")
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestGetDefaultShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())

	v := goconf.GetStringOrDefault("notexists", "default")
	assert.Equal(t, "default", v)
}

func TestCastToIntShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())

	v, ok := goconf.Cast("age", goconf.IntConversion)

	assert.True(t, ok)
	assert.Equal(t, 18, v)
}

func TestCastOrDefaultShouldPassed(t *testing.T) {
	initTestConfig(newTestReader())

	v := goconf.CastOrDefault("notexists", 18, goconf.IntConversion)
	assert.Equal(t, 18, v)
}

func TestCastNotExists(t *testing.T) {
	initTestConfig(newTestReader())

	v, ok := goconf.Cast("notexists", goconf.IntConversion)
	assert.False(t, ok)
	assert.Equal(t, nil, v)
}

func TestCastShouldPanicBeforeInitialized(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, "configuration is not initialized", r)
		}
	}()

	goconf.Cast("name", goconf.IntConversion)
}
