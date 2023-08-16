package main

import (
	"bytes"
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
	// converter struct em json
	// json.Marshal()

	// converter json em struct
	// json.Unmarshal()

	d := Dog{"Poodle", 5, "Jubiscreito"}
	fmt.Println(d)

	jsonDog, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	// irá printar um slice de uint8
	fmt.Println(jsonDog)

	// convertendo o slice de bytes em um string json
	fmt.Println(bytes.NewBuffer(jsonDog))
}
