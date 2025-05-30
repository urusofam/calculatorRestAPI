package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urusofam/calculatorRestAPI/internal/server/models/calculation"
)

type Storage struct {
	conn *pgxpool.Pool
}

func NewStorage(dbUrl string) (*Storage, error) {
	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Storage{conn: conn}, nil
}

func (s *Storage) Close() {
	s.conn.Close()
}

func (s *Storage) GetCalculations() ([]calculation.Calculation, error) {
	query := "select * from calculations"

	rows, err := s.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []calculation.Calculation{}
	for rows.Next() {
		calc := calculation.Calculation{}

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

func (s *Storage) AddCalculation(calc calculation.Calculation) error {
	query := "insert into calculations (id, expression, result) values ($1, $2, $3)"

	_, err := s.conn.Exec(context.Background(), query, calc.ID, calc.Expression, calc.Result)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateCalculation(expression, result, id string) error {
	query := "update calculations set expression = $1, result = $2 where id = $3"

	_, err := s.conn.Exec(context.Background(), query, expression, result, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteCalculation(id string) error {
	query := "delete from calculations where id = $1"

	_, err := s.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
