package main

import "fmt"

// go support
func values() (int, int) {
	return 7, 3
}
func main() {
	a, b := values()
	fmt.Println(a, b)

	//Subset of the returned values
	_, c := values()
	fmt.Println(c)

}
