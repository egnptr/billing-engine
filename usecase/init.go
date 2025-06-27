package usecase

import "fmt"

type Loan struct {
	ID          int64
	Amount      float64
	TotalWeeks  int
	Installment float64
	Payments    []bool
	CurrentWeek int
	IsDelinq    bool
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
