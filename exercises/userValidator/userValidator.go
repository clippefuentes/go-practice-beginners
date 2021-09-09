package main

import (
	"fmt"
	"os"
)

func main() {
	correctUser, correctPass, correctUser2, correctPass2 := "clyne", "1803", "test", "1111"
	user, pass := os.Args[1], os.Args[2]
	fmt.Println(user != correctUser || user != correctUser2)
	if len(os.Args) != 3 {
		fmt.Println("Usage: [Username] [Password]")
	} else if user != correctUser && user != correctUser2 {
		fmt.Printf("Denied access for %s", os.Args[1])
	} else if (pass != correctPass && user == correctUser) || (pass != correctPass2 && user == correctUser2) {
		fmt.Printf("Wrong password for \"%s\"", user)
	} else {
		fmt.Printf("Access granted to %s", user)
	}
}