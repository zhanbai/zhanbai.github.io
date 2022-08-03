package main

import "fmt"

func getDiameter() string {
	return "20"
}

func main() {
	bn := 2
	gn := 2
	fmt.Println("You have " + fmt.Sprint(bn+gn) + " children.")
	w := "hello"
	fmt.Println("The word " + w + " is " + fmt.Sprint(len([]rune(w))) + " characters long.")
	p := "20.5"
	fmt.Println("You owe " + p + " immediately.")
	fmt.Println("My circle's diameter is " + getDiameter() + " inches.")
}
