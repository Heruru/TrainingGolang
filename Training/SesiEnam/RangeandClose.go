package SesiEnam

import "fmt"

func fibonacci(length int, channel chan int) {
	number1, number2 := 1, 1
	for i := 0; i < length; i++ {
		channel <- number1
		number1, number2 = number2, number1+number2
	}
	close(channel)
}

func SesiEnamRangeandClose() {
	channel := make(chan int, 10)

	go fibonacci(cap(channel), channel)

	for i := range channel {
		fmt.Println(i)
	}
}
