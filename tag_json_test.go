package goJson

import (
	"encoding/json"
	"fmt"
	"goJson/model"
	"testing"
)

func TestJsonTagEncode(t *testing.T) {
	product := model.Product{
		Id:       "P0001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &model.Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
