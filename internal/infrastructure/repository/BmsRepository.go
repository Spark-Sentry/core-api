package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// BmsRepository handles database operations for Bms.
type BmsRepository interface {
	CreateBms(bms *entities.Bms) error
	ListBms() ([]entities.Bms, error)
	FindBmsByID(id uint) (*entities.Bms, error)
	UpdateBmsByID(id uint, updateData map[string]interface{}) error
	DeleteBmsByID(id uint) error
}

type bmsRepository struct {
	db *gorm.DB
}

// NewBmsRepository creates a new instance of BmsRepository.
func NewBmsRepository(db *gorm.DB) BmsRepository {
	return &bmsRepository{db: db}
}

func (r *bmsRepository) CreateBms(bms *entities.Bms) error {
	return r.db.Create(bms).Error
}

func (r *bmsRepository) ListBms() ([]entities.Bms, error) {
	var bmsList []entities.Bms
	err := r.db.Find(&bmsList).Error
	return bmsList, err
}

func (r *bmsRepository) FindBmsByID(id uint) (*entities.Bms, error) {
	var bms entities.Bms
	err := r.db.First(&bms, id).Error
	if err != nil {
		return nil, err
	}
	return &bms, nil
}

func (r *bmsRepository) UpdateBmsByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Bms{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *bmsRepository) DeleteBmsByID(id uint) error {
	return r.db.Delete(&entities.Bms{}, id).Error
}
