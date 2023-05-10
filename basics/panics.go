package basics

import (
	"fmt"
	"os"
)

func Panics() {
	defer fmt.Println("Defer 2")
	defer handelPanic()
	defer fmt.Println("Defer 1")
	slice := make([]int, 1)
	slice[1] = 1
	os.Exit(12)
}

func handelPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("Panic Handler")
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	panic("OH MY GOOOOOD!")
}
