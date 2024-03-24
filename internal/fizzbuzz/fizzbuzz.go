package fizzbuzz

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidParams = errors.New("invalid params")
)

// FizzBuzzService defines the Fizz-Buzz service interface.
type FizzBuzzService interface {
	GenerateFizzBuzz(p FizzBuzzParams) ([]string, error)
}

// fizzBuzzServiceImpl implements the Fizz-Buzz service.
type fizzBuzzServiceImpl struct{}

// NewFizzBuzzService creates a new instance of the Fizz-Buzz service.
func NewFizzBuzzService() FizzBuzzService {
	return &fizzBuzzServiceImpl{}
}

// GenerateFizzBuzz generates Fizz-Buzz output based on the given parameters.
func (s *fizzBuzzServiceImpl) GenerateFizzBuzz(p FizzBuzzParams) ([]string, error) {
	var result []string

	if err := p.Validate(); err != nil {
		return result, fmt.Errorf("%w: %v", ErrInvalidParams, err)
	}

	for i := 1; i <= p.Limit; i++ {
		var str string

		if i%p.Int1 == 0 {
			str += p.Str1
		}
		if i%p.Int2 == 0 {
			str += p.Str2
		}

		if str == "" {
			str = fmt.Sprintf("%d", i)
		}

		result = append(result, str)
	}
	return result, nil
}
