package app

import (
	"log"
	"net/http"

	"github.com/aditiadika/banking/domain"
	"github.com/aditiadika/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// defining route
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8002", router))
}
