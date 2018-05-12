package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		oddOrEven := "odd"
		if number%2 == 0 {
			oddOrEven = "even"
		}
		fmt.Println(number, "is", oddOrEven)
	}
}
