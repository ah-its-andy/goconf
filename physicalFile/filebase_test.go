package physicalfile_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	physicalfile "github.com/ah-its-andy/goconf/physicalFile"
	"github.com/stretchr/testify/assert"
)

func TestPhysicalFileConstructorRelativePath(t *testing.T) {
	path := "./testdata/test.conf"
	source := physicalfile.PhysicalFile(path, nil)
	execPath, _ := filepath.Abs("./physicalFile")
	assert.Equal(t, filepath.Join(filepath.Dir(execPath), path), source.FilePath)
}

func TestPhysicalFileNotExists(t *testing.T) {
	path := "./testdata/notexists.conf"
	source := physicalfile.PhysicalFile(path, nil)
	err := source.ResolveFile()
	assert.Equal(t, os.ErrNotExist, err)
}

func TestLoadYaml(t *testing.T) {
	reader := bytes.NewReader([]byte(
		`name: "test"
age: 18
# comment`))
	source := physicalfile.YamlReader(reader)
	provider := source.BuildProvider(nil)
	err := provider.Load()
	assert.Nil(t, err)
	v, ok := provider.GetString("name")
	assert.True(t, ok)
	assert.Equal(t, "test", v)
	v, ok = provider.GetString("age")
	assert.True(t, ok)
	assert.Equal(t, "18", v)
}

func TestLoadJson(t *testing.T) {
	reader := bytes.NewReader([]byte(
		`{"name": "test", "age": 18}`))
	source := physicalfile.JsonReader(reader)
	provider := source.BuildProvider(nil)
	err := provider.Load()
	assert.Nil(t, err)
	v, ok := provider.GetString("name")
	assert.True(t, ok)
	assert.Equal(t, "test", v)
	v, ok = provider.GetString("age")
	assert.True(t, ok)
	assert.Equal(t, "18", v)
}
