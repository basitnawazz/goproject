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
