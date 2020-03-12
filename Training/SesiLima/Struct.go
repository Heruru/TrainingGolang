package SesiLima

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

type Test struct {
	a, b, c float64
}

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	*Human // anonymous field
	School string
}

type Employee struct {
	*Human  // anonymous field
	Company string
}

type Member struct {
	*Human // anonymous field
	Job    string
	Build  string
}

//interface Men implemented by Human, Student and Employee
type Men interface {
	SayHi()
	Sing(title string)
	Singx(title string)
}
