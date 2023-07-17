package dto

type CreateInfoDTO struct {
	Key      string `json:"key" binding:"required" validate:"required"`
	Value    string `json:"value" binding:"required" validate:"required"`
	ReadOnly bool   `json:"read_only"`
}

func (dto *CreateInfoDTO) Validate() error {
	v := GetValidator()
	return v.Struct(dto)
}
