package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DeleteInfoDTO struct {
	ID  uuid.UUID `json:"id" binding:"required" validate:"required"`
	Key string    `json:"key" binding:"required" validate:"required"`
}

func (dto *DeleteInfoDTO) Validate() error {
	v := validator.New()
	return v.Struct(dto)
}
