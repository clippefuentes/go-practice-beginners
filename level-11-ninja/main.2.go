package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Sayings []string
}

func main () {
	p1 := person{
		First: "Clyne",
		Last:  "Fuentes",
		Sayings: []string{"I have much money", "I have a lot of money", "I am a very rich"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln("JSON did not marshall - errir", err)
		return
	}

	fmt.Println(string(bs))
}

func toJSON (v interface{}) ([]byte, error) {
	bs, err := json.Marshal(v)

	if err != nil {
		// return []byte{}, fmt.Errorf("There was an error: %v", err)
		return []byte{}, errors.New(fmt.Sprintf("There was an error: %v", err))
	}
	return bs, nil
}