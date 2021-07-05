package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age int
}

func main () {
	u1 := user{
		First: "Clyne",
		Age: 24,
	}
	u2 := user{
		First: "Sheena",
		Age: 23,
	}
	u3 := user{
		First: "Clippe",
		Age: 25,
	}
	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
