package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webapp/handlers"
	"webapp/routes"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	handlers.Render(w, "index.html", nil)
}

func customers(w http.ResponseWriter, r *http.Request) {
	handlers.Render(w, "customers.html", nil)
}

func products(w http.ResponseWriter, r *http.Request) {
	handlers.Render(w, "products.html", nil)
}

func main() {
	mux := mux.NewRouter()

	staticPath := http.FileServer(http.Dir("client/static"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticPath))

	handlers.LoadTemplates()

	mux.HandleFunc("/", index).Methods("GET")
	mux.HandleFunc("/customers", customers).Methods("GET")
	mux.HandleFunc("/products", products).Methods("GET")

	routes.CreateCustomerRoutes(mux)
	routes.CreateProductsRoutes(mux)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Server is running on port 3000")
	log.Fatal(server.ListenAndServe())
}
