package service

import (
	"strconv"

	"github.com/deltrinos/fizzbuzz_server/domain"
)

// FizzBuzzService defines the Fizz-Buzz service interface.
type FizzBuzzService interface {
	GenerateFizzBuzz(p domain.FizzBuzzParams) ([]string, error)
}

// fizzBuzzServiceImpl implements the Fizz-Buzz service.
type fizzBuzzServiceImpl struct{}

// NewFizzBuzzService creates a new instance of the Fizz-Buzz service.
func NewFizzBuzzService() FizzBuzzService {
	return &fizzBuzzServiceImpl{}
}

// GenerateFizzBuzz generates Fizz-Buzz output based on the given parameters.
func (s *fizzBuzzServiceImpl) GenerateFizzBuzz(p domain.FizzBuzzParams) ([]string, error) {
	var result []string

	if err := p.Validate(); err != nil {
		return result, err
	}

	for i := 1; i <= p.Limit; i++ {
		isInt1Multiple := i%p.Int1 == 0
		isInt2Multiple := i%p.Int2 == 0

		if isInt1Multiple && isInt2Multiple {
			result = append(result, p.Str1+p.Str2)
		} else if isInt1Multiple {
			result = append(result, p.Str1)
		} else if isInt2Multiple {
			result = append(result, p.Str2)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil
}
