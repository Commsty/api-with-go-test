package service

import (
	"calc/internal/entity"
	"calc/internal/repository"
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (entity.Calculation, error)
	GetAllCalculations() ([]entity.Calculation, error)
	GetCalculationByID(id string) (entity.Calculation, error)
	UpdateCalculation(id, expression string) (entity.Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo repository.CalculationRepository
}

func NewCalcService(r repository.CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) CreateCalculation(expression string) (entity.Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return entity.Calculation{}, err
	}

	calc := entity.Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return entity.Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) GetAllCalculations() ([]entity.Calculation, error) {
	return s.repo.GetAllCalculations()
}

func (s *calcService) GetCalculationByID(id string) (entity.Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calcService) UpdateCalculation(id, expression string) (entity.Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return entity.Calculation{}, err
	}

	calc.Expression = expression
	calc.Result, err = s.calculateExpression(expression)
	if err != nil {
		return entity.Calculation{}, err
	}

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return entity.Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}
