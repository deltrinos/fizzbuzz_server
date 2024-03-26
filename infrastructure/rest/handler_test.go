package rest

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/deltrinos/fizzbuzz_server/domain"
	"github.com/deltrinos/fizzbuzz_server/service"
)

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
			handler := NewFizzBuzzHandler(service.NewFizzBuzzService())
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
		requests     map[domain.FizzBuzzParams]int
		expectedCode int
	}{
		{
			name: "Valid Request",
			requests: map[domain.FizzBuzzParams]int{
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
