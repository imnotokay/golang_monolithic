package routes

import (
	customershandler "webapp/handlers/customersHandler"

	"github.com/gorilla/mux"
)

func CreateCustomerRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/customers", customershandler.GetAll).Methods("GET")
	mux.HandleFunc("/api/customers/{id:[0-9]+}", customershandler.GetById).Methods("GET")
	mux.HandleFunc("/api/customers", customershandler.Create).Methods("POST")
	mux.HandleFunc("/api/customers/{id:[0-9]+}", customershandler.Update).Methods("PUT")
	mux.HandleFunc("/api/customers/{id:[0-9]+}", customershandler.Delete).Methods("DELETE")
}
