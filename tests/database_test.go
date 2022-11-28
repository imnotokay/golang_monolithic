package tests

import (
	"fmt"
	"testing"
	"webapp/providers/database"
)

func TestExecute(t *testing.T) {
	defer (func() {
		err := recover()
		if err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			t.Error("An exception occurs when trying to execute the insert statement \r\nException detail: " + errorMessage)
		}
	})()
	database.ExecuteQuery("INSERT INTO customers (name, typeDoc, document, creationDate) values ('Noah Castillo', 1, '465789132', current_timestamp)")
}

func TestGetData(t *testing.T) {
	defer (func() {
		err := recover()
		if err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			t.Error("An exception occurs when trying to execute the insert statement \r\nException detail: " + errorMessage)
		}
	})()
	rows := database.GetData("SELECT name FROM customers WHERE document = '465789132'")
	if !rows.Next() {
		t.Error("No records were found")
	}
}
