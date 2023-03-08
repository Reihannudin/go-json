package goJson

import (
	"encoding/json"
	"fmt"
	"goJson/model"
	"testing"
)

func TestDecode(t *testing.T) {

	jsonRequest := `{"FirstName" : "Andrian" , "Lastname" : "Raihannudin"}`
	jsonBytes := []byte(jsonRequest)

	customer := &model.Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)

}
