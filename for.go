package main

import "fmt"

func main() {
	//simple condition(while)
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//A classic initial/condition/after
	for j := 7; j <= 10; j++ {
		fmt.Println(j)
	}

	//without a condition
	for {
		fmt.Println("loop")
		break
	}

	//with  continue in the middle of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			//jump to the next iteration of the loop
			continue
		}
		fmt.Println(n)
	}
}
