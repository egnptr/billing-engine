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
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error converting string to int64", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	loan := usecase.NewLoan(id, 5000000, 0.1, 50)
	loans[id] = loan
	json.NewEncoder(w).Encode(loan)
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Error converting string to int64", http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusNotFound)
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

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusNotFound)
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

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[id]
	if !ok {
		http.Error(w, "Loan not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(loan)
}
