package SesiEnam

import "fmt"

func SesiEnamBufferedChannels() {
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
