package main

import (
	"Greetings"
	"fmt"
)

func main() {

	var name string
	fmt.Print("Enter Your name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println(Greetings.ShowGreetings(name))
}
