package dto

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	V    *validator.Validate
	once = &sync.Once{}
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		V = validator.New()
	})

	return V
}
