package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLify(t *testing.T) {
	testCases := map[string]struct {
		input          string
		inputLenght    int
		expectedResult string
	}{
		"Mr John Smith": {
			input:          "Mr John Smith",
			inputLenght:    13,
			expectedResult: "Mr%20John%20Smith",
		},
		"Mr John Smith   ": {
			input:          "Mr John Smith   ",
			inputLenght:    13,
			expectedResult: "Mr%20John%20Smith",
		},
	}

	for name, v := range testCases {
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := URLify(tc.input, tc.inputLenght)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
