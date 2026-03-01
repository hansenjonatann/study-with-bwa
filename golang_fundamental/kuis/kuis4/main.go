package main

import (
	"errors"
	"fmt"
)

func main() {

	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	result, err := calculate(10, 2, "+")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

}

func sum(value []int) int {
	var total int
	for _, val := range value {
		total = total + val
	}
	return total
}

func calculate(number1, number2 int, operate string) (int, error) {
	var result int
	var err error
	switch operate {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		err = errors.New("Unknown operation")
	}
	return result, err
}
