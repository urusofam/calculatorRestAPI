package services

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/google/uuid"
	"github.com/urusofam/calculatorRestAPI/internal/server/models"
	"github.com/urusofam/calculatorRestAPI/internal/server/repositories"
)

type CalculationService interface {
	calculateExpression(expression string) (string, error)
	GetCalculations() ([]models.Calculation, error)
	AddCalculation(expression string) (models.Calculation, error)
	UpdateCalculation(expression, id string) (models.Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo repositories.CalculationRepository
}

func NewCalculationService(repo repositories.CalculationRepository) CalculationService {
	return &calcService{repo: repo}
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	program, err := expr.Compile(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Run(program, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}

func (s *calcService) GetCalculations() ([]models.Calculation, error) {
	return s.repo.GetCalculations()
}

func (s *calcService) AddCalculation(expression string) (models.Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return models.Calculation{}, err
	}

	calc := models.Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.AddCalculation(calc); err != nil {
		return models.Calculation{}, err
	}
	return calc, nil
}

func (s *calcService) UpdateCalculation(expression, id string) (models.Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return models.Calculation{}, err
	}

	calc := models.Calculation{
		ID:         id,
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return models.Calculation{}, err
	}
	return models.Calculation{}, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	if err := s.repo.DeleteCalculation(id); err != nil {
		return err
	}
	return nil
}
