package basics

import (
	"fmt"
	"reflect"
)

func Variables() {
	var integer int8
	integer = 1
	var someInt = 2
	const someConst string = "aaaaaaaaaaaa"
	var someRune rune = 'a'
	floatValue := 0.124
	var booleanValue bool
	booleanValue = true

	fmt.Println(integer, someInt, someConst, someRune, floatValue, booleanValue)
	fmt.Println(
		reflect.TypeOf(integer),
		reflect.TypeOf(someInt),
		reflect.TypeOf(someConst),
		reflect.TypeOf(someRune),
		reflect.TypeOf(floatValue),
		reflect.TypeOf(booleanValue),
	)

	a, b := "a", "b"
	b, a, _ = a, b, 5

	fmt.Println(a, b)

	var emptyString string
	var notDefinedInt int
	var notDefinedFloat float32
	var notDefinedRune rune
	var notDefinedBoolean bool
	bytesFromString := []byte("AaBaA")
	runesFromString := []rune("AaBaA")

	fmt.Println(emptyString, notDefinedInt, notDefinedFloat, notDefinedRune, notDefinedBoolean)
	fmt.Println(bytesFromString)
	fmt.Println(runesFromString)
}
