package repository

import (
	"fmt"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (r *ClassRepository) Create(class *model.Class) error {

	err := r.db.Create(class).Error
	if err != nil {
		return fmt.Errorf(constant.ErrCreateClass)
	}

	return nil

}

func (r *ClassRepository) GetAll() ([]model.Class, error) {

	var classes []model.Class

	err := r.db.Preload("Category").
		Find(&classes).Error
	if err != nil {
		return []model.Class{}, fmt.Errorf(constant.ErrFetchClass)
	}

	return classes, nil

}

func (r *ClassRepository) GetDetail(classId int) (model.Class, error) {

	var class model.Class

	err := r.db.Preload("Category").
		First(&class, classId).Error
	if err != nil {
		return model.Class{}, fmt.Errorf(constant.ErrFetchClass)
	}

	return class, nil

}

func (r *ClassRepository) Update(classId int, class *model.Class) error {

	err := r.db.Model(&model.Class{}).
		Where(`id = ?`, classId).
		Updates(class).
		Error
	if err != nil {
		return fmt.Errorf(constant.ErrEditClass)
	}

	return nil

}

func (r *ClassRepository) Delete(classId int) error {

	err := r.db.Unscoped().Delete(&model.Class{}, classId).Error
	if err != nil {
		return fmt.Errorf(constant.ErrDeleteClass)
	}

	return nil

}
