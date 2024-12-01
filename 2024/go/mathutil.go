package mathutil

import "errors"

// Add returns the sum of two numbers
func Add(x, y int) int {
	return x + y
}

// Subtract returns the difference of two numbers
func Subtract(x, y int) int {
	return x - y
}

// Multiply returns the product of two numbers
func Multiply(x, y int) int {
	return x * y
}

// Divide returns the quotient of two numbers or an error if dividing by zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}