package main

import "fmt"

// func name(argsType)returnType
func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	res1 := plus(1, 2)
	fmt.Println("1+2=", res1)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res2)
}
