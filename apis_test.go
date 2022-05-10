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
