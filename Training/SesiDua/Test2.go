package SesiDua

import (
	"fmt"
	"math/rand"
)

func TestIF() {
	var number = rand.Intn(20)

	fmt.Println(number)

	if number < 10 {
		fmt.Println("Nomor lebih kecil dari 10")
	} else if number == 10 {
		fmt.Println("Nomor sama dengan 10")
	} else {
		fmt.Println("Nomor lebih besar dari 10")
	}
}

func TestGoTo() {
	number := 0

testLabel:
	fmt.Println(number)
	number++
	if number < 100 {
		goto testLabel
	}
}

func TestForExpression() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println(sum)
}

func TestForTanpaExpression() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func VariadicFunc(args ...int) {
	for key, val := range args {
		fmt.Printf("Parameter index ke %d = %d\n", key, val)
	}
}

func TestArray() {
	arrayTest := [4]int{74, 13, 22, 49}
	for index, value := range arrayTest {
		fmt.Printf("Index %d = %d\n", index, value)
	}
}
