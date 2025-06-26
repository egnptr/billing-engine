package usecase

import "fmt"

func (l *Loan) MakePayment(amount float64) error {
	if l.CurrentWeek >= l.TotalWeeks {
		return fmt.Errorf("loan already completed")
	}

	switch {
	case amount == l.Installment:
		l.Payments[l.CurrentWeek] = true
	case amount == 0:
		l.Payments[l.CurrentWeek] = false
	default:
		return fmt.Errorf("must pay either 0 or exact installment: %.2f", l.Installment)
	}

	l.CurrentWeek++
	return nil
}

func (l *Loan) GetOutstanding() float64 {
	remaining := 0
	for i := 0; i < l.TotalWeeks; i++ {
		if !l.Payments[i] {
			remaining += 1
		}
	}
	return float64(remaining) * l.Installment
}

func (l *Loan) IsDelinquent() bool {
	if l.CurrentWeek < 2 {
		return l.IsDelinq
	}

	var count int
	for i := 0; i < l.CurrentWeek; i++ {
		if !l.Payments[i] {
			count++
		}
		if count > 2 {
			l.IsDelinq = true
			return l.IsDelinq
		}
	}

	return l.IsDelinq
}
