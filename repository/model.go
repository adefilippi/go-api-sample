package repository

import (
	"github.com/google/uuid"

	"example/web-service-gin/entity"
)

func GetAllModels() ([]entity.Model, error) {
	var users []entity.Model
	err := db.Model(&entity.Model{}).Preload("Images").Find(&users).Error // Use the correct field name for association
	return users, err
}

func GetModelById(id uuid.UUID) (entity.Model, error) {
	var model entity.Model
	if result := db.First(&model, id); result.Error != nil {
		return entity.Model{}, result.Error
	}
	return model, nil
}

func CreateModel(model entity.Model) (entity.Model, error) {
	u := uuid.New()
	model.ID = &u
	transaction := db.Begin()
	if err := db.Create(&model).Error; err != nil {
		transaction.Rollback()
		return entity.Model{}, err
	}

	transaction.Commit()
	return model, nil
}

func UpdateModel(model entity.Model) (entity.Model, error) {
	transaction := db.Begin()
	if err := db.Save(&model).Error; err != nil {
		transaction.Rollback()
		return entity.Model{}, err
	}
	transaction.Commit()
	return model, nil
}

func DeleteModel(id uuid.UUID) error {
	var model entity.Model
	transaction := db.Begin()
	if err := db.Where("id = ?", id).Delete(&model).Error; err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
