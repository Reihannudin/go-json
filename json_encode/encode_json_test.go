package json_encode

import (
	"encoding/json"
	"goJson/model"
	"os"
	"testing"
)

func TestEncodeJson(t *testing.T) {

	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	//	get struct data
	customer := model.Customer{
		FirstName:  "Min",
		MiddleName: "Yonggi",
		LastName:   "D",
	}

	encoder.Encode(customer)

}
