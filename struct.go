package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var kahfi Customer
	kahfi.Name = "Kahfi"
	kahfi.Address = "Indonesia"
	kahfi.Age = 21

	kahfi.sayHi("Widiana")
	kahfi.sayHuuu()

	//fmt.Println(kahfi.Name)
	//fmt.Println(kahfi.Address)
	//fmt.Println(kahfi.Age)
	//
	//widi := Customer{
	//	Name:    "Widiana",
	//	Address: "Soreang",
	//	Age:     22,
	//}
	//fmt.Println(widi)
	//
	//fauzan := Customer{"Fauzan", "Bandung", 21}
	//fmt.Println(fauzan)
}
