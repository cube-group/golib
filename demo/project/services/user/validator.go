package user

type ValidatorList struct {
	Page     uint `form:"page" binding:"omitempty"`
	PageSize uint `form:"page_size" binding:"omitempty"`
}

type ValidatorAdd struct {
	Name string `form:"name" binding:"required"`
}
