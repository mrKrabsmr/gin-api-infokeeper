package dto

import (
	"github.com/go-playground/validator/v10"
)

type CreateInfoDTO struct {
	Key      string `json:"key" binding:"required" validate:"required"`
	Value    string `json:"value" binding:"required" validate:"required"`
	ReadOnly bool   `json:"read_only"`
}

func (dto *CreateInfoDTO) Validate() error {
	v := validator.New()
	return v.Struct(dto)
}
