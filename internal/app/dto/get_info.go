package dto

import (
	"github.com/go-playground/validator/v10"
)

type GetInfoDTO struct {
	ID  string `form:"id" binding:"required" validate:"required"`
	Key string `form:"key" binding:"required" validate:"required"`
}

func (dto *GetInfoDTO) Validate() error {
	v := validator.New()
	return v.Struct(dto)
}
