package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

// Data 单个记录
type Data struct {
	gorm.Model
	Title     string
	Content   string
	UpdatedAt time.Time
}

func (d Data) TableName() string {
	return "datas"
}

// CreateData 新增一条记录
func CreateData(db *gorm.DB, title string, content string) error {
	data := &Data{
		Title:     title,
		Content:   content,
		UpdatedAt: time.Now(),
	}
	result := db.Create(data)
	if result.Error != nil {
		log.Printf("Failed to create data: %v", result.Error)
		return result.Error
	}
	return nil
}

// GetDataByID 根据 ID 获取一条记录
func GetDataByID(db *gorm.DB, id uint) (*Data, error) {
	var data Data
	result := db.First(&data, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		log.Printf("Failed to get data by ID: %v", result.Error)
		return nil, result.Error
	}
	return &data, nil
}

// UpdateDataByID 根据 ID 更新记录
func UpdateDataByID(db *gorm.DB, id uint, newTitle string, newContent string) error {
	result := db.Model(&Data{}).Where("id = ?", id).Updates(Data{
		Title:     newTitle,
		Content:   newContent,
		UpdatedAt: time.Now(),
	})
	if result.Error != nil {
		log.Printf("Failed to update data by ID: %v", result.Error)
		return result.Error
	}
	return nil
}

// DeleteDataByID 根据 ID 删除记录
func DeleteDataByID(db *gorm.DB, id uint) error {
	result := db.Unscoped().Delete(&Data{}, id)
	if result.Error != nil {
		log.Printf("Failed to delete data by ID: %v", result.Error)
		return result.Error
	}
	return nil
}

// ListAllData 查询所有记录（按更新时间降序）
func ListAllData(db *gorm.DB) ([]Data, error) {
	var datas []Data
	result := db.Order("updated_at DESC").Find(&datas)
	if result.Error != nil {
		log.Printf("Failed to list all data: %v", result.Error)
		return nil, result.Error
	}
	return datas, nil
}
