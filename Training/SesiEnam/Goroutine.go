package SesiEnam

import (
	"fmt"
	"runtime"
)

func print(text string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(text)
	}
}

func SesiEnamGoroutine() {
	go print("Hello")   // create new goroutine
	go print("Another") // existing goroutine
	print("World")      // existing goroutine
	// print("Another")  // existing goroutine

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
