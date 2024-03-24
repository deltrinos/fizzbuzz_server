package fizzbuzz

import (
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFizzBuzz(t *testing.T) {
	tests := []struct {
		name      string
		params    FizzBuzzParams
		wantError bool
		want      []string
	}{
		{
			name:      "Basic Test FizzBuzz with 3, 5, 15",
			params:    FizzBuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"},
			wantError: false,
			want:      []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name:      "Basic Test up to 20",
			params:    FizzBuzzParams{Int1: 23, Int2: 25, Limit: 20, Str1: "Fizz", Str2: "Buzz"},
			wantError: false,
			want:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &fizzBuzzServiceImpl{}
			got, gotErr := s.GenerateFizzBuzz(tt.params)

			if tt.wantError {
				assert.Error(t, gotErr)
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateFizzBuzz() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestHandleFizzBuzz(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{
			name:         "Valid Request",
			requestBody:  `{"int1": 3, "int2": 5, "limit": 15, "str1": "Fizz", "str2": "Buzz"}`,
			expectedCode: 200,
		},
		{
			name:         "Invalid Request same ints",
			requestBody:  `{"int1": 3, "int2": 3, "limit": 15, "str1": "Fizz", "str2": "Buzz"}`,
			expectedCode: 400,
		},
		{
			name:         "Invalid Request empty strs",
			requestBody:  `{"int1": 3, "int2": 3, "limit": 15}`,
			expectedCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewFizzBuzzHandler()
			req := httptest.NewRequest("POST", "/fizzbuzz", strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()
			handler.HandleFizzBuzz(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedCode {
				t.Errorf("Expected status code %d but got %d", tt.expectedCode, resp.StatusCode)
			}
		})
	}
}

func TestHandleStatistics(t *testing.T) {
	tests := []struct {
		name         string
		requests     map[FizzBuzzParams]int
		expectedCode int
	}{
		{
			name: "Valid Request",
			requests: map[FizzBuzzParams]int{
				{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}: 10,
			},
			expectedCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewStatisticsHandler()
			requestsCounter = tt.requests
			req := httptest.NewRequest("GET", "/stats", nil)
			w := httptest.NewRecorder()
			handler.HandleStatistics(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedCode {
				t.Errorf("Expected status code %d but got %d", tt.expectedCode, resp.StatusCode)
			}
		})
	}
}
