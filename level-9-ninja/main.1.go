package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main () {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go rountines\t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()

	fmt.Println("mid CPUs\t\t", runtime.NumCPU())
	fmt.Println("mid Go rountines\t\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Waiting")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go rountines\t\t", runtime.NumGoroutine())
}
