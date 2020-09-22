package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken not stirred",
			"Any last wishes?",
			"Never say never",
		},
	}

	bs, err := toJSON(p1)
	if err != nil {
		err := fmt.Errorf("JSON marshal failed: %v", err)
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	err = errors.New("test error")
	if err != nil {
		return nil, err
	}
	return bs, nil
}
