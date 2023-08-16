package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Breed string `json:"breed"`
	Years uint8  `json:"years"`
	Name  string `json:"name"`
}

func main() {
	jsonDog := `{"breed":"Poodle","years":5,"name":"Jubiscreito"}`
	var d Dog

	if err := json.Unmarshal([]byte(jsonDog), &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
}
