package service

import (
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/app/repository"
)

type ClassService struct {
	repo *repository.ClassRepository
}

func NewClassService(repository *repository.ClassRepository) *ClassService {
	return &ClassService{repo: repository}
}

func (s *ClassService) CreateClassService(dto request.ClassDto) error {

	class := model.Class{
		Name:        dto.Name,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
		Instructor:  dto.Instructor,
	}

	return s.repo.Create(&class)

}

func (s *ClassService) EditClassService(classId int, dto request.ClassDto) error {

	class := model.Class{
		Name:        dto.Name,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
		Instructor:  dto.Instructor,
	}

	return s.repo.Update(classId, &class)

}

func (s *ClassService) GetAllClassService() ([]response.ClassDto, error) {

	classes, err := s.repo.GetAll()
	if err != nil {
		return []response.ClassDto{}, err
	}

	var classesDto []response.ClassDto

	for _, c := range classes {

		class := response.ClassDto{
			Id:       c.ID,
			Name:     c.Name,
			Category: c.Category.Name,
		}

		classesDto = append(classesDto, class)

	}

	return classesDto, nil

}

func (s *ClassService) GetClassDetailService(classId int) (response.ClassDetailDto, error) {

	class, err := s.repo.GetDetail(classId)
	if err != nil {
		return response.ClassDetailDto{}, err
	}

	classDto := response.ClassDetailDto{
		Id:          class.ID,
		Name:        class.Name,
		Category:    class.Category.Name,
		Instructor:  class.Instructor,
		Description: class.Description,
	}

	return classDto, nil

}

func (s *ClassService) DeleteClassService(classId int) error {

	return s.repo.Delete(classId)

}
