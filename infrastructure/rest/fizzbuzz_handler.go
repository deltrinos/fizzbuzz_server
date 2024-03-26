package rest

import (
	"encoding/json"
	"net/http"

	"github.com/deltrinos/fizzbuzz_server/domain"
	"github.com/deltrinos/fizzbuzz_server/service"
)

// FizzBuzzHandler handles requests for Fizz-Buzz generation.
type FizzBuzzHandler struct {
	fizzBuzzService service.FizzBuzzService
}

// NewFizzBuzzHandler creates a new instance of FizzBuzzHandler.
func NewFizzBuzzHandler(fizzBuzzService service.FizzBuzzService) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		fizzBuzzService: fizzBuzzService,
	}
}

// HandleFizzBuzz handles requests to generate Fizz-Buzz output.
func (h *FizzBuzzHandler) HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	var params domain.FizzBuzzParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.fizzBuzzService.GenerateFizzBuzz(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(result)

	// update statistics
	requestsCounterMu.Lock()
	requestsCounter[params]++
	requestsCounterMu.Unlock()
}
