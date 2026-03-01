package main

import (
	"fmt"
	"kuis1/calculation"
)

func main() {
	fmt.Println("Program Perkalian")

	result := calculation.Multiply(3, 4)
	fmt.Println("Hasil perkalian:", result)
}
