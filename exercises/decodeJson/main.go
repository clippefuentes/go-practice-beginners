package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool

type user struct {
	Name     string `json:"username"`
	Permissions permissions `json:"perms,omitempty"`
}


func main() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}