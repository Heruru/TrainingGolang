package SesiSatu

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

func TestPrint() {
	fmt.Println("It's works!")
}

func Hello() string {
	return quote.Hello()
}

func TestLen() {
	fmt.Println(len("It's works!"))
}

func TestMath() {
	fmt.Println("1 + 1 =", 1+1)
}

func TestDecimal() {
	fmt.Println("1.2 + 1.3 =", 1.2+1.3)
}

func TestError() {
	err := errors.New("Data not complete!")
	if err != nil {
		fmt.Println(err)
	}
}

func TestVariable() {
	var (
		nama   = "David Maulana"
		umur   = 17
		alamat = "Jakarta"
	)

	fmt.Printf("%s berumur %d tahun\n", nama, umur)
	fmt.Println("Beralamat di ", alamat)
}

func TestStruct() {
	type Human struct {
		name string
		age  int
	}

	type Student struct {
		Human Human
		class string
	}

	//student1 := Student{Human{"Andi", 34}, "Intro to Programming"}
	student2 := Student{Human{"Andi2", 24}, "Intro to Programming 2"}

	fmt.Println("Nama saya adalah ", student2.Human.name)
	fmt.Println("Umur saya adalah ", student2.Human.age)
	fmt.Println("Saya ikut kelas ", student2.class)
}

func TestStruct1() {
	type person struct {
		name string
		age  int
	}

	var P person

	P.name = "Andi"
	P.age = 34

	fmt.Println("Nama saya adalah ", P.name)
	fmt.Println("Umur saya ", P.age)
}
