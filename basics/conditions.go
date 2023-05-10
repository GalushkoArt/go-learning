package basics

import (
	"errors"
	"fmt"
	"log"
)

func Conditions(age int) {
	printMessage(getMessage("How it going?"))
	printMessage(nil)
	link, err := tryToEnter(age)
	if err == nil {
		printMessage(getMessage("Go to localhost" + link))
	} else {
		log.Fatal(err)
	}
}

func printMessage(message *string) {
	if message == nil {
		fmt.Println("Here we go nil")
	} else {
		fmt.Println(*message)
	}
}

func getMessage(input string) *string {
	message := "It's Johnny! " + input
	return &message
}

func tryToEnter(age int) (string, error) {
	if age >= 75 {
		return "https://youtube.com", errors.New("you cannot enter. you're too old. try something else")
	} else if age >= 18 {
		return "/index.html", nil
	} else {
		return "https://google.com", errors.New("you cannot enter. you're too young. please avoid this site")
	}
}
