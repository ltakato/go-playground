package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoringTimes = 5
const monitoringDelay = 5

func main() {
	// showNames()

	_, age := returnNameAndAge()
	// fmt.Println("I'm", name, "and I'm", age, "years")
	// we can do like this to ignore first return value
	fmt.Println("I'm", age, "years")

	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		// with if, else, else if
		// if command == 1 {
		// 	fmt.Println("Monitoring...")
		// } else if command == 2 {
		// 	fmt.Println("Showing logs...")
		// } else if command == 0 {
		// 	fmt.Println("Exiting program...")
		// } else {
		// 	fmt.Println("Command not recognized")
		// }

		// with switch case
		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Command not recognized")
			os.Exit(-1)
		}
	}
}

func messageOutput(s string) string {
	return s
}

func returnNameAndAge() (string, int) {
	name := "Takatittos"
	age := 28
	return name, age
}

func showIntroduction() {
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
}

func readCommand() int {
	var command int

	// with memory allocating using scanf
	// fmt.Scanf("%d", &command)

	// scan ignores if I type a char/string and assign 0 as default (because command is int)
	fmt.Scan(&command)

	fmt.Println("Chosen command is:", command)
	fmt.Println("Address of my variable command: ", &command)
	fmt.Println("The type of chosen command is", reflect.TypeOf(command))

	return command
}

func showMenu() {
	fmt.Println("Start program, choose an option...")
	fmt.Println("1- start monitoring")
	fmt.Println("2- show logs")
	fmt.Println("0- exit program")
}

func initMonitoring() {
	fmt.Println("Monitoring...")

	// array
	// var sites [4]string

	// sites[0] = "https://random-status-code.herokuapp.com/"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"

	// slice
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	fmt.Println(sites)

	// classical for looping
	// for i := 0; i < len(sites); i++ {
	// 	site := sites[i]
	//}

	for i := 0; i < monitoringTimes; i++ {
		// iterate each item (like for each)
		for i, site := range sites {
			testSite(site, i)
		}
		time.Sleep(monitoringDelay * time.Second)
	}

	fmt.Println("")
}

func testSite(site string, index int) {
	resp, _ := http.Get(site)
	statusCode := resp.StatusCode

	fmt.Println("Testing site", index, ":", site)

	if statusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully!")
	} else {
		fmt.Println("Site:", site, "has problems. Status code:", statusCode)
	}
}

func showNames() {
	// it's a Slice! it's an array, but we don't need to worry about length size
	names := []string{"Anya", "Yor", "Loid"}

	fmt.Println(names)

	fmt.Println("The size of slice is:", len(names), "and has capacity:", cap(names))

	names = append(names, "Bond")

	fmt.Println(names)
	// slice just get doubled in capacity!
	fmt.Println("The size of slice is:", len(names), "and has capacity:", cap(names))
	fmt.Println(reflect.TypeOf(names))
}
