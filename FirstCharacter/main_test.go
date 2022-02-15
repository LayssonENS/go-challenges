package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepeatingCharacter(t *testing.T) {
	testCases := map[string]struct {
		input          string
		expectedResult int
	}{
		"empty string returns -1": {
			input:          "",
			expectedResult: -1,
		},
		"aabb returns -1": {
			input:          "aabb",
			expectedResult: -1,
		},
		"a returns 0": {
			input:          "a",
			expectedResult: 0,
		},
		"aba returns 1": {
			input:          "aba",
			expectedResult: 1,
		},
		"ababcded returns 4": {
			input:          "ababcded",
			expectedResult: 4,
		},
		"aaab returns 3": {
			input:          "aaab",
			expectedResult: 3,
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := FirstNonRepeatingCharacter(tc.input)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

func TestFirstNonRepeatingCharacterOld(t *testing.T) {
	testCases := map[string]struct {
		input          string
		expectedResult int
	}{
		"empty string returns -1": {
			input:          "",
			expectedResult: -1,
		},
		"aabb returns -1": {
			input:          "aabb",
			expectedResult: -1,
		},
		"a returns 0": {
			input:          "a",
			expectedResult: 0,
		},
		"aba returns 1": {
			input:          "aba",
			expectedResult: 1,
		},
		"ababcded returns 4": {
			input:          "ababcded",
			expectedResult: 4,
		},
		"aaab returns 3": {
			input:          "aaab",
			expectedResult: 3,
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := FirstNonRepeatingCharacterOld(tc.input)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
