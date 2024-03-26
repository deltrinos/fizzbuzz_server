package domain_test

import (
	"testing"

	"github.com/deltrinos/fizzbuzz_server/domain"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzParamsValidation(t *testing.T) {
	tests := []struct {
		name     string
		params   domain.FizzBuzzParams
		expected bool
	}{
		{
			name:     "ValidParams",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"},
			expected: true,
		},
		{
			name:     "MissingStrs",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15},
			expected: false,
		},
		{
			name:     "MissingInt1",
			params:   domain.FizzBuzzParams{Int2: 5, Limit: 15},
			expected: false,
		},
		{
			name:     "MissingInt2",
			params:   domain.FizzBuzzParams{Int1: 3, Limit: 15},
			expected: false,
		},
		{
			name:     "MissingLimit",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5},
			expected: false,
		},
		{
			name:     "InvalidLimit",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5, Limit: -1},
			expected: false,
		},
		{
			name:     "ZeroParams",
			params:   domain.FizzBuzzParams{},
			expected: false,
		},
		{
			name:     "EqualIntegers",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 3, Limit: 15, Str1: "Fizz", Str2: "Buzz"},
			expected: false,
		},
		{
			name:     "NegativeIntegers",
			params:   domain.FizzBuzzParams{Int1: -3, Int2: 5, Limit: 15},
			expected: false,
		},
		{
			name:     "EmptyStrings",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "", Str2: ""},
			expected: false,
		},
		{
			name:     "EmptyStr1",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str2: "buzz"},
			expected: false,
		},
		{
			name:     "EmptyStr2",
			params:   domain.FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz"},
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
