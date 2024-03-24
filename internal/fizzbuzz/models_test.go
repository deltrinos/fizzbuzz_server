package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzParamsValidation(t *testing.T) {
	tests := []struct {
		name     string
		params   FizzBuzzParams
		expected bool
	}{
		{
			name:     "ValidParams",
			params:   FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"},
			expected: true,
		},
		{
			name:     "MissingStrs",
			params:   FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15},
			expected: false,
		},
		{
			name:     "MissingInt1",
			params:   FizzBuzzParams{Int2: 5, Limit: 15},
			expected: false,
		},
		{
			name:     "MissingInt2",
			params:   FizzBuzzParams{Int1: 3, Limit: 15},
			expected: false,
		},
		{
			name:     "MissingLimit",
			params:   FizzBuzzParams{Int1: 3, Int2: 5},
			expected: false,
		},
		{
			name:     "InvalidLimit",
			params:   FizzBuzzParams{Int1: 3, Int2: 5, Limit: -1},
			expected: false,
		},
		{
			name:     "ZeroParams",
			params:   FizzBuzzParams{},
			expected: false,
		},
		{
			name:     "EqualIntegers",
			params:   FizzBuzzParams{Int1: 3, Int2: 3, Limit: 15, Str1: "Fizz", Str2: "Buzz"},
			expected: false,
		},
		{
			name:     "NegativeIntegers",
			params:   FizzBuzzParams{Int1: -3, Int2: 5, Limit: 15},
			expected: false,
		},
		{
			name:     "EmptyStrings",
			params:   FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "", Str2: ""},
			expected: false,
		},
		{
			name:     "EmptyStr1",
			params:   FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str2: "buzz"},
			expected: false,
		},
		{
			name:     "EmptyStr2",
			params:   FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestFizzBuzzGeneration(t *testing.T) {
	tests := []struct {
		name           string
		params         FizzBuzzParams
		expectedError  bool
		expectedOutput []string
	}{
		{
			name: "BasicFizzBuzz",
			params: FizzBuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedError:  false,
			expectedOutput: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			name: "NoMultiples",
			params: FizzBuzzParams{
				Int1:  4,
				Int2:  7,
				Limit: 10,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedError:  false,
			expectedOutput: []string{"1", "2", "3", "fizz", "5", "6", "buzz", "fizz", "9", "10"},
		},
		{
			name: "EmptyStrings",
			params: FizzBuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "",
				Str2:  "",
			},
			expectedError:  true,
			expectedOutput: []string{"1", "2", "", "4", "", "fizz", "", "7", "8", "fizz", "", "buzz", "fizz", "11", "", "fizzbuzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewFizzBuzzService()
			result, err := service.GenerateFizzBuzz(tt.params)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectedOutput, result)
			}
		})
	}
}
