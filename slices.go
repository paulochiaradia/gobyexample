package main

import "fmt"

// "Dynamic data implementation"
func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)     // make a empty slice of string size 3
	s[0] = "a"                 // set a as 0 element of the slice
	s[1] = "b"                 // set b as 1 element of the slice
	s[2] = "c"                 // set c as 2 element of the slice
	fmt.Println("set:", s)     //print the slice
	fmt.Println("get:", s[2])  //print the element in index 2 in the slice
	fmt.Println("len", len(s)) //print length of slice s

	s = append(s, "d")      //append insert the element in end of the slice and return a new slice
	s = append(s, "e", "f") //append insert the element in end of the slice and return a new slice
	fmt.Println("apd:", s)  //print the new slice with more elements in the end

	c := make([]string, len(s)) // make a new slice with the length of the slice s
	copy(c, s)                  //copy the elements of the slice s to slice c
	fmt.Println("cpy:", c)      //print the c slice

	l := s[2:5]            //slice of the slice
	fmt.Println("slc1", l) // print the slice od the slice

	j := s[:5]             //slice of the slice
	fmt.Println("slc2", j) // print the slice od the slice e
}
