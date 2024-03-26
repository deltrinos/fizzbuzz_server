package repository

import (
	"sync"

	"github.com/deltrinos/fizzbuzz_server/domain"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Prometheus metrics
	fizzBuzzCalls  prometheus.Counter
	mostUsedParams *prometheus.GaugeVec
)

type StatisticsRepository struct {
	requests map[domain.FizzBuzzParams]int
	mu       sync.Mutex

	fizzBuzzCalls  prometheus.Counter
	mostUsedParams *prometheus.GaugeVec
}

func init() {
	fizzBuzzCalls = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fizzbuzz_calls_total",
		Help: "Total number of FizzBuzz calls.",
	})
	mostUsedParams = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "most_used_params_map",
		Help: "Most used FizzBuzz parameters with counts.",
	}, []string{"params"})
}

func NewStatisticsRepository() *StatisticsRepository {
	return &StatisticsRepository{
		requests:       make(map[domain.FizzBuzzParams]int),
		fizzBuzzCalls:  fizzBuzzCalls,
		mostUsedParams: mostUsedParams,
	}
}

func NewStatisticsRepositoryWithRequests(requests map[domain.FizzBuzzParams]int) *StatisticsRepository {
	return &StatisticsRepository{
		requests:       requests,
		fizzBuzzCalls:  fizzBuzzCalls,
		mostUsedParams: mostUsedParams,
	}
}

func (repo *StatisticsRepository) LogRequest(params domain.FizzBuzzParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.requests[params]++

	// Increment Prometheus counters
	repo.fizzBuzzCalls.Inc()
	repo.mostUsedParams.WithLabelValues(params.String()).Inc()
}

func (repo *StatisticsRepository) GetMostFrequentRequest() (params domain.FizzBuzzParams, hits int) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	maxHits := 0
	for p, h := range repo.requests {
		if h > maxHits {
			params = p
			hits = h
			maxHits = h
		}
	}

	return
}
