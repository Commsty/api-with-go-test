package repository

import (
	"calc/internal/entity"

	"gorm.io/gorm"
)

// main CRUD methods

type CalculationRepository interface {
	CreateCalculation(calc entity.Calculation) error
	GetAllCalculations() ([]entity.Calculation, error)
	GetCalculationByID(id string) (entity.Calculation, error)
	UpdateCalculation(calc entity.Calculation) error
	DeleteCalculation(id string) error
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalcRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}
}

func (r *calcRepository) CreateCalculation(calc entity.Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *calcRepository) GetAllCalculations() ([]entity.Calculation, error) {
	var calculations []entity.Calculation
	err := r.db.Find(&calculations).Error
	return calculations, err
}

func (r *calcRepository) GetCalculationByID(id string) (entity.Calculation, error) {
	var calc entity.Calculation
	err := r.db.First(&calc, "id = ?", id).Error
	return calc, err
}

func (r *calcRepository) UpdateCalculation(calc entity.Calculation) error {
	return r.db.Save(&calc).Error
}

func (r *calcRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&entity.Calculation{}, "id = ?", id).Error
}
