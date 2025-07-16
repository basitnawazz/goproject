// this the'Main' package
// Evry Go program must haev a main package
// this is the entry point for the application
package main

import (
	"fmt"
	"learingo/basicmaths"
	"learingo/stringman"
)

func main() {
	sum := basicmaths.AddTwoNumber(34, 9685)
	fmt.Println(sum)

	sub := basicmaths.SubtwoNumbers(44, 66)
	fmt.Println(sub)

	fmt.Println(sub + sum)

	mul := basicmaths.MulTwoNumber(8, 88)
	fmt.Println(mul)

	fullname := stringman.AddLastnameFirstname("basit", "nawaz")
	fmt.Println(fullname)
}
