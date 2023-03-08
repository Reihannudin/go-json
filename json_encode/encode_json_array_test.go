package json_encode

import (
	"encoding/json"
	model2 "goJson/model"
	"os"
	"testing"
)

func TestEncodeJsonArray(t *testing.T) {

	writer, _ := os.Create("UserOut.json")
	encoder := json.NewEncoder(writer)

	user := model2.User{
		UserName: "reihannudin",
		Email:    "reedbulls91@gmail.com",
		Hobbies:  []string{"Gaming", "Coding", "Singing"},
	}

	encoder.Encode(user)

	//bytes , _ := json.Marshal(user)
	//fmt.Println(string(bytes))

}

func TestEncodeArrayComplex(t *testing.T) {

	writer, _ := os.Create("FullDataUser.json")
	encoder := json.NewEncoder(writer)

	user := model2.User{
		UserName: "reihannudin",
		Email:    "reedbulls91@gmail.com",
		Hobbies:  []string{"Gaming", "Coding", "Singing"},
		Addresses: []model2.Address{
			{
				Street:     "st.bangtan 07",
				District:   "Seoul",
				PostalCode: "1307213",
				Country:    "South Korea",
			},
		},
	}

	encoder.Encode(user)

	//bytes , _ := json.Marshal(user)
	//fmt.Println(string(bytes))
}
