package basicmaths

import (
	"fmt"
	"math/rand"
)

// this function add two number and return its sum
func AddTwoNumber(a, b int) int {
	return a + b
}

func SubtwoNumbers(x, y int) (int, error) {
	if x <= 0 {
		return 0, fmt.Errorf("x must be greater than 0")
	}
	if y >= 99 {
		return 0, fmt.Errorf("y must be less than 99")
	}
	if x <= y {
		return 0, fmt.Errorf("x must be greater than y")
	}
	return x - y, nil
}

func MulTwoNumber(a, b int) (int, error) {
	if a <= 0 {
		return 0, fmt.Errorf("a must be greater than 0")
	}
	if b <= 0 {
		return 0, fmt.Errorf("b must be greater than 0")
	}
	return a * b, nil
}

func DivTwoNumber(a, b int) (int, error) {
	return fmt.Println(a / b)
}

func GanerateRandNumber(a int) int {
	return rand.Intn(a)
}

func NumberToWords(i int) string {
	fmt.Println("writr ", i, "as")
	var wrd string
	switch i {
	case 1:
		wrd = "One"
		fmt.Println(wrd)
	case 2:
		wrd = "two"
		fmt.Println(wrd)
	case 3:
		wrd = "three"
		fmt.Println(wrd)
	case 4:
		wrd = "four"
		fmt.Println(wrd)
	case 5:
		wrd = "five"
		fmt.Println(wrd)
	case 6:
		wrd = "six"
		fmt.Println(wrd)
	case 7:
		wrd = "seven"
		fmt.Println(wrd)
	case 8:
		fmt.Println(wrd)
	case 9:
		wrd = "nine"
		fmt.Println(wrd)
	case 10:
		wrd = "ten"
		fmt.Println(wrd)
	}
	return wrd
}
