package rest

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/deltrinos/fizzbuzz_server/domain"
)

var (
	requestsCounterMu sync.Mutex
	requestsCounter   = make(map[domain.FizzBuzzParams]int)
)

// StatisticsHandler handles requests for statistics.
type StatisticsHandler struct{}

// NewStatisticsHandler creates a new instance of StatisticsHandler.
func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{}
}

// HandleStatistics handles requests to get statistics.
func (h *StatisticsHandler) HandleStatistics(w http.ResponseWriter, r *http.Request) {
	requestsCounterMu.Lock()
	defer requestsCounterMu.Unlock()

	var (
		mostCommonRequest domain.FizzBuzzParams
		maxHits           int
	)
	for req, hits := range requestsCounter {
		if hits > maxHits {
			maxHits = hits
			mostCommonRequest = req
		}
	}

	if maxHits == 0 {
		http.Error(w, "empty", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"most_common_request": mostCommonRequest,
		"hits":                maxHits,
	})
}
