package customershandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/entities/customer_entity"
	"webapp/providers"

	"github.com/gorilla/mux"
)

type CustomersList []customer_entity.Customers

func GetAll(w http.ResponseWriter, r *http.Request) {

	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	resp.Data = customer_entity.GetAll()
	resp.Send()
}

func GetById(w http.ResponseWriter, r *http.Request) {
	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	values := mux.Vars(r)
	id, _ := strconv.Atoi(values["id"])
	resp.Data = customer_entity.GetById(id)
	resp.Send()
}

func Create(w http.ResponseWriter, r *http.Request) {

	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	encoder := json.NewDecoder(r.Body)
	customer := customer_entity.Customers{}
	encoder.Decode(&customer)
	customer_entity.Create(customer)
	resp.Data = customer
	resp.Send()
}

func Update(w http.ResponseWriter, r *http.Request) {

	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	values := mux.Vars(r)
	id, _ := strconv.Atoi(values["id"])
	encoder := json.NewDecoder(r.Body)
	customer := customer_entity.Customers{}
	encoder.Decode(&customer)
	customer_entity.Update(id, customer)
	resp.Data = customer
	resp.Send()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	resp := providers.DefaultResponse(w)
	defer providers.CatchExceptions(resp)
	values := mux.Vars(r)
	id, _ := strconv.Atoi(values["id"])
	customer_entity.Delete(id)
	resp.Send()
}
