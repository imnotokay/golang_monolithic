package routes

import (
	productshandler "webapp/handlers/productsHandler"

	"github.com/gorilla/mux"
)

func CreateProductsRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/products", productshandler.GetAll).Methods("GET")
	mux.HandleFunc("/api/products/{id:[0-9]+}", productshandler.GetById).Methods("GET")
	mux.HandleFunc("/api/products", productshandler.Create).Methods("POST")
	mux.HandleFunc("/api/products/{id:[0-9]+}", productshandler.Update).Methods("PUT")
	mux.HandleFunc("/api/products/{id:[0-9]+}", productshandler.Delete).Methods("DELETE")
}
