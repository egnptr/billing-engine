package usecase

import (
	"fmt"
	"log"
	"math"
)

func (l *Loan) MakePayment(amount float64) (err error) {
	// check no payment and the last week
	if amount == 0 && l.CurrentWeek == l.TotalWeeks-1 {
		err = fmt.Errorf("payment must be completed on the last week")
		log.Println(err)
		return
	}

	// check payment complete
	if l.CurrentWeek >= l.TotalWeeks {
		err = fmt.Errorf("loan already completed")
		log.Println(err)
		return
	}

	// check no payment
	if amount == 0 {
		l.CurrentWeek++
		return nil
	}

	// check exact amount
	if math.Mod(amount, l.Installment) != 0 {
		err = fmt.Errorf("payment must be multipler of %.f", l.Installment)
		log.Println(err)
		return
	}

	// check for unpaid weeks
	weeksUnpaid := 0
	for i := 0; i <= l.CurrentWeek; i++ {
		if !l.Payments[i] {
			weeksUnpaid++
		}
	}

	// check amount is the exact amount
	if weeksUnpaid > 1 {
		newInstallment := l.Installment * float64(weeksUnpaid) // make 2 or more payment
		if amount != newInstallment {
			err = fmt.Errorf("%d weeks pending, payment must be %.f", weeksUnpaid, newInstallment)
			log.Println(err)
			return
		}
		for i := 0; i <= l.CurrentWeek; i++ {
			l.Payments[i] = true
		}
	} else {
		if amount != l.Installment {
			err = fmt.Errorf("payment must be the exact amount %.f", l.Installment)
			log.Println(err)
			return
		}
		l.Payments[l.CurrentWeek] = true
	}

	l.Amount -= amount
	l.CurrentWeek++
	return nil
}

func (l *Loan) GetOutstanding() float64 {
	return l.Amount
}

func (l *Loan) IsDelinquent() bool {
	if l.CurrentWeek < 2 {
		return l.IsDelinq
	}

	l.IsDelinq = !l.Payments[l.CurrentWeek-1] && !l.Payments[l.CurrentWeek-2]
	return l.IsDelinq
}
