package uv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUVVersion(t *testing.T) {
	rawTOML, err := os.ReadFile("pyproject.toml")
	assert.NoError(t, err)
	p := parse(string(rawTOML))
	tag := getUVVersion(p)
	assert.Equal(t, "v0.1.0", tag.String())
}

func TestUpdate(t *testing.T) {
	rawTOML, err := os.ReadFile("pyproject.toml")
	assert.NoError(t, err)

	actual := UpdateUvVersion(string(rawTOML), "patch")
	os.WriteFile("cuh.toml", []byte(actual), 0644)
	p := parse(actual)
	tag := getUVVersion(p)
	assert.Equal(t, "v0.1.1", tag.String())
}
