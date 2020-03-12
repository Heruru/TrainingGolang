package SesiLima

import "fmt"

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, nama saya %s dan kalian dapat menghubungi saya di no %s\n", h.Name, h.Phone)
}

func SesiLimaMethodInheritance() {
	upin := Employee{&Human{"Upin", 22, "123-456-789"}, "Google Inc"}
	ipin := Student{&Human{"Ipin", 22, "987-654-321"}, "MIT"}

	upin.SayHi()
	ipin.SayHi()
}
