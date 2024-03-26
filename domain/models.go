package domain

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidParams = errors.New("invalid params")
)

var validate = validator.New()

// FizzBuzzParams represents the parameters for generating Fizz-Buzz.
type FizzBuzzParams struct {
	Int1  int    `json:"int1" validate:"required,min=1"`
	Int2  int    `json:"int2" validate:"required,min=1"`
	Limit int    `json:"limit" validate:"required,min=1,max=9999999"`
	Str1  string `json:"str1"  validate:"required"`
	Str2  string `json:"str2"  validate:"required"`
}

// Validate validates the FizzBuzzParams.
func (f *FizzBuzzParams) Validate() error {
	if err := validate.Struct(f); err != nil {
		return err
	}

	if f.Int1 == f.Int2 {
		return fmt.Errorf("%w: invalid int1 is equal to int2", ErrInvalidParams)
	}

	if f.Str1 == "" || f.Str2 == "" {
		return fmt.Errorf("%w: invalid input strs must be not empty", ErrInvalidParams)
	}

	// no error: params are valids
	return nil
}
