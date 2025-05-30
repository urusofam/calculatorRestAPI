package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urusofam/calculatorRestAPI/internal/server/models"
)

type CalculationRepository interface {
	GetCalculations() ([]models.Calculation, error)
	AddCalculation(calc models.Calculation) error
	UpdateCalculation(calc models.Calculation) error
	DeleteCalculation(id string) error
}

type calcRepository struct {
	conn *pgxpool.Pool
}

func NewCalculationRepository(conn *pgxpool.Pool) CalculationRepository {
	return &calcRepository{conn: conn}
}

func (r *calcRepository) GetCalculations() ([]models.Calculation, error) {
	query := "select * from calculations"

	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []models.Calculation{}
	for rows.Next() {
		calc := models.Calculation{}

		err := rows.Scan(&calc.ID, &calc.Expression, &calc.Result)
		if err != nil {
			return nil, err
		}
		result = append(result, calc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (repos *calcRepository) AddCalculation(calc models.Calculation) error {
	query := "insert into calculations (id, expression, result) values ($1, $2, $3)"

	_, err := repos.conn.Exec(context.Background(), query, calc.ID, calc.Expression, calc.Result)
	if err != nil {
		return err
	}
	return nil
}

func (r *calcRepository) UpdateCalculation(calc models.Calculation) error {
	query := "update calculations set expression = $1, result = $2 where id = $3"

	_, err := r.conn.Exec(context.Background(), query, calc.Expression, calc.Result, calc.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *calcRepository) DeleteCalculation(id string) error {
	query := "delete from calculations where id = $1"

	_, err := r.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
