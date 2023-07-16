package dto

import "github.com/google/uuid"

type UpdateInfoDTO struct {
	ID uuid.UUID `json:"id" binding:"required" validate:"required"`
	*CreateInfoDTO
}
