package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)


func main() {
	names := os.Args[1:]
	if len(names) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}
	sort.Strings(names)

	var data []byte

	for i, file := range names {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, file...)
		data = append(data, '\n')
	}

	err := ioutil.WriteFile("Test.txt", data, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}