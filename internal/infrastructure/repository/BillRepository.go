package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// BillRepository handles database operations for bills.
type BillRepository interface {
	CreateBill(bill *entities.Bill) error
	ListAllBills() ([]entities.Bill, error)
	FindBillByID(id uint) (*entities.Bill, error)
	UpdateBillByID(id uint, updateData map[string]interface{}) error
	DeleteBillByID(id uint) error
}

type billRepository struct {
	db *gorm.DB
}

// NewBillRepository creates a new instance of BillRepository.
func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{db: db}
}

func (r *billRepository) CreateBill(bill *entities.Bill) error {
	return r.db.Create(bill).Error
}

func (r *billRepository) ListAllBills() ([]entities.Bill, error) {
	var bills []entities.Bill
	err := r.db.Find(&bills).Error
	return bills, err
}

func (r *billRepository) FindBillByID(id uint) (*entities.Bill, error) {
	var bill entities.Bill
	err := r.db.First(&bill, id).Error
	if err != nil {
		return nil, err
	}
	return &bill, nil
}

func (r *billRepository) UpdateBillByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Bill{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *billRepository) DeleteBillByID(id uint) error {
	return r.db.Delete(&entities.Bill{}, id).Error
}
