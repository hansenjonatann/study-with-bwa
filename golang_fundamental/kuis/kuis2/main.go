package main

import "fmt"

func main() {
	title := "Golang the best language"

	for index, value := range title {
		if index%2 == 0 {
			fmt.Println("Letter: ", string(value))
		}
		switch value {
		case 'a', 'i', 'u', 'e', 'o':
			fmt.Println("Huruf Vokal :", string(value))

		}
	}
}
