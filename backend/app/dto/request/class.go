package request

type ClassDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Instructor  string `json:"instructor" validate:"required"`
	CategoryId  uint   `json:"categoryId" validate:"required"`
}
