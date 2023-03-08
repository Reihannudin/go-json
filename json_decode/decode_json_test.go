package json_decode

import (
	"encoding/json"
	"fmt"
	"goJson/model"
	"testing"
)

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"UserName":"Reihan","Email":"reedbulls91@gmail.com","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &model.User{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.UserName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"UserName":"reihannudin","Email":"reedbulls91@gmail.com","Addresses":[{"Street":"st.bangtan 07","District":"Seoul","PostalCode":"1307213","Country":"South Korea"}],"Hobbies":["Gaming","Coding","Singing"]}`
	jsonBytes := []byte(jsonString)

	customer := &model.User{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.UserName)
	fmt.Println(customer.Addresses)
}
