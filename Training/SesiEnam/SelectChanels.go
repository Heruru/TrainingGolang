package SesiEnam

import "fmt"

func fibonacciSelect(channel, quit chan int) {
	number1, number2 := 1, 1
	for {
		select {
		case channel <- number1:
			number1, number2 = number2, number1+number2
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func SesiEnamSelectChannels() {
	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()
	fibonacciSelect(channel, quit)
}
