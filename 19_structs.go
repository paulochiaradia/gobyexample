package main

import "fmt"

type person struct {
	name string
	age  int
} //This person struct type has name and age fields.

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

// newPerson constructs a new person struct with the given name and age.
func main() {
	fmt.Println(person{"bob", 20})      //This syntax creates a new struct.
	fmt.Println(person{name: "Philly"}) //Omitted fields will be zero-valued.
	fmt.Println(&person{
		name: "Paulo",
		age:  21,
	}) //An & prefix yields a pointer to the struct.
	fmt.Println(newPerson("Jota", 20))
	test := newPerson("Test", 30)
	fmt.Println(test.name, test.age) //Access struct fields with a dot.

	test2 := &test //pointer to pointer
	fmt.Println((*test2).age)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	//Structs are mutable.
}
