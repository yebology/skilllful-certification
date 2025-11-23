package service

import (
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
)

// type ClassServiceInterface interface {
// 	CreateClassService(dto request.ClassDto) error
// 	EditClassService(id int, dto request.ClassDto) error
// 	GetAllClassService() (interface{}, error)
// 	GetClassDetailService(id int) (interface{}, error)
// 	DeleteClassService(id int) error
// }

type ClassServiceInterface interface {
	CreateClassService(dto request.ClassDto) error
	EditClassService(id int, dto request.ClassDto) error
	GetAllClassService() ([]response.ClassDto, error)
	GetClassDetailService(id int) (response.ClassDetailDto, error)
	DeleteClassService(id int) error
}
