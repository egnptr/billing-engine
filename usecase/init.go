package usecase

type Loan struct {
	ID           int64
	InitalAmount float64
	InterestRate float64
	TotalWeeks   int
	Installment  float64
	Payments     []bool
	CurrentWeek  int
	IsDelinq     bool
}

func NewLoan(id int64, amount float64, interestRate float64, weeks int) *Loan {
	total := amount + (amount * interestRate)
	installment := total / float64(weeks)

	return &Loan{
		ID:           id,
		InitalAmount: amount,
		InterestRate: interestRate,
		TotalWeeks:   weeks,
		Installment:  installment,
		Payments:     make([]bool, weeks),
		CurrentWeek:  0,
		IsDelinq:     false,
	}
}
