package rest

import (
	"encoding/json"
	"net/http"

	"github.com/deltrinos/fizzbuzz_server/repository"
)

// StatisticsHandler handles requests for statistics.
type StatisticsHandler struct {
	stats *repository.StatisticsRepository
}

// NewStatisticsHandler creates a new instance of StatisticsHandler.
func NewStatisticsHandler(stats *repository.StatisticsRepository) *StatisticsHandler {
	return &StatisticsHandler{
		stats: stats,
	}
}

// HandleStatistics handles requests to get statistics.
func (h *StatisticsHandler) HandleStatistics(w http.ResponseWriter, r *http.Request) {
	mostCommonRequest, maxHits := h.stats.GetMostFrequentRequest()

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
