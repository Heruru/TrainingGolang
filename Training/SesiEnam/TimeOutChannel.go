package SesiEnam

import (
	"fmt"
	"time"
)

func TimeOut(channel1 chan int, channel2 chan bool) {
	for {
		select {
		case value := <-channel1:
			fmt.Println(value)
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			channel2 <- true
			break
		}
	}
}

func SesiEnamTimeOutChannels() {
	channel1 := make(chan int)
	channel2 := make(chan bool)
	go func() {
		TimeOut(channel1, channel2)
	}()
	<-channel2
}
