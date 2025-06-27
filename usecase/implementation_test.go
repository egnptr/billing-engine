package usecase

import (
	"testing"
)

func TestLoan_MakePayment(t *testing.T) {
	type fields struct {
		ID          int64
		Amount      float64
		TotalWeeks  int
		Installment float64
		Payments    []bool
		CurrentWeek int
		IsDelinq    bool
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{

			name: "case sucess",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				Payments:    make([]bool, 50),
			},
			args: args{
				amount: 110000,
			},
			wantErr: false,
		},
		{

			name: "case sucess amount 0",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				Payments:    make([]bool, 50),
			},
			args: args{
				amount: 0,
			},
			wantErr: false,
		},
		{

			name: "case sucess with missed payment",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				CurrentWeek: 3,
				Payments:    []bool{true, false, false, false},
			},
			args: args{
				amount: 220000,
			},
			wantErr: false,
		},
		{
			name: "case error with missed payment",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				CurrentWeek: 3,
				Payments:    []bool{true, false, false, false},
			},
			args: args{
				amount: 110000,
			},
			wantErr: true,
		},
		{
			name: "case error not exact amount",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				Payments:    make([]bool, 50),
			},
			args: args{
				amount: 10000,
			},
			wantErr: true,
		},
		{
			name: "case error payment complete",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  3,
				Installment: 110000,
				CurrentWeek: 3,
				Payments:    make([]bool, 50),
			},
			args: args{
				amount: 110000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loan{
				ID:          tt.fields.ID,
				Amount:      tt.fields.Amount,
				TotalWeeks:  tt.fields.TotalWeeks,
				Installment: tt.fields.Installment,
				Payments:    tt.fields.Payments,
				CurrentWeek: tt.fields.CurrentWeek,
				IsDelinq:    tt.fields.IsDelinq,
			}
			if err := l.MakePayment(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Loan.MakePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoan_GetOutstanding(t *testing.T) {
	type fields struct {
		ID          int64
		Amount      float64
		TotalWeeks  int
		Installment float64
		Payments    []bool
		CurrentWeek int
		IsDelinq    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "case sucess",
			fields: fields{
				ID:          1,
				Amount:      5060000,
				TotalWeeks:  50,
				Installment: 110000,
				CurrentWeek: 5,
				Payments:    make([]bool, 50),
			},
			want: 5060000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loan{
				ID:          tt.fields.ID,
				Amount:      tt.fields.Amount,
				TotalWeeks:  tt.fields.TotalWeeks,
				Installment: tt.fields.Installment,
				Payments:    tt.fields.Payments,
				CurrentWeek: tt.fields.CurrentWeek,
				IsDelinq:    tt.fields.IsDelinq,
			}
			if got := l.GetOutstanding(); got != tt.want {
				t.Errorf("Loan.GetOutstanding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoan_IsDelinquent(t *testing.T) {
	type fields struct {
		ID          int64
		Amount      float64
		TotalWeeks  int
		Installment float64
		Payments    []bool
		CurrentWeek int
		IsDelinq    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case is delinquent",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				CurrentWeek: 4,
				Payments:    []bool{true, true, false, false, true},
			},
			want: true,
		},
		{
			name: "case is not delinquent",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				CurrentWeek: 4,
				Payments:    []bool{true, false, true, true, true},
			},
			want: false,
		},
		{
			name: "case is not delinquent (less than 2 weeks)",
			fields: fields{
				ID:          1,
				Amount:      5000000,
				TotalWeeks:  50,
				Installment: 110000,
				CurrentWeek: 1,
				Payments:    []bool{true},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loan{
				ID:          tt.fields.ID,
				Amount:      tt.fields.Amount,
				TotalWeeks:  tt.fields.TotalWeeks,
				Installment: tt.fields.Installment,
				Payments:    tt.fields.Payments,
				CurrentWeek: tt.fields.CurrentWeek,
				IsDelinq:    tt.fields.IsDelinq,
			}
			if got := l.IsDelinquent(); got != tt.want {
				t.Errorf("Loan.IsDelinquent() = %v, want %v", got, tt.want)
			}
		})
	}
}
