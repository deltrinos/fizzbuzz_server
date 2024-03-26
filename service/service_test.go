package service

import (
	"testing"

	"github.com/deltrinos/fizzbuzz_server/domain"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzGeneration(t *testing.T) {
	tests := []struct {
		name           string
		params         domain.FizzBuzzParams
		expectedError  bool
		expectedOutput []string
	}{
		{
			name: "BasicFizzBuzz",
			params: domain.FizzBuzzParams{
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
			params: domain.FizzBuzzParams{
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
			params: domain.FizzBuzzParams{
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
			fbService := NewFizzBuzzService()
			result, err := fbService.GenerateFizzBuzz(tt.params)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectedOutput, result)
			}
		})
	}
}
