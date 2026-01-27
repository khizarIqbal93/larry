package uv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawTOML = `[project]
name = "nesoai-utils-api"
version = "0.1.0"
description = "Add your description here"
readme = "README.md"
requires-python = ">=3.14"
dependencies = [
    "aiohttp>=3.13.2",
    "azure-storage-blob>=12.27.1",
    "fastapi[standard]>=0.125.0",
    "python-dotenv>=1.2.1",
]
`

func TestGetUVVersion(t *testing.T) {
	p := parse(rawTOML)
	tag := getUVVersion(p)
	assert.Equal(t, "v0.1.0", tag.String())
}

func TestUpdate(t *testing.T) {
	const expected = `[project]
name = "nesoai-utils-api"
version = "0.1.1"
description = "Add your description here"
readme = "README.md"
requires-python = ">=3.14"
dependencies = [
    "aiohttp>=3.13.2",
    "azure-storage-blob>=12.27.1",
    "fastapi[standard]>=0.125.0",
    "python-dotenv>=1.2.1",
]
`
	actual := UpdateUvVersion(rawTOML, "patch")
	os.WriteFile("bruh.toml", []byte(actual), 0667)
	assert.Equal(t, expected, actual)
}
