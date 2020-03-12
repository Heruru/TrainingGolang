package SesiLima

import "fmt"

// method
func (h *Human) Sing(title string) {
	fmt.Println("Saya suka menyanyikan lagu ", title)
}

func (h *Human) Singx(title string) {
	fmt.Println("Saya suka menyanyikan lagu ", title)
}

func SesiLimaInterface() {
	upin := &Employee{&Human{"Upin", 22, "123-456-789"}, "Google Inc"}
	ipin := Student{&Human{"Ipin", 22, "987-654-321"}, "MIT"}

	// define interface iName
	var iName Men

	// iName store employee
	iName = upin
	iName.SayHi()
	iName.Sing("November Rain")

	// iName store student
	iName = ipin
	iName.SayHi()
	iName.Sing("Born to be Wild")
	iName.Singx("OM Telolet OM")
}
