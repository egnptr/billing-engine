# billing-engine

## Features

- Create a loan with initial amount, interest, and duration
- Track weekly payments and outstanding balance
- Handle missed payments and enforce repayment catch-up 
- Check if borrower is delinquent (missed 2 consecutive payments)

## Restriction

- Total week duration are not expandable
- Loan must be completed on designated week

## Getting Started

### Install Dependencies

```bash
go mod tidy
```

### Run the Server

```
go run ./cmd/main.go
```

or by building and running the binary file from Makefile:

```
make run
```

Server will start at http://localhost:8080

## API endpoints

### POST /loan

Create a new loan.

**Request Body**

```
{
  "id": 1,
  "initial_amount": 5000000,
  "interest_rate": 0.10,
  "weeks": 50
}
```

### POST /loan/{id}/payment

Make weekly payment of certain amount.

**Request Body**

```
{
  "amount": 110000
}
```

---

### GET /loan/{id}/outstanding

Get outstanding balance of the loan (defined as pending amount the borrower needs to pay at any point).

**Response Body**

```
{
  "outstanding": 5280000
}
```

---

### GET /loan/{id}/delinquent

Check if loan have missed 2 continuous repayments.

**Response Body**

```
{
  "delinquent": true
}
```

---

### GET /loan/{id}

Get loan details.

**Response Body**

```
{
  "ID": 1,
  "Amount": 5500000,
  "TotalWeeks": 5,
  "Installment": 1100000,
  "Payments": [
    false,
    false,
    false,
    false,
    false
  ],
  "CurrentWeek": 0,
  "IsDelinq": false
}
```
