package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPermutation(t *testing.T) {
	testCases := map[string]struct {
		inputOne       string
		inputTwo       string
		expectedResult bool
	}{
		"empty string returns true": {
			inputOne:       "",
			inputTwo:       "",
			expectedResult: true,
		},
		"aabb, aabc returns false": {
			inputOne:       "aabb",
			inputTwo:       "aabc",
			expectedResult: false,
		},
		"aabbc, aabc returns false": {
			inputOne:       "aabbc",
			inputTwo:       "aabc",
			expectedResult: false,
		},
		"aabb, aabb returns true": {
			inputOne:       "aabb",
			inputTwo:       "aabb",
			expectedResult: true,
		},
		"aabb, abab returns true": {
			inputOne:       "aabb",
			inputTwo:       "abab",
			expectedResult: true,
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := CheckPermutation(tc.inputOne, tc.inputTwo)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
