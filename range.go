package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	//range provides the index and the value for each entry
	for _, num := range nums {
		sum += num
		//in this case we didn't use the index just the value
	}
	fmt.Println("sum:", sum)

	//Using the index and the value of each element
	for i, num := range nums {
		if num == 2 {
			fmt.Printf("Index: %d Value: %d\n", i, num)
		}
	}

	//Using range in map structure
	keyValueString := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range keyValueString {
		fmt.Printf("%s ->  %s\n", k, v)
	}
	//Iterates over the key value of each element in the map structure
	for key := range keyValueString {
		fmt.Println("Key:", key)
	}

	//Iterates over Unicode in strings
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
