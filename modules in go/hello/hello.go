package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Glad")
	fmt.Println(message)
}
