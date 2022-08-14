package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (Person Person) GetName() string {
	return Person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var kahfi Person
	kahfi.Name = "Kahfi"

	SayHello(kahfi)

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}
