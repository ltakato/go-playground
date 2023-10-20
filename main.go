package main

import "fmt"

func main() {
	var name string = "Takatittos"
	var age int = 28
	var version float32 = 1.21

	fmt.Println("Hello, mr.", name, "your age is", age)
	fmt.Println("This program is in version", version)
}

func messageOutput(s string) string {
	return s
}
