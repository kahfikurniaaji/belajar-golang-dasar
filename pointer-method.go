package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {
	kahfi := Man{"Kahfi"}
	kahfi.Married()

	fmt.Println(kahfi.Name)
}
