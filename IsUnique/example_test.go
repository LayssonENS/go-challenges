package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique(t *testing.T) {
	testCases := map[string]struct {
		input          string
		expectedResult bool
	}{
		"empty string returns false": {
			input:          "",
			expectedResult: false,
		},
		"aabb returns false": {
			input:          "aabb",
			expectedResult: false,
		},
		"aabbc returns true": {
			input:          "aabbc",
			expectedResult: true,
		},
		"adasdasdasdasda4asdasd returns true": {
			input:          "adasdasdasdasda4asdasd",
			expectedResult: true,
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := IsUnique(tc.input)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
