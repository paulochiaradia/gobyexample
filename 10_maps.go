package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int) //make(map[key-type]val-type)
	m["k1"] = 10
	m["k2"] = 20
	fmt.Println("map:", m) //printing the map

	v1 := m["k1"] //attributing the value of the pair with the key k1
	fmt.Println("V1:", v1)

	delete(m, "k2") //deleting an element from m map
	fmt.Println(m)

	value, present := m["k1"] //return the value and if the element present or not in the map
	fmt.Println("present:", present)
	fmt.Println("value:", value)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
