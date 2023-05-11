package basics

import "fmt"

func Basics() {
	fmt.Println("Hello world")
	Variables()
	Conditions(25)
	Iterations()
	counter := Counter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	Pointers()
	Slices()
	Maps()
	Panics()
	Structures()
	Interfaces()
	month := May
	Enums(&month)
	Generics()
	Caches()
}
