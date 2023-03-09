package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	testCases := map[string]struct {
		input          string
		expectedResult bool
	}{
		"Return true for Tact Coa": {
			input:          "Tact Coa",
			expectedResult: true,
		},
		"Return true for racecar": {
			input:          "racecar",
			expectedResult: true,
		},
		"Return true for A man a plan a canal Panama": {
			input:          "A man a plan a canal Panama",
			expectedResult: true,
		},
		"Return true for Never odd or even": {
			input:          "Never odd or even",
			expectedResult: true,
		},
		"Return false for Hello world": {
			input:          "Hello world",
			expectedResult: false,
		},
		"Return false for Tact Coal": {
			input:          "Tact Coal",
			expectedResult: false,
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := IsPermutationOfPalindrome(tc.input)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
