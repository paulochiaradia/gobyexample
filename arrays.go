package main

import "fmt"

// "Static data implementation"
func main() {
	var a [5]int
	fmt.Println("emp:", a) //print "a" array

	a[4] = 100
	fmt.Println("set:", a)    //print array with a[4] insertion
	fmt.Println("get:", a[4]) //print the element a[4]

	fmt.Println("len:", len(a)) //print the length of the array
}
