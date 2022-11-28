package product_entities

import (
	"time"
	"webapp/providers/database"
)

type Products struct {
	Id               uint32
	Name             string
	UnitPrice        float32
	Tax              float32
	CreationDate     time.Time
	ModificationDate time.Time
}

type ProductList []Products

func GetAll() ProductList {
	query := "SELECT id, name, unitPrice, tax, creationDate, modificationDate FROM products"
	rows := database.GetData(query)
	customers := ProductList{}
	var tmpcustomer Products
	for rows.Next() {
		tmpcustomer = Products{}
		rows.Scan(&tmpcustomer.Id, &tmpcustomer.Name, &tmpcustomer.UnitPrice, &tmpcustomer.Tax, &tmpcustomer.CreationDate, &tmpcustomer.ModificationDate)
		customers = append(customers, tmpcustomer)
	}
	return customers
}

func GetById(id int) Products {
	query := "SELECT id, name, unitPrice, tax, creationDate, modificationDate FROM products WHERE id = ?"
	rows := database.GetData(query, id)
	var tmpcustomer Products = Products{}
	if rows.Next() {
		rows.Scan(&tmpcustomer.Id, &tmpcustomer.Name, &tmpcustomer.UnitPrice, &tmpcustomer.Tax, &tmpcustomer.CreationDate, &tmpcustomer.ModificationDate)
	}
	return tmpcustomer
}

func Create(record Products) {
	query := "INSERT INTO products (name, unitPrice, tax, creationDate) values (?,?,?,current_timestamp)"
	database.ExecuteQuery(query, record.Name, record.UnitPrice, record.Tax)
}

func Update(id int, record Products) {
	query := "UPDATE products SET name = ?, unitPrice = ?, tax = ?, modificationDate = current_timestamp WHERE id = ?"
	database.ExecuteQuery(query, record.Name, record.UnitPrice, record.Tax, id)
}

func Delete(id int) {
	query := "DELETE FROM products WHERE id = ?"
	database.ExecuteQuery(query, id)
}
