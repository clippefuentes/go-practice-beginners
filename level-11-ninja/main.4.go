package main

import (
	// "errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat string
	long string
	err error
}

func (ce sqrtError) Error() string {
	return fmt.Sprintf("math error: %s %s %s", ce.lat, ce.long, ce.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt (f float64) (float64, error) {
	if f < 0 {
		// e := errors.New("more money needed")
		e :=  fmt.Errorf("more money needed value was %v", f)
		return 0, sqrtError{lat: "negative", long: "number", err: e}
	}
	return 42, nil
}