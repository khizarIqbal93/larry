package gitutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	actual := ParseTag("v1.0.0")
	assert.Equal(t, 1, actual.Major)
	assert.Equal(t, 0, actual.Minor)
	assert.Equal(t, 0, actual.Patch)
}

func TestTagSort(t *testing.T) {
	tests := []struct {
		Name     string
		Input    []string
		Expected []string
	}{
		{
			Name:     "Already in order",
			Input:    []string{"v0.0.1", "v0.0.2", "vsomething"},
			Expected: []string{"vsomething", "v0.0.1", "v0.0.2"},
		},
		{
			Name:     "Simple unordered tag list",
			Input:    []string{"v0.0.2", "vsomething", "v0.0.1"},
			Expected: []string{"vsomething", "v0.0.1", "v0.0.2"},
		},
		{
			Name:     "Patch tag > 9",
			Input:    []string{"v0.0.11", "v0.0.1", "v0.0.2"},
			Expected: []string{"v0.0.1", "v0.0.2", "v0.0.11"},
		},
		{
			Name:     "Minor tag > 9",
			Input:    []string{"v0.10.1", "v0.1.1"},
			Expected: []string{"v0.1.1", "v0.10.1"},
		},
		{
			Name:     "Major tag > 9",
			Input:    []string{"v10.1.1", "v9.1.1"},
			Expected: []string{"v9.1.1", "v10.1.1"},
		},
		{
			Name:     "Major & Minor & Patch tag > 9",
			Input:    []string{"v9.1.1", "v9.11.2", "v10.1.1"},
			Expected: []string{"v9.1.1", "v9.11.2", "v10.1.1"},
		},
		{
			Name:     "Final Boss Scenario",
			Input:    []string{"v1.1.1", "v0.3.11", "v0.3.10", "v0.3.9", "v0.2.9", "v0.3.8", "v0.3.7"},
			Expected: []string{"v0.2.9", "v0.3.7", "v0.3.8", "v0.3.9", "v0.3.10", "v0.3.11", "v1.1.1"},
		},
	}

	for _, tc := range tests {
		actual := tagSort(tc.Input)
		assert.Equal(t, tc.Expected, actual, "test name: %v", tc.Name)
	}
}
