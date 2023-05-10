package basics

import "fmt"

func Pointers() {
	input := "Some String"
	fmt.Println(input, &input)
	printCopyPointer(input)
	changeString(&input)
	fmt.Println(input, &input)
	number := 5
	var pointer *int = &number
	fmt.Println(number, &number, pointer)
	*pointer = 12
	fmt.Println(number, &number, pointer)
	array := [...]int{1, 2, 3}
	printCopiedArray(array)
	fmt.Println(array)
	slice := []string{"a", "b", "c"}
	changeAndPrintSlice(slice)
	fmt.Println(slice)
}

func printCopyPointer(input string) {
	fmt.Println(&input)
}

func changeString(input *string) {
	fmt.Println(input)
	*input += " was changed!"
}

func printCopiedArray(array [3]int) {
	array[2] = 5
	fmt.Println(array)
}

func changeAndPrintSlice(slice []string) {
	slice[0] = "1"
	fmt.Println(slice)
}
