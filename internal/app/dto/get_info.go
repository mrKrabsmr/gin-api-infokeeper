package dto

type GetInfoDTO struct {
	ID  string `form:"id" binding:"required" validate:"required"`
	Key string `form:"key" binding:"required" validate:"required"`
}

func (dto *GetInfoDTO) Validate() error {
	v := GetValidator()
	return v.Struct(dto)
}
