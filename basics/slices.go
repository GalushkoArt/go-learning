package basics

import "fmt"

func Slices() {
	slice := make([]int, 2)
	fmt.Println(slice, len(slice), cap(slice))
	sliceWithCap := make([]int, 1, 2)
	fmt.Println(sliceWithCap, len(sliceWithCap), cap(sliceWithCap))
	sliceWithCap = append(sliceWithCap, 3)
	fmt.Println(sliceWithCap, len(sliceWithCap), cap(sliceWithCap))
	sliceWithCap = append(sliceWithCap, 4)
	fmt.Println(sliceWithCap, len(sliceWithCap), cap(sliceWithCap))
	matrix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, len(matrix))
		for j := i; j < len(matrix); j++ {
			matrix[i][j] = 1
		}
		fmt.Println(matrix[i])
	}
}
