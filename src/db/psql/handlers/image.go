package handlers

import (
	"image-resizer-api/src/db/psql"
	"image-resizer-api/src/db/psql/models"
)

// структура таблицы в СУБД
type ImageStore struct {
	models.ImageEntity
}

// Описание методов
type ImageStoreInterface interface {
	Create() (ImageStore, error)
	Update() (ImageStore, error)
}

// создать запись
func (vs *ImageStore) Create(content ImageStore) (ImageStore, error) {
	result := psql.DB.Create(&content)
	if result.Error != nil {
		return content, result.Error
	}
	return content, nil
}

// обновить запись
func (vs *ImageStore) Update() error {
	return psql.DB.Model(&vs).Updates(&vs).Error
}

func (vs *ImageStore) SelectFirst() error {
	if errSelect := psql.DB.First(&vs).Error; errSelect != nil {
		return errSelect
	}
	return nil
}
