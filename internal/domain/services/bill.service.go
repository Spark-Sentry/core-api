package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// BillService defines operations related to bills.
type BillService interface {
	CreateBill(req dto.CreateBillRequest) (*entities.Bill, error)
	ListBills() ([]entities.Bill, error)
	GetBillByID(id uint) (*entities.Bill, error)
	UpdateBillByID(id uint, req dto.UpdateBillRequest) error
	DeleteBillByID(id uint) error
}

type billService struct {
	repo repository.BillRepository
}

// NewBillService creates a new instance of BillService.
func NewBillService(repo repository.BillRepository) BillService {
	return &billService{
		repo: repo,
	}
}

func (s *billService) CreateBill(req dto.CreateBillRequest) (*entities.Bill, error) {
	bill := &entities.Bill{
		MeterID:  req.MeterID,
		Start:    req.Start,
		Stop:     req.Stop,
		Quantity: req.Quantity,
		Cost:     req.Cost,
	}
	if err := s.repo.CreateBill(bill); err != nil {
		return nil, fmt.Errorf("failed to create bill: %w", err)
	}
	return bill, nil
}

func (s *billService) ListBills() ([]entities.Bill, error) {
	return s.repo.ListAllBills()
}

func (s *billService) GetBillByID(id uint) (*entities.Bill, error) {
	return s.repo.FindBillByID(id)
}

func (s *billService) UpdateBillByID(id uint, req dto.UpdateBillRequest) error {
	updateData := map[string]interface{}{
		"start":    req.Start,
		"stop":     req.Stop,
		"quantity": req.Quantity,
		"cost":     req.Cost,
	}
	if err := s.repo.UpdateBillByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update bill: %w", err)
	}
	return nil
}

func (s *billService) DeleteBillByID(id uint) error {
	return s.repo.DeleteBillByID(id)
}
