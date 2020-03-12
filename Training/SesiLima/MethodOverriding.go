package SesiLima

import "fmt"

// define a method in Employee
func (e *Employee) SayHi() {
	fmt.Printf("Hi, nama saya %s dan umur saya %d tahun, saya bekerja di %s\n", e.Name, e.Age, e.Company)
}

// define a method in Employee
func (e *Member) SayHi() {
	fmt.Printf("Hi, nama saya %s dan umur saya %d tahun, job saya %s, build %s\n", e.Name, e.Age, e.Job, e.Build)
}

func SesiLimaMethodOverriding() {
	upin := Employee{&Human{"Upin", 22, "123-456-789"}, "Google Inc"}
	ipin := Student{&Human{"Ipin", 22, "987-654-321"}, "MIT"}
	heruru := Member{&Human{"heruru", 22, "987-654-321"}, "Dragon Knight", "Bash"}

	upin.SayHi()
	ipin.SayHi()
	heruru.SayHi()
}
