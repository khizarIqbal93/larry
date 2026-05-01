package uv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUVVersion(t *testing.T) {
	rawTOML, err := os.ReadFile("pyproject.toml")
	assert.NoError(t, err)
	p, err := parse(string(rawTOML))
	assert.NoError(t, err)
	tag, err := getUVVersion(p)
	assert.NoError(t, err)
	assert.Equal(t, "v0.1.0", tag.String())
}

func TestUpdate(t *testing.T) {
	rawTOML, err := os.ReadFile("pyproject.toml")
	assert.NoError(t, err)

	actual := UpdateUvVersion(string(rawTOML), "v0.1.1")
	p, err := parse(actual)
	assert.NoError(t, err)
	tag, err := getUVVersion(p)
	assert.NoError(t, err)
	assert.Equal(t, "v0.1.1", tag.String())
}
