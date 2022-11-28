package customer_entity

import (
	"time"
	"webapp/providers/database"
)

type Customers struct {
	Id               uint32
	Name             string
	TypeDoc          uint8
	Document         string
	CreationDate     time.Time
	ModificationDate time.Time
}

type CustomersList []Customers

func GetAll() CustomersList {
	query := "SELECT id, name, typeDoc, document, creationDate, modificationDate FROM customers"
	rows := database.GetData(query)
	customers := CustomersList{}
	var tmpcustomer Customers
	for rows.Next() {
		tmpcustomer = Customers{}
		rows.Scan(&tmpcustomer.Id, &tmpcustomer.Name, &tmpcustomer.TypeDoc, &tmpcustomer.Document, &tmpcustomer.CreationDate, &tmpcustomer.ModificationDate)
		customers = append(customers, tmpcustomer)
	}
	return customers
}

func GetById(id int) Customers {
	query := "SELECT id, name, typeDoc, document, creationDate, modificationDate FROM customers WHERE id = ?"
	rows := database.GetData(query, id)
	var tmpcustomer Customers = Customers{}
	if rows.Next() {
		rows.Scan(&tmpcustomer.Id, &tmpcustomer.Name, &tmpcustomer.TypeDoc, &tmpcustomer.Document, &tmpcustomer.CreationDate, &tmpcustomer.ModificationDate)
	}
	return tmpcustomer
}

func Create(record Customers) {
	query := "INSERT INTO customers (name, typeDoc, document, creationDate) values (?,?,?,current_timestamp)"
	database.ExecuteQuery(query, record.Name, record.TypeDoc, record.Document)
}

func Update(id int, record Customers) {
	query := "UPDATE customers SET name = ?, typeDoc = ?, document = ?, modificationDate = current_timestamp WHERE id = ?"
	database.ExecuteQuery(query, record.Name, record.TypeDoc, record.Document, id)
}

func Delete(id int) {
	query := "DELETE FROM customers WHERE id = ?"
	database.ExecuteQuery(query, id)
}
