package productshandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	product_entities "webapp/entities/product_entity"
	"webapp/providers"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {

	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	resp.Data = product_entities.GetAll()
	resp.Send()
}

func GetById(w http.ResponseWriter, r *http.Request) {
	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	values := mux.Vars(r)
	id, _ := strconv.Atoi(values["id"])
	resp.Data = product_entities.GetById(id)
	resp.Send()
}

func Create(w http.ResponseWriter, r *http.Request) {

	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	encoder := json.NewDecoder(r.Body)
	customer := product_entities.Products{}
	encoder.Decode(&customer)
	product_entities.Create(customer)
	resp.Data = customer
	resp.Send()
}

func Update(w http.ResponseWriter, r *http.Request) {

	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	values := mux.Vars(r)
	id, _ := strconv.Atoi(values["id"])
	encoder := json.NewDecoder(r.Body)
	customer := product_entities.Products{}
	encoder.Decode(&customer)
	product_entities.Update(id, customer)
	resp.Data = customer
	resp.Send()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	values := mux.Vars(r)
	id, _ := strconv.Atoi(values["id"])
	product_entities.Delete(id)
	resp.Send()
}
