package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/egnptr/billing-engine/usecase"
	"github.com/gorilla/mux"
)

var (
	loans = make(map[int64]*usecase.Loan)
	mu    sync.Mutex
)

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID           int64   `json:"id"`
		InitalAmount float64 `json:"initial_amount"`
		InterestRate float64 `json:"interest_rate"`
		Weeks        int     `json:"weeks"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	// parse loan param
	if payload.ID <= 0 {
		http.Error(w, "Loan ID must be greater than 0", http.StatusBadRequest)
		return
	}
	if payload.InitalAmount == 0 {
		payload.InitalAmount = 5000000 // default initial amount
	}
	if payload.InterestRate == 0 {
		payload.InterestRate = 0.1 // default interest rate
	}
	if payload.Weeks == 0 {
		payload.Weeks = 50 // default weeks
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := loans[payload.ID]; ok {
		http.Error(w, "Loan ID already exist", http.StatusBadRequest)
		return
	}

	loan, err := usecase.NewLoan(payload.ID, payload.InitalAmount, payload.InterestRate, payload.Weeks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	loans[payload.ID] = loan

	json.NewEncoder(w).Encode(loan)
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error converting string to int64", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, "Loan ID must be greater than 0", http.StatusBadRequest)
		return
	}

	var body struct {
		Amount float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusNotFound)
		return
	}

	if err := loan.MakePayment(body.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "payment recorded"})
}

func GetOutstanding(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error converting string to int64", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, "Loan ID must be greater than 0", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"outstanding": loan.GetOutstanding()})
}

func IsDelinquent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error converting string to int64", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, "Loan ID must be greater than 0", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"delinquent": loan.IsDelinquent()})
}

func GetLoan(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error converting string to int64", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, "Loan ID must be greater than 0", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(loan)
}
