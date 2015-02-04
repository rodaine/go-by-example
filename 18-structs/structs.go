package main

import "fmt"

// they used lowercase, but I prefer my types to be upper-camel
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51 // auto-dereference :O magic
	fmt.Println(s.age)
}
