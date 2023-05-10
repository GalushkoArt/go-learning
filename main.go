package main

import "fmt"
import "go-learinng/basics"

func main() {
	fmt.Println("Hello world")
	basics.Variables()
	basics.Conditions(25)
	basics.Iterations()
	counter := basics.Counter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	basics.Pointers()
	basics.Slices()
	basics.Maps()
	basics.Panics()
	basics.Structures()
	basics.Interfaces()
	month := basics.May
	basics.Enums(&month)
	basics.Generics()
}
