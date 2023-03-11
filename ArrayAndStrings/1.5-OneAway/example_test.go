package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneAway(t *testing.T) {
	testCases := map[string]struct {
		inputOne       string
		inputTwo       string
		expectedResult bool
	}{
		"Return true: remove": {
			inputOne:       "pale",
			inputTwo:       "ple",
			expectedResult: true,
		},
		"Return true: insert": {
			inputOne:       "pale",
			inputTwo:       "pales",
			expectedResult: true,
		},
		"Return true: edits": {
			inputOne:       "pale",
			inputTwo:       "bale",
			expectedResult: true,
		},
		"Return true: insert more": {
			inputOne:       "pales",
			inputTwo:       "pale",
			expectedResult: true,
		},
		"Return true: bake": {
			inputOne:       "pale",
			inputTwo:       "bake",
			expectedResult: false,
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := OneAway(tc.inputOne, tc.inputTwo)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
