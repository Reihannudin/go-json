package goJson

import (
	"encoding/json"
	"fmt"
	"goJson/model"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := model.Customer{
		FirstName:  "Min",
		MiddleName: "Yonggi",
		LastName:   "D",
	}

	encoder.Encode(customer)
}

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("CustomerOut.json")
	decoder := json.NewDecoder(reader)

	customer := &model.Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)

}
