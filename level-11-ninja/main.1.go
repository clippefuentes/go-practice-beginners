package main

import (
	"encoding/json"
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

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatalln("JSON did not marshall - errir", err)
	}

	fmt.Println(string(bs))
}