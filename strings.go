package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Kahfi Kurnia", "Kahfi"))
	fmt.Println(strings.Contains("Kahfi Kurnia", "Fauzan"))

	fmt.Println(strings.Split("Kahfi Kurnia Aji", " "))

	fmt.Println(strings.ToLower("Kahfi Kurnia Aji"))
	fmt.Println(strings.ToUpper("Kahfi Kurnia Aji"))
	fmt.Println(strings.ToTitle("kahfi kurnia aji"))

	fmt.Println(strings.Trim("          Kahfi Kurnia        ", " "))
	fmt.Println(strings.ReplaceAll("Kahfi Widi Kahfi", "Kahfi", "Fauzan"))
}
