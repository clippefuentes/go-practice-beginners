package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	json1 := `[{"First":"Clyne","Age":24},{"First":"Sheena","Age":23},{"First":"Clippe","Age":25}]`
	fmt.Println(json1)

	var people []person

	err := json.Unmarshal([]byte(json1), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("First: ", person.First)
		fmt.Println("Age: ", person.Age)
	}
}
