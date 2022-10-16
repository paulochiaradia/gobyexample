package main

import "fmt"

// Variadic function can be called with any number of trailing arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 5, 6, 9)

	numeros := []int{1, 2, 3, 7, 7, 7, 9}
	sum(numeros...)
}
