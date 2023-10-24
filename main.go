package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string = "Takatittos"
	var name = "Takatittos"    // with type inference
	short_name := "Takatittos" // short variable declaration operator
	var age int = 28
	var version float32 = 1.21

	fmt.Println("Hello, mr.", name, "your age is", age)
	fmt.Println("This program is in version", version)
	fmt.Println("The value of short_name is:", short_name)

	// getting infered type
	fmt.Println("The type of variable age is", reflect.TypeOf(age))
	fmt.Println("The type of variable version is", reflect.TypeOf(version))

	fmt.Println("Start program, choose an option...")
	fmt.Println("1- start monitoring")
	fmt.Println("2- show logs")
	fmt.Println("0- exit program")

	var command int

	// with memory allocating using scanf
	// fmt.Scanf("%d", &command)

	// scan ignores if I type a char/string and assign 0 as default (because command is int)
	fmt.Scan(&command)

	fmt.Println("Chosen command is:", command)
	fmt.Println("Address of my variable command: ", &command)
	fmt.Println("The type of chosen command is", reflect.TypeOf(command))

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing logs...")
	} else if command == 0 {
		fmt.Println("Exiting program...")
	} else {
		fmt.Println("Command not recognized")
	}
}

func messageOutput(s string) string {
	return s
}
