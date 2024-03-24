package fizzbuzz

import (
	"encoding/json"
	"errors"
	"net/http"
)

// FizzBuzzHandler handles requests for Fizz-Buzz generation.
type FizzBuzzHandler struct {
	fizzBuzzService FizzBuzzService
}

// NewFizzBuzzHandler creates a new instance of FizzBuzzHandler.
func NewFizzBuzzHandler() *FizzBuzzHandler {
	return &FizzBuzzHandler{
		fizzBuzzService: NewFizzBuzzService(),
	}
}

// HandleFizzBuzz handles requests to generate Fizz-Buzz output.
func (h *FizzBuzzHandler) HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	var params FizzBuzzParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.fizzBuzzService.GenerateFizzBuzz(params)
	if errors.Is(err, ErrInvalidParams) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(result)

	// update statistics
	requestsCounterMu.Lock()
	requestsCounter[params]++
	requestsCounterMu.Unlock()
}
