package calculatorService

import "gorm.io/gorm"

type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type CalcRepository struct {
	db *gorm.DB
}

func NewCalcRepository(db *gorm.DB) *CalcRepository {
	return &CalcRepository{db: db}
}
func (r *CalcRepository) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *CalcRepository) GetAllCalculations() ([]Calculation, error) {
	var calcs []Calculation
	err := r.db.Find(&calcs).Error
	return calcs, err
}
func (r *CalcRepository) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation
	err := r.db.First(&calc, id).Error
	return calc, err
}
func (r *CalcRepository) UpdateCalculation(calc Calculation) error {
	return r.db.Save(&calc).Error
}
func (r *CalcRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, id).Error
}
