package app

import (
	"encoding/json"
	"net/http"

	"github.com/aditiadika/banking/service"
)

type Customer struct {
	ID   int
	Name string `json:"full_name"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
