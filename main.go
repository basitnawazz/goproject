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

	// basicmaths.SubtwoNumbers(3, 78)
	// // fmt.Println(sub)

	// fmt.Println(sub + sum)

	fmt.Println(basicmaths.MulTwoNumber(0, 88))

	fullname := stringman.AddLastnameFirstname("basit", "nawaz")
	fmt.Println(fullname)

	basicmaths.DivTwoNumber(7, 3)

	fmt.Println(basicmaths.GanerateRandNumber(70))

	fmt.Println(basicmaths.SubtwoNumbers(70, 100))

	fmt.Println(basicmaths.MulTwoNumber(4, 55))

}
