package dto

import (
	"github.com/google/uuid"
)

type DeleteInfoDTO struct {
	ID  uuid.UUID `json:"id" binding:"required" validate:"required"`
	Key string    `json:"key" binding:"required" validate:"required"`
}

func (dto *DeleteInfoDTO) Validate() error {
	v := GetValidator()
	return v.Struct(dto)
}
