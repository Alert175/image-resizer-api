package models

import "gorm.io/gorm"

// описание модели продукта
type ImageEntity struct {
	ID                uint   `gorm:"primarykey" json:"id"`
	ParentUrl         string `json:"parentUrl"`
	CompressImageName string `json:"compressImageName"`
	Size              int    `json:"size"`
	gorm.Model
}

// TableName переопределяет название таблицы для ImageEntity на `images`
func (ve ImageEntity) TableName() string {
	return "images"
}
