package usecase

import "fmt"

type Loan struct {
	ID          int64   `json:"id"`
	Amount      float64 `json:"amount"`
	TotalWeeks  int     `json:"total_weeks"`
	Installment float64 `json:"installment"`
	Payments    []bool  `json:"payments"`
	CurrentWeek int     `json:"current_week"`
	IsDelinq    bool    `json:"delinquent"`
}

func NewLoan(id int64, amount float64, interestRate float64, weeks int) (*Loan, error) {
	total := amount + (amount * interestRate)
	installment := total / float64(weeks)

	if total == 0 || installment == 0 {
		return &Loan{}, fmt.Errorf("invalid loan")
	}

	return &Loan{
		ID:          id,
		Amount:      total,
		TotalWeeks:  weeks,
		Installment: installment,
		Payments:    make([]bool, weeks),
		CurrentWeek: 0,
		IsDelinq:    false,
	}, nil
}
