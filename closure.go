package main

import "fmt"

func main() {
	name := "Kahfi"
	counter := 0

	increment := func() {
		name := "Fauzan"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
