package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Kahfi")
	//helper.sayGoodbye("Kahfi") // error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // error
}
