package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomInteger(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description           string
		queryStringParameters map[string]string
		expectedIntMin        int
		expectedIntMax        int
		expectedError         bool
	}{
		{
			"no text parameter",
			map[string]string{},
			0,
			100,
			false,
		},
		{
			"empty text parameter",
			map[string]string{
				"text": "",
			},
			0,
			100,
			false,
		},
		{
			"'die'",
			map[string]string{
				"text": "die",
			},
			1,
			6,
			false,
		},
		{
			"'dice'",
			map[string]string{
				"text": "dice",
			},
			1,
			6,
			false,
		},
		{
			"10",
			map[string]string{
				"text": "10",
			},
			0,
			10,
			false,
		},
		{
			"not a number",
			map[string]string{
				"text": "notanumber",
			},
			0,
			0,
			true,
		},
		{
			"range (5, 10)",
			map[string]string{
				"text": "5 10",
			},
			5,
			10,
			false,
		},
		{
			"range (1000, 200000)",
			map[string]string{
				"text": "1000 200000",
			},
			1000,
			200000,
			false,
		},
		{
			"range (notanumber, 200000)",
			map[string]string{
				"text": "notanumber 200000",
			},
			0,
			0,
			true,
		},
		{
			"range (1000, notanumber)",
			map[string]string{
				"text": "1000 notanumber",
			},
			0,
			0,
			true,
		},
		{
			"more than two parameters",
			map[string]string{
				"text": "5 10 15",
			},
			0,
			0,
			true,
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		for j := 0; j < 100; j++ {
			t.Run(testCase.description, func(t *testing.T) {
				t.Parallel()

				r, err := generateRandomInteger(testCase.queryStringParameters)
				if testCase.expectedError {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				assert.True(t, r >= testCase.expectedIntMin)
				assert.True(t, r <= testCase.expectedIntMax)
			})
		}
	}
}
