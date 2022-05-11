package physicalfile_test

import (
	"strings"
	"testing"

	physicalfile "github.com/ah-its-andy/goconf/physicalFile"
	"github.com/stretchr/testify/assert"
)

func TestYamlLoad(t *testing.T) {
	data := `k1: v1
section1:
   key1: value1
   key2: value2
   key3: value3`

	reader := strings.NewReader(data)
	provider := physicalfile.YamlReader(reader)
	ret := provider.BuildProvider(nil)
	err := ret.Load()
	assert.Nil(t, err)
	k, ok := ret.GetString("k1")
	assert.True(t, ok)
	assert.Equal(t, "v1", k)

	k, ok = ret.GetString("section1.key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", k)

	k, ok = ret.GetString("section1.key2")
	assert.True(t, ok)
	assert.Equal(t, "value2", k)

	k, ok = ret.GetString("section1.key3")
	assert.True(t, ok)
	assert.Equal(t, "value3", k)

}
