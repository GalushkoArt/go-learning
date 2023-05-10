package basics

import (
	"errors"
	"fmt"
	"log"
)

func Iterations() {
	numbers := []int{0, 12, 1, -887, 5492352, -7348328}
	min, err := funcMin(numbers...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(min)
	infiniteLoop()
	value := 0
	for value < 5 { //while
		fmt.Println("it's true")
		value++
	}
	fmt.Println(value)
	for i := range numbers {
		fmt.Println(i, numbers[i])
	}
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

func funcMin(numbers ...int) (int, error) {
	if len(numbers) < 1 {
		return 0, errors.New("no numbers passed")
	}
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min, nil
}

func infiniteLoop() {
	message := ""
	for {
		message += "A"
		fmt.Println(message + "!")
		if len(message) > 20 {
			break
		}
	}
}
