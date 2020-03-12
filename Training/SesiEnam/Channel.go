package SesiEnam

import "fmt"

func sum(arrayInt []int, channelInt chan int) {
	total := 0
	for _, value := range arrayInt {
		total += value
	}

	channelInt <- total // send total to channelInt
}

func SesiEnamChannel() {
	arrayInt := []int{7, 2, 8, -9, 4, 0}

	channelInt := make(chan int)
	go sum(arrayInt[:len(arrayInt)/2], channelInt)
	go sum(arrayInt[len(arrayInt)/2:], channelInt)
	result1, result2 := <-channelInt, <-channelInt // receive from channelInt

	fmt.Println(result1, result2, result1+result2)
}
